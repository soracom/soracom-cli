package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(OperatorCmd)
}

// OperatorCmd defines 'operator' subcommand
var OperatorCmd = &cobra.Command{
	Use:   "operator",
	Short: TRCLI("cli.operator.summary"),
	Long:  TRCLI(`cli.operator.description`),
}
