package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxCmd.AddCommand(SandboxCouponsCmd)
}

// SandboxCouponsCmd defines 'coupons' subcommand
var SandboxCouponsCmd = &cobra.Command{
	Use:   "coupons",
	Short: TRCLI("cli.sandbox.coupons.summary"),
	Long:  TRCLI(`cli.sandbox.coupons.description`),
}
