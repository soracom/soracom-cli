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
	Short: TR("shipping_addresses.cli.summary"),
	Long:  TR(`shipping_addresses.cli.description`),
}
