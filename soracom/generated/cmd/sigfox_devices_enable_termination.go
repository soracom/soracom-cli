// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SigfoxDevicesEnableTerminationCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesEnableTerminationCmdDeviceId string

func InitSigfoxDevicesEnableTerminationCmd() {
	SigfoxDevicesEnableTerminationCmd.Flags().StringVar(&SigfoxDevicesEnableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesEnableTerminationCmd.RunE = SigfoxDevicesEnableTerminationCmdRunE

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesEnableTerminationCmd)
}

// SigfoxDevicesEnableTerminationCmd defines 'enable-termination' subcommand
var SigfoxDevicesEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/sigfox_devices/{device_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/enable_termination:post:description`) + "\n\n" + createLinkToAPIReference("SigfoxDevice", "enableTerminationOnSigfoxDevice"),
}

func SigfoxDevicesEnableTerminationCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSigfoxDevicesEnableTerminationCmdParams(ac)
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

func collectSigfoxDevicesEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SigfoxDevicesEnableTerminationCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesEnableTerminationCmd("/sigfox_devices/{device_id}/enable_termination"),
		query:  buildQueryForSigfoxDevicesEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesEnableTerminationCmd(path string) string {

	escapedDeviceId := url.PathEscape(SigfoxDevicesEnableTerminationCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
