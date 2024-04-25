package dapp

import (
	"fmt"

	"github.com/KiraCore/ryokai/internal/core/orchestration"
	"github.com/KiraCore/ryokai/internal/core/orchestration/dockerCompose"
	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
)

func NewDApp(dappType, URL string, ID int) (*types.DApp, error) {
	dapp := &types.DApp{}
	switch dappType {

	case orchestration.DOCKER_COMPOSE:
		orchestrator := dockerCompose.NewDockerComposeOrchestrator(dapp)
		dapp = dappConstructor(dapp, dappType, URL, ID, orchestrator)
	default:
		return nil, fmt.Errorf("unsupported dapp-type: %v", dappType)
	}

	return dapp, nil
}

func dappConstructor(dapp *types.DApp, dappType, URL string, ID int, orchestrator orchestration.Orchestration) *types.DApp {
	dapp.Orchestrator = orchestrator
	dapp.Type = dappType
	dapp.ID = ID
	dapp.URL = URL
	return dapp
}
