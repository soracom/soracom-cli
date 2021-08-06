// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// RolesDeleteCmdOperatorId holds value of 'operator_id' option
var RolesDeleteCmdOperatorId string

// RolesDeleteCmdRoleId holds value of 'role_id' option
var RolesDeleteCmdRoleId string

func init() {
	RolesDeleteCmd.Flags().StringVar(&RolesDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesDeleteCmd.Flags().StringVar(&RolesDeleteCmdRoleId, "role-id", "", TRAPI("role_id"))
	RolesCmd.AddCommand(RolesDeleteCmd)
}

// RolesDeleteCmd defines 'delete' subcommand
var RolesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:delete:description`),
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

		param, err := collectRolesDeleteCmdParams(ac)
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

func collectRolesDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if RolesDeleteCmdOperatorId == "" {
		RolesDeleteCmdOperatorId = ac.OperatorID
	}

	if RolesDeleteCmdRoleId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "role-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForRolesDeleteCmd("/operators/{operator_id}/roles/{role_id}"),
		query:  buildQueryForRolesDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(RolesDeleteCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	return path
}

func buildQueryForRolesDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
