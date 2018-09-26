package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersAuthKeysDeleteCmdAuthKeyId holds value of 'auth_key_id' option
var UsersAuthKeysDeleteCmdAuthKeyId string

// UsersAuthKeysDeleteCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysDeleteCmdOperatorId string

// UsersAuthKeysDeleteCmdUserName holds value of 'user_name' option
var UsersAuthKeysDeleteCmdUserName string

func init() {
	UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdAuthKeyId, "auth-key-id", "", TRAPI("auth_key_id"))

	UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAuthKeysCmd.AddCommand(UsersAuthKeysDeleteCmd)
}

// UsersAuthKeysDeleteCmd defines 'delete' subcommand
var UsersAuthKeysDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}:delete:description`),
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

		param, err := collectUsersAuthKeysDeleteCmdParams(ac)
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

func collectUsersAuthKeysDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersAuthKeysDeleteCmdOperatorId == "" {
		UsersAuthKeysDeleteCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersAuthKeysDeleteCmd("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}"),
		query:  buildQueryForUsersAuthKeysDeleteCmd(),
	}, nil
}

func buildPathForUsersAuthKeysDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"auth_key_id"+"}", UsersAuthKeysDeleteCmdAuthKeyId, -1)

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersAuthKeysDeleteCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersAuthKeysDeleteCmdUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
