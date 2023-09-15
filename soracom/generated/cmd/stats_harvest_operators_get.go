// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsHarvestOperatorsGetCmdOperatorId holds value of 'operator_id' option
var StatsHarvestOperatorsGetCmdOperatorId string

// StatsHarvestOperatorsGetCmdYearMonth holds value of 'year_month' option
var StatsHarvestOperatorsGetCmdYearMonth string

func InitStatsHarvestOperatorsGetCmd() {
	StatsHarvestOperatorsGetCmd.Flags().StringVar(&StatsHarvestOperatorsGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	StatsHarvestOperatorsGetCmd.Flags().StringVar(&StatsHarvestOperatorsGetCmdYearMonth, "year-month", "", TRAPI("Year/Month in 'YYYYMM' format."))

	StatsHarvestOperatorsGetCmd.RunE = StatsHarvestOperatorsGetCmdRunE

	StatsHarvestOperatorsCmd.AddCommand(StatsHarvestOperatorsGetCmd)
}

// StatsHarvestOperatorsGetCmd defines 'get' subcommand
var StatsHarvestOperatorsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/harvest/operators/{operator_id}:get:summary"),
	Long:  TRAPI(`/stats/harvest/operators/{operator_id}:get:description`) + "\n\n" + createLinkToAPIReference("Stats", "getHarvestExportedDataStats"),
}

func StatsHarvestOperatorsGetCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectStatsHarvestOperatorsGetCmdParams(ac)
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

func collectStatsHarvestOperatorsGetCmdParams(ac *apiClient) (*apiParams, error) {
	if StatsHarvestOperatorsGetCmdOperatorId == "" {
		StatsHarvestOperatorsGetCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsHarvestOperatorsGetCmd("/stats/harvest/operators/{operator_id}"),
		query:  buildQueryForStatsHarvestOperatorsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsHarvestOperatorsGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(StatsHarvestOperatorsGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForStatsHarvestOperatorsGetCmd() url.Values {
	result := url.Values{}

	if StatsHarvestOperatorsGetCmdYearMonth != "" {
		result.Add("year_month", StatsHarvestOperatorsGetCmdYearMonth)
	}

	return result
}
