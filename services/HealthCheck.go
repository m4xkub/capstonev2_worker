package services

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {

	// to do
	// extract output to see if it ready to work

	cmd := exec.Command("/usr/sbin/drbdadm", "status")
	// cmd := exec.Command("cat", "./test.txt")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	raw_output := strings.Split("/n", string(output))
	for i, e := range raw_output {
		raw_output[i] = strings.TrimSpace(e)
	}

	status := make(map[string]string)

	for _, e := range raw_output {
		if e == "" {
			continue
		}

		tmpElement := strings.Split(" ", e)
		for _, e := range tmpElement {
			if !strings.Contains(e, ":") {
				continue
			}

			x := strings.Split(e, ":")
			status[x[0]] = x[1]
		}

	}

	// fmt.Println(string(output))
	// fmt.Println(status["role"])
	fmt.Println(status["disk"])
	fmt.Println(status["peer-disk"])
	c.JSON(http.StatusOK, gin.H{"message": string(output)})
}
