package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services"
)

func GetRootController() *gin.Engine {
	r := gin.Default()

	r.GET("/healthCheck", services.HealthCheck)
	r.GET("/promote", services.Promote)
	r.GET("/demote", services.Demote)
	return r
}
