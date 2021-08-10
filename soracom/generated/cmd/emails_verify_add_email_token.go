// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// EmailsVerifyAddEmailTokenCmdToken holds value of 'token' option
var EmailsVerifyAddEmailTokenCmdToken string

// EmailsVerifyAddEmailTokenCmdBody holds contents of request body to be sent
var EmailsVerifyAddEmailTokenCmdBody string

func init() {
	EmailsVerifyAddEmailTokenCmd.Flags().StringVar(&EmailsVerifyAddEmailTokenCmdToken, "token", "", TRAPI(""))

	EmailsVerifyAddEmailTokenCmd.Flags().StringVar(&EmailsVerifyAddEmailTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	EmailsCmd.AddCommand(EmailsVerifyAddEmailTokenCmd)
}

// EmailsVerifyAddEmailTokenCmd defines 'verify-add-email-token' subcommand
var EmailsVerifyAddEmailTokenCmd = &cobra.Command{
	Use:   "verify-add-email-token",
	Short: TRAPI("/operators/add_email_token/verify:post:summary"),
	Long:  TRAPI(`/operators/add_email_token/verify:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
		}

		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectEmailsVerifyAddEmailTokenCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectEmailsVerifyAddEmailTokenCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForEmailsVerifyAddEmailTokenCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if EmailsVerifyAddEmailTokenCmdToken == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "token")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForEmailsVerifyAddEmailTokenCmd("/operators/add_email_token/verify"),
		query:       buildQueryForEmailsVerifyAddEmailTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEmailsVerifyAddEmailTokenCmd(path string) string {

	return path
}

func buildQueryForEmailsVerifyAddEmailTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForEmailsVerifyAddEmailTokenCmd() (string, error) {
	var result map[string]interface{}

	if EmailsVerifyAddEmailTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(EmailsVerifyAddEmailTokenCmdBody, "@") {
			fname := strings.TrimPrefix(EmailsVerifyAddEmailTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if EmailsVerifyAddEmailTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(EmailsVerifyAddEmailTokenCmdBody)
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

	if EmailsVerifyAddEmailTokenCmdToken != "" {
		result["token"] = EmailsVerifyAddEmailTokenCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
