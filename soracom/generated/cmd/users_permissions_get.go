package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPermissionsGetCmdOperatorId holds value of 'operator_id' option
var UsersPermissionsGetCmdOperatorId string

// UsersPermissionsGetCmdUserName holds value of 'user_name' option
var UsersPermissionsGetCmdUserName string

func init() {
	UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPermissionsCmd.AddCommand(UsersPermissionsGetCmd)
}

// UsersPermissionsGetCmd defines 'get' subcommand
var UsersPermissionsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/permission:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/permission:get:description`),
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

		param, err := collectUsersPermissionsGetCmdParams()
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

func collectUsersPermissionsGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersPermissionsGetCmd("/operators/{operator_id}/users/{user_name}/permission"),
		query:  buildQueryForUsersPermissionsGetCmd(),
	}, nil
}

func buildPathForUsersPermissionsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersPermissionsGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersPermissionsGetCmdUserName, -1)

	return path
}

func buildQueryForUsersPermissionsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
