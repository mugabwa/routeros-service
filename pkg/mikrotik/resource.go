package mikrotik

import (
	"log"

	"github.com/go-routeros/routeros"
)

func FetchSystemResources(client *routeros.Client) ([]M, error) {
	cmd := "/system/resource/print"
	resp, err := client.Run(cmd)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	var data []M
	for _, res := range resp.Re{
		data = append(data, res.Map)
	}
	return data, nil
}

func FetchSystemIdentity(client *routeros.Client) (M, error) {
	cmd := "/system/identity/print"
	resp, err := client.Run(cmd)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	return resp.Re[len(resp.Re)-1].Map, nil
}

func PatchSystemIdentity(client *routeros.Client, hostname string) error {
	cmd := "/system/identity/set"
	payload := "=name=" + hostname
	_, err := client.Run(cmd, payload)
	if err != nil {
		log.Printf("Error running command: %v", err)
		return err
	}
	return nil
}