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
	r.GET("/system/resources", handler.GetSystemResources)
	r.GET("/system/identity", handler.GetSystemIdentity)
	r.PATCH("/system/identity", handler.PatchSystemIdentity)
	r.GET("/ip/dhcp", handler.GetDHCPClients)
	r.POST("/ip/dhcp", handler.PostDHCPClient)
	r.DELETE("/ip/dhcp/:id", handler.DeleteDHCPClient)
	r.GET("/ip/firewall-rules", handler.GetFirewallRules)
	r.POST("/ip/firewall-rules", handler.PostFirewallRules)
	r.DELETE("/ip/firewall-rules/:id", handler.RemoveFirewallRule)
	return r
}