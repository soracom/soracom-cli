package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersGetCmdOperatorId holds value of 'operator_id' option
var UsersGetCmdOperatorId string

// UsersGetCmdUserName holds value of 'user_name' option
var UsersGetCmdUserName string

func init() {
	UsersGetCmd.Flags().StringVar(&UsersGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersGetCmd.Flags().StringVar(&UsersGetCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersCmd.AddCommand(UsersGetCmd)
}

// UsersGetCmd defines 'get' subcommand
var UsersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:get:description`),
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

		param, err := collectUsersGetCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectUsersGetCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersGetCmdOperatorId == "" {
		UsersGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersGetCmd("/operators/{operator_id}/users/{user_name}"),
		query:  buildQueryForUsersGetCmd(),
	}, nil
}

func buildPathForUsersGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersGetCmdUserName, -1)

	return path
}

func buildQueryForUsersGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
