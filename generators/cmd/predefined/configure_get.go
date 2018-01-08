package cmd

import "github.com/spf13/cobra"

func init() {
	ConfigureCmd.AddCommand(ConfigureGetCmd)
}

// ConfigureGetCmd defineds 'get' subcommand
var ConfigureGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRCLI("cli.configure.get.summary"),
	Long:  TRCLI("cli.configure.get.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pn := getSpecifiedProfileName()
		if pn == "" {
			pn = "default"
		}

		p, err := loadProfile(pn)
		if err != nil {
			return err
		}

		return prettyPrintObjectAsJSON(p)
	},
}
