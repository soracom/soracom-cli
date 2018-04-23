package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// RolesListUsersCmdOperatorId holds value of 'operator_id' option
var RolesListUsersCmdOperatorId string

// RolesListUsersCmdRoleId holds value of 'role_id' option
var RolesListUsersCmdRoleId string

func init() {
	RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesCmd.AddCommand(RolesListUsersCmd)
}

// RolesListUsersCmd defines 'list-users' subcommand
var RolesListUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}/users:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}/users:get:description`),
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

		param, err := collectRolesListUsersCmdParams(ac)
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

func collectRolesListUsersCmdParams(ac *apiClient) (*apiParams, error) {

	if RolesListUsersCmdOperatorId == "" {
		RolesListUsersCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListUsersCmd("/operators/{operator_id}/roles/{role_id}/users"),
		query:  buildQueryForRolesListUsersCmd(),
	}, nil
}

func buildPathForRolesListUsersCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesListUsersCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"role_id"+"}", RolesListUsersCmdRoleId, -1)

	return path
}

func buildQueryForRolesListUsersCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
