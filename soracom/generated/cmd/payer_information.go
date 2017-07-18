package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PayerInformationCmd)
}

// PayerInformationCmd defines 'payer-information' subcommand
var PayerInformationCmd = &cobra.Command{
	Use:   "payer-information",
	Short: TRCLI("cli.payer-information.summary"),
	Long:  TRCLI(`cli.payer-information.description`),
}
