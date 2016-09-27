package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CredentialsCmd)
}

// CredentialsCmd defines 'credentials' subcommand
var CredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: TR("credentials.cli.summary"),
	Long:  TR(`credentials.cli.description`),
}
