// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SystemNotificationsListCmdOperatorId holds value of 'operator_id' option
var SystemNotificationsListCmdOperatorId string

func init() {
	SystemNotificationsListCmd.Flags().StringVar(&SystemNotificationsListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	SystemNotificationsCmd.AddCommand(SystemNotificationsListCmd)
}

// SystemNotificationsListCmd defines 'list' subcommand
var SystemNotificationsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/system_notifications:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/system_notifications:get:description`),
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

		param, err := collectSystemNotificationsListCmdParams(ac)
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

func collectSystemNotificationsListCmdParams(ac *apiClient) (*apiParams, error) {
	if SystemNotificationsListCmdOperatorId == "" {
		SystemNotificationsListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSystemNotificationsListCmd("/operators/{operator_id}/system_notifications"),
		query:  buildQueryForSystemNotificationsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSystemNotificationsListCmd(path string) string {

	escapedOperatorId := url.PathEscape(SystemNotificationsListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForSystemNotificationsListCmd() url.Values {
	result := url.Values{}

	return result
}
