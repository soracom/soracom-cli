// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomcamSettingsGetNightVisionCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomcamSettingsGetNightVisionCmdDeviceId string

func InitSoraCamDevicesAtomcamSettingsGetNightVisionCmd() {
	SoraCamDevicesAtomcamSettingsGetNightVisionCmd.Flags().StringVar(&SoraCamDevicesAtomcamSettingsGetNightVisionCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesAtomcamSettingsGetNightVisionCmd.RunE = SoraCamDevicesAtomcamSettingsGetNightVisionCmdRunE

	SoraCamDevicesAtomcamSettingsCmd.AddCommand(SoraCamDevicesAtomcamSettingsGetNightVisionCmd)
}

// SoraCamDevicesAtomcamSettingsGetNightVisionCmd defines 'get-night-vision' subcommand
var SoraCamDevicesAtomcamSettingsGetNightVisionCmd = &cobra.Command{
	Use:   "get-night-vision",
	Short: TRAPI("/sora_cam/devices/{device_id}/atomcam/settings/night_vision:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atomcam/settings/night_vision:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettingsNightVision"),
}

func SoraCamDevicesAtomcamSettingsGetNightVisionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomcamSettingsGetNightVisionCmdParams(ac)
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

func collectSoraCamDevicesAtomcamSettingsGetNightVisionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomcamSettingsGetNightVisionCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomcamSettingsGetNightVisionCmd("/sora_cam/devices/{device_id}/atomcam/settings/night_vision"),
		query:  buildQueryForSoraCamDevicesAtomcamSettingsGetNightVisionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomcamSettingsGetNightVisionCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomcamSettingsGetNightVisionCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomcamSettingsGetNightVisionCmd() url.Values {
	result := url.Values{}

	return result
}
