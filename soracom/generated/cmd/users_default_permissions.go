package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersDefaultPermissionsCmd)
}

// UsersDefaultPermissionsCmd defines 'default-permissions' subcommand
var UsersDefaultPermissionsCmd = &cobra.Command{
	Use:   "default-permissions",
	Short: TRCLI("cli.users.default-permissions.summary"),
	Long:  TRCLI(`cli.users.default-permissions.description`),
}
