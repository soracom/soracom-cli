package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DataCmd)
}

// DataCmd defines 'data' subcommand
var DataCmd = &cobra.Command{
	Use:   "data",
	Short: TR("data.cli.summary"),
	Long:  TR(`data.cli.description`),
}
