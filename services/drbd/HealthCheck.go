package drbd

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	cmd := exec.Command("/usr/sbin/drbdadm", "status")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	//raw_output := strings.Split("\n",string(output))
	raw_output := strings.Split(string(output), "\n")

	fmt.Println("output : ", string(output))
	fmt.Println("raw_output : ", raw_output)

	for i, e := range raw_output {
		raw_output[i] = strings.TrimSpace(e)
	}

	status := make(map[string]string)

	for _, e := range raw_output {
		if e == "" {
			continue
		}

		if strings.Contains(e, "mydrbd role") {
			x := strings.Split(e, ":")
			status["role"] = x[1]
		} else if strings.Contains(e, "disk") {
			x := strings.Split(e, ":")
			status["disk-status"] = x[1]
			break
		}

	}

	fmt.Println(status["disk"])
	fmt.Println(status["peer-disk"])
	c.JSON(http.StatusOK, status)
}
