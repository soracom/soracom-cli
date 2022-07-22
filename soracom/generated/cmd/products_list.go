// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	ProductsCmd.AddCommand(ProductsListCmd)
}

// ProductsListCmd defines 'list' subcommand
var ProductsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/products:get:summary"),
	Long:  TRAPI(`/products:get:description`) + "\n\n" + createLinkToAPIReference("Order", "listProducts"),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			return fmt.Errorf("unexpected arguments passed => %v", args)
		}

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

		param, err := collectProductsListCmdParams(ac)
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

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectProductsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForProductsListCmd("/products"),
		query:  buildQueryForProductsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForProductsListCmd(path string) string {

	return path
}

func buildQueryForProductsListCmd() url.Values {
	result := url.Values{}

	return result
}
