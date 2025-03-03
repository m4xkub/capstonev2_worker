package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Green"})
	})

	r.Run(":8080")
}
