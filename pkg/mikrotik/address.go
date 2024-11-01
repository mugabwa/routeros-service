package mikrotik

import (
	"log"

	"github.com/go-routeros/routeros"
)

type M map[string]string

func FetchIPAddresses(client *routeros.Client) ([]M, error) {
	cmd := "/ip/address/print"
	res, err := client.Run(cmd)
	
	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	var ips []M
	for _, re := range res.Re{
		ips = append(ips, re.Map)
	}
	return ips, nil
}

