package services

import (
	"log"
	"router_os/pkg/mikrotik"

	"github.com/go-routeros/routeros"
)


func FetchSystemResourses(client *routeros.Client) ([]mikrotik.M, error) {
	data, err := mikrotik.FetchSystemResources(client)
	if err != nil {
		log.Printf("Failed to fetch system resources: %v", err)
		return nil, err
	}
	return data, nil
}

func FetchSystemIdentity(client *routeros.Client) (mikrotik.M, error) {
	data, err := mikrotik.FetchSystemIdentity(client)
	if err != nil {
		log.Printf("Failed to fetch system identity: %v", err)
		return nil, err
	}
	return data, nil
}

func PatchSystemIdentity(client *routeros.Client, hostname string) error {
	err := mikrotik.PatchSystemIdentity(client, hostname)
	if err != nil {
		log.Printf("Failed to patch system identity: %v", err)
		return err
	}
	return err
}