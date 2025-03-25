// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomCamSettingsGetSoundCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsGetSoundCmdDeviceId string

func InitSoraCamDevicesAtomCamSettingsGetSoundCmd() {
	SoraCamDevicesAtomCamSettingsGetSoundCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsGetSoundCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsGetSoundCmd.RunE = SoraCamDevicesAtomCamSettingsGetSoundCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsGetSoundCmd)
}

// SoraCamDevicesAtomCamSettingsGetSoundCmd defines 'get-sound' subcommand
var SoraCamDevicesAtomCamSettingsGetSoundCmd = &cobra.Command{
	Use:   "get-sound",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/sound:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/sound:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettingsSound"),
}

func SoraCamDevicesAtomCamSettingsGetSoundCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsGetSoundCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsGetSoundCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsGetSoundCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomCamSettingsGetSoundCmd("/sora_cam/devices/{device_id}/atom_cam/settings/sound"),
		query:  buildQueryForSoraCamDevicesAtomCamSettingsGetSoundCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsGetSoundCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsGetSoundCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsGetSoundCmd() url.Values {
	result := url.Values{}

	return result
}
