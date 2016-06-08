package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(VpgCmd)
}

var VpgCmd = &cobra.Command{
  Use: "vpg",
  Short: TR("virtual_private_gateway.cli.summary"),
  Long: TR(`virtual_private_gateway.cli.description`),
}
