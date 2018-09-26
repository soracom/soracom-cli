package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesGetInstanceCmdDeviceId holds value of 'device_id' option
var DevicesGetInstanceCmdDeviceId string

// DevicesGetInstanceCmdInstance holds value of 'instance' option
var DevicesGetInstanceCmdInstance string

// DevicesGetInstanceCmdObject holds value of 'object' option
var DevicesGetInstanceCmdObject string

// DevicesGetInstanceCmdModel holds value of 'model' option
var DevicesGetInstanceCmdModel bool

func init() {
	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesGetInstanceCmd.Flags().StringVar(&DevicesGetInstanceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesGetInstanceCmd.Flags().BoolVar(&DevicesGetInstanceCmdModel, "model", false, TRAPI("Whether or not to add model information"))

	DevicesCmd.AddCommand(DevicesGetInstanceCmd)
}

// DevicesGetInstanceCmd defines 'get-instance' subcommand
var DevicesGetInstanceCmd = &cobra.Command{
	Use:   "get-instance",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}:get:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}:get:description`),
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

		param, err := collectDevicesGetInstanceCmdParams(ac)
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

func collectDevicesGetInstanceCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetInstanceCmd("/devices/{device_id}/{object}/{instance}"),
		query:  buildQueryForDevicesGetInstanceCmd(),
	}, nil
}

func buildPathForDevicesGetInstanceCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesGetInstanceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesGetInstanceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesGetInstanceCmdObject, -1)

	return path
}

func buildQueryForDevicesGetInstanceCmd() string {
	result := []string{}

	if DevicesGetInstanceCmdModel != false {
		result = append(result, sprintf("%s=%t", "model", DevicesGetInstanceCmdModel))
	}

	return strings.Join(result, "&")
}
