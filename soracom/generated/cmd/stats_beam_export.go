package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StatsBeamExportCmdExportMode holds value of 'export_mode' option
var StatsBeamExportCmdExportMode string

// StatsBeamExportCmdOperatorId holds value of 'operator_id' option
var StatsBeamExportCmdOperatorId string

// StatsBeamExportCmdPeriod holds value of 'period' option
var StatsBeamExportCmdPeriod string

// StatsBeamExportCmdFrom holds value of 'from' option
var StatsBeamExportCmdFrom int64

// StatsBeamExportCmdTo holds value of 'to' option
var StatsBeamExportCmdTo int64

// StatsBeamExportCmdBody holds contents of request body to be sent
var StatsBeamExportCmdBody string

func init() {
	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdOperatorId, "operator-id", "", TRAPI("operator ID"))

	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdPeriod, "period", "", TRAPI(""))

	StatsBeamExportCmd.Flags().Int64Var(&StatsBeamExportCmdFrom, "from", 0, TRAPI(""))

	StatsBeamExportCmd.Flags().Int64Var(&StatsBeamExportCmdTo, "to", 0, TRAPI(""))

	StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	StatsBeamCmd.AddCommand(StatsBeamExportCmd)
}

// StatsBeamExportCmd defines 'export' subcommand
var StatsBeamExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/stats/beam/operators/{operator_id}/export:post:summary"),
	Long:  TRAPI(`/stats/beam/operators/{operator_id}/export:post:description`),
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

		param, err := collectStatsBeamExportCmdParams()
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

func collectStatsBeamExportCmdParams() (*apiParams, error) {

	body, err := buildBodyForStatsBeamExportCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForStatsBeamExportCmd("/stats/beam/operators/{operator_id}/export"),
		query:       buildQueryForStatsBeamExportCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForStatsBeamExportCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", StatsBeamExportCmdOperatorId, -1)

	return path
}

func buildQueryForStatsBeamExportCmd() string {
	result := []string{}

	if StatsBeamExportCmdExportMode != "" {
		result = append(result, sprintf("%s=%s", "export_mode", StatsBeamExportCmdExportMode))
	}

	return strings.Join(result, "&")
}

func buildBodyForStatsBeamExportCmd() (string, error) {
	if StatsBeamExportCmdBody != "" {
		if strings.HasPrefix(StatsBeamExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsBeamExportCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if StatsBeamExportCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return StatsBeamExportCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if StatsBeamExportCmdPeriod != "" {
		result["period"] = StatsBeamExportCmdPeriod
	}

	if StatsBeamExportCmdFrom != 0 {
		result["from"] = StatsBeamExportCmdFrom
	}

	if StatsBeamExportCmdTo != 0 {
		result["to"] = StatsBeamExportCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
