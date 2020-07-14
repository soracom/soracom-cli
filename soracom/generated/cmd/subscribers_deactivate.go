// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

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

		param, err := collectSubscribersDeactivateCmdParams(ac)
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

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersDeactivateCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersDeactivateCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersDeactivateCmd("/subscribers/{imsi}/deactivate"),
		query:  buildQueryForSubscribersDeactivateCmd(),
	}, nil
}

func buildPathForSubscribersDeactivateCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersDeactivateCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersDeactivateCmd() url.Values {
	result := url.Values{}

	return result
}
