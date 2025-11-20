# 🌐 WhoIs IP CLI

A fast and minimal CLI tool written in Go to fetch **Whois information for domains** and **IP geolocation details** directly from the terminal. Supports JSON output and versioning. Powered by public APIs.

![Go Version](https://img.shields.io/badge/Go-1.24+-blue?logo=go)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🌍 Lookup domain Whois info via `api.whois.vu`
- 📍 Get IP geolocation and provider info via `ip-api.com`
- 🔎 Detects input type (IP or domain) automatically
- 📦 JSON output option for automation or scripting
- 🔒 Proxy support (HTTP, HTTPS, SOCKS5) for secure and private lookups
- 🛠️ Custom dialer support for advanced network configurations
- 🧾 Simple flags: `-json`, `-v`, `-h`, `-proxy`
- ⚙️ Written in pure Go with minimal dependencies

---

## 📦 Installation

### Option 1: Go Install

```bash
go install github.com/hadzicni/whois-ip-cli/cmd/whoiscli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build (Windows, Linux, macOS)

#### 🪟 Windows (PowerShell oder CMD)

```powershell
git clone https://github.com/hadzicni/whois-ip-cli.git
cd whois-ip-cli/cmd/whoiscli
go build -o whoiscli.exe
```

#### 🐧 Linux / 🍏 macOS

```bash
git clone https://github.com/hadzicni/whois-ip-cli.git
cd whois-ip-cli/cmd/whoiscli
go build -o whoiscli
```

---

## 🚀 Usage

```bash
whoiscli [flags] <domain|ip>
```

### Available Flags

| Flag      | Description                                              |
| --------- | -------------------------------------------------------- |
| `-json`   | Output as JSON                                           |
| `-v`      | Show version info                                        |
| `-h`      | Show help message                                        |
| `-proxy`  | Proxy URL (supports http://, https://, socks5://)       |

---

## 🔧 Examples

Check a domain:

```bash
whoiscli example.com
```

Check an IP address:

```bash
whoiscli 8.8.8.8
```

Output as JSON:

```bash
whoiscli -json example.com
```

Show version:

```bash
whoiscli -v
```

---

## 🔐 Proxy Support

The CLI supports HTTP, HTTPS, and SOCKS5 proxies for secure and private lookups.

### Using Proxy via Command-Line Flag

Use the `-proxy` flag to specify a proxy URL:

```bash
# SOCKS5 proxy
whoiscli -proxy socks5://127.0.0.1:1080 example.com

# SOCKS5 proxy with authentication
whoiscli -proxy socks5://username:password@127.0.0.1:1080 example.com

# HTTP proxy
whoiscli -proxy http://proxy.example.com:8080 8.8.8.8

# HTTPS proxy
whoiscli -proxy https://proxy.example.com:443 example.com
```

### Using Proxy via Environment Variables

The CLI also respects standard proxy environment variables:

```bash
# Set proxy via environment variable
export HTTPS_PROXY=http://proxy.example.com:8080
whoiscli example.com

# Or use HTTP_PROXY
export HTTP_PROXY=socks5://127.0.0.1:1080
whoiscli 8.8.8.8
```

Supported environment variables (in priority order):
- `HTTPS_PROXY` / `https_proxy`
- `HTTP_PROXY` / `http_proxy`

### Using Custom Dialers (Library Usage)

For advanced use cases, you can use the library programmatically with custom dialers:

```go
import (
    "whois-ip-cli/internal/whois"
    "golang.org/x/net/proxy"
)

// Option 1: Using a custom dialer
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
config := whois.DefaultClientConfig()
config.CustomDialer = dialer
whois.SetDefaultClientConfig(config)

// Option 2: Using a dialer creation function
config.DialerFunc = func() (proxy.Dialer, error) {
    // Your custom dialer logic here
    return dialer, nil
}

// Option 3: Using a proxy URL
config.ProxyURL = "socks5://127.0.0.1:1080"
whois.SetDefaultClientConfig(config)
```

See the [examples/custom_dialer.go](examples/custom_dialer.go) file for more detailed usage examples.

---

## 👨‍💻 Author

Made by **Nikola Hadzic**  
GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
