# Proxy Support Implementation Summary

## Overview
This implementation adds comprehensive proxy support to the whois-ip-cli tool, including HTTP, HTTPS, and SOCKS5 proxies with authentication support. It also provides the ability to pass custom dialers or dialer creation functions for advanced use cases.

## Features Implemented

### 1. Proxy URL Support
- **HTTP Proxy**: `http://proxy.example.com:8080`
- **HTTPS Proxy**: `https://proxy.example.com:443`
- **SOCKS5 Proxy**: `socks5://127.0.0.1:1080`
- **SOCKS5 with Auth**: `socks5://username:password@127.0.0.1:1080`

### 2. Configuration Methods

#### Command-Line Flag
```bash
whoiscli -proxy socks5://127.0.0.1:1080 example.com
```

#### Environment Variables
The tool respects standard proxy environment variables:
- `HTTPS_PROXY` / `https_proxy`
- `HTTP_PROXY` / `http_proxy`

Priority: Command-line flag > Environment variables

### 3. Custom Dialer Support (Library API)

#### Method 1: Custom Dialer
```go
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
config := whois.DefaultClientConfig()
config.CustomDialer = dialer
whois.SetDefaultClientConfig(config)
```

#### Method 2: Dialer Creation Function
```go
config := whois.DefaultClientConfig()
config.DialerFunc = func() (proxy.Dialer, error) {
    // Custom logic (e.g., rotating proxies, load balancing)
    return dialer, nil
}
whois.SetDefaultClientConfig(config)
```

#### Method 3: Proxy URL
```go
config := whois.DefaultClientConfig()
config.ProxyURL = "socks5://127.0.0.1:1080"
whois.SetDefaultClientConfig(config)
```

#### Method 4: SOCKS5 Authentication via ProxyAuth
```go
config := whois.DefaultClientConfig()
config.ProxyURL = "socks5://127.0.0.1:1080"
config.ProxyAuth = &proxy.Auth{
    User:     "username",
    Password: "password",
}
whois.SetDefaultClientConfig(config)
```

## Implementation Details

### New Files
- **internal/whois/client.go**: Core client configuration and HTTP client factory
- **examples/custom_dialer.go**: Example code demonstrating all usage patterns

### Modified Files
- **cmd/whoiscli/main.go**: Added `-proxy` flag and environment variable support
- **internal/whois/ip.go**: Modified to use configurable HTTP client
- **internal/whois/domain.go**: Modified to use configurable HTTP client
- **README.md**: Added comprehensive proxy documentation
- **go.mod**: Added `golang.org/x/net` dependency

### Architecture

```
┌─────────────────────┐
│   CLI / User Code   │
│  (main.go or lib)   │
└──────────┬──────────┘
           │
           │ Sets config
           ▼
┌─────────────────────┐
│   ClientConfig      │
│  - ProxyURL         │
│  - CustomDialer     │
│  - DialerFunc       │
│  - ProxyAuth        │
└──────────┬──────────┘
           │
           │ Creates
           ▼
┌─────────────────────┐
│   HTTP Client       │
│  with Transport     │
└──────────┬──────────┘
           │
           │ Uses
           ▼
┌─────────────────────┐
│  Proxy Dialer       │
│  (SOCKS5/HTTP)      │
└─────────────────────┘
```

### Priority Order
When multiple configuration methods are used, the priority is:
1. CustomDialer (highest)
2. DialerFunc
3. ProxyURL
4. Direct connection (lowest)

## Security

- **CodeQL Analysis**: Passed with 0 vulnerabilities
- **Authentication**: SOCKS5 authentication credentials can be embedded in URL or passed separately via ProxyAuth
- **No Credentials in Logs**: Sensitive proxy credentials are not logged

## Backward Compatibility

The implementation is fully backward compatible:
- Existing code continues to work without modifications
- Default behavior (direct connection) remains unchanged
- No breaking changes to existing APIs

## Testing

The implementation includes:
- Successful builds of CLI and examples
- Comprehensive example code in `examples/custom_dialer.go`
- Documentation with usage examples in README.md

## Usage Examples in README

See the updated README.md for detailed usage examples including:
- Command-line proxy usage
- Environment variable configuration
- Programmatic library usage with custom dialers
- SOCKS5 authentication examples

## Dependencies

- `golang.org/x/net/proxy`: Standard Go networking package for SOCKS5 support
- No other external dependencies added

## Future Enhancements (Optional)

Potential future improvements not implemented in this PR:
- Proxy auto-discovery (PAC files)
- Proxy connection pooling
- Proxy failover/retry logic
- HTTP proxy authentication (currently only SOCKS5 auth is supported)
