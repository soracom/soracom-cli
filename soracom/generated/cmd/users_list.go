package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersListCmdOperatorId holds value of 'operator_id' option
var UsersListCmdOperatorId string

func init() {
	UsersListCmd.Flags().StringVar(&UsersListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersCmd.AddCommand(UsersListCmd)
}

// UsersListCmd defines 'list' subcommand
var UsersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/users:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users:get:description`),
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

		param, err := collectUsersListCmdParams(ac)
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

func collectUsersListCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersListCmdOperatorId == "" {
		UsersListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersListCmd("/operators/{operator_id}/users"),
		query:  buildQueryForUsersListCmd(),
	}, nil
}

func buildPathForUsersListCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersListCmdOperatorId, -1)

	return path
}

func buildQueryForUsersListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
