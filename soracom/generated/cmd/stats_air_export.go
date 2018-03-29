package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StatsAirExportCmdExportMode holds value of 'export_mode' option
var StatsAirExportCmdExportMode string

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
	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdPeriod, "period", "", TRAPI(""))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdFrom, "from", 0, TRAPI(""))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdTo, "to", 0, TRAPI(""))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	StatsAirCmd.AddCommand(StatsAirExportCmd)
}

// StatsAirExportCmd defines 'export' subcommand
var StatsAirExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/stats/air/operators/{operator_id}/export:post:summary"),
	Long:  TRAPI(`/stats/air/operators/{operator_id}/export:post:description`),
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

		param, err := collectStatsAirExportCmdParams()
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

	if StatsAirExportCmdExportMode != "" {
		result = append(result, sprintf("%s=%s", "export_mode", StatsAirExportCmdExportMode))
	}

	return strings.Join(result, "&")
}

func buildBodyForStatsAirExportCmd() (string, error) {
	if StatsAirExportCmdBody != "" {
		if strings.HasPrefix(StatsAirExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsAirExportCmdBody, "@")
			// #nosec
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
