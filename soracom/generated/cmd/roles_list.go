// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// RolesListCmdOperatorId holds value of 'operator_id' option
var RolesListCmdOperatorId string

func init() {
	RolesListCmd.Flags().StringVar(&RolesListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	RolesCmd.AddCommand(RolesListCmd)
}

// RolesListCmd defines 'list' subcommand
var RolesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/roles:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles:get:description`),
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

		param, err := collectRolesListCmdParams(ac)
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

func collectRolesListCmdParams(ac *apiClient) (*apiParams, error) {
	if RolesListCmdOperatorId == "" {
		RolesListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListCmd("/operators/{operator_id}/roles"),
		query:  buildQueryForRolesListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesListCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForRolesListCmd() url.Values {
	result := url.Values{}

	return result
}
