// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdDeviceId string

func InitSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd() {
	SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd.RunE = SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd)
}

// SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd defines 'get-motion-sensitivity' subcommand
var SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd = &cobra.Command{
	Use:   "get-motion-sensitivity",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/motion_sensitivity:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/motion_sensitivity:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceAtomCamSettingsMotionSensitivity"),
}

func SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd("/sora_cam/devices/{device_id}/atom_cam/settings/motion_sensitivity"),
		query:  buildQueryForSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsGetMotionSensitivityCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd() url.Values {
	result := url.Values{}

	return result
}
