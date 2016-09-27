package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StatsAirExportCmdOperatorId holds value of 'operator_id' option
var StatsAirExportCmdOperatorId string

// StatsAirExportCmdPeriod holds value of 'period' option
var StatsAirExportCmdPeriod string

// StatsAirExportCmdFrom holds value of 'from' option
var StatsAirExportCmdFrom int64

// StatsAirExportCmdTo holds value of 'to' option
var StatsAirExportCmdTo int64

// StatsAirExportCmdBody holds contents of request body to be sent
var StatsAirExportCmdBody string

func init() {
	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdOperatorId, "operator-id", "", TR("stats.export_air_stats.post.parameters.operator_id.description"))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdPeriod, "period", "", TR(""))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdFrom, "from", 0, TR(""))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdTo, "to", 0, TR(""))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	StatsAirCmd.AddCommand(StatsAirExportCmd)
}

// StatsAirExportCmd defines 'export' subcommand
var StatsAirExportCmd = &cobra.Command{
	Use:   "export",
	Short: TR("stats.export_air_stats.post.summary"),
	Long:  TR(`stats.export_air_stats.post.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectStatsAirExportCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectStatsAirExportCmdParams() (*apiParams, error) {

	body, err := buildBodyForStatsAirExportCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForStatsAirExportCmd("/stats/air/operators/{operator_id}/export"),
		query:       buildQueryForStatsAirExportCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForStatsAirExportCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", StatsAirExportCmdOperatorId, -1)

	return path
}

func buildQueryForStatsAirExportCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForStatsAirExportCmd() (string, error) {
	if StatsAirExportCmdBody != "" {
		if strings.HasPrefix(StatsAirExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsAirExportCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if StatsAirExportCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return StatsAirExportCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if StatsAirExportCmdPeriod != "" {
		result["period"] = StatsAirExportCmdPeriod
	}

	if StatsAirExportCmdFrom != 0 {
		result["from"] = StatsAirExportCmdFrom
	}

	if StatsAirExportCmdTo != 0 {
		result["to"] = StatsAirExportCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
