package orchestration

import "context"

type Orchestration interface {
	Pull(ctx context.Context, url string) error
	Build(ctx context.Context, composeFile string) (string, error)
	Up(ctx context.Context, composeFile string) error
	Down(ctx context.Context, composeFile string) error
	Prune(ctx context.Context, composeFile string) error
}
