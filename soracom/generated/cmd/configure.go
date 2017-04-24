package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(ConfigureCmd)
	RootCmd.AddCommand(UnconfigureCmd)
}

// ConfigureCmd defines 'configure' subcommand
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: TR("configure.cli.summary"),
	Long:  TR("configure.cli.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pn := getSpecifiedProfileName()
		if pn == "" {
			pn = "default"
		}

		profile, err := collectProfileInfo(pn)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		err = saveProfile(pn, profile)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}

// UnconfigureCmd defines 'unconfigure' subcommand
var UnconfigureCmd = &cobra.Command{
	Use:   "unconfigure",
	Short: TR("unconfigure.cli.summary"),
	Long:  TR("unconfigure.cli.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pn := getSpecifiedProfileName()
		if pn == "" {
			pn = "default"
		}

		if confirmDeleteProfile(pn) {
			err := deleteProfile(pn)
			if err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
