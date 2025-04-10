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
	DiskName   string `json:"disk_name"`
}

func escapeSlashes(input string) string {
	return strings.ReplaceAll(input, "/", "\\/")
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing %s: %s\n", name, err.Error())
	}
	return err
}

func manageResFile(instanceNumber string, privateIp string, diskname string) {
	parts := strings.Split(privateIp, ".")

	hostname := "ip" + "-" + strings.Join(parts, "-")

	escapedDiskname := escapeSlashes(diskname)

	// hostname
	runCommand("sed", "-i", fmt.Sprintf("s/hostname%s/%s/g", instanceNumber, hostname), "./mydrbd.res")

	// disk
	runCommand("sed", "-i", fmt.Sprintf("s/ec2disk%s/%s/g", instanceNumber, escapedDiskname), "./mydrbd.res")

	// private ip
	runCommand("sed", "-i", fmt.Sprintf("s/privateIp%s/%s/g", instanceNumber, privateIp), "./mydrbd.res")
}

func InitializeConfigFile(c *gin.Context) {

	var req InitializationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating DRBD configuration file...")

	manageResFile("1", req.PrivateIp1, req.DiskName)
	manageResFile("2", req.PrivateIp2, req.DiskName)

	runCommand("ln", "./mydrbd.res", "/etc/drbd.d/mydrbd.res")
	fmt.Println("Initialization complete!")
}

func InitializeMetaData(c *gin.Context) {

	fmt.Println("Creating DRBD Meta Data")

	runCommand("sudo", "drbdadm", "create-md", "mydrbd")
	runCommand("sudo", "drbdadm", "up", "mydrbd")
	runCommand("sudo", "drbdadm", "secondary", "mydrbd")

	fmt.Println("DRBD Meta Data Created")

}
