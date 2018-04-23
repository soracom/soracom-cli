package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// BillsGetCmdYyyyMM holds value of 'yyyyMM' option
var BillsGetCmdYyyyMM string

func init() {
	BillsGetCmd.Flags().StringVar(&BillsGetCmdYyyyMM, "yyyy-mm", "", TRAPI("year and month"))

	BillsCmd.AddCommand(BillsGetCmd)
}

// BillsGetCmd defines 'get' subcommand
var BillsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/bills/{yyyyMM}:get:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}:get:description`),
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

		param, err := collectBillsGetCmdParams(ac)
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

func collectBillsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetCmd("/bills/{yyyyMM}"),
		query:  buildQueryForBillsGetCmd(),
	}, nil
}

func buildPathForBillsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"yyyyMM"+"}", BillsGetCmdYyyyMM, -1)

	return path
}

func buildQueryForBillsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
