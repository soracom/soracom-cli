package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersMfaGetCmdOperatorId holds value of 'operator_id' option
var UsersMfaGetCmdOperatorId string

// UsersMfaGetCmdUserName holds value of 'user_name' option
var UsersMfaGetCmdUserName string

func init() {
	UsersMfaGetCmd.Flags().StringVar(&UsersMfaGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaGetCmd.Flags().StringVar(&UsersMfaGetCmdUserName, "user-name", "", TRAPI("SAM user name"))

	UsersMfaCmd.AddCommand(UsersMfaGetCmd)
}

// UsersMfaGetCmd defines 'get' subcommand
var UsersMfaGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersMfaGetCmdParams(ac)
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

func collectUsersMfaGetCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersMfaGetCmdOperatorId == "" {
		UsersMfaGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersMfaGetCmd("/operators/{operator_id}/users/{user_name}/mfa"),
		query:  buildQueryForUsersMfaGetCmd(),
	}, nil
}

func buildPathForUsersMfaGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersMfaGetCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersMfaGetCmdUserName, -1)

	return path
}

func buildQueryForUsersMfaGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
