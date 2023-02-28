// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BillsSummariesGetBillItemsCmdOutputJSONL indicates to output with jsonl format
var BillsSummariesGetBillItemsCmdOutputJSONL bool

func init() {
	BillsSummariesGetBillItemsCmd.Flags().BoolVar(&BillsSummariesGetBillItemsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	BillsSummariesCmd.AddCommand(BillsSummariesGetBillItemsCmd)
}

// BillsSummariesGetBillItemsCmd defines 'get-bill-items' subcommand
var BillsSummariesGetBillItemsCmd = &cobra.Command{
	Use:   "get-bill-items",
	Short: TRAPI("/bills/summaries/bill_items:get:summary"),
	Long:  TRAPI(`/bills/summaries/bill_items:get:description`) + "\n\n" + createLinkToAPIReference("Billing", "getBillingSummaryOfBillItems"),
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

		param, err := collectBillsSummariesGetBillItemsCmdParams(ac)
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
			if BillsSummariesGetBillItemsCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectBillsSummariesGetBillItemsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsSummariesGetBillItemsCmd("/bills/summaries/bill_items"),
		query:  buildQueryForBillsSummariesGetBillItemsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsSummariesGetBillItemsCmd(path string) string {

	return path
}

func buildQueryForBillsSummariesGetBillItemsCmd() url.Values {
	result := url.Values{}

	return result
}
