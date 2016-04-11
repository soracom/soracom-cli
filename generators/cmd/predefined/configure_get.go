package cmd

import "github.com/spf13/cobra"

func init() {
	ConfigureCmd.AddCommand(ConfigureGetCmd)
}

var ConfigureGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("configure.get.cli.summary"),
	Long:  TR("configure.get.cli.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pn := getSpecifiedProfileName()
		if pn == "" {
			pn = "default"
		}

		p, err := loadProfile(pn)
		if err != nil {
			return err
		}

		prettyPrintObjectAsJSON(p)
		return nil
	},
}
