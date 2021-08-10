// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
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

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
		}

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

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersTerminateCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersTerminateCmd("/subscribers/{imsi}/terminate"),
		query:  buildQueryForSubscribersTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersTerminateCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersTerminateCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersTerminateCmd() url.Values {
	result := url.Values{}

	return result
}
