package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxStatsCmd.AddCommand(SandboxStatsAirCmd)
}

// SandboxStatsAirCmd defines 'air' subcommand
var SandboxStatsAirCmd = &cobra.Command{
	Use:   "air",
	Short: TRCLI("cli.sandbox.stats.air.summary"),
	Long:  TRCLI(`cli.sandbox.stats.air.description`),
}
