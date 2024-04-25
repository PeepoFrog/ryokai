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

	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
	"github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/git"
	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"
)

type DockerComposeOrchestrator struct {
	ComposeFilePath string
	*types.DApp
}

func NewDockerComposeOrchestrator(dapp *types.DApp) *DockerComposeOrchestrator {
	defaultComposeFilePath := "sekin-compose.yml"
	return &DockerComposeOrchestrator{DApp: dapp, ComposeFilePath: defaultComposeFilePath}
}

func (o *DockerComposeOrchestrator) Run(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFilePath), "up", "-d")
	log.Printf("Running: %v", cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing <%v>: %w", cmd, err)
	}
	log.Println(output)
	return nil
}

func (o DockerComposeOrchestrator) Pull(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.ID)
	check := osutils.PathExists(dappSrcFolder)
	if check {
		err := os.RemoveAll(dappSrcFolder)
		if err != nil {
			return err
		}
	}
	err := git.CloneRepo(o.URL, dappSrcFolder)
	if err != nil {
		return fmt.Errorf("error while cloning <%v> repo: %w", o.URL, err)
	}
	return nil
}

func (o *DockerComposeOrchestrator) Build(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFilePath), "build")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing <%v>: %w", cmd, err)
	}
	log.Println(output)
	return nil
}

func (o *DockerComposeOrchestrator) Down(ctx context.Context) error {
	dappSrcFolder := getDAppSrcFolder(o.ID)
	cmd := exec.Command("docker-compose", "-f", filepath.Join(dappSrcFolder, o.ComposeFilePath), "down")
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
