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
	var result map[string]interface{}

	if OperatorIssueMfaRevokeTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorIssueMfaRevokeTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorIssueMfaRevokeTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorIssueMfaRevokeTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorIssueMfaRevokeTokenCmdBody)
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
