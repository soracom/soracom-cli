package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersDeleteSessionCmdImsi holds value of 'imsi' option
var SubscribersDeleteSessionCmdImsi string

func init() {
	SubscribersDeleteSessionCmd.Flags().StringVar(&SubscribersDeleteSessionCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersDeleteSessionCmd)
}

// SubscribersDeleteSessionCmd defines 'delete-session' subcommand
var SubscribersDeleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: TRAPI("/subscribers/{imsi}/delete_session:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/delete_session:post:description`),
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

		param, err := collectSubscribersDeleteSessionCmdParams()
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

func collectSubscribersDeleteSessionCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersDeleteSessionCmd("/subscribers/{imsi}/delete_session"),
		query:  buildQueryForSubscribersDeleteSessionCmd(),
	}, nil
}

func buildPathForSubscribersDeleteSessionCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersDeleteSessionCmdImsi, -1)

	return path
}

func buildQueryForSubscribersDeleteSessionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
