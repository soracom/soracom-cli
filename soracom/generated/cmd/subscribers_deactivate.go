package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersDeactivateCmdImsi holds value of 'imsi' option
var SubscribersDeactivateCmdImsi string

func init() {
	SubscribersDeactivateCmd.Flags().StringVar(&SubscribersDeactivateCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersDeactivateCmd)
}

// SubscribersDeactivateCmd defines 'deactivate' subcommand
var SubscribersDeactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: TRAPI("/subscribers/{imsi}/deactivate:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/deactivate:post:description`),
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

		param, err := collectSubscribersDeactivateCmdParams()
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

func collectSubscribersDeactivateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersDeactivateCmd("/subscribers/{imsi}/deactivate"),
		query:  buildQueryForSubscribersDeactivateCmd(),
	}, nil
}

func buildPathForSubscribersDeactivateCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersDeactivateCmdImsi, -1)

	return path
}

func buildQueryForSubscribersDeactivateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
