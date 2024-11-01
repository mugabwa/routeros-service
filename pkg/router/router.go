package router

import (
	"router_os/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/go-routeros/routeros"
)

func SetupRouter(client *routeros.Client) *gin.Engine {
	r := gin.Default()
	handler := &handlers.Handler{Client: client}

	r.GET("/ip", handler.GetIPAddresses)
	r.GET("/resources", handler.GetSystemResources)
	r.GET("/identity", handler.GetSystemIdentity)
	r.PATCH("/identity", handler.PatchSystemIdentity)
	r.GET("/dhcp", handler.GetDHCPClients)
	r.POST("/dhcp", handler.PostDHCPClient)
	return r
}