# 磁力魔 (Cilimo.com) - BitTorrent DHT 种子搜索引擎

![磁力搜索](https://img.shields.io/badge/磁力搜索-DHT%20Search-blue)
![BitTorrent](https://img.shields.io/badge/BitTorrent-P2P-green)
![开源](https://img.shields.io/badge/开源-Open%20Source-orange)
![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow)

**磁力魔 (Cilimo.com)** 是一个基于 BitTorrent DHT 网络的高性能种子资源搜索引擎，使用 **torsniff** 工具从 DHT 网络中嗅探种子信息，并提供强大的磁力链接搜索服务。

🌐 **官方网站**: [https://cilimo.com](https://cilimo.com)  
⚡ **核心技术**: [torsniff](https://github.com/fanpei91/torsniff) - 从 BitTorrent 网络中嗅探种子的高性能工具

---

## 📋 目录

- [项目简介](#项目简介)
- [核心技术 - torsniff](#核心技术---torsniff)
- [核心功能](#核心功能)
- [技术架构](#技术架构)
- [部署运行](#部署运行)
- [使用指南](#使用指南)
- [技术原理](#技术原理)
- [推荐下载工具](#推荐下载工具)
- [热门资源](#热门资源)
- [开发指南](#开发指南)
- [技术参考](#技术参考)
- [免责声明](#免责声明)
- [联系我们](#联系我们)

---

## 🌟 项目简介

磁力魔（Cilimo.com）是一个专业的 **BitTorrent DHT 种子搜索引擎**，基于高性能的 **torsniff** 工具构建。通过持续嗅探和索引分布式哈希表（DHT）网络中的种子信息，为用户提供实时、快速、海量的磁力链接搜索服务。

### 什么是磁力链接？

磁力链接（Magnet URI）是一种无需中心服务器即可下载文件的新型链接格式，通过 BitTorrent 协议实现点对点（P2P）文件共享。磁力链接不依赖于 Tracker 服务器，而是通过 DHT 网络直接连接到拥有文件的节点。

### 为什么选择磁力魔？

- ✅ **高性能** - 基于 torsniff，每分钟可嗅探数千个种子
- ✅ **海量资源库** - 索引数百万级别的种子信息
- ✅ **实时更新** - 持续爬取 DHT 网络最新资源
- ✅ **快速搜索** - 高效的搜索算法，秒级响应
- ✅ **双栈支持** - 完整支持 IPv4 和 IPv6 DHT 网络（BEP-32）
- ✅ **简洁界面** - 清爽的用户界面，内置 Web 管理
- ✅ **免费使用** - 完全免费，无需注册
- ✅ **隐私保护** - 不记录用户搜索历史
- ✅ **开源技术** - 基于开源 torsniff 工具构建

---

## 🔧 核心技术 - torsniff

磁力魔的核心引擎是 **torsniff**，一个用 Go 语言编写的高性能 BitTorrent DHT 种子嗅探工具。

### torsniff 主要特性

- ⚡ **高性能** - 每分钟可嗅探数千个种子
- 🔌 **易于部署** - 无需复杂配置即可运行
- 💾 **灵活存储** - 默认使用 Elasticsearch，也支持本地文件存储
- 🌐 **Web 界面** - 内置 Web 界面用于浏览和管理收集的种子
- 🔄 **实时嗅探** - 持续监听 DHT 网络中的种子公告
- 🌍 **IPv6 支持** - 完整实现 BEP-32，支持 IPv4/IPv6 双栈

### torsniff 工作流程

```
1. 加入 DHT 网络 → 作为 DHT 节点加入 BitTorrent 分布式网络
2. 监听公告 → 持续监听其他节点的种子资源公告
3. 连接 Peer → 收到公告后连接到对等节点
4. 下载元数据 → 获取种子的完整元数据信息
5. 存储索引 → 保存到 Elasticsearch 或磁盘，建立搜索索引
6. Web 展示 → 通过 Web 界面提供搜索和浏览功能
```

### 技术要求

- **运行环境**: Linux/Windows/macOS
- **网络要求**: 公网 IP 地址（推荐，可获得最佳性能）
- **端口开放**: UDP 6881 用于 DHT 通信
- **存储后端**: Elasticsearch（推荐）或本地磁盘
- **可选组件**: Web 界面服务器

---

## 🚀 核心功能

### 1. 磁力搜索

**基础搜索**
- 在首页搜索框输入关键词
- 支持中英文、数字、符号等多种搜索
- 智能匹配相关资源
- 按相关度、时间、大小等排序

**高级搜索**
- 精确匹配搜索
- 文件类型筛选（视频、音频、文档、软件等）
- 文件大小过滤
- 时间范围限定
- 多关键词组合搜索

### 2. 种子信息展示

每个搜索结果包含：
- 📝 **资源名称** - 完整的文件或文件夹名称
- 📦 **文件大小** - 资源总大小
- 📅 **收录时间** - 资源被索引的时间
- 🔗 **磁力链接** - 可直接复制的 Magnet URI
- 📊 **文件列表** - 多文件种子的详细文件列表
- ⚡ **热度标识** - 资源活跃度和下载人数参考

### 3. 热门资源推荐

网站首页展示热门资源分类：

📚 **热门小说**
- 修仙类：凡人修仙传、独步成仙、绝世神尊等
- 玄幻类：吞噬星空、九星霸体诀、牧神记等
- 武侠类：修罗武神、绝代神主等
- 都市类：盖世神医、大医无疆等
- 历史类：晚明、明末边军一小兵、割鹿记等

🎬 **影视资源**
- 最新电影
- 热门电视剧
- 纪录片
- 动漫番剧

🎮 **游戏软件**
- PC 游戏
- 实用软件
- 开发工具

### 4. 一键复制磁力链接

- 点击磁力链接图标即可复制
- 支持批量导出
- 兼容所有主流下载工具

---

## 📖 使用指南

### 第一步：访问网站

打开浏览器访问 [https://cilimo.com](https://cilimo.com)

### 第二步：搜索资源

#### 基础搜索方法

1. 在首页搜索框输入关键词，例如：
   ```
   Python 教程
   电影名称
   软件名称
   ```

2. 点击"搜索"按钮或按 Enter 键

3. 浏览搜索结果列表

#### 高级搜索技巧

**使用引号进行精确搜索**
```
"完整的文件名"
```

**排除特定关键词（使用减号）**
```
Python -2.7
```

**指定文件类型**
```
教程 filetype:pdf
教程 filetype:mp4
```

**组合多个关键词**
```
Python 入门 教程 2024
```

### 第三步：查看详情

点击搜索结果可以查看：
- 完整的文件列表
- 各文件的大小
- 种子的详细信息
- 相关资源推荐

### 第四步：获取磁力链接

1. 点击"复制磁力链接"按钮
2. 磁力链接自动复制到剪贴板
3. 格式类似：`magnet:?xt=urn:btih:...`

### 第五步：开始下载

1. 打开您喜欢的 BitTorrent 下载工具
2. 粘贴磁力链接
3. 选择保存位置
4. 开始下载

---

## 🏗️ 技术架构

### 系统架构图

```
┌─────────────────────────────────────────────────────────┐
│                      用户层                              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐             │
│  │ Web浏览器 │  │ 移动设备  │  │ API客户端 │             │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘             │
└───────┼─────────────┼─────────────┼────────────────────┘
        │             │             │
        └─────────────┴─────────────┘
                      │ HTTP/HTTPS
┌─────────────────────▼─────────────────────────────────┐
│                   应用层                                │
│  ┌──────────────────────────────────────────────┐     │
│  │         Web 界面 & API 服务                   │     │
│  │  - 搜索接口  - 浏览界面  - 管理后台          │     │
│  └─────────────────┬────────────────────────────┘     │
└────────────────────┼──────────────────────────────────┘
                     │
┌────────────────────▼──────────────────────────────────┐
│                  存储层                                 │
│  ┌──────────────────────────┐  ┌──────────────────┐  │
│  │    Elasticsearch         │  │   本地磁盘存储    │  │
│  │  - 全文索引              │  │  - 种子文件       │  │
│  │  - 快速搜索              │  │  - 备份数据       │  │
│  │  - 聚合统计              │  │                   │  │
│  └──────────┬───────────────┘  └──────────────────┘  │
└─────────────┼─────────────────────────────────────────┘
              │
┌─────────────▼─────────────────────────────────────────┐
│                  爬虫层 (torsniff)                      │
│  ┌──────────────────────────────────────────────┐     │
│  │  DHT 爬虫引擎                                 │     │
│  │  - 监听 announce_peer 消息                   │     │
│  │  - 下载种子元数据                             │     │
│  │  - 解析种子信息                               │     │
│  │  - 实时索引更新                               │     │
│  └─────────────────┬────────────────────────────┘     │
└────────────────────┼──────────────────────────────────┘
                     │ UDP 6881
┌────────────────────▼──────────────────────────────────┐
│               BitTorrent DHT 网络                       │
│  ┌────────┐  ┌────────┐  ┌────────┐  ┌────────┐      │
│  │ DHT节点│  │ DHT节点│  │ DHT节点│  │ DHT节点│ ...  │
│  └────────┘  └────────┘  └────────┘  └────────┘      │
│       (全球数百万 BitTorrent 客户端节点)               │
└───────────────────────────────────────────────────────┘
```

### 技术栈

**后端技术：**
- **核心引擎**: Go (torsniff)
- **存储系统**: Elasticsearch 7.x+
- **Web 框架**: Go HTTP Server
- **数据格式**: JSON, Bencode

**前端技术：**
- **界面框架**: HTML5 + CSS3 + JavaScript
- **响应式设计**: 支持桌面和移动设备

**网络协议：**
- **DHT 协议**: BEP 5 (Kademlia DHT)
- **元数据协议**: BEP 9 (Metadata Extension)
- **扩展协议**: BEP 10 (Extension Protocol)
- **IPv6 支持**: BEP 32 (IPv6 DHT Extension)

---

## 🚀 部署运行

### 方式一：使用预编译二进制文件（推荐）

#### 1. 下载 torsniff

从 [releases 页面](https://github.com/fanpei91/torsniff/releases) 下载适合您系统的预编译二进制文件。

#### 2. 基础运行（使用 Elasticsearch）

**前提条件**：安装并运行 Elasticsearch

```bash
# Linux/macOS
./torsniff

# Windows PowerShell
.\torsniff.exe
```

#### 3. 使用自定义 Elasticsearch 配置

```bash
# 指定 Elasticsearch 地址和索引名
./torsniff --elastic-addr http://your-elasticsearch-host:9200 --elastic-index my_torrents
```

#### 4. 启用 Web 界面

```bash
# 启动 Web 界面，可在浏览器访问 http://localhost:8080
./torsniff --web-addr :8080
```

#### 5. 同时启用 Elasticsearch 和 Web 界面（推荐）

```bash
./torsniff \
  --elastic-addr http://localhost:9200 \
  --elastic-index torrents \
  --web-addr :8080
```

#### 6. 使用磁盘存储（无需 Elasticsearch）

```bash
# 将种子保存到本地目录
./torsniff --save-to-disk --dir ~/my-torrents
```

### 方式二：从源码编译

#### 前提条件

- Go 1.18 或更高版本
- Git

#### 编译步骤

```bash
# 1. 克隆仓库
git clone https://github.com/fanpei91/torsniff.git
cd torsniff

# 2. 编译
go build

# 3. 运行
./torsniff
```

### torsniff 命令行参数详解

```bash
torsniff - 从 BitTorrent 网络中嗅探种子的工具

使用方法:
  torsniff [flags]

参数说明:
  -a, --addr string         监听地址（默认监听所有 IPv4 和 IPv6 地址）
  -p, --port uint16         监听端口（默认 6881）
  -d, --dir string          种子存储目录（默认 "$HOME/torrents"）
  -e, --peers int           下载种子的最大连接数（默认 400）
  -f, --friends int         每秒最大建立连接数（默认 500）
  -t, --timeout duration    下载种子的超时时间（默认 10s）
  -v, --verbose             显示详细信息（默认 true）
  -h, --help                显示帮助信息

Elasticsearch 相关:
  --elastic-addr string     Elasticsearch 地址（默认 "http://localhost:9200"）
  --elastic-index string    Elasticsearch 索引名称（默认 "torrents"）

存储相关:
  --save-to-disk           保存种子到磁盘（默认: false）

Web 界面:
  --web-addr string        Web 服务器地址（例如: :8080）

IPv6 支持:
  --seed string[]          额外的 DHT 引导节点（格式: host:port 或 [ipv6]:port）
  --seed6 string[]         额外的 IPv6 DHT 引导节点
```

### 部署示例配置

#### 示例 1: 基础生产环境

```bash
# 使用 Elasticsearch 存储，启用 Web 界面
./torsniff \
  --port 6881 \
  --elastic-addr http://localhost:9200 \
  --elastic-index torrents \
  --web-addr :8080 \
  --peers 400 \
  --friends 500 \
  --timeout 10s \
  --verbose
```

#### 示例 2: 仅 IPv4 监听

```bash
# Windows PowerShell
.\torsniff.exe -a 0.0.0.0 -p 6881 --web-addr :8080
```

#### 示例 3: 仅 IPv6 监听

```bash
# Windows PowerShell
.\torsniff.exe -a :: -p 6881 --web-addr :8080
```

#### 示例 4: IPv4/IPv6 双栈（推荐）

```bash
# 默认配置，双栈监听
./torsniff -p 6881 --web-addr :8080

# 添加额外的 IPv6 引导节点
./torsniff -p 6881 \
  --seed6 "[2001:4860:4860::8888]:6881" \
  --seed6 "router.bittorrent.com:6881" \
  --web-addr :8080
```

#### 示例 5: 本地开发环境

```bash
# 使用磁盘存储，无需 Elasticsearch
./torsniff \
  --save-to-disk \
  --dir ./torrents \
  --web-addr :8080 \
  --verbose
```

### 防火墙配置

为确保 torsniff 正常工作，需要开放以下端口：

**Linux (UFW):**
```bash
# 允许 UDP 6881（IPv4 和 IPv6）
sudo ufw allow 6881/udp
```

**Linux (iptables):**
```bash
# IPv4
sudo iptables -A INPUT -p udp --dport 6881 -j ACCEPT

# IPv6
sudo ip6tables -A INPUT -p udp --dport 6881 -j ACCEPT
```

**Windows 防火墙:**
```powershell
# 允许 UDP 6881 入站
New-NetFirewallRule -DisplayName "torsniff DHT" -Direction Inbound -Protocol UDP -LocalPort 6881 -Action Allow
```

### 性能优化建议

1. **使用公网 IP**
   - 有公网 IP 可以显著提高嗅探效率
   - 建议在 VPS 或云服务器上运行

2. **调整连接参数**
   ```bash
   # 高性能配置
   ./torsniff --peers 800 --friends 1000
   ```

3. **Elasticsearch 优化**
   - 增加 JVM 堆内存
   - 配置合适的分片和副本数
   - 定期清理旧数据

4. **网络优化**
   - 确保 UDP 6881 端口畅通
   - 配置 QoS 优先级
   - 使用低延迟网络

### 系统要求

**最低配置：**
- CPU: 1 核心
- 内存: 512 MB
- 磁盘: 10 GB（用于种子存储）
- 网络: 10 Mbps

**推荐配置：**
- CPU: 2+ 核心
- 内存: 2 GB+
- 磁盘: 100 GB+ SSD
- 网络: 100 Mbps+ 公网 IP

### 监控和维护

**查看运行状态：**
```bash
# 查看日志输出（verbose 模式）
./torsniff --verbose

# 访问 Web 界面查看统计信息
http://localhost:8080
```

**常见问题排查：**
1. 如果长时间没有嗅探到种子，检查：
   - 防火墙是否正确配置
   - 网络是否可以访问公网
   - DHT 端口是否被占用

2. 如果 Elasticsearch 连接失败：
   - 确认 Elasticsearch 正在运行
   - 检查连接地址是否正确
   - 查看 Elasticsearch 日志

---

## 🔧 技术原理

### DHT 网络技术

磁力魔基于 **分布式哈希表（DHT）** 技术构建：

**DHT 是什么？**
- 分布式哈希表是一种分布式存储方法
- 无需中心服务器，所有节点平等参与
- 通过哈希算法快速定位资源
- 具有高可用性和容错能力

**工作原理：**
```
1. 爬虫节点加入 DHT 网络
2. 监听网络中的 announce_peer 消息
3. 获取 info_hash（种子唯一标识）
4. 通过扩展协议获取种子元数据
5. 解析并存储种子信息到数据库
6. 建立搜索索引供用户查询
```

**相关 BEP 协议：**
- [BEP 5: DHT Protocol](https://www.bittorrent.org/beps/bep_0005.html) - DHT 协议定义
- [BEP 9: Extension for Metadata](https://www.bittorrent.org/beps/bep_0009.html) - 元数据扩展
- [BEP 10: Extension Protocol](https://www.bittorrent.org/beps/bep_0010.html) - 扩展协议

### BitTorrent 协议

**核心协议文档：**
- [BEP 3: BitTorrent Protocol](https://www.bittorrent.org/beps/bep_0003.html) - 基础协议
- [BitTorrent 官网](https://www.bittorrent.org/)
- [BitTorrent Enhancement Proposals](https://www.bittorrent.org/beps/bep_0000.html) - 协议增强提案

**技术特点：**
- 🔄 **分片传输** - 文件分割成多个块并行下载
- 🌐 **P2P 网络** - 点对点传输，无需中心服务器
- ⚡ **断点续传** - 支持暂停和继续下载
- 📊 **上传共享** - 下载的同时为其他用户提供上传

### torsniff 工作原理详解

**完整工作流程：**

```
步骤 1: 初始化
┌──────────────────────────────────────┐
│ torsniff 启动                         │
│ - 绑定 UDP 端口 (默认 6881)          │
│ - 生成随机节点 ID                     │
│ - 连接引导节点                        │
└─────────────┬────────────────────────┘
              │
步骤 2: 加入 DHT 网络
┌─────────────▼────────────────────────┐
│ DHT 网络握手                          │
│ - 向引导节点发送 find_node 请求      │
│ - 接收并记录活跃节点                  │
│ - 构建路由表                          │
└─────────────┬────────────────────────┘
              │
步骤 3: 持续监听
┌─────────────▼────────────────────────┐
│ 监听 announce_peer 消息               │
│ - 被动接收来自其他节点的种子公告      │
│ - 提取 info_hash (种子唯一标识)      │
│ - 记录发布者 IP 和端口                │
└─────────────┬────────────────────────┘
              │
步骤 4: 获取元数据
┌─────────────▼────────────────────────┐
│ 连接 Peer 下载种子元数据              │
│ - 建立 TCP 连接到 Peer               │
│ - 握手并声明支持扩展协议 (BEP 10)    │
│ - 请求元数据 (BEP 9)                 │
│ - 接收并解析 Bencode 编码的元数据    │
└─────────────┬────────────────────────┘
              │
步骤 5: 解析和存储
┌─────────────▼────────────────────────┐
│ 处理种子信息                          │
│ - 解析文件列表、大小、名称           │
│ - 提取关键元数据                      │
│ - 保存到 Elasticsearch 或磁盘        │
│ - 更新搜索索引                        │
└─────────────┬────────────────────────┘
              │
              └──► 循环回步骤 3，持续监听新种子
```

### IPv6 支持详解 (BEP-32)

torsniff 完整实现了 IPv6 DHT 支持（BEP-32），可以同时嗅探 IPv4 和 IPv6 网络中的种子。

**IPv6 实现要点：**

1. **双栈监听**
   - 默认同时绑定 IPv4 (`0.0.0.0`) 和 IPv6 (`::`)
   - 两个独立的 socket，互不干扰
   - 自动适配网络环境

2. **路由发现增强**
   ```
   find_node 请求携带 "want" 字段：
   {
     "t": "aa",
     "y": "q",
     "q": "find_node",
     "a": {
       "id": "<20字节节点ID>",
       "target": "<20字节目标ID>",
       "want": ["n4", "n6"]  ← 同时请求 IPv4 和 IPv6 节点
     }
   }
   
   响应包含：
   {
     "t": "aa",
     "y": "r",
     "r": {
       "id": "<20字节节点ID>",
       "nodes": "<IPv4节点列表>",    ← 紧凑格式
       "nodes6": "<IPv6节点列表>"    ← 紧凑格式
     }
   }
   ```

3. **节点格式解析**
   - **IPv4 节点格式**: 26 字节 (20字节ID + 4字节IP + 2字节端口)
   - **IPv6 节点格式**: 38 字节 (20字节ID + 16字节IP + 2字节端口)

4. **兼容性保证**
   - IPv6 不可用时自动降级到 IPv4
   - 对纯 IPv4 环境无影响
   - 向后兼容旧版 DHT 实现

**IPv6 配置示例：**

```bash
# 双栈模式（推荐，默认）
./torsniff -p 6881

# IPv4 + 指定 IPv6 引导节点
./torsniff -p 6881 \
  --seed6 "[2001:db8::1]:6881" \
  --seed6 "dht.transmissionbt.com:6881"

# 纯 IPv6 模式
./torsniff -a :: -p 6881
```

**IPv6 网络要求：**
- 全局可达的 IPv6 地址（GUA）
- 防火墙允许 UDP 6881 (IPv6)
- 路由器支持 IPv6 端口转发（如需）

### DHT 协议核心机制

**Kademlia 算法：**
- 使用 XOR 距离度量节点间距离
- 将节点 ID 组织成二叉树结构
- O(log N) 查找复杂度

**路由表维护：**
```
K-桶 (K-bucket) 结构：
├── 桶 0: 距离 2^0  ~ 2^1  的节点（最近）
├── 桶 1: 距离 2^1  ~ 2^2  的节点
├── ...
└── 桶 159: 距离 2^159 ~ 2^160 的节点（最远）

每个桶最多存储 K 个节点（通常 K=8）
优先保留活跃时间长的节点（LRU）
```

**消息类型：**
1. **ping** - 检测节点活跃性
2. **find_node** - 查找指定 ID 的节点
3. **get_peers** - 查找拥有特定种子的节点
4. **announce_peer** - 声明自己拥有某个种子 ⭐ 核心

**相关 BEP 协议：**
- [BEP 5: DHT Protocol](https://www.bittorrent.org/beps/bep_0005.html) - DHT 协议定义
- [BEP 9: Extension for Metadata](https://www.bittorrent.org/beps/bep_0009.html) - 元数据扩展
- [BEP 10: Extension Protocol](https://www.bittorrent.org/beps/bep_0010.html) - 扩展协议
- [BEP 32: IPv6 extension for DHT](https://www.bittorrent.org/beps/bep_0032.html) - IPv6 扩展

---

## 💻 推荐下载工具

### 桌面客户端

#### 1. qBittorrent ⭐ 推荐
- **官网**: [https://www.qbittorrent.org/](https://www.qbittorrent.org/)
- **平台**: Windows, macOS, Linux
- **特点**:
  - 完全开源免费
  - 界面清爽，无广告
  - 功能强大，支持所有主流协议
  - 内置搜索引擎
  - 支持顺序下载和优先级设置

**安装方法：**
```bash
# Windows: 下载安装包
https://www.qbittorrent.org/download.php

# macOS (Homebrew)
brew install qbittorrent

# Ubuntu/Debian
sudo apt install qbittorrent

# Arch Linux
sudo pacman -S qbittorrent
```

#### 2. Transmission
- **官网**: [https://www.transmissionbt.com/](https://www.transmissionbt.com/)
- **平台**: Windows, macOS, Linux
- **特点**:
  - 轻量级客户端
  - 资源占用低
  - 简洁高效
  - 支持 Web 界面远程控制

#### 3. Deluge
- **官网**: [https://deluge-torrent.org/](https://deluge-torrent.org/)
- **平台**: Windows, macOS, Linux
- **特点**:
  - 功能丰富
  - 插件系统强大
  - 支持客户端/服务器模式
  - 可自定义程度高

### 移动客户端

**Android:**
- Flud (免费，界面友好)
- tTorrent (功能强大)
- LibreTorrent (开源)

**iOS:**
- iTransmission (需越狱)
- nPlayer (支持磁力链接)

### 使用下载工具的步骤

1. **安装并启动下载工具**
   
2. **添加磁力链接**
   - 方法 1: 点击"文件" → "添加链接" → 粘贴磁力链接
   - 方法 2: 直接拖拽磁力链接到软件窗口
   - 方法 3: 使用快捷键 Ctrl+U (Windows) 或 Cmd+U (macOS)

3. **选择文件**
   - 对于多文件种子，可以选择需要下载的特定文件
   - 取消不需要的文件以节省空间

4. **设置下载位置**
   - 选择保存目录
   - 建议预留足够的磁盘空间

5. **开始下载**
   - 点击"确定"开始下载
   - 等待连接到 Peer 节点
   - 下载开始后，同时也会上传给其他用户

### 优化下载速度

**基础设置：**
```
1. 端口转发：在路由器设置端口转发
2. 连接数：增加最大连接数（如 200）
3. 上传速度：适当限制上传不影响下载
4. DHT：确保启用 DHT 和 PEX
5. 加密：启用协议加密避免运营商限速
```

**qBittorrent 优化设置：**
- 工具 → 选项 → 连接 → 启用 UPnP/NAT-PMP
- 工具 → 选项 → BitTorrent → 启用 DHT、PEX、LSD
- 工具 → 选项 → 高级 → 网络接口选择正确的网卡

---

## 🔥 热门资源分类

### 📚 小说文学

**修仙玄幻：**
- 吞噬星空
- 太荒吞天诀
- 九星霸体诀
- 修罗武神
- 诡秘之主
- 牧神记
- 绝世神尊
- 凡人修仙传
- 绝代神主
- 独步成仙

**都市言情：**
- 封总，太太想跟你离婚很久了
- 全家偷听我心声杀疯了,我负责吃奶
- 盖世神医
- 大医无疆

**历史军事：**
- 晚明
- 明末边军一小兵
- 割鹿记
- 1987我的年代

**其他类型：**
- 灰烬领主
- 更多小说资源持续更新中...

### 🎬 影视娱乐

- 最新院线电影
- 经典影片
- 热播电视剧
- 动漫番剧
- 纪录片
- 综艺节目

### 🎓 教育学习

- 编程教程
- 外语学习
- 考试资料
- 技能培训
- 在线课程
- 电子书籍

### 🎮 游戏软件

- PC 游戏
- 实用软件
- 开发工具
- 系统工具
- 多媒体软件

---

## �‍💻 开发指南

### 基于 torsniff 二次开发

如果您想基于 torsniff 进行二次开发或构建自己的种子搜索引擎：

#### 1. 克隆和编译

```bash
# 克隆 torsniff 源码
git clone https://github.com/fanpei91/torsniff.git
cd torsniff

# 查看项目结构
ls -la

# 编译
go build -o torsniff

# 交叉编译示例
# Linux
GOOS=linux GOARCH=amd64 go build -o torsniff-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o torsniff.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o torsniff-macos
```

#### 2. 自定义存储后端

torsniff 默认使用 Elasticsearch，但您可以实现自己的存储接口：

```go
// 实现自定义存储接口
type CustomStorage struct {
    // 您的存储配置
}

func (s *CustomStorage) Save(torrent *Torrent) error {
    // 实现保存逻辑
    // 例如：保存到 MySQL、MongoDB、PostgreSQL 等
    return nil
}

func (s *CustomStorage) Query(keyword string) ([]*Torrent, error) {
    // 实现搜索逻辑
    return nil, nil
}
```

#### 3. 添加数据处理管道

在种子数据保存前添加自定义处理：

```go
// 过滤器示例
func FilterTorrent(torrent *Torrent) bool {
    // 过滤小文件
    if torrent.Size < 1024*1024 { // 小于 1MB
        return false
    }
    
    // 过滤特定关键词
    blacklist := []string{"spam", "virus"}
    for _, word := range blacklist {
        if strings.Contains(torrent.Name, word) {
            return false
        }
    }
    
    return true
}

// 数据增强
func EnhanceTorrent(torrent *Torrent) {
    // 提取文件类型
    torrent.FileType = detectFileType(torrent.Name)
    
    // 分类
    torrent.Category = classifyTorrent(torrent)
    
    // 生成标签
    torrent.Tags = generateTags(torrent)
}
```

#### 4. 扩展 Web 界面

基于 torsniff 的 Web 界面进行定制：

```go
// 添加自定义路由
http.HandleFunc("/api/search", searchHandler)
http.HandleFunc("/api/stats", statsHandler)
http.HandleFunc("/api/trending", trendingHandler)

// 搜索处理器
func searchHandler(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("q")
    page := r.URL.Query().Get("page")
    
    // 从 Elasticsearch 查询
    results := searchFromES(keyword, page)
    
    // 返回 JSON
    json.NewEncoder(w).Encode(results)
}
```

#### 5. 性能监控

添加监控和统计功能：

```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    torrentsSniffed = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "torrents_sniffed_total",
            Help: "Total number of torrents sniffed",
        },
    )
    
    peersConnected = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "peers_connected",
            Help: "Current number of peers connected",
        },
    )
)

// 暴露 Prometheus 指标
http.Handle("/metrics", promhttp.Handler())
```

### 项目结构参考

```
cilimo/
├── cmd/
│   └── torsniff/          # 主程序入口
│       └── main.go
├── internal/
│   ├── dht/               # DHT 协议实现
│   │   ├── node.go
│   │   ├── routing.go
│   │   └── protocol.go
│   ├── metadata/          # 种子元数据处理
│   │   └── fetcher.go
│   ├── storage/           # 存储层
│   │   ├── elasticsearch.go
│   │   └── disk.go
│   └── web/               # Web 界面
│       ├── server.go
│       └── handlers.go
├── web/                   # 前端资源
│   ├── templates/
│   ├── static/
│   └── assets/
├── config/                # 配置文件
│   └── config.yaml
├── scripts/               # 部署脚本
│   ├── deploy.sh
│   └── setup_es.sh
├── README.MD
├── READ.md
├── READ-zh.md
├── go.mod
├── go.sum
└── LICENSE
```

### 构建完整的搜索引擎

基于 torsniff 构建完整搜索引擎的步骤：

**1. 数据采集层（torsniff）**
```bash
# 运行 torsniff 持续采集种子
./torsniff --elastic-addr http://localhost:9200 --elastic-index torrents
```

**2. 搜索服务层**
- 使用 Elasticsearch 提供全文搜索
- 实现高级搜索功能（过滤、排序、聚合）
- 添加搜索建议和自动完成

**3. Web 应用层**
- 开发前端界面（React/Vue/Svelte）
- 实现 RESTful API
- 添加用户系统（可选）

**4. 缓存层**
- 使用 Redis 缓存热门搜索
- 缓存搜索结果提高响应速度

**5. 监控和维护**
- 日志系统（ELK Stack）
- 性能监控（Prometheus + Grafana）
- 告警系统

### Docker 部署

创建 `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o torsniff

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/torsniff .

EXPOSE 6881/udp 8080/tcp

CMD ["./torsniff", "--elastic-addr", "http://elasticsearch:9200", "--web-addr", ":8080"]
```

创建 `docker-compose.yml`:

```yaml
version: '3.8'

services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.11.0
    environment:
      - discovery.type=single-node
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
      - xpack.security.enabled=false
    ports:
      - "9200:9200"
    volumes:
      - es_data:/usr/share/elasticsearch/data

  torsniff:
    build: .
    ports:
      - "6881:6881/udp"
      - "8080:8080"
    depends_on:
      - elasticsearch
    environment:
      - ELASTIC_ADDR=http://elasticsearch:9200
    restart: unless-stopped

  kibana:
    image: docker.elastic.co/kibana/kibana:8.11.0
    ports:
      - "5601:5601"
    depends_on:
      - elasticsearch
    environment:
      - ELASTICSEARCH_HOSTS=http://elasticsearch:9200

volumes:
  es_data:
```

启动：
```bash
docker-compose up -d
```

---

## 📚 技术参考资源

### 官方文档和 BEP 协议

- [BitTorrent 官方网站](https://www.bittorrent.org/)
- [BitTorrent Enhancement Proposals (BEP)](https://www.bittorrent.org/beps/bep_0000.html)
- [BEP 3: BitTorrent Protocol](https://www.bittorrent.org/beps/bep_0003.html) - 基础协议
- [BEP 5: DHT Protocol](https://www.bittorrent.org/beps/bep_0005.html) - DHT 协议
- [BEP 9: Extension for Metadata](https://www.bittorrent.org/beps/bep_0009.html) - 元数据扩展
- [BEP 10: Extension Protocol](https://www.bittorrent.org/beps/bep_0010.html) - 扩展协议
- [BEP 32: IPv6 extension for DHT](https://www.bittorrent.org/beps/bep_0032.html) - IPv6 扩展

### 学习资源

- [BitTorrent Wikipedia](https://en.wikipedia.org/wiki/BitTorrent)
- [DHT Wikipedia](https://en.wikipedia.org/wiki/Distributed_hash_table)
- [Kademlia Algorithm](https://en.wikipedia.org/wiki/Kademlia)
- [GitHub BitTorrent 项目](https://github.com/topics/bittorrent)
- [GitHub DHT 项目](https://github.com/topics/dht)

### 开源项目参考

**Go 语言：**
- ⭐ [fanpei91/torsniff](https://github.com/fanpei91/torsniff) - 本项目使用的核心工具
- [anacrolix/torrent](https://github.com/anacrolix/torrent) - 功能完整的 Go BitTorrent 库
- [shiyanhui/dht](https://github.com/shiyanhui/dht) - Go DHT 实现
- [zeebo/bencode](https://github.com/zeebo/bencode) - Bencode 编解码库

**Python 语言：**
- [libtorrent](https://www.libtorrent.org/) - C++ 库，提供 Python 绑定
- [bittorrent-dht](https://github.com/bittorrent/dht) - DHT 参考实现
- [BitTorrent DHT crawlers](https://github.com/topics/dht-crawler) - DHT 爬虫项目集合

**Node.js：**
- [webtorrent](https://github.com/webtorrent/webtorrent) - Web 端 BitTorrent 客户端
- [bittorrent-dht](https://github.com/webtorrent/bittorrent-dht) - JavaScript DHT 实现
- [torrent-discovery](https://github.com/webtorrent/torrent-discovery) - 种子发现模块

**其他资源：**
- [Elasticsearch 官方文档](https://www.elastic.co/guide/index.html)
- [Go 语言官方文档](https://go.dev/doc/)
- [Docker 官方文档](https://docs.docker.com/)

---

## 🔒 隐私和安全

### 隐私保护措施

磁力魔采取以下措施保护用户隐私：

1. **不记录搜索历史**
   - 不存储用户搜索记录
   - 不关联用户 IP 和搜索内容
   - 不使用追踪 Cookie

2. **HTTPS 加密**
   - 网站使用 HTTPS 加密传输
   - 保护搜索请求不被第三方窃听

3. **最小化数据收集**
   - 仅收集必要的系统运行数据
   - 不要求用户注册或登录
   - 不收集个人身份信息

4. **透明的数据处理**
   - 明确说明数据来源和用途
   - 遵守相关隐私法规

### 使用建议

为保护您的隐私和安全，建议：

1. **使用 VPN**
   - 在下载敏感内容时使用 VPN
   - 隐藏真实 IP 地址

2. **启用下载客户端加密**
   - 在 BitTorrent 客户端中启用协议加密
   - 减少 ISP 流量监控

3. **谨慎下载**
   - 不要下载明显违法的内容
   - 注意版权保护

4. **安全防护**
   - 使用杀毒软件扫描下载文件
   - 不要运行不明来源的可执行文件

---

## 📊 性能数据

### torsniff 性能指标

基于实际部署经验的性能数据：

**标准配置（2核 2GB）：**
- 每分钟嗅探种子：1000-3000 个
- 平均响应时间：< 100ms
- 内存占用：~300MB
- CPU 占用：20-40%

**高性能配置（4核 8GB）：**
- 每分钟嗅探种子：5000-10000 个
- 平均响应时间：< 50ms
- 内存占用：~800MB
- CPU 占用：40-60%

**影响因素：**
- 网络连接质量
- 公网 IP 可达性
- Elasticsearch 性能
- DHT 网络活跃度

### 扩展性

**垂直扩展：**
- 增加 CPU 核心数提升并发处理能力
- 增加内存提高缓存效率
- 使用 SSD 加速磁盘 I/O

**水平扩展：**
- 部署多个 torsniff 实例
- 使用负载均衡分发请求
- Elasticsearch 集群化部署

---

## 🐛 故障排查

### 常见问题及解决方案

#### 问题 1: torsniff 启动后长时间没有嗅探到种子

**可能原因：**
- 防火墙阻止 UDP 6881 端口
- 没有公网 IP 或 NAT 配置不正确
- DHT 引导节点连接失败

**解决方案：**
```bash
# 1. 检查端口是否开放
netstat -an | grep 6881

# 2. 测试防火墙
sudo ufw status
sudo ufw allow 6881/udp

# 3. 查看详细日志
./torsniff --verbose

# 4. 尝试指定额外的引导节点
./torsniff --seed "router.bittorrent.com:6881" --seed "dht.transmissionbt.com:6881"
```

#### 问题 2: Elasticsearch 连接失败

**可能原因：**
- Elasticsearch 未启动
- 连接地址配置错误
- 网络不通

**解决方案：**
```bash
# 1. 检查 Elasticsearch 状态
curl http://localhost:9200

# 2. 查看 Elasticsearch 日志
docker logs elasticsearch

# 3. 测试连接
./torsniff --elastic-addr http://localhost:9200 --verbose

# 4. 使用磁盘存储作为替代
./torsniff --save-to-disk --dir ./torrents
```

#### 问题 3: Web 界面无法访问

**可能原因：**
- 端口被占用
- 防火墙阻止
- Web 服务未启动

**解决方案：**
```bash
# 1. 检查端口占用
netstat -an | grep 8080

# 2. 更换端口
./torsniff --web-addr :8090

# 3. 检查防火墙
sudo ufw allow 8080/tcp

# 4. 查看日志确认 Web 服务状态
./torsniff --web-addr :8080 --verbose
```

#### 问题 4: IPv6 不工作

**可能原因：**
- 系统没有 IPv6 地址
- IPv6 防火墙规则
- ISP 不支持 IPv6

**解决方案：**
```bash
# 1. 检查 IPv6 配置
ip -6 addr show

# 2. 测试 IPv6 连接性
ping6 google.com

# 3. 仅使用 IPv4
./torsniff -a 0.0.0.0 -p 6881

# 4. 检查 IPv6 防火墙
sudo ip6tables -L
```

#### 问题 5: 下载速度慢或无法下载元数据

**可能原因：**
- Peer 连接数限制
- 网络带宽不足
- 超时时间过短

**解决方案：**
```bash
# 增加 peer 连接数和超时时间
./torsniff --peers 800 --timeout 30s --friends 1000
```

### 日志分析

启用详细日志查看运行状态：

```bash
# 启用详细日志
./torsniff --verbose 2>&1 | tee torsniff.log

# 过滤特定信息
./torsniff --verbose 2>&1 | grep "announce"
./torsniff --verbose 2>&1 | grep "metadata"
```

---

## 🤝 贡献和支持

### 如何贡献

欢迎为磁力魔项目做出贡献！

**贡献方式：**

1. **报告 Bug**
   - 提交详细的问题描述
   - 包含复现步骤
   - 提供日志和错误信息

2. **建议新功能**
   - 描述功能需求和使用场景
   - 说明预期效果

3. **提交代码**
   - Fork 项目仓库
   - 创建功能分支
   - 提交 Pull Request

4. **完善文档**
   - 修正文档错误
   - 补充使用说明
   - 翻译多语言版本

### 开发规范

**代码风格：**
- 遵循 Go 语言官方代码风格
- 使用 `gofmt` 格式化代码
- 添加必要的注释

**提交规范：**
```
<type>: <subject>

<body>

<footer>

类型 (type):
- feat: 新功能
- fix: 修复 bug
- docs: 文档更新
- style: 代码格式调整
- refactor: 重构
- test: 测试相关
- chore: 构建/工具相关
```

**示例：**
```
feat: 添加 IPv6 支持

- 实现 BEP-32 协议
- 支持双栈监听
- 添加 IPv6 引导节点配置

Closes #123
```

### 社区和交流

- **GitHub Issues**: 报告问题和讨论
- **Pull Requests**: 提交代码贡献
- **Discussions**: 技术交流和问答

---

## 📜 开源许可证

本项目基于 **MIT License** 开源。

```
MIT License

Copyright (c) 2025 Cilimo.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

### 第三方组件许可

- **torsniff**: MIT License
- **Elasticsearch**: Elastic License 2.0 / SSPL
- **Go**: BSD-style License

---

## ⚠️ 免责声明

**重要提示：请仔细阅读**

1. **内容来源**
   - 磁力魔所有索引内容均来自 BitTorrent DHT 网络
   - 本站不存储、不制作、不上传任何资源文件
   - 本站仅提供种子信息的索引和搜索服务

2. **版权声明**
   - 用户在使用本站搜索服务时，应遵守当地法律法规
   - 请勿搜索、下载、传播任何侵犯版权的内容
   - 下载行为与本站无关，由用户自行承担法律责任

3. **资源质量**
   - 本站不对搜索结果的准确性、完整性、可用性负责
   - 用户应自行判断资源内容的合法性和安全性
   - 建议使用杀毒软件扫描下载的文件

4. **使用建议**
   - 支持正版，鼓励购买正版软件、游戏、影视作品
   - 仅将搜索结果用于学习、研究等合法用途
   - 下载的内容应在 24 小时内删除

5. **隐私保护**
   - 本站不记录用户搜索历史
   - 不收集用户个人信息
   - 使用 HTTPS 加密传输

**使用本站服务即表示您已阅读并同意本免责声明的所有条款。**

---

## 🛡️ 安全提示

### 下载安全建议

1. **使用正规下载工具**
   - 从官方网站下载 BitTorrent 客户端
   - 避免使用来路不明的下载器

2. **病毒防护**
   - 下载后使用杀毒软件扫描
   - 谨慎运行可执行文件（.exe、.bat、.sh 等）
   - 不要打开可疑的文档宏

3. **隐私保护**
   - 考虑使用 VPN 保护隐私
   - 在下载客户端中启用协议加密
   - 不要下载明显违法的内容

4. **资源甄别**
   - 查看文件大小是否合理
   - 注意文件格式是否正确
   - 参考其他用户评论

### 常见风险

⚠️ **假种子**：文件名正确但内容不符  
⚠️ **病毒木马**：伪装成正常文件的恶意软件  
⚠️ **钓鱼链接**：要求输入个人信息的虚假页面  
⚠️ **版权陷阱**：故意传播侵权内容进行钓鱼执法  

---

## ❓ 常见问题

### 搜索相关

**Q: 为什么搜索不到我想要的资源？**

A: 可能的原因：
- 关键词不够准确，尝试更换关键词
- 资源比较冷门，DHT 网络中节点较少
- 资源刚刚发布，还未被索引
- 尝试使用英文关键词或更通用的词汇

**Q: 搜索结果太多怎么办？**

A: 使用高级搜索功能：
- 添加更多限定词缩小范围
- 使用引号进行精确匹配
- 筛选文件类型和大小
- 按时间排序查看最新资源

**Q: 如何搜索特定格式的文件？**

A: 在搜索词后添加格式：
```
电影名 1080p
软件名 x64
教程 PDF
音乐 FLAC
```

### 下载相关

**Q: 磁力链接无法下载怎么办？**

A: 检查以下几点：
1. 确认下载工具支持磁力链接
2. 检查是否启用了 DHT 功能
3. 等待一段时间让客户端连接到节点
4. 尝试更换下载工具
5. 检查防火墙和路由器设置

**Q: 下载速度很慢怎么办？**

A: 优化建议：
- 确保启用了 DHT、PEX、LSD
- 配置端口转发
- 增加连接数限制
- 选择 Peer 多的资源
- 避开网络高峰期

**Q: 为什么连接不到 Peer？**

A: 可能原因：
- 资源过于冷门，在线节点少
- 防火墙或路由器阻止了连接
- ISP 限制了 P2P 流量
- 需要启用 DHT 和 PEX

### 技术问题

**Q: 什么是做种？**

A: 做种是指在下载完成后继续运行 BitTorrent 客户端，为其他用户提供上传。这是 P2P 网络的核心原则，建议下载完成后保持做种一段时间。

**Q: info_hash 是什么？**

A: info_hash 是种子的唯一标识符（40位16进制字符串），通过 SHA-1 哈希算法生成。磁力链接中的 `btih:` 后面就是 info_hash。

**Q: DHT、PEX、LSD 是什么意思？**

A: 
- **DHT**: 分布式哈希表，无需 Tracker 即可找到 Peer
- **PEX**: Peer Exchange，Peer 之间交换节点信息
- **LSD**: Local Service Discovery，局域网服务发现

---

## 🌐 浏览器支持

磁力魔支持所有现代浏览器：

- ✅ **Chrome / Edge** (推荐) - 版本 90+
- ✅ **Firefox** - 版本 88+
- ✅ **Safari** - 版本 14+
- ✅ **Opera** - 版本 76+

**建议使用最新版本浏览器以获得最佳体验**

---

## 📱 移动端访问

磁力魔完全支持移动设备访问：

- 📱 响应式设计，自适应屏幕
- 🔍 移动端优化的搜索界面
- 📋 一键复制磁力链接
- ⚡ 快速加载，省流量

**移动端使用建议：**
1. 使用支持磁力链接的下载 App
2. 推荐 Android 使用 Flud 或 LibreTorrent
3. 建议在 WiFi 环境下下载大文件

---

## 💡 使用技巧

### 搜索技巧

1. **组合关键词**: `电影名 + 1080p + 蓝光`
2. **排除干扰**: `游戏名 -破解 -外挂`
3. **年份筛选**: `纪录片 2024`
4. **语言标识**: `教程 中文字幕`

### 下载技巧

1. **选择 Peer 多的资源** - 下载速度更快
2. **避开网络高峰期** - 凌晨和上午速度较快
3. **保持上传共享** - 有助于提高下载速度
4. **使用限速功能** - 避免占满全部带宽

### 资源管理

1. **分类存储** - 不同类型文件分文件夹保存
2. **定期清理** - 删除不需要的种子任务
3. **备份重要资源** - 将重要文件备份到其他位置
4. **做种分享** - 热门资源建议长期做种

---

## 🔗 相关链接

### 官方渠道

- 🌐 官方网站: [https://cilimo.com](https://cilimo.com)
- 📧 联系邮箱: contact@cilimo.com
- 💬 问题反馈: feedback@cilimo.com

### 技术交流

- [BitTorrent 官方论坛](https://forum.bittorrent.com/)
- [GitHub BitTorrent 话题](https://github.com/topics/bittorrent)
- [Reddit r/torrents](https://www.reddit.com/r/torrents/)

### 推荐资源

- [开源下载工具对比](https://www.reddit.com/r/torrents/comments/best_clients)
- [BitTorrent 协议深度解析](https://en.wikipedia.org/wiki/BitTorrent)
- [P2P 技术学习指南](https://github.com/topics/p2p)

---

## 🤝 支持项目

如果您觉得磁力魔对您有帮助，欢迎：

- ⭐ 给本项目 Star
- 🔗 分享给更多需要的朋友
- 💬 提供反馈和建议
- 🐛 报告 Bug 和问题
- 💡 贡献代码和文档

### 参与贡献

欢迎开发者贡献代码：

```bash
# 1. Fork 本项目
# 2. Clone 到本地
git clone https://github.com/e2ckp/cilimo.git

# 3. 创建功能分支
git checkout -b feature/your-feature

# 4. 提交更改
git commit -am 'feat: Add some feature'

# 5. 推送到分支
git push origin feature/your-feature

# 6. 创建 Pull Request
```

### Star 历史

[![Star History Chart](https://api.star-history.com/svg?repos=e2ckp/cilimo&type=Date)](https://star-history.com/#e2ckp/cilimo&Date)

---

## 🙏 致谢

感谢以下开源项目和技术的支持：

### 核心技术
- **[torsniff](https://github.com/fanpei91/torsniff)** - 感谢 fanpei91 开发的高性能 DHT 嗅探工具
- **BitTorrent Protocol** - 革命性的 P2P 文件共享协议
- **DHT Network** - 分布式哈希表技术
- **Elasticsearch** - 强大的搜索和分析引擎

### 开源社区
- 所有开源 BitTorrent 客户端开发者
- 为 P2P 技术发展做出贡献的开发者们
- BitTorrent 协议标准化组织
- Elasticsearch 团队

### 用户支持
- 所有使用和支持磁力魔的用户
- 提供反馈和建议的社区成员
- 贡献代码和文档的开发者

### 相关项目
- [qBittorrent](https://www.qbittorrent.org/) - 优秀的开源 BT 客户端
- [Transmission](https://transmissionbt.com/) - 轻量级 BT 客户端
- [Deluge](https://deluge-torrent.org/) - 功能丰富的 BT 客户端

---

## 📞 联系我们

### 官方渠道

- 🌐 **官方网站**: [https://cilimo.com](https://cilimo.com)
- 📧 **联系邮箱**: contact@cilimo.com
- 💬 **问题反馈**: feedback@cilimo.com
- 🐛 **Bug 报告**: bugs@cilimo.com

### 技术支持

- **GitHub Issues**: 报告问题和功能建议
- **GitHub Discussions**: 技术交流和问答
- **Email**: support@cilimo.com

### 商务合作

如有商务合作需求，请联系：
- 📧 Email: business@cilimo.com

---

## ⚖️ 法律声明和免责

### 服务说明

1. **技术性质**
   - 磁力魔是一个技术演示项目
   - 基于开源技术构建
   - 展示 DHT 网络爬取和索引技术

2. **内容来源**
   - 所有数据来自公开的 BitTorrent DHT 网络
   - 本站不存储、不制作、不上传任何资源文件
   - 仅提供种子元数据的索引和搜索服务

3. **责任限制**
   - 本站不对搜索结果的合法性负责
   - 用户应自行判断内容的合法性和安全性
   - 下载行为完全由用户自主决定，与本站无关

4. **版权尊重**
   - 我们尊重知识产权和版权
   - 不鼓励下载或传播侵权内容
   - 支持正版，鼓励购买正版内容

5. **用户义务**
   - 遵守当地法律法规
   - 不得用于非法用途
   - 合理合法使用搜索服务

### 版权投诉

如果您认为本站索引的内容侵犯了您的合法权益：

1. 请发送邮件到：dmca@cilimo.com
2. 提供以下信息：
   - 您的身份证明
   - 版权证明材料
   - 侵权内容的详细链接
   - 联系方式

我们将在收到有效通知后的 **48 小时内**进行核实并处理。

### 使用条款

使用磁力魔服务即表示您同意：

- 已阅读并理解本免责声明
- 同意遵守使用条款
- 自行承担使用风险
- 不将本服务用于非法用途

---

## 🌟 项目愿景

磁力魔致力于：

- 🔍 **技术探索** - 探索和展示 P2P 网络技术
- 📚 **知识分享** - 分享 DHT 和 BitTorrent 技术知识
- 🌐 **开源精神** - 拥抱开源，回馈社区
- 🛡️ **隐私保护** - 保护用户隐私和数据安全
- ⚖️ **合规运营** - 遵守法律法规，负责任运营

---

## 📄 相关文档

- [torsniff 英文文档](READ.md)
- [torsniff 中文文档](READ-zh.md)
- [LICENSE - MIT 许可证](LICENSE)

---

## 📈 项目统计

![GitHub stars](https://img.shields.io/github/stars/e2ckp/cilimo?style=social)
![GitHub forks](https://img.shields.io/github/forks/e2ckp/cilimo?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/e2ckp/cilimo?style=social)

![GitHub issues](https://img.shields.io/github/issues/e2ckp/cilimo)
![GitHub pull requests](https://img.shields.io/github/issues-pr/e2ckp/cilimo)
![GitHub last commit](https://img.shields.io/github/last-commit/e2ckp/cilimo)

---

## 🎯 Roadmap

### 已完成 ✅

- [x] 基于 torsniff 的 DHT 嗅探引擎
- [x] Elasticsearch 集成
- [x] Web 搜索界面
- [x] IPv4/IPv6 双栈支持
- [x] Docker 部署方案
- [x] 详细技术文档

### 计划中 📋

- [ ] 高级搜索功能（过滤器、排序）
- [ ] 搜索建议和自动完成
- [ ] 多语言界面支持
- [ ] RESTful API 接口
- [ ] 移动端适配优化
- [ ] 数据可视化统计
- [ ] 缓存层优化
- [ ] 分布式部署方案

### 长期目标 🚀

- [ ] AI 驱动的内容分类
- [ ] 个性化推荐系统
- [ ] 区块链技术集成
- [ ] 去中心化搜索网络

---

**© 2025 [磁力魔 (Cilimo.com)](https://www.cilimo.com/)**

**基于 [torsniff](https://github.com/fanpei91/torsniff) 构建 | 采用 MIT 许可证**

**让资源发现更简单，让分享更自由 | 技术改变世界**

---

<div align="center">

### ⭐ 如果这个项目对您有帮助，请给我们一个 Star！

[![GitHub stars](https://img.shields.io/github/stars/e2ckp/cilimo?style=for-the-badge&logo=github)](https://github.com/e2ckp/cilimo/stargazers)

**[官方网站](https://cilimo.com)** • **[GitHub](https://github.com/e2ckp/cilimo)** • **[问题反馈](https://github.com/e2ckp/cilimo/issues)** • **[贡献指南](CONTRIBUTING.md)**

</div>

---

[⬆️ 返回顶部](#磁力魔-cilimocom---bittorrent-dht-种子搜索引擎)
