package cmd

import (
	"os"
	"strings"

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

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectDevicesUnobserveResourceCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnobserveResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}/unobserve"),
		query:  buildQueryForDevicesUnobserveResourceCmd(),
	}, nil
}

func buildPathForDevicesUnobserveResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesUnobserveResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesUnobserveResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesUnobserveResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesUnobserveResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesUnobserveResourceCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
