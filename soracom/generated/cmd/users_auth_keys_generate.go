package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersAuthKeysGenerateCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysGenerateCmdOperatorId string

// UsersAuthKeysGenerateCmdUserName holds value of 'user_name' option
var UsersAuthKeysGenerateCmdUserName string

func init() {
	UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAuthKeysCmd.AddCommand(UsersAuthKeysGenerateCmd)
}

// UsersAuthKeysGenerateCmd defines 'generate' subcommand
var UsersAuthKeysGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys:post:description`),
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

		param, err := collectUsersAuthKeysGenerateCmdParams()
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

func collectUsersAuthKeysGenerateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForUsersAuthKeysGenerateCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
		query:  buildQueryForUsersAuthKeysGenerateCmd(),
	}, nil
}

func buildPathForUsersAuthKeysGenerateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersAuthKeysGenerateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersAuthKeysGenerateCmdUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysGenerateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
