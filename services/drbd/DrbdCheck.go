package drbd

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// func DrbdCheck(c *gin.Context) {
// 	cmd := exec.Command("drbdadm", "status")

// 	// Run and get output
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println("Error:", err.Error())
// 		return
// 	}
// 	fmt.Println(string(output))
// 	if string(output) == "no resources defined!" {
// 		c.JSON(http.StatusNotImplemented, gin.H{
// 			"message": "drbd is not init yet",
// 		})
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok",
// 	})
// }

func DrbdCheck(c *gin.Context) {
	cmd := exec.Command("drbdadm", "status")

	output, err := cmd.CombinedOutput() // output includes both stdout + stderr

	// Always print the full output, even if there's an error
	fmt.Println("DRBD command output:\n", string(output))

	if err != nil {
		// Return output + error as JSON
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"details": string(output),
		})
		return
	}

	// Handle special case: "no resources defined!"
	if string(output) == "no resources defined!" {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "drbd is not init yet",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"output":  string(output),
	})
}
