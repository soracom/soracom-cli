package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// AuthIssuePasswordResetTokenCmdEmail holds value of 'email' option
var AuthIssuePasswordResetTokenCmdEmail string

// AuthIssuePasswordResetTokenCmdBody holds contents of request body to be sent
var AuthIssuePasswordResetTokenCmdBody string

func init() {
	AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdEmail, "email", "", TRAPI(""))

	AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	AuthCmd.AddCommand(AuthIssuePasswordResetTokenCmd)
}

// AuthIssuePasswordResetTokenCmd defines 'issue-password-reset-token' subcommand
var AuthIssuePasswordResetTokenCmd = &cobra.Command{
	Use:   "issue-password-reset-token",
	Short: TRAPI("/auth/password_reset_token/issue:post:summary"),
	Long:  TRAPI(`/auth/password_reset_token/issue:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectAuthIssuePasswordResetTokenCmdParams(ac)
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

func collectAuthIssuePasswordResetTokenCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForAuthIssuePasswordResetTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForAuthIssuePasswordResetTokenCmd("/auth/password_reset_token/issue"),
		query:       buildQueryForAuthIssuePasswordResetTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForAuthIssuePasswordResetTokenCmd(path string) string {

	return path
}

func buildQueryForAuthIssuePasswordResetTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForAuthIssuePasswordResetTokenCmd() (string, error) {
	var result map[string]interface{}

	if AuthIssuePasswordResetTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(AuthIssuePasswordResetTokenCmdBody, "@") {
			fname := strings.TrimPrefix(AuthIssuePasswordResetTokenCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if AuthIssuePasswordResetTokenCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(AuthIssuePasswordResetTokenCmdBody)
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

	if AuthIssuePasswordResetTokenCmdEmail != "" {
		result["email"] = AuthIssuePasswordResetTokenCmdEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
