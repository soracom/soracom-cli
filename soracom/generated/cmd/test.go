package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(TestCmd)
	TestCmd.AddCommand(Test500Cmd)
}

// TestCmd defines 'test' subcommand
var TestCmd = &cobra.Command{
	Use:    "test",
	Short:  TRCLI("cli.test.summary"),
	Long:   TRCLI("cli.test.description"),
	Hidden: true,
}

// Test500Cmd defines 'test 500' subcommand
var Test500Cmd = &cobra.Command{
	Use:   "500",
	Short: TRCLI("cli.test._500.summary"),
	Long:  TRCLI("cli.test._500.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param := &apiParams{
			method:      "POST",
			path:        "500",
			contentType: "application/json",
			body:        `{"expect":"500 Internal Server Error"}`,
		}

		_, err := ac.callAPI(param)

		return err
	},
}
