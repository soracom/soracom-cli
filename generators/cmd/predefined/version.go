package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VersionCmd)
}

var version string

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: TR("version.cli.summary"),
	Long:  TR("version.cli.description"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SORACOM API client v%s\n", version)
	},
}
