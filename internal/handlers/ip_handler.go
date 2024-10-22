package handlers

import (
	"net/http"
	"router_os/internal/services"

	"github.com/gin-gonic/gin"
)


func GetIPAddresses(ctx *gin.Context) {
	ips, err := services.FetchIPAddresses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ips": ips})
}