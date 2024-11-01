package handlers

import (
	"net/http"
	"router_os/internal/services"

	"github.com/gin-gonic/gin"
)

type IdentityRequest struct {
	Hostname string `json:"hostname"`
}

func (h *Handler) GetSystemResources(ctx *gin.Context) {
	data, err := services.FetchSystemResourses(h.Client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) GetSystemIdentity(ctx *gin.Context) {
	data, err := services.FetchSystemIdentity(h.Client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) PatchSystemIdentity(ctx *gin.Context) {
	var request IdentityRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.PatchSystemIdentity(h.Client, request.Hostname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"data": "Hostname updated to " + request.Hostname})
}
