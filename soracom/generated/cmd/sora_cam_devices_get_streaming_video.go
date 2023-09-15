// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesGetStreamingVideoCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesGetStreamingVideoCmdDeviceId string

// SoraCamDevicesGetStreamingVideoCmdFrom holds value of 'from' option
var SoraCamDevicesGetStreamingVideoCmdFrom int64

// SoraCamDevicesGetStreamingVideoCmdTo holds value of 'to' option
var SoraCamDevicesGetStreamingVideoCmdTo int64

func InitSoraCamDevicesGetStreamingVideoCmd() {
	SoraCamDevicesGetStreamingVideoCmd.Flags().StringVar(&SoraCamDevicesGetStreamingVideoCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesGetStreamingVideoCmd.Flags().Int64Var(&SoraCamDevicesGetStreamingVideoCmdFrom, "from", 0, TRAPI("Start time of recorded video (UNIX time in milliseconds).- Omit both 'from' and 'to' to get information for downloading real-time video."))

	SoraCamDevicesGetStreamingVideoCmd.Flags().Int64Var(&SoraCamDevicesGetStreamingVideoCmdTo, "to", 0, TRAPI("End time of recorded video (UNIX time in milliseconds).- Omit both 'from' and 'to' to get information for downloading real-time video.- The maximum viewing time for a single API call is 900 seconds (15 minutes). Make sure the difference between 'from' and 'to' does not exceed 900 seconds."))

	SoraCamDevicesGetStreamingVideoCmd.RunE = SoraCamDevicesGetStreamingVideoCmdRunE

	SoraCamDevicesCmd.AddCommand(SoraCamDevicesGetStreamingVideoCmd)
}

// SoraCamDevicesGetStreamingVideoCmd defines 'get-streaming-video' subcommand
var SoraCamDevicesGetStreamingVideoCmd = &cobra.Command{
	Use:   "get-streaming-video",
	Short: TRAPI("/sora_cam/devices/{device_id}/stream:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/stream:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "getSoraCamDeviceStreamingVideo"),
}

func SoraCamDevicesGetStreamingVideoCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesGetStreamingVideoCmdParams(ac)
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

func collectSoraCamDevicesGetStreamingVideoCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesGetStreamingVideoCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesGetStreamingVideoCmd("/sora_cam/devices/{device_id}/stream"),
		query:  buildQueryForSoraCamDevicesGetStreamingVideoCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesGetStreamingVideoCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesGetStreamingVideoCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesGetStreamingVideoCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesGetStreamingVideoCmdFrom != 0 {
		result.Add("from", sprintf("%d", SoraCamDevicesGetStreamingVideoCmdFrom))
	}

	if SoraCamDevicesGetStreamingVideoCmdTo != 0 {
		result.Add("to", sprintf("%d", SoraCamDevicesGetStreamingVideoCmdTo))
	}

	return result
}
