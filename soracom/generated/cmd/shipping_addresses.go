// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ShippingAddressesCmd)
}

// ShippingAddressesCmd defines 'shipping-addresses' subcommand
var ShippingAddressesCmd = &cobra.Command{
	Use:   "shipping-addresses",
	Short: TRCLI("cli.shipping-addresses.summary"),
	Long:  TRCLI(`cli.shipping-addresses.description`),
}
