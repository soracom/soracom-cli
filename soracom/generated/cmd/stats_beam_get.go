package cmd

import (
	"os"
	"strings"

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

		param, err := collectStatsBeamGetCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectStatsBeamGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsBeamGetCmd("/stats/beam/subscribers/{imsi}"),
		query:  buildQueryForStatsBeamGetCmd(),
	}, nil
}

func buildPathForStatsBeamGetCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", StatsBeamGetCmdImsi, -1)

	return path
}

func buildQueryForStatsBeamGetCmd() string {
	result := []string{}

	if StatsBeamGetCmdPeriod != "" {
		result = append(result, sprintf("%s=%s", "period", StatsBeamGetCmdPeriod))
	}

	if StatsBeamGetCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", StatsBeamGetCmdFrom))
	}

	if StatsBeamGetCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", StatsBeamGetCmdTo))
	}

	return strings.Join(result, "&")
}
