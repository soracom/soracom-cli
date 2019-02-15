package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	LagoonUsersCmd.AddCommand(LagoonUsersListCmd)
}

// LagoonUsersListCmd defines 'list' subcommand
var LagoonUsersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lagoon/users:get:summary"),
	Long:  TRAPI(`/lagoon/users:get:description`),
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

		param, err := collectLagoonUsersListCmdParams(ac)
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

func collectLagoonUsersListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLagoonUsersListCmd("/lagoon/users"),
		query:  buildQueryForLagoonUsersListCmd(),
	}, nil
}

func buildPathForLagoonUsersListCmd(path string) string {

	return path
}

func buildQueryForLagoonUsersListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
