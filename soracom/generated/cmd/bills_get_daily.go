// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BillsGetDailyCmdYyyyMM holds value of 'yyyyMM' option
var BillsGetDailyCmdYyyyMM string

func InitBillsGetDailyCmd() {
	BillsGetDailyCmd.Flags().StringVar(&BillsGetDailyCmdYyyyMM, "yyyy-mm", "", TRAPI("Target year and month"))

	BillsGetDailyCmd.RunE = BillsGetDailyCmdRunE

	BillsCmd.AddCommand(BillsGetDailyCmd)
}

// BillsGetDailyCmd defines 'get-daily' subcommand
var BillsGetDailyCmd = &cobra.Command{
	Use:   "get-daily",
	Short: TRAPI("/bills/{yyyyMM}/daily:get:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}/daily:get:description`) + "\n\n" + createLinkToAPIReference("Billing", "getBillingPerDay"),
}

func BillsGetDailyCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectBillsGetDailyCmdParams(ac)
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
}

func collectBillsGetDailyCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("yyyyMM", "yyyy-mm", "path", parsedBody, BillsGetDailyCmdYyyyMM)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetDailyCmd("/bills/{yyyyMM}/daily"),
		query:  buildQueryForBillsGetDailyCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsGetDailyCmd(path string) string {

	escapedYyyyMM := url.PathEscape(BillsGetDailyCmdYyyyMM)

	path = strReplace(path, "{"+"yyyyMM"+"}", escapedYyyyMM, -1)

	return path
}

func buildQueryForBillsGetDailyCmd() url.Values {
	result := url.Values{}

	return result
}
