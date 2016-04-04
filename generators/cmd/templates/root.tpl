package cmd

import (
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use: "soracom",
  Short: "{{.Short}}",
  Long: `{{.Long}}`,
}

var specifiedProfileName string

func init() {
  RootCmd.PersistentFlags().StringVar(&specifiedProfileName, "profile", "", "Specify profile name")

	setDefaultEndpoint("{{.Endpoint}}")
}
