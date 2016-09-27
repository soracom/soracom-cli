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
	Short: TR("stats.beam.cli.summary"),
	Long:  TR(`stats.beam.cli.description`),
}
