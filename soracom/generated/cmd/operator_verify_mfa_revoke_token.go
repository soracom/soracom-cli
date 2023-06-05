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

// OperatorVerifyMfaRevokeTokenCmdBackupCode holds value of 'backupCode' option
var OperatorVerifyMfaRevokeTokenCmdBackupCode string

// OperatorVerifyMfaRevokeTokenCmdEmail holds value of 'email' option
var OperatorVerifyMfaRevokeTokenCmdEmail string

// OperatorVerifyMfaRevokeTokenCmdPassword holds value of 'password' option
var OperatorVerifyMfaRevokeTokenCmdPassword string

// OperatorVerifyMfaRevokeTokenCmdToken holds value of 'token' option
var OperatorVerifyMfaRevokeTokenCmdToken string

// OperatorVerifyMfaRevokeTokenCmdBody holds contents of request body to be sent
var OperatorVerifyMfaRevokeTokenCmdBody string

func InitOperatorVerifyMfaRevokeTokenCmd() {
	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdBackupCode, "backup-code", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdEmail, "email", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdPassword, "password", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdToken, "token", "", TRAPI(""))

	OperatorVerifyMfaRevokeTokenCmd.Flags().StringVar(&OperatorVerifyMfaRevokeTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorVerifyMfaRevokeTokenCmd.RunE = OperatorVerifyMfaRevokeTokenCmdRunE

	OperatorCmd.AddCommand(OperatorVerifyMfaRevokeTokenCmd)
}

// OperatorVerifyMfaRevokeTokenCmd defines 'verify-mfa-revoke-token' subcommand
var OperatorVerifyMfaRevokeTokenCmd = &cobra.Command{
	Use:   "verify-mfa-revoke-token",
	Short: TRAPI("/operators/mfa_revoke_token/verify:post:summary"),
	Long:  TRAPI(`/operators/mfa_revoke_token/verify:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "verifyMFARevokingToken"),
}

func OperatorVerifyMfaRevokeTokenCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorVerifyMfaRevokeTokenCmdParams(ac)
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

func collectOperatorVerifyMfaRevokeTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForOperatorVerifyMfaRevokeTokenCmd()
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
		path:        buildPathForOperatorVerifyMfaRevokeTokenCmd("/operators/mfa_revoke_token/verify"),
		query:       buildQueryForOperatorVerifyMfaRevokeTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorVerifyMfaRevokeTokenCmd(path string) string {

	return path
}

func buildQueryForOperatorVerifyMfaRevokeTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorVerifyMfaRevokeTokenCmd() (string, error) {
	var result map[string]interface{}

	if OperatorVerifyMfaRevokeTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorVerifyMfaRevokeTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorVerifyMfaRevokeTokenCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if OperatorVerifyMfaRevokeTokenCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorVerifyMfaRevokeTokenCmdBody)
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

	if OperatorVerifyMfaRevokeTokenCmdBackupCode != "" {
		result["backupCode"] = OperatorVerifyMfaRevokeTokenCmdBackupCode
	}

	if OperatorVerifyMfaRevokeTokenCmdEmail != "" {
		result["email"] = OperatorVerifyMfaRevokeTokenCmdEmail
	}

	if OperatorVerifyMfaRevokeTokenCmdPassword != "" {
		result["password"] = OperatorVerifyMfaRevokeTokenCmdPassword
	}

	if OperatorVerifyMfaRevokeTokenCmdToken != "" {
		result["token"] = OperatorVerifyMfaRevokeTokenCmdToken
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
