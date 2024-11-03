package services

import (
	"log"
	"router_os/pkg/mikrotik"
	"router_os/utility"

	"github.com/go-routeros/routeros"
)


func FetchIPAddresses(client *routeros.Client) (interface{}, error) {
	data, err := mikrotik.FetchIPAddresses(client)
	if err != nil {
		log.Printf("Failed to fetch IP addresses: %v", err)
		return nil, err
	}
	return data, nil
}

func FetchDHCPClient(client *routeros.Client) (interface{}, error) {
	data, err := mikrotik.FetchDHCPClient(client)
	if err != nil {
		log.Printf("Failed to fetch DHCP clients: %v", err)
		return nil, err
	}
	return data, nil
}

func AddDHCPClient(client *routeros.Client, payload utility.M) (interface{}, error) {
	data, err := mikrotik.AddDHCPClient(client, payload)
	if err != nil {
		log.Printf("Failed to add DCHP clients: %v", err)
		return nil, err
	}
	return data, nil
}

func DeleteDHCPClient(client *routeros.Client, id string) error {
	err := mikrotik.DeleteDHCPClient(client, id)
	if err != nil {
		log.Printf("Failed to delete DHCP clients: %v with id %v", err, id)
		return err
	}
	return nil
}

func FetchFirewallRules(client *routeros.Client) (interface{}, error) {
	data, err := mikrotik.FetchFirewallRules(client)
	if err != nil {
		log.Printf("Failed to fetch firewall rules: %v", err)
		return nil, err
	}
	return data, nil
}

func AddFirewallRules(client *routeros.Client, payload utility.M) (interface{}, error) {
	data, err := mikrotik.AddFirewallRules(client, payload)
	if err != nil {
		log.Printf("Failed to add firewall rules: %v", err)
		return nil, err
	}
	return data, nil
}

func DeleteFirewallRule(client *routeros.Client, id string) error {
	err := mikrotik.DeleteFirewallRule(client, id)
	if err != nil {
		log.Printf("Failed to delete firewall rule: %v with id %v", err, id)
		return err
	}
	return nil
}
