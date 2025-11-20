package whois

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

// ClientConfig holds configuration for HTTP client including proxy settings
type ClientConfig struct {
	// ProxyURL is the proxy server URL (supports http, https, socks5)
	ProxyURL string
	
	// CustomDialer allows passing a custom dialer
	CustomDialer proxy.Dialer
	
	// DialerFunc allows passing a custom dialer creation function
	DialerFunc func() (proxy.Dialer, error)
	
	// Timeout for HTTP requests
	Timeout time.Duration
}

// DefaultClientConfig returns a default configuration
func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		Timeout: 30 * time.Second,
	}
}

// NewHTTPClient creates an HTTP client with the given configuration
func (c *ClientConfig) NewHTTPClient() (*http.Client, error) {
	var dialer proxy.Dialer
	var err error

	// Priority: CustomDialer > DialerFunc > ProxyURL > Default
	if c.CustomDialer != nil {
		dialer = c.CustomDialer
	} else if c.DialerFunc != nil {
		dialer, err = c.DialerFunc()
		if err != nil {
			return nil, err
		}
	} else if c.ProxyURL != "" {
		proxyURL, err := url.Parse(c.ProxyURL)
		if err != nil {
			return nil, err
		}

		switch proxyURL.Scheme {
		case "socks5":
			// Create SOCKS5 dialer
			dialer, err = proxy.SOCKS5("tcp", proxyURL.Host, nil, proxy.Direct)
			if err != nil {
				return nil, err
			}
		case "http", "https":
			// For HTTP/HTTPS proxies, use standard net dialer and configure transport
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
			}
			return &http.Client{
				Timeout:   c.Timeout,
				Transport: transport,
			}, nil
		default:
			// Use direct connection for unknown schemes
			dialer = proxy.Direct
		}
	} else {
		// Use default dialer
		dialer = proxy.Direct
	}

	// Create HTTP transport with the dialer
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	return &http.Client{
		Timeout:   c.Timeout,
		Transport: transport,
	}, nil
}

// Global default client config
var defaultConfig = DefaultClientConfig()

// SetDefaultClientConfig sets the global default client configuration
func SetDefaultClientConfig(config *ClientConfig) {
	defaultConfig = config
}

// GetDefaultHTTPClient returns an HTTP client with the default configuration
func GetDefaultHTTPClient() (*http.Client, error) {
	return defaultConfig.NewHTTPClient()
}
