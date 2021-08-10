// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OrdersCancelCmdOrderId holds value of 'order_id' option
var OrdersCancelCmdOrderId string

func init() {
	OrdersCancelCmd.Flags().StringVar(&OrdersCancelCmdOrderId, "order-id", "", TRAPI("order_id"))
	OrdersCmd.AddCommand(OrdersCancelCmd)
}

// OrdersCancelCmd defines 'cancel' subcommand
var OrdersCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: TRAPI("/orders/{order_id}/cancel:put:summary"),
	Long:  TRAPI(`/orders/{order_id}/cancel:put:description`),
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

		param, err := collectOrdersCancelCmdParams(ac)
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

func collectOrdersCancelCmdParams(ac *apiClient) (*apiParams, error) {
	if OrdersCancelCmdOrderId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "order-id")
	}

	return &apiParams{
		method: "PUT",
		path:   buildPathForOrdersCancelCmd("/orders/{order_id}/cancel"),
		query:  buildQueryForOrdersCancelCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOrdersCancelCmd(path string) string {

	escapedOrderId := url.PathEscape(OrdersCancelCmdOrderId)

	path = strReplace(path, "{"+"order_id"+"}", escapedOrderId, -1)

	return path
}

func buildQueryForOrdersCancelCmd() url.Values {
	result := url.Values{}

	return result
}
