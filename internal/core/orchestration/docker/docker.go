package docker

import (
	"github.com/docker/docker/client"
)

type DockerOrchestrator struct {
	cli *client.Client
}

func NewDockerOrchestrator() (*DockerOrchestrator, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerOrchestrator{cli: cli}, nil
}
