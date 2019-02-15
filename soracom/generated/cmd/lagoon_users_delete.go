package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUsersDeleteCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersDeleteCmdLagoonUserId int64

func init() {
	LagoonUsersDeleteCmd.Flags().Int64Var(&LagoonUsersDeleteCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersCmd.AddCommand(LagoonUsersDeleteCmd)
}

// LagoonUsersDeleteCmd defines 'delete' subcommand
var LagoonUsersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}:delete:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}:delete:description`),
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

		param, err := collectLagoonUsersDeleteCmdParams(ac)
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

func collectLagoonUsersDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLagoonUsersDeleteCmd("/lagoon/users/{lagoon_user_id}"),
		query:  buildQueryForLagoonUsersDeleteCmd(),
	}, nil
}

func buildPathForLagoonUsersDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUsersDeleteCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUsersDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
