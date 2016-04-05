package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  UsersCmd.AddCommand(UsersPasswordCmd)
}

var UsersPasswordCmd = &cobra.Command{
  Use: "password",
  Short: TR("users.password.cli.summary"),
  Long: TR(`users.password.cli.description`),
}