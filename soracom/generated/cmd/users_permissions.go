package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersPermissionsCmd)
}

// UsersPermissionsCmd defines 'permissions' subcommand
var UsersPermissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: TR("users.permissions.cli.summary"),
	Long:  TR(`users.permissions.cli.description`),
}
