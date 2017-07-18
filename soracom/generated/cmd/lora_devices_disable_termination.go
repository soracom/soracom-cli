package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesDisableTerminationCmdDeviceId holds value of 'device_id' option
var LoraDevicesDisableTerminationCmdDeviceId string

func init() {
	LoraDevicesDisableTerminationCmd.Flags().StringVar(&LoraDevicesDisableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesCmd.AddCommand(LoraDevicesDisableTerminationCmd)
}

// LoraDevicesDisableTerminationCmd defines 'disable-termination' subcommand
var LoraDevicesDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TRAPI("/lora_devices/{device_id}/disable_termination:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/disable_termination:post:description`),
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

		param, err := collectLoraDevicesDisableTerminationCmdParams()
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

func collectLoraDevicesDisableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesDisableTerminationCmd("/lora_devices/{device_id}/disable_termination"),
		query:  buildQueryForLoraDevicesDisableTerminationCmd(),
	}, nil
}

func buildPathForLoraDevicesDisableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesDisableTerminationCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesDisableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
