package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxStatsCmd.AddCommand(SandboxStatsBeamCmd)
}

// SandboxStatsBeamCmd defines 'beam' subcommand
var SandboxStatsBeamCmd = &cobra.Command{
	Use:   "beam",
	Short: TRCLI("cli.sandbox.stats.beam.summary"),
	Long:  TRCLI(`cli.sandbox.stats.beam.description`),
}
