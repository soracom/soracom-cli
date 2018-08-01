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
var specifiedCoverageType string
var providedAPIKey string
var providedAPIToken string

func init() {
	RootCmd.PersistentFlags().StringVar(&specifiedProfileName, "profile", "", "Specify profile name")
	RootCmd.PersistentFlags().StringVar(&specifiedCoverageType, "coverage-type", "", "Specify coverage type, 'g' for Global, 'jp' for Japan")
	RootCmd.PersistentFlags().StringVar(&providedAPIKey, "api-key", "", "Specify API key otherwise soracom-cli performs authentication on behalf of you")
	RootCmd.PersistentFlags().StringVar(&providedAPIToken, "api-token", "", "Specify API token otherwise soracom-cli performs authentication on behalf of you")
}
