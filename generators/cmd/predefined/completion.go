package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CompletionCmd)
	CompletionCmd.AddCommand(CompletionBashCmd)
	CompletionCmd.AddCommand(CompletionZshCmd)
}

// CompletionCmd defines 'completion' subcommand
var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: TRCLI("cli.completion.summary"),
	Long:  TRCLI("cli.completion.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		// For backward compatibility
		err := RootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	},
}

var CompletionBashCmd = &cobra.Command{
	Use:   "bash",
	Short: TRCLI("cli.completion.bash.summary"),
	Long:  TRCLI("cli.completion.bash.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := RootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	},
}

var CompletionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: TRCLI("cli.completion.zsh.summary"),
	Long:  TRCLI("cli.completion.zsh.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := RootCmd.GenZshCompletion(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	},
}
