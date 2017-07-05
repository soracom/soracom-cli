package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesObserveResourceCmdDeviceId holds value of 'deviceId' option
var DevicesObserveResourceCmdDeviceId string

// DevicesObserveResourceCmdInstance holds value of 'instance' option
var DevicesObserveResourceCmdInstance string

// DevicesObserveResourceCmdObject holds value of 'object' option
var DevicesObserveResourceCmdObject string

// DevicesObserveResourceCmdResource holds value of 'resource' option
var DevicesObserveResourceCmdResource string

// DevicesObserveResourceCmdModel holds value of 'model' option
var DevicesObserveResourceCmdModel bool

func init() {
	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdDeviceId, "device-id", "", TR("devices.observe_resource.parameters.deviceId.description"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdInstance, "instance", "", TR("devices.observe_resource.parameters.instance.description"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdObject, "object", "", TR("devices.observe_resource.parameters.object.description"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdResource, "resource", "", TR("devices.observe_resource.parameters.resource.description"))

	DevicesObserveResourceCmd.Flags().BoolVar(&DevicesObserveResourceCmdModel, "model", false, TR("devices.observe_resource.parameters.model.description"))

	DevicesCmd.AddCommand(DevicesObserveResourceCmd)
}

// DevicesObserveResourceCmd defines 'observe-resource' subcommand
var DevicesObserveResourceCmd = &cobra.Command{
	Use:   "observe-resource",
	Short: TR("devices.observe_resource.summary"),
	Long:  TR(`devices.observe_resource.description`),
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

		param, err := collectDevicesObserveResourceCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectDevicesObserveResourceCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesObserveResourceCmd("/devices/{deviceId}/{object}/{instance}/{resource}/observe"),
		query:  buildQueryForDevicesObserveResourceCmd(),
	}, nil
}

func buildPathForDevicesObserveResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesObserveResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesObserveResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesObserveResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesObserveResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesObserveResourceCmd() string {
	result := []string{}

	if DevicesObserveResourceCmdModel != false {
		result = append(result, sprintf("%s=%t", "model", DevicesObserveResourceCmdModel))
	}

	return strings.Join(result, "&")
}
