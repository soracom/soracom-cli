package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StatsHarvestGetCmdOperatorId holds value of 'operator_id' option
var StatsHarvestGetCmdOperatorId string

// StatsHarvestGetCmdYearMonth holds value of 'year_month' option
var StatsHarvestGetCmdYearMonth string

func init() {
	StatsHarvestGetCmd.Flags().StringVar(&StatsHarvestGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	StatsHarvestGetCmd.Flags().StringVar(&StatsHarvestGetCmdYearMonth, "year-month", "", TRAPI("Year/Month in 'YYYYMM' format."))

	StatsHarvestCmd.AddCommand(StatsHarvestGetCmd)
}

// StatsHarvestGetCmd defines 'get' subcommand
var StatsHarvestGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/harvest/operators/{operator_id}:get:summary"),
	Long:  TRAPI(`/stats/harvest/operators/{operator_id}:get:description`),
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

		param, err := collectStatsHarvestGetCmdParams(ac)
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

func collectStatsHarvestGetCmdParams(ac *apiClient) (*apiParams, error) {

	if StatsHarvestGetCmdOperatorId == "" {
		StatsHarvestGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsHarvestGetCmd("/stats/harvest/operators/{operator_id}"),
		query:  buildQueryForStatsHarvestGetCmd(),
	}, nil
}

func buildPathForStatsHarvestGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", StatsHarvestGetCmdOperatorId, -1)

	return path
}

func buildQueryForStatsHarvestGetCmd() string {
	result := []string{}

	if StatsHarvestGetCmdYearMonth != "" {
		result = append(result, sprintf("%s=%s", "year_month", StatsHarvestGetCmdYearMonth))
	}

	return strings.Join(result, "&")
}
