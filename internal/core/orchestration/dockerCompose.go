package orchestration

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

type DockerComposeOrchestrator struct {
	Dapp *DApp
}

func NewDockerComposeOrchestrator(Dapp *DApp) *DockerComposeOrchestrator {
	return &DockerComposeOrchestrator{Dapp: Dapp}
}

func (o *DockerComposeOrchestrator) Run(ctx context.Context, composeFilePath string) error {
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println(output)
	return nil
}

func (o *DockerComposeOrchestrator) Pull(ctx context.Context, composeFilePath string) error {
	fmt.Println(o.Dapp.ID)
	return nil
}

func (o *DockerComposeOrchestrator) Build(ctx context.Context, composeFilePath string) error {
	return nil
}

func (o *DockerComposeOrchestrator) Down(ctx context.Context, composeFilePath string) error {
	return nil
}

func (o *DockerComposeOrchestrator) Prune(ctx context.Context, composeFilePath string) error {
	return nil
}
