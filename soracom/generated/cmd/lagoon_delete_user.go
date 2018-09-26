package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonDeleteUserCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonDeleteUserCmdLagoonUserId int64

func init() {
	LagoonDeleteUserCmd.Flags().Int64Var(&LagoonDeleteUserCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonCmd.AddCommand(LagoonDeleteUserCmd)
}

// LagoonDeleteUserCmd defines 'delete-user' subcommand
var LagoonDeleteUserCmd = &cobra.Command{
	Use:   "delete-user",
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

		param, err := collectLagoonDeleteUserCmdParams(ac)
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

func collectLagoonDeleteUserCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLagoonDeleteUserCmd("/lagoon/users/{lagoon_user_id}"),
		query:  buildQueryForLagoonDeleteUserCmd(),
	}, nil
}

func buildPathForLagoonDeleteUserCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonDeleteUserCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonDeleteUserCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
