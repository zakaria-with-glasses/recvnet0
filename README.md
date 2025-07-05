# recvnet0

A fast, simple, and extensible network reconnaissance tool written in Go.

## 📖 Overview

`recvnet0` is a command-line tool for scanning networks via TCP and UDP. It supports multi-threaded scanning, custom port ranges, output formats, and configurable timeouts—ideal for both educational and practical security purposes.

---

## ⚙️ Installation

### 🧩 Requirements

- Go 1.18 or later

### 🔧 Build from Source

```bash
git clone https://github.com/zakaria-with-glasses/recvnet0.git
cd recvnet0
go build -o recvnet0
```

### 📁 Download Binary

Precompiled binaries will be available under [Releases](https://github.com/zakaria-with-glasses/recvnet0/releases).

---

## 🚀 Usage

```bash
./recvnet0 -host <target> -ports <port range> [options]
```

### 🧪 Example Commands

Scan TCP ports 1–1024 on a local IP:

```bash
./recvnet0 -host 192.168.1.1 -ports 1-1024
```

Full TCP/UDP scan on all ports with JSON output:

```bash
./recvnet0 -host example.com -ports 1-65535 -udp -format json -output scan.json
```

Fast scan using 500 workers and a shorter timeout:

```bash
./recvnet0 -host 10.10.11.1 -ports 20-1000 -workers 500 -timeout 200ms
```

---

## 📥 Command-Line Flags

| Flag       | Type     | Description                                          | Default    |
| ---------- | -------- | ---------------------------------------------------- | ---------- |
| `-host`    | string   | Target hostname or IP address                        | (required) |
| `-ports`   | string   | Port list/range (e.g. `22,80` or `1-1024`)           | `1-1024`   |
| `-udp`     | bool     | Enable UDP scan in addition to TCP                   | `false`    |
| `-workers` | int      | Number of concurrent goroutines                      | `100`      |
| `-timeout` | duration | Timeout per connection attempt (e.g., `500ms`, `1s`) | `300ms`    |
| `-retries` | int      | Number of retries for failed attempts                | `1`        |
| `-format`  | string   | Output format: `text`, `json`, `csv`                 | `text`     |
| `-output`  | string   | Output file name; prints to stdout if omitted        | (none)     |

---

## 📄 Output Examples

### Text Output (default)

```
[*] Target: 192.168.1.50
[+] TCP 22 OPEN (ssh)
[+] TCP 80 OPEN (http)
[+] UDP 53 OPEN (dns)
```

### JSON Output

```json
{
  "host": "192.168.1.50",
  "scan_time": "2.3s",
  "results": [
    {"protocol":"tcp", "port":22, "status":"open", "service":"ssh"},
    {"protocol":"tcp", "port":80, "status":"open", "service":"http"},
    {"protocol":"udp", "port":53, "status":"open", "service":"dns"}
  ]
}
```

### CSV Output

```csv
protocol,port,status,service
tcp,22,open,ssh
tcp,80,open,http
udp,53,open,dns
```

---

## 🧪 Testing

Run unit tests using:

```bash
go test ./...
```

---

## 🤝 Contributing

Pull requests are welcome. To contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes
4. Push and open a Pull Request

Please write tests for any new features.

---

## 📝 License

This project is licensed under the MIT License.

---

## 🙏 Acknowledgements

Inspired by classic tools like `nmap`, and built using Go’s powerful concurrency model.

