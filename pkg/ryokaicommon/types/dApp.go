package types

type DApp struct {
	ID               string
	Name             string
	RepositoryUrl    string
	DockerComposeURL string
	Running          bool
	Resources        SystemResources
}
