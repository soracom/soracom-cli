package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	OperatorCmd.AddCommand(OperatorAuthKeysCmd)
}

// OperatorAuthKeysCmd defines 'auth-keys' subcommand
var OperatorAuthKeysCmd = &cobra.Command{
	Use:   "auth-keys",
	Short: TRCLI("cli.operator.auth-keys.summary"),
	Long:  TRCLI(`cli.operator.auth-keys.description`),
}
