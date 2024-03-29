package cmd

import "github.com/spf13/cobra"

var configureOverwrite bool

func init() {
	ConfigureCmd.Flags().BoolVar(&configureOverwrite, "overwrite", false, TRCLI("cli.configure.overwrite"))
	RootCmd.AddCommand(ConfigureCmd)
	RootCmd.AddCommand(UnconfigureCmd)
}

// ConfigureCmd defines 'configure' subcommand
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: TRCLI("cli.configure.summary"),
	Long:  TRCLI("cli.configure.description"),
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

		err = saveProfile(pn, profile, configureOverwrite)
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
	Short: TRCLI("cli.unconfigure.summary"),
	Long:  TRCLI("cli.unconfigure.description"),
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
