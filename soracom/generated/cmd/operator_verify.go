package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorVerifyCmdToken holds value of 'token' option
var OperatorVerifyCmdToken string

// OperatorVerifyCmdBody holds contents of request body to be sent
var OperatorVerifyCmdBody string

func init() {
	OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdToken, "token", "", TRAPI(""))

	OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorVerifyCmd)
}

// OperatorVerifyCmd defines 'verify' subcommand
var OperatorVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: TRAPI("/operators/verify:post:summary"),
	Long:  TRAPI(`/operators/verify:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectOperatorVerifyCmdParams(ac)
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

func collectOperatorVerifyCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForOperatorVerifyCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorVerifyCmd("/operators/verify"),
		query:       buildQueryForOperatorVerifyCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorVerifyCmd(path string) string {

	return path
}

func buildQueryForOperatorVerifyCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorVerifyCmd() (string, error) {
	var result map[string]interface{}

	if OperatorVerifyCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorVerifyCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorVerifyCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorVerifyCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorVerifyCmdBody)
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

	if OperatorVerifyCmdToken != "" {
		result["token"] = OperatorVerifyCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
