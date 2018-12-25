package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	StatsCmd.AddCommand(StatsHarvestCmd)
}

// StatsHarvestCmd defines 'harvest' subcommand
var StatsHarvestCmd = &cobra.Command{
	Use:   "harvest",
	Short: TRCLI("cli.stats.harvest.summary"),
	Long:  TRCLI(`cli.stats.harvest.description`),
}
