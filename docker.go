package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type DockerWatcher struct {
	image     string
	container string
	port      string
}

func NewDockerWatcher(image, port string) *DockerWatcher {
	return &DockerWatcher{image: image, port: port}
}

func (d *DockerWatcher) Build() error {
	fmt.Println("[hotreload] Building Docker image...")
	cmd := exec.Command("docker", "build", "-t", d.image, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (d *DockerWatcher) Start() error {
	d.Stop()
	args := []string{"run", "-d", "--name", d.containerName()}
	if d.port != "" {
		args = append(args, "-p", d.port)
	}
	args = append(args, d.image)
	cmd := exec.Command("docker", args...)
	out, err := cmd.Output()
	if err != nil { return err }
	d.container = strings.TrimSpace(string(out))
	fmt.Printf("[hotreload] Container started: %s\n", d.container[:12])
	return nil
}

func (d *DockerWatcher) Stop() {
	exec.Command("docker", "rm", "-f", d.containerName()).Run()
}

func (d *DockerWatcher) containerName() string {
	return "hotreload-" + d.image
}
