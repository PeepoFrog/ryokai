package orchestration

import (
	"context"
	"fmt"

	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
)

// consts for dapp types
const (
	DOCKER_COMPOSE = "docker-compose"
)

type DApp struct {
	Type          string // docker compose, etc
	ID            int
	Name          string
	RepositoryUrl string
	Running       bool
	Resources     types.SystemResources
	Orchestrator  Orchestration
}

func NewDApp(dappType, URL string, ID int) (*DApp, error) {
	dapp := &DApp{}
	switch dappType {
	case DOCKER_COMPOSE:
		orchestrator := NewDockerComposeOrchestrator(dapp)
		dapp = dappConstructor(dapp, dappType, URL, ID, orchestrator)
	default:
		return nil, fmt.Errorf("unsupported dapp-type: %v", dappType)
	}

	return dapp, nil
}

func dappConstructor(dapp *DApp, dappType, URL string, ID int, orchestrator Orchestration) *DApp {
	dapp.Orchestrator = orchestrator
	dapp.Type = dappType
	dapp.RepositoryUrl = URL
	dapp.ID = ID
	return dapp
}

type Orchestration interface {
	Pull(ctx context.Context, url string) error
	Build(ctx context.Context, composeFile string) error
	Run(ctx context.Context, composeFile string) error
	Down(ctx context.Context, composeFile string) error
	Prune(ctx context.Context, composeFile string) error
}
