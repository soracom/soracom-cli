// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersDefaultPermissionsDeleteCmdOperatorId holds value of 'operator_id' option
var UsersDefaultPermissionsDeleteCmdOperatorId string

func InitUsersDefaultPermissionsDeleteCmd() {
	UsersDefaultPermissionsDeleteCmd.Flags().StringVar(&UsersDefaultPermissionsDeleteCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersDefaultPermissionsDeleteCmd.RunE = UsersDefaultPermissionsDeleteCmdRunE

	UsersDefaultPermissionsCmd.AddCommand(UsersDefaultPermissionsDeleteCmd)
}

// UsersDefaultPermissionsDeleteCmd defines 'delete' subcommand
var UsersDefaultPermissionsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/users/default_permissions:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/default_permissions:delete:description`) + "\n\n" + createLinkToAPIReference("User", "deleteDefaultPermissions"),
}

func UsersDefaultPermissionsDeleteCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectUsersDefaultPermissionsDeleteCmdParams(ac)
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

func collectUsersDefaultPermissionsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersDefaultPermissionsDeleteCmdOperatorId == "" {
		UsersDefaultPermissionsDeleteCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersDefaultPermissionsDeleteCmd("/operators/{operator_id}/users/default_permissions"),
		query:  buildQueryForUsersDefaultPermissionsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersDefaultPermissionsDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersDefaultPermissionsDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForUsersDefaultPermissionsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
