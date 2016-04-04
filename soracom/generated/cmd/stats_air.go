package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  StatsCmd.AddCommand(StatsAirCmd)
}

var StatsAirCmd = &cobra.Command{
  Use: "air",
  Short: TR("stats.air.cli.summary"),
  Long: TR(`stats.air.cli.description`),
}
