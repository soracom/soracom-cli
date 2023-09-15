// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EventHandlersDeleteCmdHandlerId holds value of 'handler_id' option
var EventHandlersDeleteCmdHandlerId string

func InitEventHandlersDeleteCmd() {
	EventHandlersDeleteCmd.Flags().StringVar(&EventHandlersDeleteCmdHandlerId, "handler-id", "", TRAPI("handler ID"))

	EventHandlersDeleteCmd.RunE = EventHandlersDeleteCmdRunE

	EventHandlersCmd.AddCommand(EventHandlersDeleteCmd)
}

// EventHandlersDeleteCmd defines 'delete' subcommand
var EventHandlersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/event_handlers/{handler_id}:delete:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}:delete:description`) + "\n\n" + createLinkToAPIReference("EventHandler", "deleteEventHandler"),
}

func EventHandlersDeleteCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
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

	param, err := collectEventHandlersDeleteCmdParams(ac)
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
}

func collectEventHandlersDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("handler_id", "handler-id", "path", parsedBody, EventHandlersDeleteCmdHandlerId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForEventHandlersDeleteCmd("/event_handlers/{handler_id}"),
		query:  buildQueryForEventHandlersDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEventHandlersDeleteCmd(path string) string {

	escapedHandlerId := url.PathEscape(EventHandlersDeleteCmdHandlerId)

	path = strReplace(path, "{"+"handler_id"+"}", escapedHandlerId, -1)

	return path
}

func buildQueryForEventHandlersDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
