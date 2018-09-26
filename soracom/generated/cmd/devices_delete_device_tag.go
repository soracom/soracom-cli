package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesDeleteDeviceTagCmdDeviceId holds value of 'device_id' option
var DevicesDeleteDeviceTagCmdDeviceId string

// DevicesDeleteDeviceTagCmdTagName holds value of 'tag_name' option
var DevicesDeleteDeviceTagCmdTagName string

func init() {
	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesDeleteDeviceTagCmd.Flags().StringVar(&DevicesDeleteDeviceTagCmdTagName, "tag-name", "", TRAPI("Name of tag to delete"))

	DevicesCmd.AddCommand(DevicesDeleteDeviceTagCmd)
}

// DevicesDeleteDeviceTagCmd defines 'delete-device-tag' subcommand
var DevicesDeleteDeviceTagCmd = &cobra.Command{
	Use:   "delete-device-tag",
	Short: TRAPI("/devices/{device_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/devices/{device_id}/tags/{tag_name}:delete:description`),
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

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectDevicesDeleteDeviceTagCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteDeviceTagCmd("/devices/{device_id}/tags/{tag_name}"),
		query:  buildQueryForDevicesDeleteDeviceTagCmd(),
	}, nil
}

func buildPathForDevicesDeleteDeviceTagCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesDeleteDeviceTagCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", DevicesDeleteDeviceTagCmdTagName, -1)

	return path
}

func buildQueryForDevicesDeleteDeviceTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
