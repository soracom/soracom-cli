package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PaymentHistoryCmd)
}

// PaymentHistoryCmd defines 'payment-history' subcommand
var PaymentHistoryCmd = &cobra.Command{
	Use:   "payment-history",
	Short: TR("payment_history.cli.summary"),
	Long:  TR(`payment_history.cli.description`),
}
