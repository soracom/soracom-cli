// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SoraCamDevicesCmd.AddCommand(SoraCamDevicesAtomCamCmd)
}

// SoraCamDevicesAtomCamCmd defines 'atom-cam' subcommand
var SoraCamDevicesAtomCamCmd = &cobra.Command{
	Use:   "atom-cam",
	Short: TRCLI("cli.sora-cam.devices.atom-cam.summary"),
	Long:  TRCLI(`cli.sora-cam.devices.atom-cam.description`),
}
