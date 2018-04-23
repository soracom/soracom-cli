package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OrdersListSubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var OrdersListSubscribersCmdLastEvaluatedKey string

// OrdersListSubscribersCmdOrderId holds value of 'order_id' option
var OrdersListSubscribersCmdOrderId string

// OrdersListSubscribersCmdLimit holds value of 'limit' option
var OrdersListSubscribersCmdLimit int64

func init() {
	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Serial number of the last subscriber in the previous page that is set to response header with X-Soracom-Next-Key."))

	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdOrderId, "order-id", "", TRAPI("order_id"))

	OrdersListSubscribersCmd.Flags().Int64Var(&OrdersListSubscribersCmdLimit, "limit", 0, TRAPI("Max number of subscribers in a response."))

	OrdersCmd.AddCommand(OrdersListSubscribersCmd)
}

// OrdersListSubscribersCmd defines 'list-subscribers' subcommand
var OrdersListSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: TRAPI("/orders/{order_id}/subscribers:get:summary"),
	Long:  TRAPI(`/orders/{order_id}/subscribers:get:description`),
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

		param, err := collectOrdersListSubscribersCmdParams(ac)
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

func collectOrdersListSubscribersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersListSubscribersCmd("/orders/{order_id}/subscribers"),
		query:  buildQueryForOrdersListSubscribersCmd(),
	}, nil
}

func buildPathForOrdersListSubscribersCmd(path string) string {

	path = strings.Replace(path, "{"+"order_id"+"}", OrdersListSubscribersCmdOrderId, -1)

	return path
}

func buildQueryForOrdersListSubscribersCmd() string {
	result := []string{}

	if OrdersListSubscribersCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", OrdersListSubscribersCmdLastEvaluatedKey))
	}

	if OrdersListSubscribersCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", OrdersListSubscribersCmdLimit))
	}

	return strings.Join(result, "&")
}
