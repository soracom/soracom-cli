package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(BillsCmd)
}

// BillsCmd defines 'bills' subcommand
var BillsCmd = &cobra.Command{
	Use:   "bills",
	Short: TR("bills.cli.summary"),
	Long:  TR(`bills.cli.description`),
}
