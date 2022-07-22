// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesUnobserveResourcesCmdDeviceId holds value of 'device_id' option
var DevicesUnobserveResourcesCmdDeviceId string

// DevicesUnobserveResourcesCmdInstance holds value of 'instance' option
var DevicesUnobserveResourcesCmdInstance string

// DevicesUnobserveResourcesCmdObject holds value of 'object' option
var DevicesUnobserveResourcesCmdObject string

func init() {
	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesUnobserveResourcesCmd.Flags().StringVar(&DevicesUnobserveResourcesCmdObject, "object", "", TRAPI("Object ID"))
	DevicesCmd.AddCommand(DevicesUnobserveResourcesCmd)
}

// DevicesUnobserveResourcesCmd defines 'unobserve-resources' subcommand
var DevicesUnobserveResourcesCmd = &cobra.Command{
	Use:   "unobserve-resources",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/unobserve:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/unobserve:post:description`) + "\n\n" + createLinkToAPIReference("Device", "unobserveDeviceResources"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectDevicesUnobserveResourcesCmdParams(ac)
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
	},
}

func collectDevicesUnobserveResourcesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesUnobserveResourcesCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("instance", "instance", "path", parsedBody, DevicesUnobserveResourcesCmdInstance)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("object", "object", "path", parsedBody, DevicesUnobserveResourcesCmdObject)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnobserveResourcesCmd("/devices/{device_id}/{object}/{instance}/unobserve"),
		query:  buildQueryForDevicesUnobserveResourcesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesUnobserveResourcesCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesUnobserveResourcesCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesUnobserveResourcesCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesUnobserveResourcesCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	return path
}

func buildQueryForDevicesUnobserveResourcesCmd() url.Values {
	result := url.Values{}

	return result
}
