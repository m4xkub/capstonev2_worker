package drbd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type InitializationRequest struct {
	PrivateIp1 string `json:"private_ip_1"`
	PrivateIp2 string `json:"private_ip_2"`
	PrivateIp3 string `json:"private_ip_3"`
	DiskName   string `json:"disk_name"`
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing %s: %v\n", name, err)
	}
	return err
}

func InitializeInstance(c *gin.Context) {

	var req InitializationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating DRBD configuration file...")
	runCommand("ln", "/etc/drbd.d/mydrbd.res", "./mydrbd.res")

	manageResFile("1", req.PrivateIp1, req.DiskName)
	manageResFile("2", req.PrivateIp2, req.DiskName)
	manageResFile("3", req.PrivateIp3, req.DiskName)

	fmt.Println("Initialization complete!")
}

func manageResFile(instanceNumber string, privateIp string, diskname string) {
	parts := strings.Split(privateIp, ".")

	hostname := "ip" + "-" + strings.Join(parts, "-")

	// hostname
	runCommand("sed", "-i", fmt.Sprintf("s/hostname%s/%s/g", instanceNumber, hostname), "./mydrbd.res")

	// disk
	runCommand("sed", "-i", fmt.Sprintf("s/ec2disk%s/%s/g", instanceNumber, diskname), "./mydrbd.res")

	// private ip
	runCommand("sed", "-i", fmt.Sprintf("s/privateIp%s/%s/g", instanceNumber, privateIp), "./mydrbd.res")
}
