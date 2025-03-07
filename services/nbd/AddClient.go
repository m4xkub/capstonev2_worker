package nbd

import (
	"os/exec"
	"github.com/gin-gonic/gin"
)

type AddClientRequest struct {
	IP string `json:"ip"`
}

func AddClient(c *gin.Context)  {
	allowListPath := "/etc/nbd-server/allowlist"
	var req AddClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	cmd := exec.Command("echo", req.IP, ">>", allowListPath)
	output, err := cmd.Output()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": string(output)})
}