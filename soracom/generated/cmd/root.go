package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd defines 'soracom' command
var RootCmd = &cobra.Command{
	Use:   "soracom",
	Short: "soracom command",
	Long:  `A command line tool to invoke SORACOM API`,
}

var specifiedProfileName string

func init() {
	RootCmd.PersistentFlags().StringVar(&specifiedProfileName, "profile", "", "Specify profile name")

	setDefaultEndpoint("https://api.soracom.io")
}
