package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PaymentStatementsCmd)
}

// PaymentStatementsCmd defines 'payment-statements' subcommand
var PaymentStatementsCmd = &cobra.Command{
	Use:   "payment-statements",
	Short: TRCLI("cli.payment-statements.summary"),
	Long:  TRCLI(`cli.payment-statements.description`),
}
