package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	PaymentMethodsCmd.AddCommand(PaymentMethodsWebpayCmd)
}

// PaymentMethodsWebpayCmd defines 'webpay' subcommand
var PaymentMethodsWebpayCmd = &cobra.Command{
	Use:   "webpay",
	Short: TR("payment_methods.webpay.cli.summary"),
	Long:  TR(`payment_methods.webpay.cli.description`),
}
