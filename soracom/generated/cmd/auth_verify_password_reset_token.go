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

// AuthVerifyPasswordResetTokenCmdPassword holds value of 'password' option
var AuthVerifyPasswordResetTokenCmdPassword string

// AuthVerifyPasswordResetTokenCmdToken holds value of 'token' option
var AuthVerifyPasswordResetTokenCmdToken string

// AuthVerifyPasswordResetTokenCmdBody holds contents of request body to be sent
var AuthVerifyPasswordResetTokenCmdBody string

func init() {
	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdPassword, "password", "", TRAPI(""))

	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdToken, "token", "", TRAPI(""))

	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	AuthCmd.AddCommand(AuthVerifyPasswordResetTokenCmd)
}

// AuthVerifyPasswordResetTokenCmd defines 'verify-password-reset-token' subcommand
var AuthVerifyPasswordResetTokenCmd = &cobra.Command{
	Use:   "verify-password-reset-token",
	Short: TRAPI("/auth/password_reset_token/verify:post:summary"),
	Long:  TRAPI(`/auth/password_reset_token/verify:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectAuthVerifyPasswordResetTokenCmdParams(ac)
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

func collectAuthVerifyPasswordResetTokenCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForAuthVerifyPasswordResetTokenCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if AuthVerifyPasswordResetTokenCmdPassword == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "password")
		}

	}

	if AuthVerifyPasswordResetTokenCmdToken == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "token")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForAuthVerifyPasswordResetTokenCmd("/auth/password_reset_token/verify"),
		query:       buildQueryForAuthVerifyPasswordResetTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAuthVerifyPasswordResetTokenCmd(path string) string {

	return path
}

func buildQueryForAuthVerifyPasswordResetTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForAuthVerifyPasswordResetTokenCmd() (string, error) {
	var result map[string]interface{}

	if AuthVerifyPasswordResetTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(AuthVerifyPasswordResetTokenCmdBody, "@") {
			fname := strings.TrimPrefix(AuthVerifyPasswordResetTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if AuthVerifyPasswordResetTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(AuthVerifyPasswordResetTokenCmdBody)
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

	if AuthVerifyPasswordResetTokenCmdPassword != "" {
		result["password"] = AuthVerifyPasswordResetTokenCmdPassword
	}

	if AuthVerifyPasswordResetTokenCmdToken != "" {
		result["token"] = AuthVerifyPasswordResetTokenCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
