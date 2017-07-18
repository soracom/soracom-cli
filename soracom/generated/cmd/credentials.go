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
	Short: TRCLI("cli.credentials.summary"),
	Long:  TRCLI(`cli.credentials.description`),
}
