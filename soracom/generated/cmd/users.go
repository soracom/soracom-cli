package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(UsersCmd)
}

var UsersCmd = &cobra.Command{
  Use: "users",
  Short: TR("users.cli.summary"),
  Long: TR(`users.cli.description`),
}
