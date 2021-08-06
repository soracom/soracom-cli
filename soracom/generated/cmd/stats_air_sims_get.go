// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsAirSimsGetCmdPeriod holds value of 'period' option
var StatsAirSimsGetCmdPeriod string

// StatsAirSimsGetCmdSimId holds value of 'simId' option
var StatsAirSimsGetCmdSimId string

// StatsAirSimsGetCmdFrom holds value of 'from' option
var StatsAirSimsGetCmdFrom int64

// StatsAirSimsGetCmdTo holds value of 'to' option
var StatsAirSimsGetCmdTo int64

func init() {
	StatsAirSimsGetCmd.Flags().StringVar(&StatsAirSimsGetCmdPeriod, "period", "", TRAPI("Units of aggregate data. For minutes, the interval is around 5 minutes."))

	StatsAirSimsGetCmd.Flags().StringVar(&StatsAirSimsGetCmdSimId, "sim-id", "", TRAPI("SIM ID"))

	StatsAirSimsGetCmd.Flags().Int64Var(&StatsAirSimsGetCmdFrom, "from", 0, TRAPI("Start time in unixtime for the aggregate data."))

	StatsAirSimsGetCmd.Flags().Int64Var(&StatsAirSimsGetCmdTo, "to", 0, TRAPI("End time in unixtime for the aggregate data."))
	StatsAirSimsCmd.AddCommand(StatsAirSimsGetCmd)
}

// StatsAirSimsGetCmd defines 'get' subcommand
var StatsAirSimsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/air/sims/{simId}:get:summary"),
	Long:  TRAPI(`/stats/air/sims/{simId}:get:description`),
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

		param, err := collectStatsAirSimsGetCmdParams(ac)
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

func collectStatsAirSimsGetCmdParams(ac *apiClient) (*apiParams, error) {
	if StatsAirSimsGetCmdPeriod == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "period")
	}

	if StatsAirSimsGetCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	if StatsAirSimsGetCmdFrom == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "from")
	}

	if StatsAirSimsGetCmdTo == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "to")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsAirSimsGetCmd("/stats/air/sims/{simId}"),
		query:  buildQueryForStatsAirSimsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsAirSimsGetCmd(path string) string {

	escapedSimId := url.PathEscape(StatsAirSimsGetCmdSimId)

	path = strReplace(path, "{"+"simId"+"}", escapedSimId, -1)

	return path
}

func buildQueryForStatsAirSimsGetCmd() url.Values {
	result := url.Values{}

	if StatsAirSimsGetCmdPeriod != "" {
		result.Add("period", StatsAirSimsGetCmdPeriod)
	}

	if StatsAirSimsGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsAirSimsGetCmdFrom))
	}

	if StatsAirSimsGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsAirSimsGetCmdTo))
	}

	return result
}
