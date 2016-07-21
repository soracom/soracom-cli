package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(OrdersCmd)
}

var OrdersCmd = &cobra.Command{
	Use:   "orders",
	Short: TR("orders.cli.summary"),
	Long:  TR(`orders.cli.description`),
}
