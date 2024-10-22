package router

import (
	"router_os/internal/handlers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ip", handlers.GetIPAddresses)
	return r
}