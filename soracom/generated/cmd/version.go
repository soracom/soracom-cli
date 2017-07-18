package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VersionCmd)
}

var version string

// VersionCmd defines 'version' subcommand
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: TRCLI("cli.version.summary"),
	Long:  TRCLI("cli.version.description"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SORACOM API client v%s\n", version)
	},
}
