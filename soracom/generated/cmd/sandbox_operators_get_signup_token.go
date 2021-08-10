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

// SandboxOperatorsGetSignupTokenCmdAuthKey holds value of 'authKey' option
var SandboxOperatorsGetSignupTokenCmdAuthKey string

// SandboxOperatorsGetSignupTokenCmdAuthKeyId holds value of 'authKeyId' option
var SandboxOperatorsGetSignupTokenCmdAuthKeyId string

// SandboxOperatorsGetSignupTokenCmdEmail holds value of 'email' option
var SandboxOperatorsGetSignupTokenCmdEmail string

// SandboxOperatorsGetSignupTokenCmdBody holds contents of request body to be sent
var SandboxOperatorsGetSignupTokenCmdBody string

func init() {
	SandboxOperatorsGetSignupTokenCmd.Flags().StringVar(&SandboxOperatorsGetSignupTokenCmdAuthKey, "auth-key", "", TRAPI(""))

	SandboxOperatorsGetSignupTokenCmd.Flags().StringVar(&SandboxOperatorsGetSignupTokenCmdAuthKeyId, "auth-key-id", "", TRAPI(""))

	SandboxOperatorsGetSignupTokenCmd.Flags().StringVar(&SandboxOperatorsGetSignupTokenCmdEmail, "email", "", TRAPI("email"))

	SandboxOperatorsGetSignupTokenCmd.Flags().StringVar(&SandboxOperatorsGetSignupTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SandboxOperatorsCmd.AddCommand(SandboxOperatorsGetSignupTokenCmd)
}

// SandboxOperatorsGetSignupTokenCmd defines 'get-signup-token' subcommand
var SandboxOperatorsGetSignupTokenCmd = &cobra.Command{
	Use:   "get-signup-token",
	Short: TRAPI("/sandbox/operators/token/{email}:post:summary"),
	Long:  TRAPI(`/sandbox/operators/token/{email}:post:description`),
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

		param, err := collectSandboxOperatorsGetSignupTokenCmdParams(ac)
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

func collectSandboxOperatorsGetSignupTokenCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSandboxOperatorsGetSignupTokenCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SandboxOperatorsGetSignupTokenCmdEmail == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "email")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxOperatorsGetSignupTokenCmd("/sandbox/operators/token/{email}"),
		query:       buildQueryForSandboxOperatorsGetSignupTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxOperatorsGetSignupTokenCmd(path string) string {

	escapedEmail := url.PathEscape(SandboxOperatorsGetSignupTokenCmdEmail)

	path = strReplace(path, "{"+"email"+"}", escapedEmail, -1)

	return path
}

func buildQueryForSandboxOperatorsGetSignupTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxOperatorsGetSignupTokenCmd() (string, error) {
	var result map[string]interface{}

	if SandboxOperatorsGetSignupTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxOperatorsGetSignupTokenCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxOperatorsGetSignupTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxOperatorsGetSignupTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxOperatorsGetSignupTokenCmdBody)
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

	if SandboxOperatorsGetSignupTokenCmdAuthKey != "" {
		result["authKey"] = SandboxOperatorsGetSignupTokenCmdAuthKey
	}

	if SandboxOperatorsGetSignupTokenCmdAuthKeyId != "" {
		result["authKeyId"] = SandboxOperatorsGetSignupTokenCmdAuthKeyId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
