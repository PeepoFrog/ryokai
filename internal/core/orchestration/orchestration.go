package orchestration

import (
	"context"
)

// consts for dapp types
const (
	DOCKER_COMPOSE = "docker-compose"
)

type Orchestration interface {
	Pull(ctx context.Context) error
	Build(ctx context.Context) error
	Run(ctx context.Context) error
	Down(ctx context.Context) error
	Prune(ctx context.Context) error
}
