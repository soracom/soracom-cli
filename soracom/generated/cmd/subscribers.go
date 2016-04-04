package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(SubscribersCmd)
}

var SubscribersCmd = &cobra.Command{
  Use: "subscribers",
  Short: TR("subscribers.cli.summary"),
  Long: TR(`subscribers.cli.description`),
}
