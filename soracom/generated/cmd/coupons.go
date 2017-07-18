package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CouponsCmd)
}

// CouponsCmd defines 'coupons' subcommand
var CouponsCmd = &cobra.Command{
	Use:   "coupons",
	Short: TRCLI("cli.coupons.summary"),
	Long:  TRCLI(`cli.coupons.description`),
}
