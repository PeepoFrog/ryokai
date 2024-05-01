package dockerCompose

import (
	"context"
	"fmt"

	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
)

type DockerComposeProjectManager struct {
	Project *types.Project
}

func NewDockerComposeManager(dockerComposeFile string) (*DockerComposeProjectManager, error) {
	p, err := GetDockerComposeProjectV2(dockerComposeFile)
	if err != nil {
		return nil, err
	}
	return &DockerComposeProjectManager{Project: p}, nil
}

func GetDockerComposeFile(ctx context.Context, url, savePath, fileName string) error {
	err := osutils.DownloadFile(url, savePath, fileName)
	if err != nil {
		return err
	}
	return nil
}

// func (dcm *DockerComposeProjectManager) GetDockerComposeProject(ctx context.Context, dockerComposeFile string) (*types.Project, error) {
// 	options, err := cli.NewProjectOptions(
// 		[]string{dockerComposeFile},
// 		cli.WithOsEnv,
// 		cli.WithDotEnv,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to verify docker-compose file: %w", err)
// 	}
// 	project, err := cli.ProjectFromOptions(ctx, options)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to create project from file: %w", err)
// 	}
// 	return project, nil
// }

func GetDockerComposeProjectV2(dockerComposeFile string) (*types.Project, error) {
	project, err := loader.Load(types.ConfigDetails{
		// ConfigFiles: []types.ConfigFile{{Content: content, Filename: dockerComposeFile}},
		ConfigFiles: []types.ConfigFile{{Filename: dockerComposeFile}},
	}, loader.WithSkipValidation)
	// WorkingDir:  getDAppSrcFolder(dcm.Dapp.ID),
	if err != nil {
		return nil, fmt.Errorf("error loading <%v> compose file: %w", dockerComposeFile, err)
	}
	return project, nil
}

func (dcm *DockerComposeProjectManager) PrepareProject(ctx context.Context, id string) (*types.Project, error) {
	dcm.prepareCPU()
	dcm.prepareRam()

	for _, s := range dcm.Project.Services {
		fmt.Println(s.MemLimit)
	}

	return dcm.Project, nil
}

func (dcm *DockerComposeProjectManager) prepareCPU() {
	servicesWithOutCPU := []types.ServiceConfig{}
	for _, s := range dcm.Project.Services {
		if s.CPUS == 0 {
			servicesWithOutCPU = append(servicesWithOutCPU, s)
		}
		dcm.Project.Services[s.Name] = s
	}
	notUsed := len(servicesWithOutCPU)
	defaultCPUValue := 1
	if notUsed > 0 {
		var CPUshare float32 = float32(defaultCPUValue) / float32(notUsed)
		for _, s := range servicesWithOutCPU {
			s.CPUS = float32(CPUshare)
			dcm.Project.Services[s.Name] = s
		}
	}
}

func (dcm *DockerComposeProjectManager) prepareRam() {
	servicesWithOutRAM := []types.ServiceConfig{}
	for _, s := range dcm.Project.Services {
		if s.MemLimit == 0 {
			servicesWithOutRAM = append(servicesWithOutRAM, s)
		}
	}
	notUsed := len(servicesWithOutRAM)
	defaultRamValueInBytes := 1073741824
	if notUsed > 0 {
		var ramShare float64 = float64(defaultRamValueInBytes) / float64(notUsed)
		fmt.Println(ramShare, notUsed)
		for _, s := range servicesWithOutRAM {
			s.MemLimit = types.UnitBytes(ramShare)
			dcm.Project.Services[s.Name] = s
		}
	}
}
