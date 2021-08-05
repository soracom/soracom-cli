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

// DiagnosticsSendRequestCmdResourceId holds value of 'resourceId' option
var DiagnosticsSendRequestCmdResourceId string

// DiagnosticsSendRequestCmdResourceType holds value of 'resourceType' option
var DiagnosticsSendRequestCmdResourceType string

// DiagnosticsSendRequestCmdService holds value of 'service' option
var DiagnosticsSendRequestCmdService string

// DiagnosticsSendRequestCmdFrom holds value of 'from' option
var DiagnosticsSendRequestCmdFrom int64

// DiagnosticsSendRequestCmdTo holds value of 'to' option
var DiagnosticsSendRequestCmdTo int64

// DiagnosticsSendRequestCmdBody holds contents of request body to be sent
var DiagnosticsSendRequestCmdBody string

func init() {
	DiagnosticsSendRequestCmd.Flags().StringVar(&DiagnosticsSendRequestCmdResourceId, "resource-id", "", TRAPI(""))

	DiagnosticsSendRequestCmd.Flags().StringVar(&DiagnosticsSendRequestCmdResourceType, "resource-type", "", TRAPI(""))

	DiagnosticsSendRequestCmd.Flags().StringVar(&DiagnosticsSendRequestCmdService, "service", "", TRAPI(""))

	DiagnosticsSendRequestCmd.Flags().Int64Var(&DiagnosticsSendRequestCmdFrom, "from", 0, TRAPI(""))

	DiagnosticsSendRequestCmd.Flags().Int64Var(&DiagnosticsSendRequestCmdTo, "to", 0, TRAPI(""))

	DiagnosticsSendRequestCmd.Flags().StringVar(&DiagnosticsSendRequestCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DiagnosticsCmd.AddCommand(DiagnosticsSendRequestCmd)
}

// DiagnosticsSendRequestCmd defines 'send-request' subcommand
var DiagnosticsSendRequestCmd = &cobra.Command{
	Use:   "send-request",
	Short: TRAPI("/diagnostics:post:summary"),
	Long:  TRAPI(`/diagnostics:post:description`),
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

		param, err := collectDiagnosticsSendRequestCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDiagnosticsSendRequestCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForDiagnosticsSendRequestCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if DiagnosticsSendRequestCmdResourceId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "resource-id")
		}

	}

	if DiagnosticsSendRequestCmdResourceType == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "resource-type")
		}

	}

	if DiagnosticsSendRequestCmdService == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "service")
		}

	}

	if DiagnosticsSendRequestCmdFrom == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "from")
		}

	}

	if DiagnosticsSendRequestCmdTo == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "to")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDiagnosticsSendRequestCmd("/diagnostics"),
		query:       buildQueryForDiagnosticsSendRequestCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDiagnosticsSendRequestCmd(path string) string {

	return path
}

func buildQueryForDiagnosticsSendRequestCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDiagnosticsSendRequestCmd() (string, error) {
	var result map[string]interface{}

	if DiagnosticsSendRequestCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DiagnosticsSendRequestCmdBody, "@") {
			fname := strings.TrimPrefix(DiagnosticsSendRequestCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DiagnosticsSendRequestCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DiagnosticsSendRequestCmdBody)
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

	if DiagnosticsSendRequestCmdResourceId != "" {
		result["resourceId"] = DiagnosticsSendRequestCmdResourceId
	}

	if DiagnosticsSendRequestCmdResourceType != "" {
		result["resourceType"] = DiagnosticsSendRequestCmdResourceType
	}

	if DiagnosticsSendRequestCmdService != "" {
		result["service"] = DiagnosticsSendRequestCmdService
	}

	if DiagnosticsSendRequestCmdFrom != 0 {
		result["from"] = DiagnosticsSendRequestCmdFrom
	}

	if DiagnosticsSendRequestCmdTo != 0 {
		result["to"] = DiagnosticsSendRequestCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
