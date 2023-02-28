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

// OperatorGenerateApiTokenCmdOperatorId holds value of 'operator_id' option
var OperatorGenerateApiTokenCmdOperatorId string

// OperatorGenerateApiTokenCmdTokenTimeoutSeconds holds value of 'tokenTimeoutSeconds' option
var OperatorGenerateApiTokenCmdTokenTimeoutSeconds int64

// OperatorGenerateApiTokenCmdBody holds contents of request body to be sent
var OperatorGenerateApiTokenCmdBody string

func init() {
	OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	OperatorGenerateApiTokenCmd.Flags().Int64Var(&OperatorGenerateApiTokenCmdTokenTimeoutSeconds, "token-timeout-seconds", 86400, TRAPI("New API token expiry duration in seconds.Default: 86400 [seconds] (24 hours)Maximum: 172800 [seconds] (48 hours)"))

	OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	OperatorCmd.AddCommand(OperatorGenerateApiTokenCmd)
}

// OperatorGenerateApiTokenCmd defines 'generate-api-token' subcommand
var OperatorGenerateApiTokenCmd = &cobra.Command{
	Use:   "generate-api-token",
	Short: TRAPI("/operators/{operator_id}/token:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/token:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "generateAuthToken"),
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

		param, err := collectOperatorGenerateApiTokenCmdParams(ac)
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

func collectOperatorGenerateApiTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if OperatorGenerateApiTokenCmdOperatorId == "" {
		OperatorGenerateApiTokenCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForOperatorGenerateApiTokenCmd()
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
		path:        buildPathForOperatorGenerateApiTokenCmd("/operators/{operator_id}/token"),
		query:       buildQueryForOperatorGenerateApiTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorGenerateApiTokenCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorGenerateApiTokenCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorGenerateApiTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorGenerateApiTokenCmd() (string, error) {
	var result map[string]interface{}

	if OperatorGenerateApiTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorGenerateApiTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorGenerateApiTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorGenerateApiTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorGenerateApiTokenCmdBody)
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

	if OperatorGenerateApiTokenCmdTokenTimeoutSeconds != 86400 {

		result["tokenTimeoutSeconds"] = OperatorGenerateApiTokenCmdTokenTimeoutSeconds

	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
