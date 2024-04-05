package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const serviceName = "rkvs.service"
const binaryName = "rkvs"
const destinationPath = "/usr/local/bin/"

var serviceTemplate = `[Unit]
Description=Remote Key Value Store
After=network.target

[Service]
User=ubuntu
ExecStart={{.DestinationPath}}{{.BinaryName}} -ip {{.IP}} -port {{.Port}} -http_port {{.HTTPPort}}
Restart=always

[Install]
WantedBy=multi-user.target
`

type Config struct {
	IP              string
	Port            string
	HTTPPort        string
	BinaryName      string
	DestinationPath string
}

func main() {
	// Parse command-line flags
	ip := flag.String("ip", "127.0.0.1", "The IP address the service will listen on")
	port := flag.String("port", "8080", "The port the service will listen on for TCP connections")
	httpPort := flag.String("http_port", "80", "The HTTP port the service will listen on")
	flag.Parse()

	config := Config{
		IP:              *ip,
		Port:            *port,
		HTTPPort:        *httpPort,
		BinaryName:      binaryName,
		DestinationPath: destinationPath,
	}

	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root. (sudo)")
		os.Exit(1)
	}

	// Assuming binary is in the current directory, adjust if necessary
	execPath, err := filepath.Abs(binaryName)
	if err != nil {
		fmt.Printf("Failed to find the binary path: %v\n", err)
		os.Exit(1)
	}

	dest := filepath.Join(destinationPath, binaryName)
	if err := exec.Command("cp", execPath, dest).Run(); err != nil {
		fmt.Printf("Failed to copy the binary to %s: %v\n", destinationPath, err)
		os.Exit(1)
	}

	// Create a service file from the template
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error parsing service template:", err)
		os.Exit(1)
	}

	serviceFilePath := filepath.Join("/etc/systemd/system/", serviceName)
	serviceFile, err := os.OpenFile(serviceFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening service file:", err)
		os.Exit(1)
	}
	defer serviceFile.Close()

	// Execute the template and write to the service file
	if err := tmpl.Execute(serviceFile, config); err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}

	// Reload systemd, enable and start the service
	systemctlCommands := [][]string{
		{"systemctl", "daemon-reload"},
		{"systemctl", "enable", serviceName},
		{"systemctl", "start", serviceName},
	}

	for _, cmd := range systemctlCommands {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			fmt.Printf("Failed to run %s: %v\n", cmd[0], err)
			os.Exit(1)
		}
	}

	fmt.Println("Service installed and started successfully.")
}
