package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/drbd"
)

func GetRootController() *gin.Engine {
	r := gin.Default()

	r.GET("/healthCheck", drbd.HealthCheck)
	r.GET("/promote", drbd.Promote)
	r.GET("/demote", drbd.Demote)
	return r
}
