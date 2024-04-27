package types

import "github.com/KiraCore/ryokai/internal/core/orchestration"

const DAppsSrcFolder string = "/tmp/srcFolder"

type DApp struct {
	Type         string // docker compose, etc
	ID           int
	Name         string
	Running      bool
	Resources    SystemResources
	Orchestrator orchestration.Orchestration
	URL          string
}
