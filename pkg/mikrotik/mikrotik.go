package mikrotik

import (
	"log"

	"github.com/go-routeros/routeros"
)



func SetupClient(address, username, password string) (*routeros.Client, error) {
	client, err := routeros.Dial(address, username, password)
	if err != nil {
		log.Printf("Failed to connect to Mikrotik: %v", err)
		return nil, err
	}
	return client, nil
}