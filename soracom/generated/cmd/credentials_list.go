package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	CredentialsCmd.AddCommand(CredentialsListCmd)
}

var CredentialsListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("credentials.list_credentials.get.summary"),
	Long:  TR(`credentials.list_credentials.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectCredentialsListCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectCredentialsListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForCredentialsListCmd("/credentials"),
		query:  buildQueryForCredentialsListCmd(),
	}, nil
}

func buildPathForCredentialsListCmd(path string) string {

	return path
}

func buildQueryForCredentialsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
