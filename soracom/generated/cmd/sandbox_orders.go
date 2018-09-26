package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxCmd.AddCommand(SandboxOrdersCmd)
}

// SandboxOrdersCmd defines 'orders' subcommand
var SandboxOrdersCmd = &cobra.Command{
	Use:   "orders",
	Short: TRCLI("cli.sandbox.orders.summary"),
	Long:  TRCLI(`cli.sandbox.orders.description`),
}
