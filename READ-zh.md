# torsniff - 一个从 BitTorrent 网络中嗅探种子的工具

torsniff 是一个可以从 BitTorrent 网络中嗅探种子并默认将其保存到 Elasticsearch 的工具。

## 特性

- 高性能，每分钟可以嗅探数千个种子
- 易于使用，无需配置
- 默认将种子保存到 Elasticsearch（可自定义）
- 可选择将种子保存在本地目录
- 内置 Web 界面用于浏览收集的种子

## 安装

从 releases 页面下载预编译的二进制文件。

## 用法

```
torsniff - 一个从 BitTorrent 网络中嗅探种子的工具。
Usage:
  torsniff [flags]

Flags:
  -a, --addr string         监听的地址 (默认所有 IPv4 和 IPv6 地址)
  -d, --dir string          存储种子的目录 (默认 "$HOME/torrents")
  -e, --peers int           下载种子的最大连接数 (默认 400)
  -f, --friends int         每秒最大建立连接数 (默认 500)
  -h, --help                帮助信息
  -p, --port uint16         监听的端口 (默认 6881)
  -t, --timeout duration    下载种子的超时时间 (默认 10s)
  -v, --verbose             是否显示详细信息 (默认开启)
      --elastic-addr string Elasticsearch 地址 (默认 "http://localhost:9200")
      --elastic-index string Elasticsearch 索引名称 (默认 "torrents")
      --save-to-disk        保存种子到磁盘 (默认: false)
      --web-addr string     Web 服务器地址 (例如: :8080)
  --seed string[]       额外的 DHT 引导节点 (host:port 或 [ipv6]:port)
  --seed6 string[]      额外的 IPv6 DHT 引导节点 (host:port 或 [ipv6]:port)
```

### 示例

#### 使用默认设置运行（Elasticsearch）

```
torsniff
```

#### 使用自定义 Elasticsearch 设置运行

```
torsniff --elastic-addr http://your-elasticsearch-host:9200 --elastic-index my_torrents
```

#### 使用 Web 界面运行

```
torsniff --web-addr :8080
```

#### 同时启用 Elasticsearch 和 Web 界面

```
torsniff --elastic-addr http://your-elasticsearch-host:9200 --elastic-index my_torrents --web-addr :8080
```

#### 使用磁盘存储而非 Elasticsearch

```
torsniff --save-to-disk --dir ~/my-torrents
```

## 工作原理

1. torsniff 作为一个 DHT 节点加入 BitTorrent 网络
2. 它监听来自其他节点的可用种子公告
3. 当收到公告时，它会连接到对等方并下载种子元数据
4. 种子信息默认保存到 Elasticsearch（或根据配置保存到磁盘）
5. 内置的 Web 界面允许浏览和搜索收集的种子

## 要求

- 公共 IP 地址可以获得最佳性能 (可选但建议)
- 端口 6881 (UDP) 开放以进行 DHT 通信
- Elasticsearch 实例（除非使用磁盘存储）

## IPv6 支持

本项目已实现对 IPv6 DHT 的支持（BEP-32）。要点如下：

- 双栈监听：默认在地址未指定时同时绑定 IPv4 与 IPv6（分别绑定 `0.0.0.0:PORT` 与 `[::]:PORT`）。
- 路由发现：在 `find_node` 请求中携带 `want: ["n4","n6"]`，可从对端同时获取 `nodes`（IPv4）与 `nodes6`（IPv6）。
- 解析扩展：支持解析 `nodes6` 中的 IPv6 compact 节点并加入路由拓展。
- 兼容性：对 IPv4 行为无影响，IPv6 不可用时将自动继续使用 IPv4。

使用建议（Windows PowerShell）：

```powershell
# 仅IPv4监听
./torsniff.exe -a 0.0.0.0 -p 6881

# 仅IPv6监听
./torsniff.exe -a :: -p 6881

# 双栈（默认，addr留空即双栈）
./torsniff.exe -p 6881

# 指定额外IPv6引导节点
./torsniff.exe -p 6881 --seed6 "[2001:4860:4860::8888]:6881" --seed6 "router.bittorrent.com:6881"
```

防火墙与网络：
- 请在系统与网络设备上放行 UDP 6881（IPv4/IPv6）。
- 建议在具有全局可达 IPv6 地址的环境中运行，以获得最佳连通性。

## 注意事项

- 可能需要几分钟才能开始接收种子
- 能够嗅探到的种子越多，性能越好
- 收集的种子默认存储在 Elasticsearch 中（可通过 --save-to-disk 更改）
- Web 界面提供了便捷的方式来浏览和搜索收集的种子

## 许可证

MIT