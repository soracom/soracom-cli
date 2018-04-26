package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesObserveResourcesCmdDeviceId holds value of 'device_id' option
var DevicesObserveResourcesCmdDeviceId string

// DevicesObserveResourcesCmdInstance holds value of 'instance' option
var DevicesObserveResourcesCmdInstance string

// DevicesObserveResourcesCmdObject holds value of 'object' option
var DevicesObserveResourcesCmdObject string

// DevicesObserveResourcesCmdModel holds value of 'model' option
var DevicesObserveResourcesCmdModel bool

func init() {
	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdObject, "object", "", TRAPI("Object ID"))

	DevicesObserveResourcesCmd.Flags().BoolVar(&DevicesObserveResourcesCmdModel, "model", false, TRAPI("Whether or not to add model information"))

	DevicesCmd.AddCommand(DevicesObserveResourcesCmd)
}

// DevicesObserveResourcesCmd defines 'observe-resources' subcommand
var DevicesObserveResourcesCmd = &cobra.Command{
	Use:   "observe-resources",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/observe:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/observe:post:description`),
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

		param, err := collectDevicesObserveResourcesCmdParams(ac)
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

func collectDevicesObserveResourcesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesObserveResourcesCmd("/devices/{device_id}/{object}/{instance}/observe"),
		query:  buildQueryForDevicesObserveResourcesCmd(),
	}, nil
}

func buildPathForDevicesObserveResourcesCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesObserveResourcesCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesObserveResourcesCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesObserveResourcesCmdObject, -1)

	return path
}

func buildQueryForDevicesObserveResourcesCmd() string {
	result := []string{}

	if DevicesObserveResourcesCmdModel != false {
		result = append(result, sprintf("%s=%t", "model", DevicesObserveResourcesCmdModel))
	}

	return strings.Join(result, "&")
}
