package whois

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type WhoisResponse struct {
	Domain    string `json:"domain"`
	Available string `json:"available"`
	Type      string `json:"type"`
	Registrar string `json:"registrar"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
	Expires   int64  `json:"expires"`
}

func LookupDomain(domain string, asJSON bool) {
	client, err := GetDefaultHTTPClient()
	if err != nil {
		fmt.Println("Fehler beim Erstellen des HTTP-Clients:", err)
		os.Exit(1)
	}

	resp, err := client.Get("https://api.whois.vu/?q=" + domain)
	if err != nil {
		fmt.Println("Fehler beim Abrufen:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data WhoisResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Fehler beim Dekodieren:", err)
		os.Exit(1)
	}

	if data.Domain == "" {
		fmt.Println("Keine gültige Whois-Antwort.")
		os.Exit(1)
	}

	if asJSON {
		json.NewEncoder(os.Stdout).Encode(data)
	} else {
		fmt.Printf("Domain:    %s\nTyp:       %s\nRegistrar: %s\nVerfügbar: %s\nErstellt:  %s\nGeändert:  %s\nAblauf:    %s\n",
			data.Domain, data.Type, data.Registrar, data.Available,
			time.Unix(data.Created, 0).Format("2006-01-02"),
			time.Unix(data.Updated, 0).Format("2006-01-02"),
			time.Unix(data.Expires, 0).Format("2006-01-02"))
	}
}
