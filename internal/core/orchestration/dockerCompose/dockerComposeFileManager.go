package dockerCompose

import (
	"context"
	"fmt"

	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"
	"github.com/compose-spec/compose-go/v2/types"

	"github.com/compose-spec/compose-go/v2/cli"
)

type DockerComposeProjectManager struct{}

func NewDockerComposeManager() *DockerComposeProjectManager {
	return &DockerComposeProjectManager{}
}

func (dcm *DockerComposeProjectManager) GetDockerComposeFile(url, savePath, fileName string) error {
	err := osutils.DownloadFile(url, savePath, fileName)
	if err != nil {
		return err
	}
	return nil
}

func (dcm *DockerComposeProjectManager) GetDockerComposeProject(ctx context.Context, dockerComposeFile string) (*types.Project, error) {
	options, err := cli.NewProjectOptions(
		[]string{dockerComposeFile},
		cli.WithOsEnv,
		cli.WithDotEnv,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to verify docker-compose file: %w", err)
	}
	project, err := cli.ProjectFromOptions(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("unable to create project from file: %w", err)
	}
	return project, nil
}
