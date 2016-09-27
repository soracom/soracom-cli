package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PaymentMethodsCmd)
}

// PaymentMethodsCmd defines 'payment-methods' subcommand
var PaymentMethodsCmd = &cobra.Command{
	Use:   "payment-methods",
	Short: TR("payment_methods.cli.summary"),
	Long:  TR(`payment_methods.cli.description`),
}
