package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	PayerInformationCmd.AddCommand(PayerInformationGetCmd)
}

// PayerInformationGetCmd defines 'get' subcommand
var PayerInformationGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/payment_statements/payer_information:get:summary"),
	Long:  TRAPI(`/payment_statements/payer_information:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
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

		param, err := collectPayerInformationGetCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectPayerInformationGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPayerInformationGetCmd("/payment_statements/payer_information"),
		query:  buildQueryForPayerInformationGetCmd(),
	}, nil
}

func buildPathForPayerInformationGetCmd(path string) string {

	return path
}

func buildQueryForPayerInformationGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
