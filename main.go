package main

import (
	"log"
	"os"
	"router_os/internal/config"
	"router_os/pkg/mikrotik"
	"router_os/pkg/router"
)

func main() {
	config.Load()
	client, err := mikrotik.SetupClient(
		os.Getenv("ROUTER_ADDRESS"), os.Getenv("ROUTER_USERNAME"),
		os.Getenv("ROUTER_PASSWORD"))
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	defer client.Close()

	r := router.SetupRouter(client)
	r.Run(":8080")
}