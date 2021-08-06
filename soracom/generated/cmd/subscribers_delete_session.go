// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

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

		param, err := collectSubscribersDeleteSessionCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersDeleteSessionCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersDeleteSessionCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersDeleteSessionCmd("/subscribers/{imsi}/delete_session"),
		query:  buildQueryForSubscribersDeleteSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersDeleteSessionCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersDeleteSessionCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersDeleteSessionCmd() url.Values {
	result := url.Values{}

	return result
}
