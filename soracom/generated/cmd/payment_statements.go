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
	Short: TR("payment_statements.cli.summary"),
	Long:  TR(`payment_statements.cli.description`),
}
