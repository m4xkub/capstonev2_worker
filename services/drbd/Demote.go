package drbd

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Demote(c *gin.Context) {
	cmd := exec.Command("sudo", "drbdadm", "secondary", "mydrbd")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": string(output)})

}
