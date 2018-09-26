package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesGetCmdDeviceId holds value of 'device_id' option
var LoraDevicesGetCmdDeviceId string

func init() {
	LoraDevicesGetCmd.Flags().StringVar(&LoraDevicesGetCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesCmd.AddCommand(LoraDevicesGetCmd)
}

// LoraDevicesGetCmd defines 'get' subcommand
var LoraDevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/lora_devices/{device_id}:get:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}:get:description`),
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

		param, err := collectLoraDevicesGetCmdParams(ac)
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

func collectLoraDevicesGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraDevicesGetCmd("/lora_devices/{device_id}"),
		query:  buildQueryForLoraDevicesGetCmd(),
	}, nil
}

func buildPathForLoraDevicesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesGetCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
