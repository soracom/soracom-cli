package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(FilesCmd)
}

// FilesCmd defines 'files' subcommand
var FilesCmd = &cobra.Command{
	Use:   "files",
	Short: TR("files.cli.summary"),
	Long:  TR(`files.cli.description`),
}
