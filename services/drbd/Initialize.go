package drbd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

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
	// fmt.Println("Updating system...")
	// runCommand("sudo", "apt", "update")
	// time.Sleep(10 * time.Second) // Sleep for 5 seconds

	// fmt.Println("Upgrading system...")
	// runCommand("sudo", "apt", "upgrade", "-y")
	// time.Sleep(10 * time.Second)

	fmt.Println("Installing drbd-utils...")
	runCommand("sudo", "apt", "install", "drbd-utils", "-y")
	time.Sleep(10 * time.Second)

	fmt.Println("Performing another upgrade...")
	runCommand("sudo", "apt", "-y", "upgrade")
	time.Sleep(10 * time.Second)

	fmt.Println("Creating DRBD configuration file...")
	runCommand("sudo", "touch", "/etc/drbd.d/mydrbd.res")

	fmt.Println("test touch file")
	runCommand("sudo", "touch", "./test.res")

	fmt.Println("Initialization complete!")
}
