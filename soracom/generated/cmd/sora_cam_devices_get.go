// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesGetCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesGetCmdDeviceId string

func InitSoraCamDevicesGetCmd() {
	SoraCamDevicesGetCmd.Flags().StringVar(&SoraCamDevicesGetCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesGetCmd.RunE = SoraCamDevicesGetCmdRunE

	SoraCamDevicesCmd.AddCommand(SoraCamDevicesGetCmd)
}

// SoraCamDevicesGetCmd defines 'get' subcommand
var SoraCamDevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/sora_cam/devices/{device_id}:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDevice"),
}

func SoraCamDevicesGetCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := authHelper(ac, cmd, args)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSoraCamDevicesGetCmdParams(ac)
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

func collectSoraCamDevicesGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesGetCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesGetCmd("/sora_cam/devices/{device_id}"),
		query:  buildQueryForSoraCamDevicesGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesGetCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesGetCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesGetCmd() url.Values {
	result := url.Values{}

	return result
}
