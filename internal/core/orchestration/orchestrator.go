package orchestration

import (
	"context"
	"log/slog"

	"github.com/KiraCore/ryokai/internal/core/orchestration/docker"
	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
)

type Orchestrator interface {
	PullImage(ctx context.Context, imageName string) error
	CreateContainer(ctx context.Context, spec types.ContainerSpec) (string, error)
	StartContainer(ctx context.Context, containerID string) error
	StopContainer(ctx context.Context, containerID string) error
	RemoveContainer(ctx context.Context, containerID string) error
}

type AppManager struct {
	orchestrator Orchestrator
	AppDeployment
	AppDestroy
}

func NewAppManager(orchestrator Orchestrator) *AppManager {
	return &AppManager{orchestrator: orchestrator}
}

type AppDeployment struct {
	Spec        types.ContainerSpec
	PreRunSteps []interface{}
}

type AppDestroy struct {
	Spec         types.ContainerSpec
	PostRunSteps []interface{}
}

func (app *AppManager) DeployApplication(ctx context.Context, deployment *AppDeployment) error {
	var err error

	switch deployment.Spec.ContainerType {
	case "lxc":
	// ...
	case "docker":
		app.orchestrator, err = docker.NewDockerOrchestrator()
		if err != nil {
			return err
		}
		// /...
	case "podman":
	}

	// err = app.orchestrator.PullImage(ctx, deployment.Spec.Image)
	// if err != nil {
	// 	return err
	// }
	slog.Info("Deploying")

	containerID, err := app.orchestrator.CreateContainer(ctx, deployment.Spec)
	if err != nil {
		return err
	}

	// PreRunSteps can be `app.orchestrator.StartContainer' with specific arg`
	// for _, step := range deployment.PreRunSteps {
	// 	if _, err := app.orchestrator.ExecCommandInContainer(ctx, containerID, step.(string)); err != nil {
	// 		return err
	// 	}
	// }

	err = app.orchestrator.StartContainer(ctx, containerID)
	if err != nil {
		return err
	}

	return nil
}
