// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// StatsFunnelExportCmdExportMode holds value of 'export_mode' option
var StatsFunnelExportCmdExportMode string

// StatsFunnelExportCmdOperatorId holds value of 'operator_id' option
var StatsFunnelExportCmdOperatorId string

// StatsFunnelExportCmdPeriod holds value of 'period' option
var StatsFunnelExportCmdPeriod string

// StatsFunnelExportCmdFrom holds value of 'from' option
var StatsFunnelExportCmdFrom int64

// StatsFunnelExportCmdTo holds value of 'to' option
var StatsFunnelExportCmdTo int64

// StatsFunnelExportCmdBody holds contents of request body to be sent
var StatsFunnelExportCmdBody string

func init() {
	StatsFunnelExportCmd.Flags().StringVar(&StatsFunnelExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	StatsFunnelExportCmd.Flags().StringVar(&StatsFunnelExportCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	StatsFunnelExportCmd.Flags().StringVar(&StatsFunnelExportCmdPeriod, "period", "", TRAPI("Degree of detail of history.- `month`: Monthly- `day`: Daily- `minutes`: Every minute"))

	StatsFunnelExportCmd.Flags().Int64Var(&StatsFunnelExportCmdFrom, "from", 0, TRAPI("Start date and time for the aggregate data (UNIX time in seconds)"))

	StatsFunnelExportCmd.Flags().Int64Var(&StatsFunnelExportCmdTo, "to", 0, TRAPI("End date and time of the period covered (UNIX time in seconds)"))

	StatsFunnelExportCmd.Flags().StringVar(&StatsFunnelExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	StatsFunnelCmd.AddCommand(StatsFunnelExportCmd)
}

// StatsFunnelExportCmd defines 'export' subcommand
var StatsFunnelExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/stats/funnel/operators/{operator_id}/export:post:summary"),
	Long:  TRAPI(`/stats/funnel/operators/{operator_id}/export:post:description`) + "\n\n" + createLinkToAPIReference("Stats", "exportFunnelStats"),
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

		param, err := collectStatsFunnelExportCmdParams(ac)
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
	},
}

func collectStatsFunnelExportCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if StatsFunnelExportCmdOperatorId == "" {
		StatsFunnelExportCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForStatsFunnelExportCmd()
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
		path:        buildPathForStatsFunnelExportCmd("/stats/funnel/operators/{operator_id}/export"),
		query:       buildQueryForStatsFunnelExportCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsFunnelExportCmd(path string) string {

	escapedOperatorId := url.PathEscape(StatsFunnelExportCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForStatsFunnelExportCmd() url.Values {
	result := url.Values{}

	if StatsFunnelExportCmdExportMode != "" {
		result.Add("export_mode", StatsFunnelExportCmdExportMode)
	}

	return result
}

func buildBodyForStatsFunnelExportCmd() (string, error) {
	var result map[string]interface{}

	if StatsFunnelExportCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(StatsFunnelExportCmdBody, "@") {
			fname := strings.TrimPrefix(StatsFunnelExportCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if StatsFunnelExportCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(StatsFunnelExportCmdBody)
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

	if StatsFunnelExportCmdPeriod != "" {
		result["period"] = StatsFunnelExportCmdPeriod
	}

	result["from"] = StatsFunnelExportCmdFrom

	result["to"] = StatsFunnelExportCmdTo

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
