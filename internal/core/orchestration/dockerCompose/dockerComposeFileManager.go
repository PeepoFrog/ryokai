package dockerCompose

import (
	"context"
	"fmt"
	"os"

	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"

	"github.com/compose-spec/compose-go/v2/cli"
)

type DockerComposeProjectManager struct{}

func NewDockerComposeManager() *DockerComposeProjectManager {
	return &DockerComposeProjectManager{}
}

func (dcm *DockerComposeProjectManager) GetDockerComposeFile(ctx context.Context, url, savePath, fileName string) error {
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

func (dcm *DockerComposeProjectManager) GetDockerComposeProjectV2(ctx context.Context, dockerComposeFile string) (*types.Project, error) {
	content, err := os.ReadFile(dockerComposeFile)
	if err != nil {
		return nil, fmt.Errorf("error reading <%v> compose file: %w", dockerComposeFile, err)
	}
	project, err := loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{Content: content}},
		// WorkingDir:  getDAppSrcFolder(dcm.Dapp.ID),
	})
	if err != nil {
		return nil, fmt.Errorf("error loading <%v> compose file: %w", dockerComposeFile, err)
	}
	return project, nil
}

// func (dcm *DockerComposeProjectManager) VerifyProject(project *types.Project) (bool, error) {

// 	return true, nil
// }

func (dcm *DockerComposeProjectManager) PrepareProject(ctx context.Context, project *types.Project, id string) (*types.Project, error) {
	// var prepared *types.Project

	for i, s := range project.Services {
		s.ContainerName = fmt.Sprintf("%v-%v", id, s.Name)

		s.CPUCount = 1
		s.CPUS = 1
		s.CPUShares = 1
		fmt.Println(s.Name, s.ContainerName)
		fmt.Println(s)
		fmt.Println(i)
		project.Services[i] = s

	}

	return project, nil
}
