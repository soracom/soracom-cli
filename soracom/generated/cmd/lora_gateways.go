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
	Short: TR("lora_gateways.cli.summary"),
	Long:  TR(`lora_gateways.cli.description`),
}
