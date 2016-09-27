package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersGetCmdHandlerId holds value of 'handler_id' option
var EventHandlersGetCmdHandlerId string

func init() {
	EventHandlersGetCmd.Flags().StringVar(&EventHandlersGetCmdHandlerId, "handler-id", "", TR("event_handlers.get_event_handler.get.parameters.handler_id.description"))

	EventHandlersCmd.AddCommand(EventHandlersGetCmd)
}

// EventHandlersGetCmd defines 'get' subcommand
var EventHandlersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("event_handlers.get_event_handler.get.summary"),
	Long:  TR(`event_handlers.get_event_handler.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectEventHandlersGetCmdParams()
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

func collectEventHandlersGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForEventHandlersGetCmd("/event_handlers/{handler_id}"),
		query:  buildQueryForEventHandlersGetCmd(),
	}, nil
}

func buildPathForEventHandlersGetCmd(path string) string {

	path = strings.Replace(path, "{"+"handler_id"+"}", EventHandlersGetCmdHandlerId, -1)

	return path
}

func buildQueryForEventHandlersGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
