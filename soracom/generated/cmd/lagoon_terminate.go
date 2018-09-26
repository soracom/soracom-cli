package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	LagoonCmd.AddCommand(LagoonTerminateCmd)
}

// LagoonTerminateCmd defines 'terminate' subcommand
var LagoonTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/lagoon/terminate:post:summary"),
	Long:  TRAPI(`/lagoon/terminate:post:description`),
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

		param, err := collectLagoonTerminateCmdParams(ac)
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

func collectLagoonTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLagoonTerminateCmd("/lagoon/terminate"),
		query:  buildQueryForLagoonTerminateCmd(),
	}, nil
}

func buildPathForLagoonTerminateCmd(path string) string {

	return path
}

func buildQueryForLagoonTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
