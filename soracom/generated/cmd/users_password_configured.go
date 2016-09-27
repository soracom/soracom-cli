package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPasswordConfiguredCmdOperatorId holds value of 'operator_id' option
var UsersPasswordConfiguredCmdOperatorId string

// UsersPasswordConfiguredCmdUserName holds value of 'user_name' option
var UsersPasswordConfiguredCmdUserName string

func init() {
	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdOperatorId, "operator-id", "", TR("operator_id"))

	UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdUserName, "user-name", "", TR("user_name"))

	UsersPasswordCmd.AddCommand(UsersPasswordConfiguredCmd)
}

// UsersPasswordConfiguredCmd defines 'configured' subcommand
var UsersPasswordConfiguredCmd = &cobra.Command{
	Use:   "configured",
	Short: TR("users.has_user_password.get.summary"),
	Long:  TR(`users.has_user_password.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectUsersPasswordConfiguredCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectUsersPasswordConfiguredCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersPasswordConfiguredCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:  buildQueryForUsersPasswordConfiguredCmd(),
	}, nil
}

func buildPathForUsersPasswordConfiguredCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersPasswordConfiguredCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersPasswordConfiguredCmdUserName, -1)

	return path
}

func buildQueryForUsersPasswordConfiguredCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
