// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesImagesListExportsForDeviceCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesImagesListExportsForDeviceCmdDeviceId string

// SoraCamDevicesImagesListExportsForDeviceCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesImagesListExportsForDeviceCmdLastEvaluatedKey string

// SoraCamDevicesImagesListExportsForDeviceCmdSort holds value of 'sort' option
var SoraCamDevicesImagesListExportsForDeviceCmdSort string

// SoraCamDevicesImagesListExportsForDeviceCmdLimit holds value of 'limit' option
var SoraCamDevicesImagesListExportsForDeviceCmdLimit int64

// SoraCamDevicesImagesListExportsForDeviceCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesImagesListExportsForDeviceCmdOutputJSONL bool

func InitSoraCamDevicesImagesListExportsForDeviceCmd() {
	SoraCamDevicesImagesListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsForDeviceCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesImagesListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsForDeviceCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Value of the x-soracom-next-key header in the response to the last listSoraCamDeviceImageExportsForDevice request. By specifying this parameter, you can continue to retrieve the list from the last request."))

	SoraCamDevicesImagesListExportsForDeviceCmd.Flags().StringVar(&SoraCamDevicesImagesListExportsForDeviceCmdSort, "sort", "desc", TRAPI("Sort order. The list in the response is sorted in ascending ('asc') or descending ('desc') order of 'requestedTime'. The default is 'desc' i.e. newer items are sorted first."))

	SoraCamDevicesImagesListExportsForDeviceCmd.Flags().Int64Var(&SoraCamDevicesImagesListExportsForDeviceCmdLimit, "limit", 10, TRAPI("Maximum number of items to retrieve in one request. Note that the response may contain fewer items than the specified limit."))

	SoraCamDevicesImagesListExportsForDeviceCmd.Flags().BoolVar(&SoraCamDevicesImagesListExportsForDeviceCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SoraCamDevicesImagesListExportsForDeviceCmd.RunE = SoraCamDevicesImagesListExportsForDeviceCmdRunE

	SoraCamDevicesImagesCmd.AddCommand(SoraCamDevicesImagesListExportsForDeviceCmd)
}

// SoraCamDevicesImagesListExportsForDeviceCmd defines 'list-exports-for-device' subcommand
var SoraCamDevicesImagesListExportsForDeviceCmd = &cobra.Command{
	Use:   "list-exports-for-device",
	Short: TRAPI("/sora_cam/devices/{device_id}/images/exports:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/images/exports:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceImageExportsForDevice"),
}

func SoraCamDevicesImagesListExportsForDeviceCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesImagesListExportsForDeviceCmdParams(ac)
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
		if SoraCamDevicesImagesListExportsForDeviceCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamDevicesImagesListExportsForDeviceCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesImagesListExportsForDeviceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesImagesListExportsForDeviceCmd("/sora_cam/devices/{device_id}/images/exports"),
		query:  buildQueryForSoraCamDevicesImagesListExportsForDeviceCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesImagesListExportsForDeviceCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesImagesListExportsForDeviceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesImagesListExportsForDeviceCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesImagesListExportsForDeviceCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraCamDevicesImagesListExportsForDeviceCmdLastEvaluatedKey)
	}

	if SoraCamDevicesImagesListExportsForDeviceCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesImagesListExportsForDeviceCmdSort)
	}

	if SoraCamDevicesImagesListExportsForDeviceCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesImagesListExportsForDeviceCmdLimit))
	}

	return result
}
