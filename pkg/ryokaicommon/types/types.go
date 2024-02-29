package types

type ContainerSpec struct {
	Image   string
	Env     []string
	Ports   map[string]string
	Volumes map[string]string
}
