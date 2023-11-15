// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesRebootCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesRebootCmdDeviceId string

func InitSoraCamDevicesRebootCmd() {
	SoraCamDevicesRebootCmd.Flags().StringVar(&SoraCamDevicesRebootCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesRebootCmd.RunE = SoraCamDevicesRebootCmdRunE

	SoraCamDevicesCmd.AddCommand(SoraCamDevicesRebootCmd)
}

// SoraCamDevicesRebootCmd defines 'reboot' subcommand
var SoraCamDevicesRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: TRAPI("/sora_cam/devices/{device_id}/reboot:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/reboot:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "rebootSoraCamDevice"),
}

func SoraCamDevicesRebootCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSoraCamDevicesRebootCmdParams(ac)
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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamDevicesRebootCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesRebootCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSoraCamDevicesRebootCmd("/sora_cam/devices/{device_id}/reboot"),
		query:  buildQueryForSoraCamDevicesRebootCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesRebootCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesRebootCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesRebootCmd() url.Values {
	result := url.Values{}

	return result
}