// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SystemNotificationsGetCmdOperatorId holds value of 'operator_id' option
var SystemNotificationsGetCmdOperatorId string

// SystemNotificationsGetCmdType holds value of 'type' option
var SystemNotificationsGetCmdType string

func InitSystemNotificationsGetCmd() {
	SystemNotificationsGetCmd.Flags().StringVar(&SystemNotificationsGetCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	SystemNotificationsGetCmd.Flags().StringVar(&SystemNotificationsGetCmdType, "type", "", TRAPI("Email address type."))

	SystemNotificationsGetCmd.RunE = SystemNotificationsGetCmdRunE

	SystemNotificationsCmd.AddCommand(SystemNotificationsGetCmd)
}

// SystemNotificationsGetCmd defines 'get' subcommand
var SystemNotificationsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/system_notifications/{type}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/system_notifications/{type}:get:description`) + "\n\n" + createLinkToAPIReference("SystemNotification", "getSystemNotification"),
}

func SystemNotificationsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSystemNotificationsGetCmdParams(ac)
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

func collectSystemNotificationsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if SystemNotificationsGetCmdOperatorId == "" {
		SystemNotificationsGetCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("type", "type", "path", parsedBody, SystemNotificationsGetCmdType)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSystemNotificationsGetCmd("/operators/{operator_id}/system_notifications/{type}"),
		query:  buildQueryForSystemNotificationsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSystemNotificationsGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(SystemNotificationsGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedType := url.PathEscape(SystemNotificationsGetCmdType)

	path = strReplace(path, "{"+"type"+"}", escapedType, -1)

	return path
}

func buildQueryForSystemNotificationsGetCmd() url.Values {
	result := url.Values{}

	return result
}
