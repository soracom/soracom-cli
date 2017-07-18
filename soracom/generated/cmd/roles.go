package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(RolesCmd)
}

// RolesCmd defines 'roles' subcommand
var RolesCmd = &cobra.Command{
	Use:   "roles",
	Short: TRCLI("cli.roles.summary"),
	Long:  TRCLI(`cli.roles.description`),
}
