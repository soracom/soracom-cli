package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var BillsExportCmdYyyyMM string

func init() {
	BillsExportCmd.Flags().StringVar(&BillsExportCmdYyyyMM, "yyyy-mm", "", TR("yyyyMM"))

	BillsCmd.AddCommand(BillsExportCmd)
}

var BillsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TR("bills.export_billing.post.summary"),
	Long:  TR(`bills.export_billing.post.description`),
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

		param, err := collectBillsExportCmdParams()
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

func collectBillsExportCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForBillsExportCmd("/bills/{yyyyMM}/export"),
		query:  buildQueryForBillsExportCmd(),
	}, nil
}

func buildPathForBillsExportCmd(path string) string {

	path = strings.Replace(path, "{"+"yyyyMM"+"}", BillsExportCmdYyyyMM, -1)

	return path
}

func buildQueryForBillsExportCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
