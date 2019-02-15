package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	LagoonCmd.AddCommand(LagoonUsersCmd)
}

// LagoonUsersCmd defines 'users' subcommand
var LagoonUsersCmd = &cobra.Command{
	Use:   "users",
	Short: TRCLI("cli.lagoon.users.summary"),
	Long:  TRCLI(`cli.lagoon.users.description`),
}
