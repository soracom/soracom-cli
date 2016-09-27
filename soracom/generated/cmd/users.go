package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(UsersCmd)
}

// UsersCmd defines 'users' subcommand
var UsersCmd = &cobra.Command{
	Use:   "users",
	Short: TR("users.cli.summary"),
	Long:  TR(`users.cli.description`),
}
