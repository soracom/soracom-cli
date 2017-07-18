package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DevicesCmd)
}

// DevicesCmd defines 'devices' subcommand
var DevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: TRCLI("cli.devices.summary"),
	Long:  TRCLI(`cli.devices.description`),
}
