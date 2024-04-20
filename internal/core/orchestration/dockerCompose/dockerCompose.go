package dockercompose

import (
	"os/exec"
)

func ComposeUp(composeFilePath string) ([]byte, error) {
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
