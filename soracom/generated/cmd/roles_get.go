package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// RolesGetCmdOperatorId holds value of 'operator_id' option
var RolesGetCmdOperatorId string

// RolesGetCmdRoleId holds value of 'role_id' option
var RolesGetCmdRoleId string

func init() {
	RolesGetCmd.Flags().StringVar(&RolesGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesGetCmd.Flags().StringVar(&RolesGetCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesCmd.AddCommand(RolesGetCmd)
}

// RolesGetCmd defines 'get' subcommand
var RolesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:get:description`),
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

		param, err := collectRolesGetCmdParams(ac)
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

func collectRolesGetCmdParams(ac *apiClient) (*apiParams, error) {

	if RolesGetCmdOperatorId == "" {
		RolesGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesGetCmd("/operators/{operator_id}/roles/{role_id}"),
		query:  buildQueryForRolesGetCmd(),
	}, nil
}

func buildPathForRolesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"role_id"+"}", RolesGetCmdRoleId, -1)

	return path
}

func buildQueryForRolesGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
