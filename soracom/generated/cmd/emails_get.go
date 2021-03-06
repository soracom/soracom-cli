// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EmailsGetCmdEmailId holds value of 'email_id' option
var EmailsGetCmdEmailId string

// EmailsGetCmdOperatorId holds value of 'operator_id' option
var EmailsGetCmdOperatorId string

func init() {
	EmailsGetCmd.Flags().StringVar(&EmailsGetCmdEmailId, "email-id", "", TRAPI("email_id"))

	EmailsGetCmd.Flags().StringVar(&EmailsGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	EmailsCmd.AddCommand(EmailsGetCmd)
}

// EmailsGetCmd defines 'get' subcommand
var EmailsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/emails/{email_id}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/emails/{email_id}:get:description`),
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

		param, err := collectEmailsGetCmdParams(ac)
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
	},
}

func collectEmailsGetCmdParams(ac *apiClient) (*apiParams, error) {
	if EmailsGetCmdOperatorId == "" {
		EmailsGetCmdOperatorId = ac.OperatorID
	}

	if EmailsGetCmdEmailId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "email-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForEmailsGetCmd("/operators/{operator_id}/emails/{email_id}"),
		query:  buildQueryForEmailsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEmailsGetCmd(path string) string {

	escapedEmailId := url.PathEscape(EmailsGetCmdEmailId)

	path = strReplace(path, "{"+"email_id"+"}", escapedEmailId, -1)

	escapedOperatorId := url.PathEscape(EmailsGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForEmailsGetCmd() url.Values {
	result := url.Values{}

	return result
}
