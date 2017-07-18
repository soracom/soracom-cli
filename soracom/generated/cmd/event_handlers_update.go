package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersUpdateCmdHandlerId holds value of 'handler_id' option
var EventHandlersUpdateCmdHandlerId string

// EventHandlersUpdateCmdBody holds contents of request body to be sent
var EventHandlersUpdateCmdBody string

func init() {
	EventHandlersUpdateCmd.Flags().StringVar(&EventHandlersUpdateCmdHandlerId, "handler-id", "", TRAPI("handler ID"))

	EventHandlersUpdateCmd.Flags().StringVar(&EventHandlersUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	EventHandlersCmd.AddCommand(EventHandlersUpdateCmd)
}

// EventHandlersUpdateCmd defines 'update' subcommand
var EventHandlersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/event_handlers/{handler_id}:put:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}:put:description`),
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

		param, err := collectEventHandlersUpdateCmdParams()
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

func collectEventHandlersUpdateCmdParams() (*apiParams, error) {

	body, err := buildBodyForEventHandlersUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForEventHandlersUpdateCmd("/event_handlers/{handler_id}"),
		query:       buildQueryForEventHandlersUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForEventHandlersUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"handler_id"+"}", EventHandlersUpdateCmdHandlerId, -1)

	return path
}

func buildQueryForEventHandlersUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForEventHandlersUpdateCmd() (string, error) {
	if EventHandlersUpdateCmdBody != "" {
		if strings.HasPrefix(EventHandlersUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(EventHandlersUpdateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if EventHandlersUpdateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return EventHandlersUpdateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
