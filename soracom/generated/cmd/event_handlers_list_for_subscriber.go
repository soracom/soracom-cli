package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersListForSubscriberCmdImsi holds value of 'imsi' option
var EventHandlersListForSubscriberCmdImsi string

func init() {
	EventHandlersListForSubscriberCmd.Flags().StringVar(&EventHandlersListForSubscriberCmdImsi, "imsi", "", TRAPI("imsi"))

	EventHandlersCmd.AddCommand(EventHandlersListForSubscriberCmd)
}

// EventHandlersListForSubscriberCmd defines 'list-for-subscriber' subcommand
var EventHandlersListForSubscriberCmd = &cobra.Command{
	Use:   "list-for-subscriber",
	Short: TRAPI("/event_handlers/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/event_handlers/subscribers/{imsi}:get:description`),
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

		param, err := collectEventHandlersListForSubscriberCmdParams(ac)
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

func collectEventHandlersListForSubscriberCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForEventHandlersListForSubscriberCmd("/event_handlers/subscribers/{imsi}"),
		query:  buildQueryForEventHandlersListForSubscriberCmd(),
	}, nil
}

func buildPathForEventHandlersListForSubscriberCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", EventHandlersListForSubscriberCmdImsi, -1)

	return path
}

func buildQueryForEventHandlersListForSubscriberCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
