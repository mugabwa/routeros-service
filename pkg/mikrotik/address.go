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

func FetchDHCPClient(client *routeros.Client) ([]M, error) {
	cmd := "/ip/dhcp-client/print"
	resp, err := client.Run(cmd)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	var data []M
	for _, re := range resp.Re {
		data = append(data, re.Map)
	}
	return data, nil
}

func AddDHCPClient(client *routeros.Client, payload M) ([]M, error) {
	cmd := "/ip/dhcp-client/add"
	str_payload := "="
	for key, val := range payload {
		str_data := key + "=" + val + ","
		str_payload += str_data
	}
	resp, err := client.Run(cmd, str_payload)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	var data []M
	for _, re := range resp.Re {
		data = append(data, re.Map)
	}
	return data, nil
}
