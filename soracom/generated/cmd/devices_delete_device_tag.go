package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesDeleteDeviceTagCmdDeviceId holds value of 'deviceId' option
var DevicesDeleteDeviceTagCmdDeviceId string

// DevicesDeleteDeviceTagCmdTagName holds value of 'tagName' option
var DevicesDeleteDeviceTagCmdTagName string

func init() {
	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdTagName, "tag-name", "", TRAPI("Name of tag to delete"))

	DevicesCmd.AddCommand(DevicesDeleteDeviceTagCmd)
}

// DevicesDeleteDeviceTagCmd defines 'delete-device-tag' subcommand
var DevicesDeleteDeviceTagCmd = &cobra.Command{
	Use:   "delete-device-tag",
	Short: TRAPI("/devices/{deviceId}/tags/{tagName}:delete:summary"),
	Long:  TRAPI(`/devices/{deviceId}/tags/{tagName}:delete:description`),
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

		param, err := collectDevicesDeleteDeviceTagCmdParams(ac)
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

func collectDevicesDeleteDeviceTagCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteDeviceTagCmd("/devices/{deviceId}/tags/{tagName}"),
		query:  buildQueryForDevicesDeleteDeviceTagCmd(),
	}, nil
}

func buildPathForDevicesDeleteDeviceTagCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesDeleteDeviceTagCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"tagName"+"}", DevicesDeleteDeviceTagCmdTagName, -1)

	return path
}

func buildQueryForDevicesDeleteDeviceTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
