package drbd

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_worker/services/utils"
)

type InitializationRequest struct {
	PrivateIp1 string `json:"private_ip_1"`
	PrivateIp2 string `json:"private_ip_2"`
	DiskName   string `json:"disk_name"`
}

func escapeSlashes(input string) string {
	return strings.ReplaceAll(input, "/", "\\/")
}

func manageResFile(instanceNumber string, privateIp string, diskname string) {
	parts := strings.Split(privateIp, ".")

	hostname := "ip" + "-" + strings.Join(parts, "-")

	escapedDiskname := escapeSlashes(diskname)

	// hostname
	utils.RunCommand("sudo", "sed", "-i", fmt.Sprintf("s/hostname%s/%s/g", instanceNumber, hostname), "./mydrbd.res")

	// disk
	utils.RunCommand("sudo", "sed", "-i", fmt.Sprintf("s/ec2disk%s/%s/g", instanceNumber, escapedDiskname), "./mydrbd.res")

	// private ip
	utils.RunCommand("sudo", "sed", "-i", fmt.Sprintf("s/privateIp%s/%s/g", instanceNumber, privateIp), "./mydrbd.res")
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

	utils.RunCommand("sudo", "ln", "./mydrbd.res", "/etc/drbd.d/mydrbd.res")
	fmt.Println("Initialization complete!")
}

func InitializeMetaData(c *gin.Context) {

	fmt.Println("Creating DRBD Meta Data")

	utils.RunCommand("sudo", "drbdadm", "create-md", "mydrbd")

	utils.RunCommand("sudo", "drbdadm", "up", "mydrbd")

	utils.RunCommand("sudo", "drbdadm", "secondary", "mydrbd")

	fmt.Println("DRBD Meta Data Created")

}

func MakeFileSystem(c *gin.Context) {
	fmt.Println("Making filesystem")

	utils.RunCommand("sudo", "mkfs.ext4", "/dev/drbd0")
	utils.RunCommand("sudo", "chown", "ubuntu:ubuntu", "/mnt")
	utils.RunCommand("sudo", "chmod", "u+w", "/mnt")

	fmt.Println("filesystem made")
}

func MountVolume() {
	fmt.Println("Mounting file")

	utils.RunCommand("sudo", "mount", "/dev/drbd0", "/mnt")

	fmt.Println("Mounted file")

}

func Unvolume() {
	fmt.Println("Unmounting file")

	utils.RunCommand("sudo", "umount", "/mnt")

	fmt.Println("unmounted file")

}
