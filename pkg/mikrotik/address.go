package mikrotik

import (
	"log"

	"github.com/go-routeros/routeros"
)


func FetchIPAddresses(client *routeros.Client) ([]string, error) {
	cmd := "/ip/address/print"
	res, err := client.Run(cmd)
	
	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	var ips []string
	for _, re := range res.Re{
		ip := re.Map["address"]
		ips = append(ips, ip)
	}
	return ips, nil
}

