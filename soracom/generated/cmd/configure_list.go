package cmd

import (
	"os"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	ConfigureCmd.AddCommand(ConfigureListCmd)
}

// ConfigureListCmd defines the 'configure list' subcommand. It prints the names
// of the configured profiles (the basenames of ~/.soracom/*.json) as a JSON
// array, WITHOUT reading or revealing their contents. This gives an agent a
// safe way to discover which profiles exist — and whether any exist at all —
// without touching the plaintext credentials stored inside each profile file.
var ConfigureListCmd = &cobra.Command{
	Use:   "list",
	Short: TRCLI("cli.configure.list.summary"),
	Long:  TRCLI("cli.configure.list.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		profiles, err := enumerateProfiles()
		if err != nil {
			// No profile directory yet means no profiles have been configured;
			// report that as an empty list rather than an error.
			if os.IsNotExist(err) {
				profiles = []string{}
			} else {
				return err
			}
		}
		if profiles == nil {
			profiles = []string{}
		}
		sort.Strings(profiles)

		return prettyPrintObjectAsJSON(profiles, os.Stdout)
	},
}
