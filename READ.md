# torsniff - A sniffer that sniffs torrents from the BitTorrent network

torsniff is a tool that sniffes torrents from the BitTorrent network and saves the torrent files to Elasticsearch by default.

## Features

- High performance, can sniff thousands of torrents per minute
- Easy to use, no configuration needed
- Save torrents to Elasticsearch by default (can be customized)
- Optionally save torrents in local directory
- Built-in web interface for browsing collected torrents

## Installation

Download a compiled binary from the releases page.

## Usage

```
torsniff - a sniffer that sniffs torrents from BitTorrent network.
Usage:
  torsniff [flags]

Flags:
  -a, --addr string         listen on given address (default all, ipv4 and ipv6)
  -d, --dir string          the directory to store the torrents (default "$HOME/torrents")
  -e, --peers int           max peers to connect to download torrents (default 400)
  -f, --friends int         max fiends to make with per second (default 500)
  -h, --help                help for torsniff
  -p, --port uint16         listen on given port (default 6881)
  -t, --timeout duration    max time allowed for downloading torrents (default 10s)
  -v, --verbose             run in verbose mode (default true)
      --elastic-addr string Elasticsearch address (default "http://localhost:9200")
      --elastic-index string Elasticsearch index name (default "torrents")
      --save-to-disk        Save torrents to disk (default: false)
      --web-addr string     Web server address (e.g., :8080)
```

### Examples

#### Run with default settings (Elasticsearch)

```
torsniff
```

#### Run with custom Elasticsearch settings

```
torsniff --elastic-addr http://your-elasticsearch-host:9200 --elastic-index my_torrents
```

#### Run with web interface

```
torsniff --web-addr :8080
```

#### Run with both Elasticsearch and web interface

```
torsniff --elastic-addr http://your-elasticsearch-host:9200 --elastic-index my_torrents --web-addr :8080
```

#### Run with disk storage instead of Elasticsearch

```
torsniff --save-to-disk --dir ~/my-torrents
```

## How it works

1. torsniff joins the BitTorrent network as a DHT node
2. It listens for announcements from other nodes about available torrents
3. When an announcement is received, it connects to peers and downloads the torrent metadata
4. Torrent information is saved to Elasticsearch by default (or to disk if configured)
5. The built-in web interface allows browsing and searching collected torrents

## Requirements

- A public IP address for best performance (optional but recommended)
- Port 6881 (UDP) open for DHT communication
- Elasticsearch instance (unless using disk storage)

## Notes

- It may take a few minutes to start receiving torrents
- The more torrents it can sniff, the better the performance
- Collected torrents are stored in Elasticsearch by default (can be changed with --save-to-disk)
- The web interface provides a convenient way to browse and search collected torrents

## License

MIT