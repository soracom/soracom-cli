// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SystemNotificationsDeleteCmdOperatorId holds value of 'operator_id' option
var SystemNotificationsDeleteCmdOperatorId string

// SystemNotificationsDeleteCmdType holds value of 'type' option
var SystemNotificationsDeleteCmdType string

func init() {
	SystemNotificationsDeleteCmd.Flags().StringVar(&SystemNotificationsDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	SystemNotificationsDeleteCmd.Flags().StringVar(&SystemNotificationsDeleteCmdType, "type", "", TRAPI("system notification type"))
	SystemNotificationsCmd.AddCommand(SystemNotificationsDeleteCmd)
}

// SystemNotificationsDeleteCmd defines 'delete' subcommand
var SystemNotificationsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/system_notifications/{type}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/system_notifications/{type}:delete:description`),
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

		param, err := collectSystemNotificationsDeleteCmdParams(ac)
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
	},
}

func collectSystemNotificationsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if SystemNotificationsDeleteCmdOperatorId == "" {
		SystemNotificationsDeleteCmdOperatorId = ac.OperatorID
	}

	if SystemNotificationsDeleteCmdType == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "type")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSystemNotificationsDeleteCmd("/operators/{operator_id}/system_notifications/{type}"),
		query:  buildQueryForSystemNotificationsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSystemNotificationsDeleteCmd(path string) string {

	escapedOperatorId := url.PathEscape(SystemNotificationsDeleteCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedType := url.PathEscape(SystemNotificationsDeleteCmdType)

	path = strReplace(path, "{"+"type"+"}", escapedType, -1)

	return path
}

func buildQueryForSystemNotificationsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}