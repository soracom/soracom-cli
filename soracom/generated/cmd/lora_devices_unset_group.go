package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesUnsetGroupCmdDeviceId holds value of 'device_id' option
var LoraDevicesUnsetGroupCmdDeviceId string

func init() {
	LoraDevicesUnsetGroupCmd.Flags().StringVar(&LoraDevicesUnsetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesCmd.AddCommand(LoraDevicesUnsetGroupCmd)
}

// LoraDevicesUnsetGroupCmd defines 'unset-group' subcommand
var LoraDevicesUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/lora_devices/{device_id}/unset_group:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/unset_group:post:description`),
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

		param, err := collectLoraDevicesUnsetGroupCmdParams(ac)
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

func collectLoraDevicesUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesUnsetGroupCmd("/lora_devices/{device_id}/unset_group"),
		query:  buildQueryForLoraDevicesUnsetGroupCmd(),
	}, nil
}

func buildPathForLoraDevicesUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesUnsetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
