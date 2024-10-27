package main

import (
	"router_os/internal/config"
	"router_os/pkg/router"
)

func main() {
	config.Load()

	r := router.SetupRouter()
	r.Run(":8080")
}