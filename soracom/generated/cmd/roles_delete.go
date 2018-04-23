package cmd

import (
	"os"
	"strings"

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

func collectRolesDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	if RolesDeleteCmdOperatorId == "" {
		RolesDeleteCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForRolesDeleteCmd("/operators/{operator_id}/roles/{role_id}"),
		query:  buildQueryForRolesDeleteCmd(),
	}, nil
}

func buildPathForRolesDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesDeleteCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"role_id"+"}", RolesDeleteCmdRoleId, -1)

	return path
}

func buildQueryForRolesDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
