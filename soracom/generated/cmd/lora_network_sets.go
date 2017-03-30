package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LoraNetworkSetsCmd)
}

// LoraNetworkSetsCmd defines 'lora-network-sets' subcommand
var LoraNetworkSetsCmd = &cobra.Command{
	Use:   "lora-network-sets",
	Short: TR("lora_network_sets.cli.summary"),
	Long:  TR(`lora_network_sets.cli.description`),
}
