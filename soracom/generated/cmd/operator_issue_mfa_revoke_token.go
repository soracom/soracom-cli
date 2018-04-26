package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorIssueMfaRevokeTokenCmdEmail holds value of 'email' option
var OperatorIssueMfaRevokeTokenCmdEmail string

// OperatorIssueMfaRevokeTokenCmdPassword holds value of 'password' option
var OperatorIssueMfaRevokeTokenCmdPassword string

// OperatorIssueMfaRevokeTokenCmdBody holds contents of request body to be sent
var OperatorIssueMfaRevokeTokenCmdBody string

func init() {
	OperatorIssueMfaRevokeTokenCmd.Flags().StringVar(&OperatorIssueMfaRevokeTokenCmdEmail, "email", "", TRAPI(""))

	OperatorIssueMfaRevokeTokenCmd.Flags().StringVar(&OperatorIssueMfaRevokeTokenCmdPassword, "password", "", TRAPI(""))

	OperatorIssueMfaRevokeTokenCmd.Flags().StringVar(&OperatorIssueMfaRevokeTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorIssueMfaRevokeTokenCmd)
}

// OperatorIssueMfaRevokeTokenCmd defines 'issue-mfa-revoke-token' subcommand
var OperatorIssueMfaRevokeTokenCmd = &cobra.Command{
	Use:   "issue-mfa-revoke-token",
	Short: TRAPI("/operators/mfa_revoke_token/issue:post:summary"),
	Long:  TRAPI(`/operators/mfa_revoke_token/issue:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectOperatorIssueMfaRevokeTokenCmdParams(ac)
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectOperatorIssueMfaRevokeTokenCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForOperatorIssueMfaRevokeTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorIssueMfaRevokeTokenCmd("/operators/mfa_revoke_token/issue"),
		query:       buildQueryForOperatorIssueMfaRevokeTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorIssueMfaRevokeTokenCmd(path string) string {

	return path
}

func buildQueryForOperatorIssueMfaRevokeTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorIssueMfaRevokeTokenCmd() (string, error) {
	if OperatorIssueMfaRevokeTokenCmdBody != "" {
		if strings.HasPrefix(OperatorIssueMfaRevokeTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorIssueMfaRevokeTokenCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if OperatorIssueMfaRevokeTokenCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return OperatorIssueMfaRevokeTokenCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if OperatorIssueMfaRevokeTokenCmdEmail != "" {
		result["email"] = OperatorIssueMfaRevokeTokenCmdEmail
	}

	if OperatorIssueMfaRevokeTokenCmdPassword != "" {
		result["password"] = OperatorIssueMfaRevokeTokenCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
