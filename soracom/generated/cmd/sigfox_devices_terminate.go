// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SigfoxDevicesTerminateCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesTerminateCmdDeviceId string

// SigfoxDevicesTerminateCmdDeleteImmediately holds value of 'delete_immediately' option
var SigfoxDevicesTerminateCmdDeleteImmediately bool

func InitSigfoxDevicesTerminateCmd() {
	SigfoxDevicesTerminateCmd.Flags().StringVar(&SigfoxDevicesTerminateCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesTerminateCmd.Flags().BoolVar(&SigfoxDevicesTerminateCmdDeleteImmediately, "delete-immediately", false, TRAPI("If the Sigfox device is deleted immediately."))

	SigfoxDevicesTerminateCmd.RunE = SigfoxDevicesTerminateCmdRunE

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesTerminateCmd)
}

// SigfoxDevicesTerminateCmd defines 'terminate' subcommand
var SigfoxDevicesTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/sigfox_devices/{device_id}/terminate:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/terminate:post:description`) + "\n\n" + createLinkToAPIReference("SigfoxDevice", "terminateSigfoxDevice"),
}

func SigfoxDevicesTerminateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSigfoxDevicesTerminateCmdParams(ac)
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

func collectSigfoxDevicesTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SigfoxDevicesTerminateCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesTerminateCmd("/sigfox_devices/{device_id}/terminate"),
		query:  buildQueryForSigfoxDevicesTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesTerminateCmd(path string) string {

	escapedDeviceId := url.PathEscape(SigfoxDevicesTerminateCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesTerminateCmd() url.Values {
	result := url.Values{}

	if SigfoxDevicesTerminateCmdDeleteImmediately != false {
		result.Add("delete_immediately", sprintf("%t", SigfoxDevicesTerminateCmdDeleteImmediately))
	}

	return result
}
