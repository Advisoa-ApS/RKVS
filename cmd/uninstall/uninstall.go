package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const serviceName = "rkvs.service"
const binaryName = "rkvs"
const destinationPath = "/usr/local/bin/"

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root. (sudo)")
		os.Exit(1)
	}

	fmt.Println("Attempting to stop and disable the service...")
	// Stop and disable the service
	systemctlCommands := [][]string{
		{"systemctl", "stop", serviceName},
		{"systemctl", "disable", serviceName},
	}

	for _, cmd := range systemctlCommands {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			fmt.Printf("Failed to run %s: %v\n", cmd[0], err)
			// Continue to try to clean up
		} else {
			fmt.Printf("Successfully ran %s.\n", cmd[0])
		}
	}

	servicePath := filepath.Join("/etc/systemd/system/", serviceName)
	fmt.Printf("Removing service file: %s\n", servicePath)
	if err := os.Remove(servicePath); err != nil {
		fmt.Printf("Failed to remove service file: %v\n", err)
	} else {
		fmt.Println("Service file removed successfully.")
	}

	// Reload systemd daemon to recognize that the service has been removed
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		fmt.Printf("Failed to reload systemd daemon: %v\n", err)
	} else {
		fmt.Println("Systemd daemon reloaded successfully.")
	}

	dest := filepath.Join(destinationPath, binaryName)
	fmt.Printf("Removing binary: %s\n", dest)
	if err := os.Remove(dest); err != nil {
		fmt.Printf("Failed to remove the binary from %s: %v\n", destinationPath, err)
		os.Exit(1)
	} else {
		fmt.Println("Binary removed successfully.")
	}

	fmt.Println("Service and binary uninstalled successfully.")
}
