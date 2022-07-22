// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BillsExportCmdExportMode holds value of 'export_mode' option
var BillsExportCmdExportMode string

// BillsExportCmdYyyyMM holds value of 'yyyyMM' option
var BillsExportCmdYyyyMM string

func init() {
	BillsExportCmd.Flags().StringVar(&BillsExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	BillsExportCmd.Flags().StringVar(&BillsExportCmdYyyyMM, "yyyy-mm", "", TRAPI("yyyyMM"))
	BillsCmd.AddCommand(BillsExportCmd)
}

// BillsExportCmd defines 'export' subcommand
var BillsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/bills/{yyyyMM}/export:post:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}/export:post:description`) + "\n\n" + createLinkToAPIReference("Billing", "exportBilling"),
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

		param, err := collectBillsExportCmdParams(ac)
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
		rawOutput = true

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectBillsExportCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("yyyyMM", "yyyy-mm", "path", parsedBody, BillsExportCmdYyyyMM)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForBillsExportCmd("/bills/{yyyyMM}/export"),
		query:  buildQueryForBillsExportCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBillsExportCmd(path string) string {

	escapedYyyyMM := url.PathEscape(BillsExportCmdYyyyMM)

	path = strReplace(path, "{"+"yyyyMM"+"}", escapedYyyyMM, -1)

	return path
}

func buildQueryForBillsExportCmd() url.Values {
	result := url.Values{}

	if BillsExportCmdExportMode != "" {
		result.Add("export_mode", BillsExportCmdExportMode)
	}

	return result
}
