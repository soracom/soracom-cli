package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	StatsCmd.AddCommand(StatsBeamCmd)
}

// StatsBeamCmd defines 'beam' subcommand
var StatsBeamCmd = &cobra.Command{
	Use:   "beam",
	Short: TRCLI("cli.stats.beam.summary"),
	Long:  TRCLI(`cli.stats.beam.description`),
}
