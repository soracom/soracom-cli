// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDeviceDeleteCmdDeviceId holds value of 'device_id' option
var SoraCamDeviceDeleteCmdDeviceId string

func InitSoraCamDeviceDeleteCmd() {
	SoraCamDeviceDeleteCmd.Flags().StringVar(&SoraCamDeviceDeleteCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDeviceDeleteCmd.RunE = SoraCamDeviceDeleteCmdRunE

	SoraCamDeviceCmd.AddCommand(SoraCamDeviceDeleteCmd)
}

// SoraCamDeviceDeleteCmd defines 'delete' subcommand
var SoraCamDeviceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/sora_cam/devices/{device_id}:delete:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}:delete:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "handleDeleteSoraCamDevice"),
}

func SoraCamDeviceDeleteCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDeviceDeleteCmdParams(ac)
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

func collectSoraCamDeviceDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDeviceDeleteCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSoraCamDeviceDeleteCmd("/sora_cam/devices/{device_id}"),
		query:  buildQueryForSoraCamDeviceDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDeviceDeleteCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDeviceDeleteCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDeviceDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
