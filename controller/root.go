package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/drbd"
	"github.com/m4xkub/capstonev2_worker/services/nbd"
)

func GetRootController() *gin.Engine {
	r := gin.Default()
	r.GET("/test", drbd.InitializeInstance)
	r.GET("/healthCheck", drbd.HealthCheck)
	r.GET("/promote", drbd.Promote)
	r.GET("/demote", drbd.Demote)
	r.POST("/addClient", nbd.AddClient)
	return r
}
