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