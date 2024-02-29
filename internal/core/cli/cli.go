package cli

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	// Command info
	use   = "ryokai"
	short = "Orchestrator manager for Kira network"
	long  = "Orchestrator manager for Kira network"
)

func NewRyokaiCLI(commands []*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   use,
		Short: short,
		Long:  long,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		},
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	return rootCmd
}

func Run() {
	commands := []*cobra.Command{}
	c := NewRyokaiCLI(commands)

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
