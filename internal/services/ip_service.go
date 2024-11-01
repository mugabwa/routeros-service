package services

import (
	"log"
	"router_os/pkg/mikrotik"

	"github.com/go-routeros/routeros"
)


func FetchIPAddresses(client *routeros.Client) ([]mikrotik.M, error) {
	ips, err := mikrotik.FetchIPAddresses(client)
	if err != nil {
		log.Printf("Failed to fetch IP addresses: %v", err)
		return nil, err
	}
	return ips, nil
}

func FetchDHCPClient(client *routeros.Client) ([]mikrotik.M, error) {
	data, err := mikrotik.FetchDHCPClient(client)
	if err != nil {
		log.Printf("Failed to fetch DHCP clients: %v", err)
		return nil, err
	}
	return data, nil
}

func AddDHCPClient(client *routeros.Client, payload mikrotik.M) ([]mikrotik.M, error) {
	data, err := mikrotik.AddDHCPClient(client, payload)
	if err != nil {
		log.Printf("Failed to add DCHP clients: %v", err)
		return nil, err
	}
	return data, nil
}
