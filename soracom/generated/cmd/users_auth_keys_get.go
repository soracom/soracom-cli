// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

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

		param, err := collectUsersAuthKeysGetCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectUsersAuthKeysGetCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersAuthKeysGetCmdOperatorId == "" {
		UsersAuthKeysGetCmdOperatorId = ac.OperatorID
	}

	if UsersAuthKeysGetCmdAuthKeyId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "auth-key-id")
	}

	if UsersAuthKeysGetCmdUserName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "user-name")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersAuthKeysGetCmd("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}"),
		query:  buildQueryForUsersAuthKeysGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersAuthKeysGetCmd(path string) string {

	escapedAuthKeyId := url.PathEscape(UsersAuthKeysGetCmdAuthKeyId)

	path = strReplace(path, "{"+"auth_key_id"+"}", escapedAuthKeyId, -1)

	escapedOperatorId := url.PathEscape(UsersAuthKeysGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersAuthKeysGetCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysGetCmd() url.Values {
	result := url.Values{}

	return result
}
