// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomCamSettingsGetStatusLightCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsGetStatusLightCmdDeviceId string

func InitSoraCamDevicesAtomCamSettingsGetStatusLightCmd() {
	SoraCamDevicesAtomCamSettingsGetStatusLightCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsGetStatusLightCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsGetStatusLightCmd.RunE = SoraCamDevicesAtomCamSettingsGetStatusLightCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsGetStatusLightCmd)
}

// SoraCamDevicesAtomCamSettingsGetStatusLightCmd defines 'get-status-light' subcommand
var SoraCamDevicesAtomCamSettingsGetStatusLightCmd = &cobra.Command{
	Use:   "get-status-light",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/status_light:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/status_light:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettingsStatusLight"),
}

func SoraCamDevicesAtomCamSettingsGetStatusLightCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectSoraCamDevicesAtomCamSettingsGetStatusLightCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsGetStatusLightCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsGetStatusLightCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomCamSettingsGetStatusLightCmd("/sora_cam/devices/{device_id}/atom_cam/settings/status_light"),
		query:  buildQueryForSoraCamDevicesAtomCamSettingsGetStatusLightCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsGetStatusLightCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsGetStatusLightCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsGetStatusLightCmd() url.Values {
	result := url.Values{}

	return result
}
