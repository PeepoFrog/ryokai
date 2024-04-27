package dockerCompose

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"

	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"

	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
)

type DockerComposeOrchestrator struct {
	*DockerComposeProjectManager
	ComposeFileName string
	Dapp            *types.DApp
	// *types.DApp
}

func NewDockerComposeOrchestrator(dapp *types.DApp) *DockerComposeOrchestrator {
	defaultComposeFilePath := "compose.yml"
	dcm := NewDockerComposeManager()
	return &DockerComposeOrchestrator{ComposeFileName: defaultComposeFilePath, DockerComposeProjectManager: dcm}
	// return &DockerComposeOrchestrator{DApp: dapp, ComposeFileName: defaultComposeFilePath, DockerComposeProjectManager: dcm}
}

func (o *DockerComposeOrchestrator) Run(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.Dapp.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFileName), "up", "-d")
	log.Printf("Running: %v", cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing <%v>: %w", cmd, err)
	}
	log.Println(output)
	return nil
}

func (o *DockerComposeOrchestrator) Pull(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.Dapp.ID)
	check := osutils.PathExists(dappSrcFolder)
	if check {
		err := os.RemoveAll(dappSrcFolder)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(dappSrcFolder, 0o777)
	if err != nil {
		return err
	}
	err = o.GetDockerComposeFile(ctx, o.Dapp.URL, getDAppSrcFolder(o.Dapp.ID), o.ComposeFileName)
	if err != nil {
		return err
	}
	return nil
}

func (o *DockerComposeOrchestrator) Build(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.Dapp.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFileName), "build")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing <%v>: %w", cmd, err)
	}
	log.Println(output)
	return nil
}

func (o *DockerComposeOrchestrator) Down(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.Dapp.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFileName), "down")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing <%v>: %w", cmd, err)
	}
	log.Println(output)
	return nil
}

func (o *DockerComposeOrchestrator) Prune(ctx context.Context) error {
	return nil
}

func getDAppSrcFolder(dappID int) string {
	dappSrcFolder := path.Join(path.Clean(types.DAppsSrcFolder), strconv.Itoa(dappID))
	return dappSrcFolder
}
