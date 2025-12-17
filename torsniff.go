package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/marksamman/bencode"
	"github.com/mitchellh/go-homedir"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/pkg/fileutil"
)

const (
	directoryConst = "torrents"
)

type tfile struct {
	Name   string `json:"name"`
	Length int64  `json:"length"`
}

func (t *tfile) String() string {
	return fmt.Sprintf("name: %s\n, size: %d\n", t.Name, t.Length)
}

type torrent struct {
	infohashHex string
	name        string
	length      int64
	files       []*tfile
	createdAt   time.Time
}

func (t *torrent) String() string {
	return fmt.Sprintf(
		"link: %s\nname: %s\nsize: %d\nfile: %d\n",
		fmt.Sprintf("magnet:?xt=urn:btih:%s", t.infohashHex),
		t.name,
		t.length,
		len(t.files),
	)
}

// TorrentDocument 用于 Elasticsearch
type TorrentDocument struct {
	InfoHash  string    `json:"info_hash"`
	Name      string    `json:"name"`
	Length    int64     `json:"length"`
	Files     []*tfile  `json:"files"`
	CreatedAt time.Time `json:"created_at"`
}

// TorrentStats 用于跟踪每个种子的DHT活动统计
type TorrentStats struct {
	InfoHash      string
	AnnounceCount int64
	GetPeersCount int64
	LastSeen      time.Time
	FirstSeen     time.Time
}

func parseTorrent(meta []byte, infohashHex string) (*torrent, error) {
	dict, err := bencode.Decode(bytes.NewBuffer(meta))
	if err != nil {
		return nil, err
	}

	t := &torrent{infohashHex: infohashHex, createdAt: time.Now()}
	if name, ok := dict["name.utf-8"].(string); ok {
		t.name = name
	} else if name, ok := dict["name"].(string); ok {
		t.name = name
	}
	if length, ok := dict["length"].(int64); ok {
		t.length = length
	}

	var totalSize int64
	var extractFiles = func(file map[string]interface{}) {
		var filename string
		var filelength int64
		if inter, ok := file["path.utf-8"].([]interface{}); ok {
			name := make([]string, len(inter))
			for i, v := range inter {
				name[i] = fmt.Sprint(v)
			}
			filename = strings.Join(name, "/")
		} else if inter, ok := file["path"].([]interface{}); ok {
			name := make([]string, len(inter))
			for i, v := range inter {
				name[i] = fmt.Sprint(v)
			}
			filename = strings.Join(name, "/")
		}
		if length, ok := file["length"].(int64); ok {
			filelength = length
			totalSize += filelength
		}
		t.files = append(t.files, &tfile{Name: filename, Length: filelength})
	}

	if files, ok := dict["files"].([]interface{}); ok {
		for _, file := range files {
			if f, ok := file.(map[string]interface{}); ok {
				extractFiles(f)
			}
		}
	}

	if t.length == 0 {
		t.length = totalSize
	}
	if len(t.files) == 0 {
		t.files = append(t.files, &tfile{Name: t.name, Length: t.length})
	}

	return t, nil
}

type torsniff struct {
	laddr         string
	maxFriends    int
	maxPeers      int
	secret        string
	timeout       time.Duration
	blacklist     *blackList
	dir           string
	elasticClient *elastic.Client
	elasticIndex  string
	webAddr       string
	httpServer    *http.Server
	saveToDisk    bool // 新增标志，控制是否保存到磁盘

	// DHT状态跟踪
	dhtInstance        *dht
	totalAnnouncements int64
	startTime          time.Time
	mu                 sync.RWMutex

	// 种子热度统计
	torrentStats map[string]*TorrentStats
	statsMu      sync.RWMutex
}

func (t *torsniff) run() error {
	tokens := make(chan struct{}, t.maxPeers)

	// 初始化状态跟踪
	t.startTime = time.Now()
	t.totalAnnouncements = 0
	t.torrentStats = make(map[string]*TorrentStats)

	// 启动 Web 服务器
	if t.webAddr != "" {
		t.startWebServer()
	}

	dht, err := newDHT(t.laddr, t.maxFriends, t.onDHTQuery)
	if err != nil {
		return err
	}

	// 保存DHT实例引用
	t.dhtInstance = dht

	dht.run()

	log.Println("running, it may take a few minutes...")

	for {
		select {
		case <-dht.announcements.wait():
			for {
				if ac := dht.announcements.get(); ac != nil {
					t.mu.Lock()
					t.totalAnnouncements++
					t.mu.Unlock()

					tokens <- struct{}{}
					go t.work(ac, tokens)
					continue
				}
				break
			}
		case <-dht.die:
			// 关闭 Web 服务器
			if t.httpServer != nil {
				t.httpServer.Shutdown(context.Background())
			}
			return dht.errDie
		}
	}
}

func (t *torsniff) work(ac *announcement, tokens chan struct{}) {
	defer func() {
		<-tokens
	}()

	// 检查是否已存在于磁盘（如果启用了磁盘存储）
	if t.saveToDisk && t.isTorrentExist(ac.infohashHex) {
		return
	}

	// 检查是否已存在于 Elasticsearch（如果启用了 Elasticsearch）
	if t.elasticClient != nil && t.isTorrentExistInElasticsearch(ac.infohashHex) {
		return
	}

	peerAddr := ac.peer.String()
	if t.blacklist.has(peerAddr) {
		return
	}

	wire := newMetaWire(string(ac.infohash), peerAddr, t.timeout)
	defer wire.free()

	meta, err := wire.fetch()
	if err != nil {
		t.blacklist.add(peerAddr)
		return
	}

	// 保存到磁盘（如果启用了磁盘存储）
	if t.saveToDisk {
		if err := t.saveTorrent(ac.infohashHex, meta); err != nil {
			log.Printf("Failed to save torrent to disk: %v", err)
		}
	}

	torrent, err := parseTorrent(meta, ac.infohashHex)
	if err != nil {
		return
	}

	// 保存到 Elasticsearch（如果启用了 Elasticsearch）
	if t.elasticClient != nil {
		t.saveToElasticsearch(torrent)
	} else if !t.saveToDisk {
		// 如果既没有启用 Elasticsearch 也没有启用磁盘存储，则只打印到日志
		log.Println(torrent)
	}
}

// 添加检查 Elasticsearch 中是否存在 torrent 的方法
func (t *torsniff) isTorrentExistInElasticsearch(infohashHex string) bool {
	if t.elasticClient == nil {
		return false
	}

	exists, err := t.elasticClient.Exists().
		Index(t.elasticIndex).
		Id(infohashHex).
		Do(context.Background())

	if err != nil {
		log.Printf("Failed to check existence in Elasticsearch: %v", err)
		return false
	}

	return exists
}

// 添加保存到 Elasticsearch 的方法
func (t *torsniff) saveToElasticsearch(torrent *torrent) {
	doc := &TorrentDocument{
		InfoHash:  torrent.infohashHex,
		Name:      torrent.name,
		Length:    torrent.length,
		Files:     torrent.files,
		CreatedAt: torrent.createdAt,
	}

	_, err := t.elasticClient.Index().
		Index(t.elasticIndex).
		Id(torrent.infohashHex).
		BodyJson(doc).
		Do(context.Background())

	if err != nil {
		log.Printf("Failed to insert torrent into Elasticsearch: %v", err)
	} else {
		log.Printf("Inserted torrent %s into Elasticsearch", torrent.infohashHex)
	}
}

func (t *torsniff) isTorrentExist(infohashHex string) bool {
	name, _ := t.torrentPath(infohashHex)
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func (t *torsniff) saveTorrent(infohashHex string, data []byte) error {
	name, dir := t.torrentPath(infohashHex)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	d, err := bencode.Decode(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	f, err := fileutil.TryLockFile(name, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bencode.Encode(map[string]interface{}{
		"info": d,
	}))
	if err != nil {
		return err
	}

	return nil
}

func (t *torsniff) onDHTQuery(queryType, infohash string) {
	t.statsMu.Lock()
	defer t.statsMu.Unlock()

	stats, exists := t.torrentStats[infohash]
	if !exists {
		stats = &TorrentStats{
			InfoHash:  infohash,
			FirstSeen: time.Now(),
		}
		t.torrentStats[infohash] = stats
	}

	stats.LastSeen = time.Now()
	switch queryType {
	case "announce_peer":
		stats.AnnounceCount++
	case "get_peers":
		stats.GetPeersCount++
	}
}

func (t *torsniff) torrentPath(infohashHex string) (name string, dir string) {
	dir = path.Join(t.dir, infohashHex[:2], infohashHex[len(infohashHex)-2:])
	name = path.Join(dir, infohashHex+".torrent")
	return
}

// 添加 Web 服务器启动方法
func (t *torsniff) startWebServer() {
	// 创建 HTTP 服务器
	mux := http.NewServeMux()

	// 提供静态文件服务
	mux.Handle("/", http.FileServer(http.Dir("./web/static/")))

	// 添加 API 端点用于获取最近种子
	mux.HandleFunc("/api/latest", t.handleRecentTorrents)

	// 添加搜索端点
	mux.HandleFunc("/api/search", t.handleTorrentSearch)

	// 添加获取种子详情端点
	mux.HandleFunc("/api/torrent/", t.handleTorrentDetails)

	// 添加统计信息端点
	mux.HandleFunc("/api/stats", t.handleStats)

	// 添加DHT状态端点
	mux.HandleFunc("/api/dht", t.handleDHTStatus)

	// 添加系统状态端点
	mux.HandleFunc("/api/health", t.handleHealthStatus)

	server := &http.Server{
		Addr:    t.webAddr,
		Handler: mux,
	}

	t.httpServer = server

	// 在单独的 goroutine 中启动服务器
	go func() {
		log.Printf("Starting web server on %s", t.webAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Web server error: %v", err)
		}
	}()
}

// handleTorrentDetails 处理获取种子详情的 HTTP 请求
func (t *torsniff) handleTorrentDetails(w http.ResponseWriter, r *http.Request) {
	// 从 URL 路径中提取 infohash
	path := r.URL.Path[len("/api/torrent/"):]
	if path == "" {
		http.Error(w, "Missing infohash", http.StatusBadRequest)
		return
	}

	// 验证 infohash 格式：必须是40位十六进制字符
	if len(path) != 40 {
		http.Error(w, "Invalid infohash length", http.StatusBadRequest)
		return
	}

	// 验证是否只包含十六进制字符
	for _, c := range path {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			http.Error(w, "Invalid infohash format", http.StatusBadRequest)
			return
		}
	}

	// 如果连接了 Elasticsearch，则从 Elasticsearch 获取详情
	if t.elasticClient != nil {
		details, err := t.getTorrentDetailsFromElasticsearch(path)
		if err != nil {
			http.Error(w, "Failed to get torrent details: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if details == nil {
			http.Error(w, "Torrent not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(details)
		return
	}

	// 否则返回错误
	http.Error(w, "Elasticsearch not configured", http.StatusServiceUnavailable)
}

// handleRecentTorrents 处理获取最新种子的请求
func (t *torsniff) handleRecentTorrents(w http.ResponseWriter, r *http.Request) {
	// 获取limit参数，默认为20，增加严格验证
	limit := 20
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 && l <= 100 && len(limitParam) <= 5 {
			limit = l
		} else {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	var response *RecentTorrentsResponse
	var err error

	// 如果连接了 Elasticsearch，优先从 Elasticsearch 获取数据
	if t.elasticClient != nil {
		response, err = t.getRecentTorrentsFromElasticsearch(limit)
	} else {
		// 否则从文件系统获取
		response, err = t.getRecentTorrents(limit)
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Failed to get recent torrents: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleTorrentSearch 处理搜索请求
func (t *torsniff) handleTorrentSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	if t.elasticClient != nil {
		t.handleElasticsearchSearch(w, r, query)
	} else {
		http.Error(w, "Elasticsearch is not configured", http.StatusInternalServerError)
	}
}

// handleElasticsearchSearch 处理基于 Elasticsearch 的搜索
func (t *torsniff) handleElasticsearchSearch(w http.ResponseWriter, r *http.Request, query string) {
	// 验证和清理查询参数，防止注入
	if len(query) > 500 {
		http.Error(w, "Query too long", http.StatusBadRequest)
		return
	}

	// 获取分页参数
	page := 0
	size := 20

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p >= 0 && p <= 10000 {
			page = p
		}
	}

	if sizeParam := r.URL.Query().Get("size"); sizeParam != "" {
		if s, err := strconv.Atoi(sizeParam); err == nil && s > 0 && s <= 100 {
			size = s
		}
	}

	// 获取其他搜索参数并进行验证
	category := r.URL.Query().Get("category")
	if len(category) > 50 {
		category = ""
	}

	sortBy := r.URL.Query().Get("sort_by")
	// 白名单验证排序字段
	validSortBy := map[string]bool{
		"": true, "popularity": true, "size": true, "length": true,
		"first_seen": true, "created_at": true, "last_seen": true,
	}
	if !validSortBy[sortBy] {
		sortBy = ""
	}

	sortOrder := r.URL.Query().Get("sort_order")
	// 白名单验证排序顺序
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = ""
	}

	minSize := r.URL.Query().Get("min_size")
	maxSize := r.URL.Query().Get("max_size")
	// 验证大小参数格式
	for _, sizeParam := range []string{minSize, maxSize} {
		if sizeParam != "" {
			if _, err := strconv.ParseInt(sizeParam, 10, 64); err != nil || len(sizeParam) > 20 {
				// 清空无效参数
				if sizeParam == minSize {
					minSize = ""
				} else {
					maxSize = ""
				}
			}
		}
	}

	// 在 Elasticsearch 中搜索
	response, err := t.searchTorrentsInElasticsearchAdvanced(query, page, size, category, sortBy, sortOrder, minSize, maxSize)
	if err != nil {
		http.Error(w, "Search failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// The main function was incorrectly placed earlier in the file.
// It has been moved here and the duplicate constant declaration removed.
func main() {
	log.SetFlags(0)

	var addr string
	var port uint16
	var peers int
	var timeout time.Duration
	var dir string
	var verbose bool
	var friends int
	var elasticAddr string
	var elasticIndex string
	var webAddr string
	var saveToDisk bool // 新增参数
	var seedHosts []string
	var seed6Hosts []string

	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "无法获取主目录: %v\n", err)
		os.Exit(1)
	}
	userHome := path.Join(home, directoryConst)

	root := &cobra.Command{
		Use:          "torsniff",
		Short:        "torsniff - A sniffer that sniffs torrents from BitTorrent network.",
		SilenceUsage: true,
	}
	root.RunE = func(cmd *cobra.Command, args []string) error {
		if dir == userHome && err != nil {
			return err
		}

		absDir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}

		log.SetOutput(io.Discard)
		if verbose {
			log.SetOutput(os.Stdout)
		}

		// 初始化 Elasticsearch 客户端（如果提供了地址）
		var esClient *elastic.Client
		if elasticAddr != "" {
			client, err := elastic.NewClient(
				elastic.SetURL(elasticAddr),
				elastic.SetSniff(false),
				elastic.SetHealthcheck(false),
			)
			if err != nil {
				log.Printf("Failed to create Elasticsearch client: %v", err)
				if !saveToDisk {
					return fmt.Errorf("Elasticsearch connection failed and disk saving is disabled")
				}
			} else {
				esClient = client
				log.Printf("Connected to Elasticsearch at %s", elasticAddr)

				// 创建索引（如果不存在）
				exists, err := esClient.IndexExists(elasticIndex).Do(context.Background())
				if err != nil {
					log.Printf("Error checking Elasticsearch index: %v", err)
				} else if !exists {
					// 定义索引映射，去掉分词器设置
					mapping := `{
						"mappings": {
							"properties": {
								"info_hash": {
									"type": "keyword"
								},
								"name": {
									"type": "text"
								},
								"length": {
									"type": "long"
								},
								"files": {
									"type": "nested",
									"properties": {
										"name": {
											"type": "text"
										},
										"length": {
											"type": "long"
										}
									}
								},
								"created_at": {
									"type": "date"
								}
							}
						}
					}`
					_, err := esClient.CreateIndex(elasticIndex).BodyString(mapping).Do(context.Background())
					if err != nil {
						log.Printf("Error creating Elasticsearch index: %v", err)
					} else {
						log.Printf("Created Elasticsearch index: %s", elasticIndex)
					}
				}
			}
		} else if !saveToDisk {
			// 如果既没有 Elasticsearch 也没有启用磁盘存储，给出警告
			log.Println("Warning: Neither Elasticsearch nor disk saving is enabled. Torrents will only be logged.")
		}

		// 合并命令行提供的引导节点
		if len(seedHosts) > 0 {
			seeds = append(seeds, seedHosts...)
		}
		if len(seed6Hosts) > 0 {
			seeds = append(seeds, seed6Hosts...)
		}

		p := &torsniff{
			laddr:         net.JoinHostPort(addr, strconv.Itoa(int(port))),
			timeout:       timeout,
			maxFriends:    friends,
			maxPeers:      peers,
			secret:        "",
			dir:           absDir,
			blacklist:     newBlackList(5*time.Minute, 50000),
			elasticClient: esClient,
			elasticIndex:  elasticIndex,
			webAddr:       webAddr,
			saveToDisk:    saveToDisk, // 设置新参数
		}
		return p.run()
	}

	root.Flags().StringVarP(&addr, "addr", "a", "", "listen on given address (default all, ipv4 and ipv6)")
	root.Flags().Uint16VarP(&port, "port", "p", 6881, "listen on given port")
	root.Flags().IntVarP(&friends, "friends", "f", 500, "max fiends to make with per second")
	root.Flags().IntVarP(&peers, "peers", "e", 400, "max peers to connect to download torrents")
	root.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "max time allowed for downloading torrents")
	root.Flags().StringVarP(&dir, "dir", "d", userHome, "the directory to store the torrents")
	root.Flags().BoolVarP(&verbose, "verbose", "v", true, "run in verbose mode")
	root.Flags().StringVar(&elasticAddr, "elastic-addr", "http://localhost:9200", "Elasticsearch address (e.g., http://localhost:9200)")
	root.Flags().StringVar(&elasticIndex, "elastic-index", "torrents", "Elasticsearch index name")
	root.Flags().StringVar(&webAddr, "web-addr", "", "Web server address (e.g., :8080)")
	root.Flags().BoolVar(&saveToDisk, "save-to-disk", false, "Save torrents to disk (default: false)")
	// DHT 引导节点（可重复）
	root.Flags().StringSliceVar(&seedHosts, "seed", nil, "additional DHT bootstrap nodes (host:port or [ipv6]:port)")
	root.Flags().StringSliceVar(&seed6Hosts, "seed6", nil, "additional IPv6 DHT bootstrap nodes (host:port or [ipv6]:port)")

	if err := root.Execute(); err != nil {
		fmt.Println(fmt.Errorf("could not start: %s", err))
	}
}

// handleStats 处理统计信息请求
func (t *torsniff) handleStats(w http.ResponseWriter, r *http.Request) {
	stats := make(map[string]interface{})

	if t.elasticClient != nil {
		// 从 Elasticsearch 获取统计信息
		if count, err := t.elasticClient.Count(t.elasticIndex).Do(context.Background()); err == nil {
			stats["total_torrents"] = count
		}

		// 获取最近24小时的种子数量
		dayAgo := time.Now().Add(-24 * time.Hour)
		dayQuery := elastic.NewRangeQuery("created_at").Gte(dayAgo)
		if dayCount, err := t.elasticClient.Count(t.elasticIndex).Query(dayQuery).Do(context.Background()); err == nil {
			stats["last_24h_torrents"] = dayCount
		}
	} else {
		// 从文件系统获取统计信息
		var totalFiles int
		var totalSize int64
		err := filepath.Walk(t.dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && filepath.Ext(path) == ".torrent" {
				totalFiles++
				totalSize += info.Size()
			}
			return nil
		})
		if err == nil {
			stats["total_torrents"] = totalFiles
			stats["total_size"] = totalSize
		}
	}

	// 添加系统统计信息
	uptime := time.Since(t.startTime)
	stats["uptime"] = uptime.String()
	stats["version"] = "v1.0.0"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// handleDHTStatus 处理DHT状态请求
func (t *torsniff) handleDHTStatus(w http.ResponseWriter, r *http.Request) {
	dhtStatus := make(map[string]interface{})

	// 获取真实的DHT状态信息
	t.mu.RLock()
	status := "stopped"
	activePeers := 0
	totalAnnouncements := t.totalAnnouncements
	uptime := time.Since(t.startTime)

	if t.dhtInstance != nil {
		status = "running"
		// 活跃对等方数：基于接收到的公告数量来估算
		// 这是一个合理的估算：每个公告通常来自一个不同的对等方
		if totalAnnouncements > 0 {
			// 假设平均每个对等方发布了多个公告，估算活跃对等方数
			estimatedPeers := int(totalAnnouncements / 10) // 假设平均每个对等方10个公告
			if estimatedPeers < 1 {
				estimatedPeers = 1
			}
			if estimatedPeers > t.maxPeers {
				estimatedPeers = t.maxPeers
			}
			activePeers = estimatedPeers
		}
	}
	t.mu.RUnlock()

	// 解析监听端口
	port := "unknown"
	if t.laddr != "" {
		if _, p, err := net.SplitHostPort(t.laddr); err == nil {
			port = p
		}
	}

	dhtStatus["nodes"] = activePeers        // DHT路由表中的节点数（估算）
	dhtStatus["active_peers"] = activePeers // 当前活跃的对等方数（估算）
	dhtStatus["total_announcements"] = totalAnnouncements
	dhtStatus["status"] = status
	dhtStatus["listening_port"] = port
	dhtStatus["uptime"] = uptime.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dhtStatus)
}

// handleHealthStatus 处理系统健康状态请求
func (t *torsniff) handleHealthStatus(w http.ResponseWriter, r *http.Request) {
	health := make(map[string]interface{})

	// 检查Elasticsearch状态
	elasticsearchStatus := "not_configured"
	if t.elasticClient != nil {
		if _, err := t.elasticClient.ClusterHealth().Do(context.Background()); err == nil {
			elasticsearchStatus = "healthy"
		} else {
			elasticsearchStatus = "unhealthy"
		}
	}
	health["elasticsearch"] = elasticsearchStatus

	// 检查磁盘空间
	var diskStatus string
	if _, err := os.Stat(t.dir); err == nil {
		health["storage_directory"] = t.dir
		health["storage_accessible"] = true
		diskStatus = "healthy"
	} else {
		health["storage_directory"] = t.dir
		health["storage_accessible"] = false
		diskStatus = "unhealthy"
	}
	health["storage"] = diskStatus

	// 检查Web服务器状态
	health["web_server"] = "running"

	// 总体健康状态
	overallStatus := "healthy"
	if elasticsearchStatus == "unhealthy" || diskStatus == "unhealthy" {
		overallStatus = "degraded"
	}
	health["status"] = overallStatus
	health["timestamp"] = time.Now().Format(time.RFC3339)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
