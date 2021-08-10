// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

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

		param, err := collectLoraDevicesDisableTerminationCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLoraDevicesDisableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if LoraDevicesDisableTerminationCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesDisableTerminationCmd("/lora_devices/{device_id}/disable_termination"),
		query:  buildQueryForLoraDevicesDisableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesDisableTerminationCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesDisableTerminationCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesDisableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
