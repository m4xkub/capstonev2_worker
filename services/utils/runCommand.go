package utils

import (
	"fmt"
	"os/exec"
)

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
