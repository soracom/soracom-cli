// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesGetCmdDeviceId holds value of 'device_id' option
var DevicesGetCmdDeviceId string

// DevicesGetCmdModel holds value of 'model' option
var DevicesGetCmdModel bool

func init() {
	DevicesGetCmd.Flags().StringVar(&DevicesGetCmdDeviceId, "device-id", "", TRAPI("Device ID"))

	DevicesGetCmd.Flags().BoolVar(&DevicesGetCmdModel, "model", false, TRAPI("Whether or not to add model information"))
	DevicesCmd.AddCommand(DevicesGetCmd)
}

// DevicesGetCmd defines 'get' subcommand
var DevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/devices/{device_id}:get:summary"),
	Long:  TRAPI(`/devices/{device_id}:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		param, err := collectDevicesGetCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesGetCmdParams(ac *apiClient) (*apiParams, error) {
	if DevicesGetCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetCmd("/devices/{device_id}"),
		query:  buildQueryForDevicesGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesGetCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesGetCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForDevicesGetCmd() url.Values {
	result := url.Values{}

	if DevicesGetCmdModel != false {
		result.Add("model", sprintf("%t", DevicesGetCmdModel))
	}

	return result
}
