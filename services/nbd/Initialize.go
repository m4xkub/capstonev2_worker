package nbd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/utils"
)

func InitializeConfigFile(c *gin.Context) {

	// install and enable nbd module
	utils.RunCommand("sudo", "apt-get", "install", "nbd-server", "nbd-client")
	utils.RunCommand("sudo", "modprobe", "nbd")

	// Your new config content
	newConfig := `[generic]
	listenaddr = 0.0.0.0
	allowlist = true
	user = root
	group = root

	[export1]
	exportname = /dev/drbd0
	readonly = false
	allowlist = true
	authfile = /etc/nbd-server/allowlist
	multifile = false
	copyonwrite = false
	`

	// Delete the old config file
	utils.RunCommand("sudo", "rm", "-f", "/etc/nbd-server/config")

	// Write new config
	utils.RunCommand("bash", "-c", fmt.Sprintf("echo '%s' | sudo tee /etc/nbd-server/config", newConfig))

	utils.RunCommand("sudo", "touch", "/etc/nbd-server/allowlist")

	utils.RunCommand("sudo", "systemctl", "restart", "nbd-server")

	utils.RunCommand("sudo", "umount", "/mnt")
	fmt.Println("NBD Config fully replaced and server restarted.")
	c.JSON(200, gin.H{"message": "NBD config replaced and server restarted successfully!"})
}
