package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesEnableTerminationCmdDeviceId holds value of 'device_id' option
var LoraDevicesEnableTerminationCmdDeviceId string

func init() {
	LoraDevicesEnableTerminationCmd.Flags().StringVar(&LoraDevicesEnableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesCmd.AddCommand(LoraDevicesEnableTerminationCmd)
}

// LoraDevicesEnableTerminationCmd defines 'enable-termination' subcommand
var LoraDevicesEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/lora_devices/{device_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/enable_termination:post:description`),
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

		param, err := collectLoraDevicesEnableTerminationCmdParams()
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

func collectLoraDevicesEnableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesEnableTerminationCmd("/lora_devices/{device_id}/enable_termination"),
		query:  buildQueryForLoraDevicesEnableTerminationCmd(),
	}, nil
}

func buildPathForLoraDevicesEnableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesEnableTerminationCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesEnableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
