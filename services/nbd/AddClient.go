package nbd

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/utils"
)

type AddClientRequest struct {
	IP string `json:"ip"`
}

func AddClient(c *gin.Context) {
	allowListPath := "/etc/nbd-server/allowlist"
	var req AddClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := utils.RunCommand("sudo", "echo", req.IP, ">>", allowListPath)

	//cmd := exec.Command("echo", req.IP, ">>", allowListPath)
	//output, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Client added successfully!"})
}
