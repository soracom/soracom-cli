// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomcamGetSettingsCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomcamGetSettingsCmdDeviceId string

func InitSoraCamDevicesAtomcamGetSettingsCmd() {
	SoraCamDevicesAtomcamGetSettingsCmd.Flags().StringVar(&SoraCamDevicesAtomcamGetSettingsCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesAtomcamGetSettingsCmd.RunE = SoraCamDevicesAtomcamGetSettingsCmdRunE

	SoraCamDevicesAtomcamCmd.AddCommand(SoraCamDevicesAtomcamGetSettingsCmd)
}

// SoraCamDevicesAtomcamGetSettingsCmd defines 'get-settings' subcommand
var SoraCamDevicesAtomcamGetSettingsCmd = &cobra.Command{
	Use:   "get-settings",
	Short: TRAPI("/sora_cam/devices/{device_id}/atomcam/settings:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atomcam/settings:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettings"),
}

func SoraCamDevicesAtomcamGetSettingsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomcamGetSettingsCmdParams(ac)
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

func collectSoraCamDevicesAtomcamGetSettingsCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomcamGetSettingsCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomcamGetSettingsCmd("/sora_cam/devices/{device_id}/atomcam/settings"),
		query:  buildQueryForSoraCamDevicesAtomcamGetSettingsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomcamGetSettingsCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomcamGetSettingsCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomcamGetSettingsCmd() url.Values {
	result := url.Values{}

	return result
}
