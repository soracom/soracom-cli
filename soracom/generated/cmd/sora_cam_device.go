// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SoraCamCmd.AddCommand(SoraCamDeviceCmd)
}

// SoraCamDeviceCmd defines 'device' subcommand
var SoraCamDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: TRCLI("cli.sora-cam.device.summary"),
	Long:  TRCLI(`cli.sora-cam.device.description`),
}