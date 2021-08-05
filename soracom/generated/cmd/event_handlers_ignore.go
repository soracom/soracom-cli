// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EventHandlersIgnoreCmdHandlerId holds value of 'handler_id' option
var EventHandlersIgnoreCmdHandlerId string

// EventHandlersIgnoreCmdImsi holds value of 'imsi' option
var EventHandlersIgnoreCmdImsi string

func init() {
	EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdHandlerId, "handler-id", "", TRAPI("handler_id"))

	EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdImsi, "imsi", "", TRAPI("imsi"))
	EventHandlersCmd.AddCommand(EventHandlersIgnoreCmd)
}

// EventHandlersIgnoreCmd defines 'ignore' subcommand
var EventHandlersIgnoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: TRAPI("/event_handlers/{handler_id}/subscribers/{imsi}/ignore:post:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}/subscribers/{imsi}/ignore:post:description`),
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

		param, err := collectEventHandlersIgnoreCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectEventHandlersIgnoreCmdParams(ac *apiClient) (*apiParams, error) {
	if EventHandlersIgnoreCmdHandlerId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "handler-id")
	}

	if EventHandlersIgnoreCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForEventHandlersIgnoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
		query:  buildQueryForEventHandlersIgnoreCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEventHandlersIgnoreCmd(path string) string {

	escapedHandlerId := url.PathEscape(EventHandlersIgnoreCmdHandlerId)

	path = strReplace(path, "{"+"handler_id"+"}", escapedHandlerId, -1)

	escapedImsi := url.PathEscape(EventHandlersIgnoreCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForEventHandlersIgnoreCmd() url.Values {
	result := url.Values{}

	return result
}
