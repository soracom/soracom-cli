package cmd

import (
	"os"
	"strings"

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

func collectOrdersCancelCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "PUT",
		path:   buildPathForOrdersCancelCmd("/orders/{order_id}/cancel"),
		query:  buildQueryForOrdersCancelCmd(),
	}, nil
}

func buildPathForOrdersCancelCmd(path string) string {

	path = strings.Replace(path, "{"+"order_id"+"}", OrdersCancelCmdOrderId, -1)

	return path
}

func buildQueryForOrdersCancelCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
