package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OrdersConfirmCmdOrderId holds value of 'order_id' option
var OrdersConfirmCmdOrderId string

func init() {
	OrdersConfirmCmd.Flags().StringVar(&OrdersConfirmCmdOrderId, "order-id", "", TRAPI("order_id"))

	OrdersCmd.AddCommand(OrdersConfirmCmd)
}

// OrdersConfirmCmd defines 'confirm' subcommand
var OrdersConfirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: TRAPI("/orders/{order_id}/confirm:put:summary"),
	Long:  TRAPI(`/orders/{order_id}/confirm:put:description`),
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

		param, err := collectOrdersConfirmCmdParams(ac)
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

func collectOrdersConfirmCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "PUT",
		path:   buildPathForOrdersConfirmCmd("/orders/{order_id}/confirm"),
		query:  buildQueryForOrdersConfirmCmd(),
	}, nil
}

func buildPathForOrdersConfirmCmd(path string) string {

	path = strings.Replace(path, "{"+"order_id"+"}", OrdersConfirmCmdOrderId, -1)

	return path
}

func buildQueryForOrdersConfirmCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
