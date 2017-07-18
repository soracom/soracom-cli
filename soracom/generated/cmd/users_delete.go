package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersDeleteCmdOperatorId holds value of 'operator_id' option
var UsersDeleteCmdOperatorId string

// UsersDeleteCmdUserName holds value of 'user_name' option
var UsersDeleteCmdUserName string

func init() {
	UsersDeleteCmd.Flags().StringVar(&UsersDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersDeleteCmd.Flags().StringVar(&UsersDeleteCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersCmd.AddCommand(UsersDeleteCmd)
}

// UsersDeleteCmd defines 'delete' subcommand
var UsersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:delete:description`),
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

		param, err := collectUsersDeleteCmdParams()
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

func collectUsersDeleteCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersDeleteCmd("/operators/{operator_id}/users/{user_name}"),
		query:  buildQueryForUsersDeleteCmd(),
	}, nil
}

func buildPathForUsersDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersDeleteCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersDeleteCmdUserName, -1)

	return path
}

func buildQueryForUsersDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
