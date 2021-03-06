// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsBeamGetCmdImsi holds value of 'imsi' option
var StatsBeamGetCmdImsi string

// StatsBeamGetCmdPeriod holds value of 'period' option
var StatsBeamGetCmdPeriod string

// StatsBeamGetCmdFrom holds value of 'from' option
var StatsBeamGetCmdFrom int64

// StatsBeamGetCmdTo holds value of 'to' option
var StatsBeamGetCmdTo int64

func init() {
	StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdImsi, "imsi", "", TRAPI("imsi"))

	StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdPeriod, "period", "", TRAPI("Units of aggregate data. For minutes, the interval is around 5 minutes."))

	StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdFrom, "from", 0, TRAPI("Start time in unixtime for the aggregate data."))

	StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdTo, "to", 0, TRAPI("End time in unixtime for the aggregate data."))
	StatsBeamCmd.AddCommand(StatsBeamGetCmd)
}

// StatsBeamGetCmd defines 'get' subcommand
var StatsBeamGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/beam/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/stats/beam/subscribers/{imsi}:get:description`),
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

		param, err := collectStatsBeamGetCmdParams(ac)
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
	},
}

func collectStatsBeamGetCmdParams(ac *apiClient) (*apiParams, error) {
	if StatsBeamGetCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	if StatsBeamGetCmdPeriod == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "period")
	}

	if StatsBeamGetCmdFrom == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "from")
	}

	if StatsBeamGetCmdTo == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "to")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsBeamGetCmd("/stats/beam/subscribers/{imsi}"),
		query:  buildQueryForStatsBeamGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsBeamGetCmd(path string) string {

	escapedImsi := url.PathEscape(StatsBeamGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForStatsBeamGetCmd() url.Values {
	result := url.Values{}

	if StatsBeamGetCmdPeriod != "" {
		result.Add("period", StatsBeamGetCmdPeriod)
	}

	if StatsBeamGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsBeamGetCmdFrom))
	}

	if StatsBeamGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsBeamGetCmdTo))
	}

	return result
}
