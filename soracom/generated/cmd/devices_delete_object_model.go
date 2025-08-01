// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesDeleteObjectModelCmdModelId holds value of 'model_id' option
var DevicesDeleteObjectModelCmdModelId string

func InitDevicesDeleteObjectModelCmd() {
	DevicesDeleteObjectModelCmd.Flags().StringVar(&DevicesDeleteObjectModelCmdModelId, "model-id", "", TRAPI("Target device object model ID"))

	DevicesDeleteObjectModelCmd.RunE = DevicesDeleteObjectModelCmdRunE

	DevicesCmd.AddCommand(DevicesDeleteObjectModelCmd)
}

// DevicesDeleteObjectModelCmd defines 'delete-object-model' subcommand
var DevicesDeleteObjectModelCmd = &cobra.Command{
	Use:   "delete-object-model",
	Short: TRAPI("/device_object_models/{model_id}:delete:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}:delete:description`) + "\n\n" + createLinkToAPIReference("DeviceObjectModel", "deleteDeviceObjectModel"),
}

func DevicesDeleteObjectModelCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectDevicesDeleteObjectModelCmdParams(ac)
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

func collectDevicesDeleteObjectModelCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("model_id", "model-id", "path", parsedBody, DevicesDeleteObjectModelCmdModelId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteObjectModelCmd("/device_object_models/{model_id}"),
		query:  buildQueryForDevicesDeleteObjectModelCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesDeleteObjectModelCmd(path string) string {

	escapedModelId := url.PathEscape(DevicesDeleteObjectModelCmdModelId)

	path = strReplace(path, "{"+"model_id"+"}", escapedModelId, -1)

	return path
}

func buildQueryForDevicesDeleteObjectModelCmd() url.Values {
	result := url.Values{}

	return result
}
