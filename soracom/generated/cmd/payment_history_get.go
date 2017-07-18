package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// PaymentHistoryGetCmdPaymentTransactionId holds value of 'payment_transaction_id' option
var PaymentHistoryGetCmdPaymentTransactionId string

func init() {
	PaymentHistoryGetCmd.Flags().StringVar(&PaymentHistoryGetCmdPaymentTransactionId, "payment-transaction-id", "", TRAPI("payment_transaction_id"))

	PaymentHistoryCmd.AddCommand(PaymentHistoryGetCmd)
}

// PaymentHistoryGetCmd defines 'get' subcommand
var PaymentHistoryGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/payment_history/transactions/{payment_transaction_id}:get:summary"),
	Long:  TRAPI(`/payment_history/transactions/{payment_transaction_id}:get:description`),
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

		param, err := collectPaymentHistoryGetCmdParams()
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

func collectPaymentHistoryGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentHistoryGetCmd("/payment_history/transactions/{payment_transaction_id}"),
		query:  buildQueryForPaymentHistoryGetCmd(),
	}, nil
}

func buildPathForPaymentHistoryGetCmd(path string) string {

	path = strings.Replace(path, "{"+"payment_transaction_id"+"}", PaymentHistoryGetCmdPaymentTransactionId, -1)

	return path
}

func buildQueryForPaymentHistoryGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
