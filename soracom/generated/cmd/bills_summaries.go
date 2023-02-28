// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	BillsCmd.AddCommand(BillsSummariesCmd)
}

// BillsSummariesCmd defines 'summaries' subcommand
var BillsSummariesCmd = &cobra.Command{
	Use:   "summaries",
	Short: TRCLI("cli.bills.summaries.summary"),
	Long:  TRCLI(`cli.bills.summaries.description`),
}
