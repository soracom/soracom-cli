// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	OrdersCmd.AddCommand(OrdersListCmd)
}

// OrdersListCmd defines 'list' subcommand
var OrdersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/orders:get:summary"),
	Long:  TRAPI(`/orders:get:description`),
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

		param, err := collectOrdersListCmdParams(ac)
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

func collectOrdersListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersListCmd("/orders"),
		query:  buildQueryForOrdersListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOrdersListCmd(path string) string {

	return path
}

func buildQueryForOrdersListCmd() url.Values {
	result := url.Values{}

	return result
}
