package cmd

import (
	"encoding/json"

	"io/ioutil"

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
	OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorGenerateApiTokenCmd.Flags().Int64Var(&OperatorGenerateApiTokenCmdTokenTimeoutSeconds, "token-timeout-seconds", 0, TRAPI(""))

	OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorGenerateApiTokenCmd)
}

// OperatorGenerateApiTokenCmd defines 'generate-api-token' subcommand
var OperatorGenerateApiTokenCmd = &cobra.Command{
	Use:   "generate-api-token",
	Short: TRAPI("/operators/{operator_id}/token:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/token:post:description`),
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

		param, err := collectOperatorGenerateApiTokenCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectOperatorGenerateApiTokenCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorGenerateApiTokenCmdOperatorId == "" {
		OperatorGenerateApiTokenCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForOperatorGenerateApiTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorGenerateApiTokenCmd("/operators/{operator_id}/token"),
		query:       buildQueryForOperatorGenerateApiTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorGenerateApiTokenCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorGenerateApiTokenCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorGenerateApiTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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

	if OperatorGenerateApiTokenCmdTokenTimeoutSeconds != 0 {
		result["tokenTimeoutSeconds"] = OperatorGenerateApiTokenCmdTokenTimeoutSeconds
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
