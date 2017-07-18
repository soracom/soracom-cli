package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesUnsetGroupCmdDeviceId holds value of 'deviceId' option
var DevicesUnsetGroupCmdDeviceId string

func init() {
	DevicesUnsetGroupCmd.Flags().StringVar(&DevicesUnsetGroupCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesCmd.AddCommand(DevicesUnsetGroupCmd)
}

// DevicesUnsetGroupCmd defines 'unset-group' subcommand
var DevicesUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/devices/{deviceId}/unset_group:post:summary"),
	Long:  TRAPI(`/devices/{deviceId}/unset_group:post:description`),
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

		param, err := collectDevicesUnsetGroupCmdParams()
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

func collectDevicesUnsetGroupCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnsetGroupCmd("/devices/{deviceId}/unset_group"),
		query:  buildQueryForDevicesUnsetGroupCmd(),
	}, nil
}

func buildPathForDevicesUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesUnsetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForDevicesUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
