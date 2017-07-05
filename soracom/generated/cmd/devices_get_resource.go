package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesGetResourceCmdDeviceId holds value of 'deviceId' option
var DevicesGetResourceCmdDeviceId string

// DevicesGetResourceCmdInstance holds value of 'instance' option
var DevicesGetResourceCmdInstance string

// DevicesGetResourceCmdObject holds value of 'object' option
var DevicesGetResourceCmdObject string

// DevicesGetResourceCmdResource holds value of 'resource' option
var DevicesGetResourceCmdResource string

// DevicesGetResourceCmdModel holds value of 'model' option
var DevicesGetResourceCmdModel bool

func init() {
	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdDeviceId, "device-id", "", TR("devices.get_resource.parameters.deviceId.description"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdInstance, "instance", "", TR("devices.get_resource.parameters.instance.description"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdObject, "object", "", TR("devices.get_resource.parameters.object.description"))

	DevicesGetResourceCmd.Flags().StringVar(&DevicesGetResourceCmdResource, "resource", "", TR("devices.get_resource.parameters.resource.description"))

	DevicesGetResourceCmd.Flags().BoolVar(&DevicesGetResourceCmdModel, "model", false, TR("devices.get_resource.parameters.model.description"))

	DevicesCmd.AddCommand(DevicesGetResourceCmd)
}

// DevicesGetResourceCmd defines 'get-resource' subcommand
var DevicesGetResourceCmd = &cobra.Command{
	Use:   "get-resource",
	Short: TR("devices.get_resource.summary"),
	Long:  TR(`devices.get_resource.description`),
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

		param, err := collectDevicesGetResourceCmdParams()
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

func collectDevicesGetResourceCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetResourceCmd("/devices/{deviceId}/{object}/{instance}/{resource}"),
		query:  buildQueryForDevicesGetResourceCmd(),
	}, nil
}

func buildPathForDevicesGetResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesGetResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesGetResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesGetResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesGetResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesGetResourceCmd() string {
	result := []string{}

	if DevicesGetResourceCmdModel != false {
		result = append(result, sprintf("%s=%t", "model", DevicesGetResourceCmdModel))
	}

	return strings.Join(result, "&")
}
