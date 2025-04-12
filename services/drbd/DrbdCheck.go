package drbd

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func DrbdCheck(c *gin.Context) {
	cmd := exec.Command("drbdadm", "status")

	// Run and get output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(output))
	if string(output) == "no resources defined!" {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "drbd is not init yet",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
