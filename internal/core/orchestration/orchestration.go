package orchestration

import (
	"context"

	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
)

type DApp struct {
	Type          string // docker compose, etc
	ID            string
	Name          string
	RepositoryUrl string
	Running       bool
	Resources     types.SystemResources
	Orchestrator  Orchestration
}

func NewDApp(dappType, URL string) {
}

type Orchestration interface {
	Pull(ctx context.Context, url string) error
	Build(ctx context.Context, composeFile string) (string, error)
	Run(ctx context.Context, composeFile string) error
	Down(ctx context.Context, composeFile string) error
	Prune(ctx context.Context, composeFile string) error
}
