// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
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

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
		}

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

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectUsersPasswordDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersPasswordDeleteCmdOperatorId == "" {
		UsersPasswordDeleteCmdOperatorId = ac.OperatorID
	}

	if UsersPasswordDeleteCmdUserName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "user-name")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForUsersPasswordDeleteCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:  buildQueryForUsersPasswordDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersPasswordDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersPasswordDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersPasswordDeleteCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersPasswordDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
