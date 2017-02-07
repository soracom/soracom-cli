package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LoraDevicesCmd)
}

// LoraDevicesCmd defines 'lora-devices' subcommand
var LoraDevicesCmd = &cobra.Command{
	Use:   "lora-devices",
	Short: TR("lora_devices.cli.summary"),
	Long:  TR(`lora_devices.cli.description`),
}
