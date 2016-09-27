package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersAuthKeysCmd)
}

// UsersAuthKeysCmd defines 'auth-keys' subcommand
var UsersAuthKeysCmd = &cobra.Command{
	Use:   "auth-keys",
	Short: TR("users.auth_keys.cli.summary"),
	Long:  TR(`users.auth_keys.cli.description`),
}
