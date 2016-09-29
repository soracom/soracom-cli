package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LogsCmd)
}

// LogsCmd defines 'logs' subcommand
var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: TR("logs.cli.summary"),
	Long:  TR(`logs.cli.description`),
}
