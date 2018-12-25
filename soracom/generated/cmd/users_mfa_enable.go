package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersMfaEnableCmdOperatorId holds value of 'operator_id' option
var UsersMfaEnableCmdOperatorId string

// UsersMfaEnableCmdUserName holds value of 'user_name' option
var UsersMfaEnableCmdUserName string

func init() {
	UsersMfaEnableCmd.Flags().StringVar(&UsersMfaEnableCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaEnableCmd.Flags().StringVar(&UsersMfaEnableCmdUserName, "user-name", "", TRAPI("SAM user name"))

	UsersMfaCmd.AddCommand(UsersMfaEnableCmd)
}

// UsersMfaEnableCmd defines 'enable' subcommand
var UsersMfaEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersMfaEnableCmdParams(ac)
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

func collectUsersMfaEnableCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersMfaEnableCmdOperatorId == "" {
		UsersMfaEnableCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForUsersMfaEnableCmd("/operators/{operator_id}/users/{user_name}/mfa"),
		query:  buildQueryForUsersMfaEnableCmd(),
	}, nil
}

func buildPathForUsersMfaEnableCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersMfaEnableCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersMfaEnableCmdUserName, -1)

	return path
}

func buildQueryForUsersMfaEnableCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
