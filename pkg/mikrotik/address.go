package mikrotik

import (
	"log"
	"router_os/utility"
	"strings"

	"github.com/go-routeros/routeros"
)

func FetchIPAddresses(client *routeros.Client) (interface{}, error) {
	cmd := "/ip/address/print"
	resp, err := client.Run(cmd)
	
	if err := utility.ErrorCommand(err); err != nil {
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func FetchDHCPClient(client *routeros.Client) (interface{}, error) {
	cmd := "/ip/dhcp-client/print"
	resp, err := client.Run(cmd)

	if err := utility.ErrorCommand(err); err != nil {
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func AddDHCPClient(client *routeros.Client, payload utility.M) (interface{}, error) {
	cmd := "/ip/dhcp-client/add"
	var str_payload string
	for key, val := range payload {
		str_data := "=" + key + "=" + val + ","
		str_payload += str_data
	}
	str_payload = strings.TrimRight(str_payload, ",")
	log.Printf("Data: %v", str_payload)
	resp, err := client.Run(cmd, str_payload)

	if err := utility.ErrorCommand(err); err != nil {
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func DeleteDHCPClient(client *routeros.Client, id string) error {
	cmd := "/ip/dhcp-client/remove"
	param := "=.id=" + id
	_, err := client.Run(cmd, param)
	if err := utility.ErrorCommand(err); err != nil {
		return err
	}
	return nil
}

func FetchFirewallRules(client *routeros.Client) (interface{}, error) {
	cmd := "/ip/firewall/filter/print"
	resp, err := client.Run(cmd)
	if err := utility.ErrorCommand(err); err != nil {
		return nil, err
	}
	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func AddFirewallRules(client *routeros.Client, payload utility.M) (interface{}, error) {
	cmd := "/ip/firewall/filter/add"
	var str_payload string
	for key, val := range payload {
		str_data := "=" + key + "=" + val + ","
		str_payload += str_data
	}
	str_payload = strings.TrimRight(str_payload, ",")
	log.Printf("Data: %v", str_payload)
	resp, err := client.Run(cmd, str_payload)

	if err := utility.ErrorCommand(err); err != nil {
		return nil, err
	}

	data, err := utility.ParseResponse(resp)
	err = utility.ErrorParser(err)
	return data, err
}

func DeleteFirewallRule(client *routeros.Client, id string) error {
	cmd := "/ip/firewall/filter/remove"
	param := "=.id=" + id
	_, err := client.Run(cmd, param)
	if err := utility.ErrorCommand(err); err != nil {
		return err
	}
	return nil
}
