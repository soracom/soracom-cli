// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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

func InitStatsAirExportCmd() {
	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdExportMode, "export-mode", "", TRAPI("Specify how to obtain the URL to download the Air Data Usage Report CSV.- 'async': Get the 'exportedFieldId' without waiting for the URL to be issued on the Soracom platform. Specify this 'exportedFieldId' in ['Files:getExportedFile API'](#/Files/getExportedFile) to get the URL. If the file size of the Air Data Usage Report CSV is huge, specify 'async'.- 'sync' (default): Wait for the URL to be issued on the Soracom platform. However, if the file size of the Air Data Usage Report CSV is huge, it may time out and the URL cannot be retrieved. If the timeout occurs, specify 'async'."))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdPeriod, "period", "", TRAPI("Degree of detail of history.- 'month': Monthly- 'day': Daily- 'minutes': Every minute"))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdFrom, "from", 0, TRAPI("Start date and time for the aggregate data (UNIX time in seconds)"))

	StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdTo, "to", 0, TRAPI("End date and time of the period covered (UNIX time in seconds)"))

	StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	StatsAirExportCmd.RunE = StatsAirExportCmdRunE

	StatsAirCmd.AddCommand(StatsAirExportCmd)
}

// StatsAirExportCmd defines 'export' subcommand
var StatsAirExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/stats/air/operators/{operator_id}/export:post:summary"),
	Long:  TRAPI(`/stats/air/operators/{operator_id}/export:post:description`) + "\n\n" + createLinkToAPIReference("Stats", "exportAirStats"),
}

func StatsAirExportCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectStatsAirExportCmdParams(ac)
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
	rawOutput = true

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectStatsAirExportCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if StatsAirExportCmdOperatorId == "" {
		StatsAirExportCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForStatsAirExportCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForStatsAirExportCmd("/stats/air/operators/{operator_id}/export"),
		query:       buildQueryForStatsAirExportCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsAirExportCmd(path string) string {

	escapedOperatorId := url.PathEscape(StatsAirExportCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForStatsAirExportCmd() url.Values {
	result := url.Values{}

	if StatsAirExportCmdExportMode != "" {
		result.Add("export_mode", StatsAirExportCmdExportMode)
	}

	return result
}

func buildBodyForStatsAirExportCmd() (string, error) {
	var result map[string]interface{}

	if StatsAirExportCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(StatsAirExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsAirExportCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if StatsAirExportCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(StatsAirExportCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if StatsAirExportCmdPeriod != "" {
		result["period"] = StatsAirExportCmdPeriod
	}

	if StatsAirExportCmd.Flags().Lookup("from").Changed {
		result["from"] = StatsAirExportCmdFrom
	}

	if StatsAirExportCmd.Flags().Lookup("to").Changed {
		result["to"] = StatsAirExportCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
