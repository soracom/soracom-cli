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

// AuthVerifyPasswordResetTokenCmdPassword holds value of 'password' option
var AuthVerifyPasswordResetTokenCmdPassword string

// AuthVerifyPasswordResetTokenCmdToken holds value of 'token' option
var AuthVerifyPasswordResetTokenCmdToken string

// AuthVerifyPasswordResetTokenCmdBody holds contents of request body to be sent
var AuthVerifyPasswordResetTokenCmdBody string

func InitAuthVerifyPasswordResetTokenCmd() {
	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdPassword, "password", "", TRAPI(""))

	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdToken, "token", "", TRAPI(""))

	AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	AuthVerifyPasswordResetTokenCmd.RunE = AuthVerifyPasswordResetTokenCmdRunE

	AuthCmd.AddCommand(AuthVerifyPasswordResetTokenCmd)
}

// AuthVerifyPasswordResetTokenCmd defines 'verify-password-reset-token' subcommand
var AuthVerifyPasswordResetTokenCmd = &cobra.Command{
	Use:   "verify-password-reset-token",
	Short: TRAPI("/auth/password_reset_token/verify:post:summary"),
	Long:  TRAPI(`/auth/password_reset_token/verify:post:description`) + "\n\n" + createLinkToAPIReference("Auth", "verifyPasswordResetToken"),
}

func AuthVerifyPasswordResetTokenCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectAuthVerifyPasswordResetTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForAuthVerifyPasswordResetTokenCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("password", "password", "body", parsedBody, AuthVerifyPasswordResetTokenCmdPassword)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("token", "token", "body", parsedBody, AuthVerifyPasswordResetTokenCmdToken)
	if err != nil {
		return nil, err
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
			b, err = os.ReadFile(fname)
		} else if AuthVerifyPasswordResetTokenCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
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
