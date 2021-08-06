// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsNapterAuditLogsGetCmdYearMonth holds value of 'year_month' option
var StatsNapterAuditLogsGetCmdYearMonth string

func init() {
	StatsNapterAuditLogsGetCmd.Flags().StringVar(&StatsNapterAuditLogsGetCmdYearMonth, "year-month", "", TRAPI("Year/Month in 'YYYYMM' format."))
	StatsNapterAuditLogsCmd.AddCommand(StatsNapterAuditLogsGetCmd)
}

// StatsNapterAuditLogsGetCmd defines 'get' subcommand
var StatsNapterAuditLogsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/napter/audit_logs:get:summary"),
	Long:  TRAPI(`/stats/napter/audit_logs:get:description`),
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

		param, err := collectStatsNapterAuditLogsGetCmdParams(ac)
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

func collectStatsNapterAuditLogsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsNapterAuditLogsGetCmd("/stats/napter/audit_logs"),
		query:  buildQueryForStatsNapterAuditLogsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsNapterAuditLogsGetCmd(path string) string {

	return path
}

func buildQueryForStatsNapterAuditLogsGetCmd() url.Values {
	result := url.Values{}

	if StatsNapterAuditLogsGetCmdYearMonth != "" {
		result.Add("year_month", StatsNapterAuditLogsGetCmdYearMonth)
	}

	return result
}
