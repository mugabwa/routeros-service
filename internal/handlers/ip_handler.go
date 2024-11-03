package handlers

import (
	"net/http"
	"router_os/internal/services"
	"router_os/utility"

	"github.com/gin-gonic/gin"
)

type DCHPRequest struct {
	AddDefaultRoute string `json:"add-default-route"`
	Comment string `json:"comment"`
	Disabled string `json:"disabled"`
	Interface string `json:"interface"`
	UsePeerDns string `json:"use-peer-dns"`
	UsePeerNtp string `json:"use-peer-ntp"`
}

type FirewallRequest struct {
	Chain string `json:"chain"`
	Action string `json:"action"`
	SrcAddress string `json:"src-action"`
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

	payload := utility.M{
	  "add-default-route": request.AddDefaultRoute,
      "comment": request.Comment,
      "disabled": request.Disabled,
      "interface": request.Interface,
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

func (h *Handler) DeleteDHCPClient(ctx *gin.Context) {
	dhcpId := ctx.Param("id")
	err := services.DeleteDHCPClient(h.Client, dhcpId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"data": "DCHP record with id %v removed"})
}

func (h *Handler) GetFirewallRules(ctx *gin.Context) {
	data, err := services.FetchFirewallRules(h.Client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) PostFirewallRules(ctx *gin.Context) {
	var request FirewallRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload := utility.M {
		"chain": request.Chain,
		"action": request.Action,
		"src-address": request.SrcAddress,
	}
	data, err := services.AddFirewallRules(h.Client, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": data})
}

func (h *Handler) RemoveFirewallRule(ctx *gin.Context) {
	ruleId := ctx.Param("id")
	err := services.DeleteFirewallRule(h.Client, ruleId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"data": "DCHP record with id %v removed"})
}
