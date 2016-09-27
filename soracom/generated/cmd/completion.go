package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CompletionCmd)
}

// CompletionCmd defines 'completion' subcommand
var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: TR("completion.cli.summary"),
	Long:  TR("completion.cli.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := RootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	},
}
