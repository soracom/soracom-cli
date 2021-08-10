// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	PaymentMethodsCmd.AddCommand(PaymentMethodsGetCurrentCmd)
}

// PaymentMethodsGetCurrentCmd defines 'get-current' subcommand
var PaymentMethodsGetCurrentCmd = &cobra.Command{
	Use:   "get-current",
	Short: TRAPI("/payment_methods/current:get:summary"),
	Long:  TRAPI(`/payment_methods/current:get:description`),
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

		param, err := collectPaymentMethodsGetCurrentCmdParams(ac)
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

func collectPaymentMethodsGetCurrentCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentMethodsGetCurrentCmd("/payment_methods/current"),
		query:  buildQueryForPaymentMethodsGetCurrentCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPaymentMethodsGetCurrentCmd(path string) string {

	return path
}

func buildQueryForPaymentMethodsGetCurrentCmd() url.Values {
	result := url.Values{}

	return result
}
