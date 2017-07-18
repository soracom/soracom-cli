package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorIssueEmailChangeTokenCmdEmail holds value of 'email' option
var OperatorIssueEmailChangeTokenCmdEmail string

// OperatorIssueEmailChangeTokenCmdBody holds contents of request body to be sent
var OperatorIssueEmailChangeTokenCmdBody string

func init() {
	OperatorIssueEmailChangeTokenCmd.Flags().StringVar(&OperatorIssueEmailChangeTokenCmdEmail, "email", "", TRAPI(""))

	OperatorIssueEmailChangeTokenCmd.Flags().StringVar(&OperatorIssueEmailChangeTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorIssueEmailChangeTokenCmd)
}

// OperatorIssueEmailChangeTokenCmd defines 'issue-email-change-token' subcommand
var OperatorIssueEmailChangeTokenCmd = &cobra.Command{
	Use:   "issue-email-change-token",
	Short: TRAPI("/operators/email_change_token/issue:post:summary"),
	Long:  TRAPI(`/operators/email_change_token/issue:post:description`),
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

		param, err := collectOperatorIssueEmailChangeTokenCmdParams()
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

func collectOperatorIssueEmailChangeTokenCmdParams() (*apiParams, error) {

	body, err := buildBodyForOperatorIssueEmailChangeTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorIssueEmailChangeTokenCmd("/operators/email_change_token/issue"),
		query:       buildQueryForOperatorIssueEmailChangeTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorIssueEmailChangeTokenCmd(path string) string {

	return path
}

func buildQueryForOperatorIssueEmailChangeTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorIssueEmailChangeTokenCmd() (string, error) {
	if OperatorIssueEmailChangeTokenCmdBody != "" {
		if strings.HasPrefix(OperatorIssueEmailChangeTokenCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorIssueEmailChangeTokenCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if OperatorIssueEmailChangeTokenCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return OperatorIssueEmailChangeTokenCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if OperatorIssueEmailChangeTokenCmdEmail != "" {
		result["email"] = OperatorIssueEmailChangeTokenCmdEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
