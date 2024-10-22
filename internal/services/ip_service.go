package services

import (
	"log"
	"os"
	"router_os/pkg/mikrotik"
)


func FetchIPAddresses() ([]string, error) {
	client, err := mikrotik.SetupClient(
		os.Getenv("ROUTER_ADDRESS"), os.Getenv("ROUTER_USERNAME"),
		os.Getenv("ROUTER_PASSWORD"))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ips, err := mikrotik.FetchIPAddresses(client)
	if err != nil {
		log.Printf("Failed to fetch IP addresses: %v", err)
		return nil, err
	}
	return ips, nil
}