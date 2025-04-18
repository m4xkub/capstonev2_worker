package migrate

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/utils"
)

type SyncRequest struct {
	PrivateIp string `json:"private_ip"`
}

func SyncData(c *gin.Context) {
	//sudo rsync -avz --exclude="lost+found" -e "ssh -i /home/ubuntu/capstonev2_worker/key.pem" /mnt/* ubuntu@43.208.136.216:/mnt

	var req SyncRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	destination := fmt.Sprintf("ubuntu@%s:/mnt", req.PrivateIp)
	utils.RunCommand("sudo", "rsync", "-avz",
		"--exclude=lost+found",
		"-e", "ssh -i /home/ubuntu/capstonev2_worker/key.pem",
		"/mnt/*",
		destination,
	)

	c.JSON(200, gin.H{"message": "complete"})
}
