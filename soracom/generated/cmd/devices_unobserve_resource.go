// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesUnobserveResourceCmdDeviceId holds value of 'device_id' option
var DevicesUnobserveResourceCmdDeviceId string

// DevicesUnobserveResourceCmdInstance holds value of 'instance' option
var DevicesUnobserveResourceCmdInstance string

// DevicesUnobserveResourceCmdObject holds value of 'object' option
var DevicesUnobserveResourceCmdObject string

// DevicesUnobserveResourceCmdResource holds value of 'resource' option
var DevicesUnobserveResourceCmdResource string

func init() {
	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdResource, "resource", "", TRAPI("Resource ID"))
	DevicesCmd.AddCommand(DevicesUnobserveResourceCmd)
}

// DevicesUnobserveResourceCmd defines 'unobserve-resource' subcommand
var DevicesUnobserveResourceCmd = &cobra.Command{
	Use:   "unobserve-resource",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/{resource}/unobserve:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}/unobserve:post:description`),
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

		param, err := collectDevicesUnobserveResourceCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesUnobserveResourceCmdParams(ac *apiClient) (*apiParams, error) {
	if DevicesUnobserveResourceCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	if DevicesUnobserveResourceCmdInstance == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "instance")
	}

	if DevicesUnobserveResourceCmdObject == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "object")
	}

	if DevicesUnobserveResourceCmdResource == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "resource")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnobserveResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}/unobserve"),
		query:  buildQueryForDevicesUnobserveResourceCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesUnobserveResourceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesUnobserveResourceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesUnobserveResourceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesUnobserveResourceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	escapedResource := url.PathEscape(DevicesUnobserveResourceCmdResource)

	path = strReplace(path, "{"+"resource"+"}", escapedResource, -1)

	return path
}

func buildQueryForDevicesUnobserveResourceCmd() url.Values {
	result := url.Values{}

	return result
}
