package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxCmd.AddCommand(SandboxSubscribersCmd)
}

// SandboxSubscribersCmd defines 'subscribers' subcommand
var SandboxSubscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: TRCLI("cli.sandbox.subscribers.summary"),
	Long:  TRCLI(`cli.sandbox.subscribers.description`),
}
