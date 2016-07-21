package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var UsersAuthKeysGetCmdAuthKeyId string

var UsersAuthKeysGetCmdOperatorId string

var UsersAuthKeysGetCmdUserName string

func init() {
	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdAuthKeyId, "auth-key-id", "", TR("auth_key_id"))

	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdOperatorId, "operator-id", "", TR("operator_id"))

	UsersAuthKeysGetCmd.Flags().StringVar(&UsersAuthKeysGetCmdUserName, "user-name", "", TR("user_name"))

	UsersAuthKeysCmd.AddCommand(UsersAuthKeysGetCmd)
}

var UsersAuthKeysGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("users.get_user_auth_key.get.summary"),
	Long:  TR(`users.get_user_auth_key.get.description`),
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

		param, err := collectUsersAuthKeysGetCmdParams()
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
