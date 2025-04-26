package nbd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/utils"
)

func InitializeConfigFile(c *gin.Context) {

	// install and enable nbd module
	utils.RunCommand("sudo", "apt-get", "install", "nbd-server", "nbd-client")
	utils.RunCommand("sudo", "modprobe", "nbd")

	// config file
	newConfig := `[generic]
	listenaddr = 0.0.0.0
	allowlist = true
	user = root
	group = root

	[export1]
	exportname = /var/nbd-disk.img
	readonly = false
	allowlist = true
	authfile = /etc/nbd-server/allowlist
	multifile = false
	copyonwrite = false
	`

	// Overwrite the config file
	err := os.WriteFile("/etc/nbd-server/config", []byte(newConfig), 0644)
	if err != nil {
		fmt.Println("Error writing new config file:", err)
		c.JSON(500, gin.H{"error": "Failed to overwrite NBD config file"})
		return
	}

	utils.RunCommand("sudo", "systemctl", "restart", "nbd-server")

	fmt.Println("NBD Config fully replaced and server restarted.")
	c.JSON(200, gin.H{"message": "NBD config replaced and server restarted successfully!"})
}
