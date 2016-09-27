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
	Short: TR("coupons.cli.summary"),
	Long:  TR(`coupons.cli.description`),
}
