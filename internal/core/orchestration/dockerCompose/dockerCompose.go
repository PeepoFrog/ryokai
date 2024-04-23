package dockercompose

import (
	"os/exec"
)

type DockerComposeOrchestrator struct{}

func NewDockerComposeOrchestrator() *DockerComposeOrchestrator {
	return &DockerComposeOrchestrator{}
}

func (o *DockerComposeOrchestrator) Run(composeFilePath string) ([]byte, error) {
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
