// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesDeleteCmdDeviceId holds value of 'device_id' option
var DevicesDeleteCmdDeviceId string

func InitDevicesDeleteCmd() {
	DevicesDeleteCmd.Flags().StringVar(&DevicesDeleteCmdDeviceId, "device-id", "", TRAPI("Device to delete"))

	DevicesDeleteCmd.RunE = DevicesDeleteCmdRunE

	DevicesCmd.AddCommand(DevicesDeleteCmd)
}

// DevicesDeleteCmd defines 'delete' subcommand
var DevicesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/devices/{device_id}:delete:summary"),
	Long:  TRAPI(`/devices/{device_id}:delete:description`) + "\n\n" + createLinkToAPIReference("Device", "deleteDevice"),
}

func DevicesDeleteCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectDevicesDeleteCmdParams(ac)
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

func collectDevicesDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesDeleteCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteCmd("/devices/{device_id}"),
		query:  buildQueryForDevicesDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesDeleteCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesDeleteCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForDevicesDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
