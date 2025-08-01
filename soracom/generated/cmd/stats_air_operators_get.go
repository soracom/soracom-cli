// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsAirOperatorsGetCmdOperatorId holds value of 'operator_id' option
var StatsAirOperatorsGetCmdOperatorId string

// StatsAirOperatorsGetCmdPeriod holds value of 'period' option
var StatsAirOperatorsGetCmdPeriod string

// StatsAirOperatorsGetCmdFrom holds value of 'from' option
var StatsAirOperatorsGetCmdFrom int64

// StatsAirOperatorsGetCmdTo holds value of 'to' option
var StatsAirOperatorsGetCmdTo int64

// StatsAirOperatorsGetCmdOutputJSONL indicates to output with jsonl format
var StatsAirOperatorsGetCmdOutputJSONL bool

func InitStatsAirOperatorsGetCmd() {
	StatsAirOperatorsGetCmd.Flags().StringVar(&StatsAirOperatorsGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	StatsAirOperatorsGetCmd.Flags().StringVar(&StatsAirOperatorsGetCmdPeriod, "period", "", TRAPI("Unit of aggregation.- 'month': Monthly.  The 'from' and 'to' should be UNIX time (in seconds) from 3 months before the current time to the current time. The actual period of interest is not the time specified by 'from' and 'to'.  | Item | Description |  |-|-|  | Actual start time | 00:00:00 on the first day of the month, including the specified UNIX time (in seconds). |  | Actual end time | 24:00:00 of the last day of the month containing the specified UNIX time (in seconds). |- 'day': Daily  The 'from' and 'to' should be UNIX time (in seconds) from 7 days before the current time to the current time. The actual period of interest is not the time specified by 'from' and 'to'.  | Item | Description |  |-|-|  | Actual start time | 00:00:00 of the day including the specified UNIX time (in seconds). |  | Actual end time | 24:00:00 of the day including the specified UNIX time (in seconds). |"))

	StatsAirOperatorsGetCmd.Flags().Int64Var(&StatsAirOperatorsGetCmdFrom, "from", 0, TRAPI("Specify the start month/day of the period to be aggregated in UNIX time in seconds."))

	StatsAirOperatorsGetCmd.Flags().Int64Var(&StatsAirOperatorsGetCmdTo, "to", 0, TRAPI("Specify the end month/day of the period to be aggregated in UNIX time in seconds."))

	StatsAirOperatorsGetCmd.Flags().BoolVar(&StatsAirOperatorsGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	StatsAirOperatorsGetCmd.RunE = StatsAirOperatorsGetCmdRunE

	StatsAirOperatorsCmd.AddCommand(StatsAirOperatorsGetCmd)
}

// StatsAirOperatorsGetCmd defines 'get' subcommand
var StatsAirOperatorsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/air/operators/{operator_id}:get:summary"),
	Long:  TRAPI(`/stats/air/operators/{operator_id}:get:description`) + "\n\n" + createLinkToAPIReference("Stats", "getAirStatsOfOperator"),
}

func StatsAirOperatorsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectStatsAirOperatorsGetCmdParams(ac)
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
		if StatsAirOperatorsGetCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectStatsAirOperatorsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if StatsAirOperatorsGetCmdOperatorId == "" {
		StatsAirOperatorsGetCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("period", "period", "query", parsedBody, StatsAirOperatorsGetCmdPeriod)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("from", "from", "query", parsedBody, StatsAirOperatorsGetCmdFrom)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("to", "to", "query", parsedBody, StatsAirOperatorsGetCmdTo)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsAirOperatorsGetCmd("/stats/air/operators/{operator_id}"),
		query:  buildQueryForStatsAirOperatorsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsAirOperatorsGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(StatsAirOperatorsGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForStatsAirOperatorsGetCmd() url.Values {
	result := url.Values{}

	if StatsAirOperatorsGetCmdPeriod != "" {
		result.Add("period", StatsAirOperatorsGetCmdPeriod)
	}

	if StatsAirOperatorsGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsAirOperatorsGetCmdFrom))
	}

	if StatsAirOperatorsGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsAirOperatorsGetCmdTo))
	}

	return result
}
