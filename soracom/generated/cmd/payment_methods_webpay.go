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
	Short: TRCLI("cli.payment-methods.webpay.summary"),
	Long:  TRCLI(`cli.payment-methods.webpay.description`),
}
