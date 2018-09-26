package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// BillsExportLatestCmdExportMode holds value of 'export_mode' option
var BillsExportLatestCmdExportMode string

func init() {
	BillsExportLatestCmd.Flags().StringVar(&BillsExportLatestCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	BillsCmd.AddCommand(BillsExportLatestCmd)
}

// BillsExportLatestCmd defines 'export-latest' subcommand
var BillsExportLatestCmd = &cobra.Command{
	Use:   "export-latest",
	Short: TRAPI("/bills/latest/export:post:summary"),
	Long:  TRAPI(`/bills/latest/export:post:description`),
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

		param, err := collectBillsExportLatestCmdParams(ac)
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

func collectBillsExportLatestCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForBillsExportLatestCmd("/bills/latest/export"),
		query:  buildQueryForBillsExportLatestCmd(),
	}, nil
}

func buildPathForBillsExportLatestCmd(path string) string {

	return path
}

func buildQueryForBillsExportLatestCmd() string {
	result := []string{}

	if BillsExportLatestCmdExportMode != "" {
		result = append(result, sprintf("%s=%s", "export_mode", BillsExportLatestCmdExportMode))
	}

	return strings.Join(result, "&")
}
