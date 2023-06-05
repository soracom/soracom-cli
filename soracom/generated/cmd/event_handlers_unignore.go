// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EventHandlersUnignoreCmdHandlerId holds value of 'handler_id' option
var EventHandlersUnignoreCmdHandlerId string

// EventHandlersUnignoreCmdImsi holds value of 'imsi' option
var EventHandlersUnignoreCmdImsi string

func InitEventHandlersUnignoreCmd() {
	EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdHandlerId, "handler-id", "", TRAPI("handler_id"))

	EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdImsi, "imsi", "", TRAPI("imsi"))

	EventHandlersUnignoreCmd.RunE = EventHandlersUnignoreCmdRunE

	EventHandlersCmd.AddCommand(EventHandlersUnignoreCmd)
}

// EventHandlersUnignoreCmd defines 'unignore' subcommand
var EventHandlersUnignoreCmd = &cobra.Command{
	Use:   "unignore",
	Short: TRAPI("/event_handlers/{handler_id}/subscribers/{imsi}/ignore:delete:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}/subscribers/{imsi}/ignore:delete:description`) + "\n\n" + createLinkToAPIReference("EventHandler", "deleteIgnoreEventHandler"),
}

func EventHandlersUnignoreCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := authHelper(ac, cmd, args)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectEventHandlersUnignoreCmdParams(ac)
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

func collectEventHandlersUnignoreCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("handler_id", "handler-id", "path", parsedBody, EventHandlersUnignoreCmdHandlerId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, EventHandlersUnignoreCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForEventHandlersUnignoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
		query:  buildQueryForEventHandlersUnignoreCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForEventHandlersUnignoreCmd(path string) string {

	escapedHandlerId := url.PathEscape(EventHandlersUnignoreCmdHandlerId)

	path = strReplace(path, "{"+"handler_id"+"}", escapedHandlerId, -1)

	escapedImsi := url.PathEscape(EventHandlersUnignoreCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForEventHandlersUnignoreCmd() url.Values {
	result := url.Values{}

	return result
}
