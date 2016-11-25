package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersListCmdTarget holds value of 'target' option
var EventHandlersListCmdTarget string

func init() {
	EventHandlersListCmd.Flags().StringVar(&EventHandlersListCmdTarget, "target", "", TR("event_handlers.list_event_handlers.get.parameters.target.description"))

	EventHandlersCmd.AddCommand(EventHandlersListCmd)
}

// EventHandlersListCmd defines 'list' subcommand
var EventHandlersListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("event_handlers.list_event_handlers.get.summary"),
	Long:  TR(`event_handlers.list_event_handlers.get.description`),
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

		param, err := collectEventHandlersListCmdParams()
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

func collectEventHandlersListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForEventHandlersListCmd("/event_handlers"),
		query:  buildQueryForEventHandlersListCmd(),
	}, nil
}

func buildPathForEventHandlersListCmd(path string) string {

	return path
}

func buildQueryForEventHandlersListCmd() string {
	result := []string{}

	if EventHandlersListCmdTarget != "" {
		result = append(result, sprintf("%s=%s", "target", EventHandlersListCmdTarget))
	}

	return strings.Join(result, "&")
}
