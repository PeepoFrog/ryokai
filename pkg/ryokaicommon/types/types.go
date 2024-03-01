package types

import "github.com/docker/docker/api/types/container"

type ContainerSpec struct {
	Image         string
	Env           []string
	Ports         map[string]string
	Volumes       map[string]string
	ContainerType string // Docker/lxc/podman
}

func GetDockerConfig(specs *ContainerSpec) (container.Config, container.HostConfig) {
	return container.Config{}, container.HostConfig{}
}

func GetPodmanConfig(specs *ContainerSpec) { // returns podmanconfig...
	// return config
}

func GetLXCConfig(specs *ContainerSpec) { // return lxc config
	// return config
}
