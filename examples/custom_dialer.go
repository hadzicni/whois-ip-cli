package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/proxy"
	"whois-ip-cli/internal/whois"
)

// Example 1: Using a SOCKS5 proxy via URL
func exampleWithProxyURL() {
	config := whois.DefaultClientConfig()
	config.ProxyURL = "socks5://127.0.0.1:1080"
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the SOCKS5 proxy
	whois.LookupIP("8.8.8.8", false)
}

// Example 1b: Using a SOCKS5 proxy with authentication via URL
func exampleWithSOCKS5Auth() {
	config := whois.DefaultClientConfig()
	// Authentication credentials can be embedded in the URL
	config.ProxyURL = "socks5://username:password@127.0.0.1:1080"
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the SOCKS5 proxy with authentication
	whois.LookupIP("8.8.8.8", false)
}

// Example 1c: Using a SOCKS5 proxy with authentication via ProxyAuth
func exampleWithSOCKS5AuthStruct() {
	config := whois.DefaultClientConfig()
	config.ProxyURL = "socks5://127.0.0.1:1080"
	config.ProxyAuth = &proxy.Auth{
		User:     "username",
		Password: "password",
	}
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the SOCKS5 proxy with authentication
	whois.LookupIP("8.8.8.8", false)
}

// Example 2: Using a custom dialer
func exampleWithCustomDialer() {
	// Create a custom SOCKS5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}

	config := whois.DefaultClientConfig()
	config.CustomDialer = dialer
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the custom dialer
	whois.LookupDomain("example.com", false)
}

// Example 3: Using a custom dialer creation function
func exampleWithDialerFunc() {
	config := whois.DefaultClientConfig()
	
	// Set a dialer creation function that returns a custom dialer
	config.DialerFunc = func() (proxy.Dialer, error) {
		// You could implement custom logic here, like rotating proxies
		// For this example, we'll create a simple direct dialer with custom timeout
		return &net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}, nil
	}
	
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the dialer from the function
	whois.LookupIP("1.1.1.1", true)
}

// Example 4: Using HTTP proxy
func exampleWithHTTPProxy() {
	config := whois.DefaultClientConfig()
	config.ProxyURL = "http://proxy.example.com:8080"
	whois.SetDefaultClientConfig(config)

	// Now all lookups will use the HTTP proxy
	whois.LookupDomain("github.com", false)
}

// Example 5: Per-request proxy configuration
func examplePerRequestProxy() {
	// Create a config for SOCKS5 proxy
	socksConfig := whois.DefaultClientConfig()
	socksConfig.ProxyURL = "socks5://127.0.0.1:1080"
	
	// Get a client with SOCKS5 proxy
	socksClient, err := socksConfig.NewHTTPClient()
	if err != nil {
		log.Fatal(err)
	}
	
	// Use the client for a specific request
	resp, err := socksClient.Get("https://api.whois.vu/?q=example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	fmt.Println("Response status:", resp.Status)
	
	// Create another config for HTTP proxy
	httpConfig := whois.DefaultClientConfig()
	httpConfig.ProxyURL = "http://proxy.example.com:8080"
	
	// Get a client with HTTP proxy
	httpClient, err := httpConfig.NewHTTPClient()
	if err != nil {
		log.Fatal(err)
	}
	
	// Use the HTTP proxy client for another request
	resp2, err := httpClient.Get("http://ip-api.com/json/8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()
	
	fmt.Println("Response status:", resp2.Status)
}

func main() {
	fmt.Println("See individual example functions for different proxy usage patterns")
	fmt.Println("Uncomment the function you want to test")
	
	// Uncomment to run examples:
	// exampleWithProxyURL()
	// exampleWithSOCKS5Auth()
	// exampleWithSOCKS5AuthStruct()
	// exampleWithCustomDialer()
	// exampleWithDialerFunc()
	// exampleWithHTTPProxy()
	// examplePerRequestProxy()
}
