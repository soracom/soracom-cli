// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsAirGroupsGetCmdGroupId holds value of 'group_id' option
var StatsAirGroupsGetCmdGroupId string

// StatsAirGroupsGetCmdPeriod holds value of 'period' option
var StatsAirGroupsGetCmdPeriod string

// StatsAirGroupsGetCmdFrom holds value of 'from' option
var StatsAirGroupsGetCmdFrom int64

// StatsAirGroupsGetCmdTo holds value of 'to' option
var StatsAirGroupsGetCmdTo int64

// StatsAirGroupsGetCmdOutputJSONL indicates to output with jsonl format
var StatsAirGroupsGetCmdOutputJSONL bool

func InitStatsAirGroupsGetCmd() {
	StatsAirGroupsGetCmd.Flags().StringVar(&StatsAirGroupsGetCmdGroupId, "group-id", "", TRAPI("Group ID"))

	StatsAirGroupsGetCmd.Flags().StringVar(&StatsAirGroupsGetCmdPeriod, "period", "", TRAPI("Unit of aggregation.- 'month': Monthly.  The 'from' and 'to' should be UNIX time (in seconds) from 3 months before the current time to the current time. The actual period of interest is not the time specified by 'from' and 'to'.  | Item | Description |  |-|-|  | Actual start time | 00:00:00 on the first day of the month, including the specified UNIX time (in seconds). |  | Actual end time | 24:00:00 of the last day of the month containing the specified UNIX time (in seconds). |- 'day': Daily  The 'from' and 'to' should be UNIX time (in seconds) from 7 days before the current time to the current time. The actual period of interest is not the time specified by 'from' and 'to'.  | Item | Description |  |-|-|  | Actual start time | 00:00:00 of the day including the specified UNIX time (in seconds). |  | Actual end time | 24:00:00 of the day including the specified UNIX time (in seconds). |"))

	StatsAirGroupsGetCmd.Flags().Int64Var(&StatsAirGroupsGetCmdFrom, "from", 0, TRAPI("Specify the start month/day of the period to be aggregated in UNIX time in seconds."))

	StatsAirGroupsGetCmd.Flags().Int64Var(&StatsAirGroupsGetCmdTo, "to", 0, TRAPI("Specify the end month/day of the period to be aggregated in UNIX time in seconds."))

	StatsAirGroupsGetCmd.Flags().BoolVar(&StatsAirGroupsGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	StatsAirGroupsGetCmd.RunE = StatsAirGroupsGetCmdRunE

	StatsAirGroupsCmd.AddCommand(StatsAirGroupsGetCmd)
}

// StatsAirGroupsGetCmd defines 'get' subcommand
var StatsAirGroupsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/air/groups/{group_id}:get:summary"),
	Long:  TRAPI(`/stats/air/groups/{group_id}:get:description`) + "\n\n" + createLinkToAPIReference("Stats", "getAirStatsOfGroup"),
}

func StatsAirGroupsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectStatsAirGroupsGetCmdParams(ac)
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
		if StatsAirGroupsGetCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectStatsAirGroupsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("group_id", "group-id", "path", parsedBody, StatsAirGroupsGetCmdGroupId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("period", "period", "query", parsedBody, StatsAirGroupsGetCmdPeriod)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("from", "from", "query", parsedBody, StatsAirGroupsGetCmdFrom)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("to", "to", "query", parsedBody, StatsAirGroupsGetCmdTo)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsAirGroupsGetCmd("/stats/air/groups/{group_id}"),
		query:  buildQueryForStatsAirGroupsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsAirGroupsGetCmd(path string) string {

	escapedGroupId := url.PathEscape(StatsAirGroupsGetCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	return path
}

func buildQueryForStatsAirGroupsGetCmd() url.Values {
	result := url.Values{}

	if StatsAirGroupsGetCmdPeriod != "" {
		result.Add("period", StatsAirGroupsGetCmdPeriod)
	}

	if StatsAirGroupsGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsAirGroupsGetCmdFrom))
	}

	if StatsAirGroupsGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsAirGroupsGetCmdTo))
	}

	return result
}
