package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersTerminateCmdImsi holds value of 'imsi' option
var SubscribersTerminateCmdImsi string

func init() {
	SubscribersTerminateCmd.Flags().StringVar(&SubscribersTerminateCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersTerminateCmd)
}

// SubscribersTerminateCmd defines 'terminate' subcommand
var SubscribersTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/subscribers/{imsi}/terminate:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/terminate:post:description`),
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

		param, err := collectSubscribersTerminateCmdParams(ac)
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

func collectSubscribersTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersTerminateCmd("/subscribers/{imsi}/terminate"),
		query:  buildQueryForSubscribersTerminateCmd(),
	}, nil
}

func buildPathForSubscribersTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersTerminateCmdImsi, -1)

	return path
}

func buildQueryForSubscribersTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
