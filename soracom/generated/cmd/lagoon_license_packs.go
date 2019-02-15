package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	LagoonCmd.AddCommand(LagoonLicensePacksCmd)
}

// LagoonLicensePacksCmd defines 'license-packs' subcommand
var LagoonLicensePacksCmd = &cobra.Command{
	Use:   "license-packs",
	Short: TRCLI("cli.lagoon.license-packs.summary"),
	Long:  TRCLI(`cli.lagoon.license-packs.description`),
}
