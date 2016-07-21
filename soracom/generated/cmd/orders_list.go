package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	OrdersCmd.AddCommand(OrdersListCmd)
}

var OrdersListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("orders.list_orders.get.summary"),
	Long:  TR(`orders.list_orders.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectOrdersListCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectOrdersListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersListCmd("/orders"),
		query:  buildQueryForOrdersListCmd(),
	}, nil
}

func buildPathForOrdersListCmd(path string) string {

	return path
}

func buildQueryForOrdersListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
