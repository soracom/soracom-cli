package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LoraGatewaysCmd)
}

// LoraGatewaysCmd defines 'lora-gateways' subcommand
var LoraGatewaysCmd = &cobra.Command{
	Use:   "lora-gateways",
	Short: TRCLI("cli.lora-gateways.summary"),
	Long:  TRCLI(`cli.lora-gateways.description`),
}
