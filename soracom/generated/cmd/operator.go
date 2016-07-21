package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(OperatorCmd)
}

var OperatorCmd = &cobra.Command{
	Use:   "operator",
	Short: TR("operator.cli.summary"),
	Long:  TR(`operator.cli.description`),
}
