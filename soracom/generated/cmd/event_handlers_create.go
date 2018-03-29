package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// EventHandlersCreateCmdDescription holds value of 'description' option
var EventHandlersCreateCmdDescription string

// EventHandlersCreateCmdName holds value of 'name' option
var EventHandlersCreateCmdName string

// EventHandlersCreateCmdStatus holds value of 'status' option
var EventHandlersCreateCmdStatus string

// EventHandlersCreateCmdTargetGroupId holds value of 'targetGroupId' option
var EventHandlersCreateCmdTargetGroupId string

// EventHandlersCreateCmdTargetImsi holds value of 'targetImsi' option
var EventHandlersCreateCmdTargetImsi string

// EventHandlersCreateCmdTargetOperatorId holds value of 'targetOperatorId' option
var EventHandlersCreateCmdTargetOperatorId string

// EventHandlersCreateCmdBody holds contents of request body to be sent
var EventHandlersCreateCmdBody string

func init() {
	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdDescription, "description", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdName, "name", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdStatus, "status", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetGroupId, "target-group-id", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetImsi, "target-imsi", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetOperatorId, "target-operator-id", "", TRAPI(""))

	EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	EventHandlersCmd.AddCommand(EventHandlersCreateCmd)
}

// EventHandlersCreateCmd defines 'create' subcommand
var EventHandlersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/event_handlers:post:summary"),
	Long:  TRAPI(`/event_handlers:post:description`),
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

		param, err := collectEventHandlersCreateCmdParams()
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

func collectEventHandlersCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForEventHandlersCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForEventHandlersCreateCmd("/event_handlers"),
		query:       buildQueryForEventHandlersCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForEventHandlersCreateCmd(path string) string {

	return path
}

func buildQueryForEventHandlersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForEventHandlersCreateCmd() (string, error) {
	if EventHandlersCreateCmdBody != "" {
		if strings.HasPrefix(EventHandlersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(EventHandlersCreateCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if EventHandlersCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return EventHandlersCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if EventHandlersCreateCmdDescription != "" {
		result["description"] = EventHandlersCreateCmdDescription
	}

	if EventHandlersCreateCmdName != "" {
		result["name"] = EventHandlersCreateCmdName
	}

	if EventHandlersCreateCmdStatus != "" {
		result["status"] = EventHandlersCreateCmdStatus
	}

	if EventHandlersCreateCmdTargetGroupId != "" {
		result["targetGroupId"] = EventHandlersCreateCmdTargetGroupId
	}

	if EventHandlersCreateCmdTargetImsi != "" {
		result["targetImsi"] = EventHandlersCreateCmdTargetImsi
	}

	if EventHandlersCreateCmdTargetOperatorId != "" {
		result["targetOperatorId"] = EventHandlersCreateCmdTargetOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
