package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	OperatorCmd.AddCommand(OperatorAuthKeysCmd)
}

var OperatorAuthKeysCmd = &cobra.Command{
	Use:   "auth-keys",
	Short: TR("operator.auth_keys.cli.summary"),
	Long:  TR(`operator.auth_keys.cli.description`),
}
