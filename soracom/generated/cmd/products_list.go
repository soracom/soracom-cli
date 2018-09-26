package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	ProductsCmd.AddCommand(ProductsListCmd)
}

// ProductsListCmd defines 'list' subcommand
var ProductsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/products:get:summary"),
	Long:  TRAPI(`/products:get:description`),
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

		param, err := collectProductsListCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectProductsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForProductsListCmd("/products"),
		query:  buildQueryForProductsListCmd(),
	}, nil
}

func buildPathForProductsListCmd(path string) string {

	return path
}

func buildQueryForProductsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
