// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersListCmdOperatorId holds value of 'operator_id' option
var UsersListCmdOperatorId string

func init() {
	UsersListCmd.Flags().StringVar(&UsersListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	UsersCmd.AddCommand(UsersListCmd)
}

// UsersListCmd defines 'list' subcommand
var UsersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/users:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users:get:description`),
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

		param, err := collectUsersListCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectUsersListCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersListCmdOperatorId == "" {
		UsersListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersListCmd("/operators/{operator_id}/users"),
		query:  buildQueryForUsersListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersListCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForUsersListCmd() url.Values {
	result := url.Values{}

	return result
}
