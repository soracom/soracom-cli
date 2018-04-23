package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	PaymentMethodsCmd.AddCommand(PaymentMethodsReactivateCurrentCmd)
}

// PaymentMethodsReactivateCurrentCmd defines 'reactivate-current' subcommand
var PaymentMethodsReactivateCurrentCmd = &cobra.Command{
	Use:   "reactivate-current",
	Short: TRAPI("/payment_methods/current/activate:post:summary"),
	Long:  TRAPI(`/payment_methods/current/activate:post:description`),
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

		param, err := collectPaymentMethodsReactivateCurrentCmdParams(ac)
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

func collectPaymentMethodsReactivateCurrentCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForPaymentMethodsReactivateCurrentCmd("/payment_methods/current/activate"),
		query:  buildQueryForPaymentMethodsReactivateCurrentCmd(),
	}, nil
}

func buildPathForPaymentMethodsReactivateCurrentCmd(path string) string {

	return path
}

func buildQueryForPaymentMethodsReactivateCurrentCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
