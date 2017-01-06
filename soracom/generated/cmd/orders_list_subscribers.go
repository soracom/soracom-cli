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
	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TR("orders.list_ordered_subscriber.get.request.last_evaluated_key.description"))

	OrdersListSubscribersCmd.Flags().StringVar(&OrdersListSubscribersCmdOrderId, "order-id", "", TR("order_id"))

	OrdersListSubscribersCmd.Flags().Int64Var(&OrdersListSubscribersCmdLimit, "limit", 0, TR("orders.list_ordered_subscriber.get.request.limit.description"))

	OrdersCmd.AddCommand(OrdersListSubscribersCmd)
}

// OrdersListSubscribersCmd defines 'list-subscribers' subcommand
var OrdersListSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: TR("orders.list_ordered_subscriber.get.summary"),
	Long:  TR(`orders.list_ordered_subscriber.get.description`),
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

		param, err := collectOrdersListSubscribersCmdParams()
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

func collectOrdersListSubscribersCmdParams() (*apiParams, error) {

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
