package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(StatsCmd)
}

var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: TR("stats.cli.summary"),
	Long:  TR(`stats.cli.description`),
}
