package handlers

import (
	"net/http"
	"router_os/internal/services"
	"router_os/pkg/mikrotik"

	"github.com/gin-gonic/gin"
)

type DCHPRequest struct {
	AddDefaultRoute string `json:"add-default-route"`
	Address string `json:"address"`
	Comment string `json:"comment"`
	DefaultRouteDistance string `json:"default-route-distance"`
	DhcpOptions string `json:"dhcp-options"`
	DhcpServer string `json:"dhcp-server"`
	Disabled string `json:"disabled"`
	Dynamic string `json:"dynamic"`
	ExpiresAfter string `json:"expires-after"`
	Gateway string `json:"gateway"`
	Interface string `json:"interface"`
	Invalid string `json:"invalid"`
	PrimaryDns string `json:"primary-dns"`
	Status string `json:"status"`
	UsePeerDns string `json:"use-peer-dns"`
	UsePeerNtp string `json:"use-peer-ntp"`
}

func (h *Handler) GetIPAddresses(ctx *gin.Context) {
	ips, err := services.FetchIPAddresses(h.Client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ips})
}

func (h *Handler) GetDHCPClients(ctx *gin.Context) {
	data, err := services.FetchDHCPClient(h.Client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) PostDHCPClient(ctx *gin.Context) {
	var request DCHPRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := mikrotik.M{
		"add-default-route": request.AddDefaultRoute,
      "address": request.Address,
      "comment": request.Comment,
      "default-route-distance": request.DefaultRouteDistance,
      "dhcp-options": request.DhcpOptions,
      "dhcp-server": request.DhcpServer,
      "disabled": request.Disabled,
      "dynamic": request.Dynamic,
      "expires-after": request.ExpiresAfter,
      "gateway": request.Gateway,
      "interface": request.Interface,
      "invalid": request.Invalid,
      "primary-dns": request.PrimaryDns,
      "status": request.Status,
      "use-peer-dns": request.UsePeerDns,
      "use-peer-ntp": request.UsePeerNtp,
	}
	data, err := services.AddDHCPClient(h.Client, payload)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": data})
}
