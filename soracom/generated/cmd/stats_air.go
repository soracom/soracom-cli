package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	StatsCmd.AddCommand(StatsAirCmd)
}

// StatsAirCmd defines 'air' subcommand
var StatsAirCmd = &cobra.Command{
	Use:   "air",
	Short: TRCLI("cli.stats.air.summary"),
	Long:  TRCLI(`cli.stats.air.description`),
}
