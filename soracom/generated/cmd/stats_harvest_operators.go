// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	StatsHarvestCmd.AddCommand(StatsHarvestOperatorsCmd)
}

// StatsHarvestOperatorsCmd defines 'operators' subcommand
var StatsHarvestOperatorsCmd = &cobra.Command{
	Use:   "operators",
	Short: TRCLI("cli.stats.harvest.operators.summary"),
	Long:  TRCLI(`cli.stats.harvest.operators.description`),
}
