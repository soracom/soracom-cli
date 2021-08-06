// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	PaymentStatementsCmd.AddCommand(PaymentStatementsListCmd)
}

// PaymentStatementsListCmd defines 'list' subcommand
var PaymentStatementsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/payment_statements:get:summary"),
	Long:  TRAPI(`/payment_statements:get:description`),
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

		param, err := collectPaymentStatementsListCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectPaymentStatementsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentStatementsListCmd("/payment_statements"),
		query:  buildQueryForPaymentStatementsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPaymentStatementsListCmd(path string) string {

	return path
}

func buildQueryForPaymentStatementsListCmd() url.Values {
	result := url.Values{}

	return result
}
