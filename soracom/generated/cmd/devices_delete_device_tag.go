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
	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdDeviceId, "device-id", "", TR("devices.delete_tag.parameters.deviceId.description"))

	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdTagName, "tag-name", "", TR("devices.delete_tag.parameters.tagName.description"))

	DevicesCmd.AddCommand(DevicesDeleteDeviceTagCmd)
}

// DevicesDeleteDeviceTagCmd defines 'delete-device-tag' subcommand
var DevicesDeleteDeviceTagCmd = &cobra.Command{
	Use:   "delete-device-tag",
	Short: TR("devices.delete_tag.summary"),
	Long:  TR(`devices.delete_tag.description`),
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

		param, err := collectDevicesDeleteDeviceTagCmdParams()
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

func collectDevicesDeleteDeviceTagCmdParams() (*apiParams, error) {

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
