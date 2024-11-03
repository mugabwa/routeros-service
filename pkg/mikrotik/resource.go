package mikrotik

import (
	"log"
	"router_os/utility"

	"github.com/go-routeros/routeros"
)

func FetchSystemResources(client *routeros.Client) (interface{}, error) {
	cmd := "/system/resource/print"
	resp, err := client.Run(cmd)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func FetchSystemIdentity(client *routeros.Client) (interface{}, error) {
	cmd := "/system/identity/print"
	resp, err := client.Run(cmd)

	if err != nil {
		log.Printf("Error running command: %v", err)
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
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