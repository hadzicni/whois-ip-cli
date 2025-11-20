package whois

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type IPResult struct {
	Query      string `json:"query"`
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	ISP        string `json:"isp"`
	Org        string `json:"org"`
	Timezone   string `json:"timezone"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Hostname   string `json:"hostname,omitempty"`
}

func IsIP(input string) bool {
	return net.ParseIP(input) != nil
}

func LookupIP(ip string, asJSON bool) {
	client, err := GetDefaultHTTPClient()
	if err != nil {
		fmt.Println("Fehler beim Erstellen des HTTP-Clients:", err)
		os.Exit(1)
	}

	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		fmt.Println("Fehler:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data IPResult
	json.NewDecoder(resp.Body).Decode(&data)

	names, err := net.LookupAddr(ip)
	if err == nil && len(names) > 0 {
		data.Hostname = names[0]
	}

	if data.Status != "success" {
		fmt.Println("Fehler:", data.Message)
		os.Exit(1)
	}

	if asJSON {
		json.NewEncoder(os.Stdout).Encode(data)
	} else {
		fmt.Printf("IP:        %s\n", data.Query)
		if data.Hostname != "" {
			fmt.Printf("Hostname:  %s\n", data.Hostname)
		}
		fmt.Printf("Land:      %s\nRegion:   %s\nStadt:     %s\nProvider:  %s\nOrganisation: %s\nZeitzone:  %s\n",
			data.Country, data.RegionName, data.City, data.ISP, data.Org, data.Timezone)
	}
}
