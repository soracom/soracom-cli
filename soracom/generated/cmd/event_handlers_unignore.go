package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersUnignoreCmdHandlerId holds value of 'handler_id' option
var EventHandlersUnignoreCmdHandlerId string

// EventHandlersUnignoreCmdImsi holds value of 'imsi' option
var EventHandlersUnignoreCmdImsi string

func init() {
	EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdHandlerId, "handler-id", "", TRAPI("handler_id"))

	EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdImsi, "imsi", "", TRAPI("imsi"))

	EventHandlersCmd.AddCommand(EventHandlersUnignoreCmd)
}

// EventHandlersUnignoreCmd defines 'unignore' subcommand
var EventHandlersUnignoreCmd = &cobra.Command{
	Use:   "unignore",
	Short: TRAPI("/event_handlers/{handler_id}/subscribers/{imsi}/ignore:delete:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}/subscribers/{imsi}/ignore:delete:description`),
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

		param, err := collectEventHandlersUnignoreCmdParams()
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

func collectEventHandlersUnignoreCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForEventHandlersUnignoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
		query:  buildQueryForEventHandlersUnignoreCmd(),
	}, nil
}

func buildPathForEventHandlersUnignoreCmd(path string) string {

	path = strings.Replace(path, "{"+"handler_id"+"}", EventHandlersUnignoreCmdHandlerId, -1)

	path = strings.Replace(path, "{"+"imsi"+"}", EventHandlersUnignoreCmdImsi, -1)

	return path
}

func buildQueryForEventHandlersUnignoreCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
