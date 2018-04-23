package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesPutDeviceTagsCmdDeviceId holds value of 'deviceId' option
var DevicesPutDeviceTagsCmdDeviceId string

func init() {
	DevicesPutDeviceTagsCmd.Flags().StringVar(&DevicesPutDeviceTagsCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesCmd.AddCommand(DevicesPutDeviceTagsCmd)
}

// DevicesPutDeviceTagsCmd defines 'put-device-tags' subcommand
var DevicesPutDeviceTagsCmd = &cobra.Command{
	Use:   "put-device-tags",
	Short: TRAPI("/devices/{deviceId}/tags:put:summary"),
	Long:  TRAPI(`/devices/{deviceId}/tags:put:description`),
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

		param, err := collectDevicesPutDeviceTagsCmdParams(ac)
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

func collectDevicesPutDeviceTagsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "PUT",
		path:   buildPathForDevicesPutDeviceTagsCmd("/devices/{deviceId}/tags"),
		query:  buildQueryForDevicesPutDeviceTagsCmd(),
	}, nil
}

func buildPathForDevicesPutDeviceTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesPutDeviceTagsCmdDeviceId, -1)

	return path
}

func buildQueryForDevicesPutDeviceTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
