// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EventHandlersListForSubscriberCmdImsi holds value of 'imsi' option
var EventHandlersListForSubscriberCmdImsi string

// EventHandlersListForSubscriberCmdOutputJSONL indicates to output with jsonl format
var EventHandlersListForSubscriberCmdOutputJSONL bool

func InitEventHandlersListForSubscriberCmd() {
	EventHandlersListForSubscriberCmd.Flags().StringVar(&EventHandlersListForSubscriberCmdImsi, "imsi", "", TRAPI("IMSI."))

	EventHandlersListForSubscriberCmd.Flags().BoolVar(&EventHandlersListForSubscriberCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	EventHandlersListForSubscriberCmd.RunE = EventHandlersListForSubscriberCmdRunE

	EventHandlersCmd.AddCommand(EventHandlersListForSubscriberCmd)
}

// EventHandlersListForSubscriberCmd defines 'list-for-subscriber' subcommand
var EventHandlersListForSubscriberCmd = &cobra.Command{
	Use:   "list-for-subscriber",
	Short: TRAPI("/event_handlers/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/event_handlers/subscribers/{imsi}:get:description`) + "\n\n" + createLinkToAPIReference("EventHandler", "listEventHandlersBySubscriber"),
}

func EventHandlersListForSubscriberCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectEventHandlersListForSubscriberCmdParams(ac)
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
		if EventHandlersListForSubscriberCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectEventHandlersListForSubscriberCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, EventHandlersListForSubscriberCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForEventHandlersListForSubscriberCmd("/event_handlers/subscribers/{imsi}"),
		query:  buildQueryForEventHandlersListForSubscriberCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEventHandlersListForSubscriberCmd(path string) string {

	escapedImsi := url.PathEscape(EventHandlersListForSubscriberCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForEventHandlersListForSubscriberCmd() url.Values {
	result := url.Values{}

	return result
}
