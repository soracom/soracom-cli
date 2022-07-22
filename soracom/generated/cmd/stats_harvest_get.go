// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsHarvestGetCmdImsi holds value of 'imsi' option
var StatsHarvestGetCmdImsi string

// StatsHarvestGetCmdPeriod holds value of 'period' option
var StatsHarvestGetCmdPeriod string

// StatsHarvestGetCmdFrom holds value of 'from' option
var StatsHarvestGetCmdFrom int64

// StatsHarvestGetCmdTo holds value of 'to' option
var StatsHarvestGetCmdTo int64

// StatsHarvestGetCmdOutputJSONL indicates to output with jsonl format
var StatsHarvestGetCmdOutputJSONL bool

func init() {
	StatsHarvestGetCmd.Flags().StringVar(&StatsHarvestGetCmdImsi, "imsi", "", TRAPI("imsi"))

	StatsHarvestGetCmd.Flags().StringVar(&StatsHarvestGetCmdPeriod, "period", "", TRAPI("Units of aggregate data. For minutes, the interval is around 5 minutes."))

	StatsHarvestGetCmd.Flags().Int64Var(&StatsHarvestGetCmdFrom, "from", 0, TRAPI("Start time in unixtime for the aggregate data."))

	StatsHarvestGetCmd.Flags().Int64Var(&StatsHarvestGetCmdTo, "to", 0, TRAPI("End time in unixtime for the aggregate data."))

	StatsHarvestGetCmd.Flags().BoolVar(&StatsHarvestGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	StatsHarvestCmd.AddCommand(StatsHarvestGetCmd)
}

// StatsHarvestGetCmd defines 'get' subcommand
var StatsHarvestGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/harvest/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/stats/harvest/subscribers/{imsi}:get:description`) + "\n\n" + createLinkToAPIReference("Stats", "getHarvestStats"),
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

		param, err := collectStatsHarvestGetCmdParams(ac)
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
			if StatsHarvestGetCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectStatsHarvestGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, StatsHarvestGetCmdImsi)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("period", "period", "query", parsedBody, StatsHarvestGetCmdPeriod)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("from", "from", "query", parsedBody, StatsHarvestGetCmdFrom)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("to", "to", "query", parsedBody, StatsHarvestGetCmdTo)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsHarvestGetCmd("/stats/harvest/subscribers/{imsi}"),
		query:  buildQueryForStatsHarvestGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsHarvestGetCmd(path string) string {

	escapedImsi := url.PathEscape(StatsHarvestGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForStatsHarvestGetCmd() url.Values {
	result := url.Values{}

	if StatsHarvestGetCmdPeriod != "" {
		result.Add("period", StatsHarvestGetCmdPeriod)
	}

	if StatsHarvestGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsHarvestGetCmdFrom))
	}

	if StatsHarvestGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsHarvestGetCmdTo))
	}

	return result
}
