// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersGetCmdImsi holds value of 'imsi' option
var SubscribersGetCmdImsi string

func init() {
	SubscribersGetCmd.Flags().StringVar(&SubscribersGetCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))
	SubscribersCmd.AddCommand(SubscribersGetCmd)
}

// SubscribersGetCmd defines 'get' subcommand
var SubscribersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}:get:description`),
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

		param, err := collectSubscribersGetCmdParams(ac)
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

func collectSubscribersGetCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersGetCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersGetCmd("/subscribers/{imsi}"),
		query:  buildQueryForSubscribersGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersGetCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersGetCmd() url.Values {
	result := url.Values{}

	return result
}
