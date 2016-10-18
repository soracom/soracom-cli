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
	Short: TR("payment.payer_information.cli.summary"),
	Long:  TR(`payment.payer_information.cli.description`),
}
