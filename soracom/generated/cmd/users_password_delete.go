package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPasswordDeleteCmdOperatorId holds value of 'operator_id' option
var UsersPasswordDeleteCmdOperatorId string

// UsersPasswordDeleteCmdUserName holds value of 'user_name' option
var UsersPasswordDeleteCmdUserName string

func init() {
	UsersPasswordDeleteCmd.Flags().StringVar(&UsersPasswordDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersPasswordDeleteCmd.Flags().StringVar(&UsersPasswordDeleteCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPasswordCmd.AddCommand(UsersPasswordDeleteCmd)
}

// UsersPasswordDeleteCmd defines 'delete' subcommand
var UsersPasswordDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/password:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/password:delete:description`),
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

		param, err := collectUsersPasswordDeleteCmdParams(ac)
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

func collectUsersPasswordDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersPasswordDeleteCmdOperatorId == "" {
		UsersPasswordDeleteCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersPasswordDeleteCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:  buildQueryForUsersPasswordDeleteCmd(),
	}, nil
}

func buildPathForUsersPasswordDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersPasswordDeleteCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersPasswordDeleteCmdUserName, -1)

	return path
}

func buildQueryForUsersPasswordDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
