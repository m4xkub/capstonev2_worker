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

func RunCommand(name string, args ...string) error {
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
	RunCommand("sed", "-i", fmt.Sprintf("s/hostname%s/%s/g", instanceNumber, hostname), "./mydrbd.res")

	// disk
	RunCommand("sed", "-i", fmt.Sprintf("s/ec2disk%s/%s/g", instanceNumber, escapedDiskname), "./mydrbd.res")

	// private ip
	RunCommand("sed", "-i", fmt.Sprintf("s/privateIp%s/%s/g", instanceNumber, privateIp), "./mydrbd.res")
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

	RunCommand("ln", "./mydrbd.res", "/etc/drbd.d/mydrbd.res")
	fmt.Println("Initialization complete!")
}

func InitializeMetaData(c *gin.Context) {

	fmt.Println("Creating DRBD Meta Data")

	RunCommand("sudo", "drbdadm", "create-md", "mydrbd")
	RunCommand("sudo", "drbdadm", "up", "mydrbd")
	RunCommand("sudo", "drbdadm", "secondary", "mydrbd")

	fmt.Println("DRBD Meta Data Created")

}

func MountVolume() {
	fmt.Println("Mounting file")

	RunCommand("sudo", "mkfs.ext4", "/dev/drbd0")
	RunCommand("sudo", "mkdir", "/mnt")
	RunCommand("sudo", "mount", "/dev/drbd0", "/mnt")

	fmt.Println("Mounted file")

}

func UnVolume() {
	fmt.Println("Unmounting file")

	RunCommand("sudo", "umount", "/mnt")

	fmt.Println("unmounted file")

}
