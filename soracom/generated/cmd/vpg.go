package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(VpgCmd)
}

var VpgCmd = &cobra.Command{
  Use: "vpg",
  Short: TR("vpg.cli.summary"),
  Long: TR(`vpg.cli.description`),
}
