package types

type DApp struct {
	ID               string
	Name             string
	RepositoryUrl    string
	DockerComposeURL string
}

type DAppConfiguration struct {
	CpuAmount uint
	RamAmount uint
}
