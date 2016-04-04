package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(PaymentHistoryCmd)
}

var PaymentHistoryCmd = &cobra.Command{
  Use: "payment-history",
  Short: TR("payment_history.cli.summary"),
  Long: TR(`payment_history.cli.description`),
}
