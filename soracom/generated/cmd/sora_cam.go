// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SoraCamCmd)
}

// SoraCamCmd defines 'sora-cam' subcommand
var SoraCamCmd = &cobra.Command{
	Use:   "sora-cam",
	Short: TRCLI("cli.sora-cam.summary"),
	Long:  TRCLI(`cli.sora-cam.description`),
}
