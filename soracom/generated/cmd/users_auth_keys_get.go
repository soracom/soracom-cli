package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersAuthKeysGetCmdAuthKeyId holds value of 'auth_key_id' option
var UsersAuthKeysGetCmdAuthKeyId string

// UsersAuthKeysGetCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysGetCmdOperatorId string

// UsersAuthKeysGetCmdUserName holds value of 'user_name' option
var UsersAuthKeysGetCmdUserName string

func init() {
	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdAuthKeyId, "auth-key-id", "", TRAPI("auth_key_id"))

	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAuthKeysCmd.AddCommand(UsersAuthKeysGetCmd)
}

// UsersAuthKeysGetCmd defines 'get' subcommand
var UsersAuthKeysGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}:get:description`),
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

		param, err := collectUsersAuthKeysGetCmdParams()
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

func collectUsersAuthKeysGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersAuthKeysGetCmd("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}"),
		query:  buildQueryForUsersAuthKeysGetCmd(),
	}, nil
}

func buildPathForUsersAuthKeysGetCmd(path string) string {

	path = strings.Replace(path, "{"+"auth_key_id"+"}", UsersAuthKeysGetCmdAuthKeyId, -1)

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersAuthKeysGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersAuthKeysGetCmdUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
