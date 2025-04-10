// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SoraCamDevicesVideosExportCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesVideosExportCmdDeviceId string

// SoraCamDevicesVideosExportCmdFrom holds value of 'from' option
var SoraCamDevicesVideosExportCmdFrom int64

// SoraCamDevicesVideosExportCmdTo holds value of 'to' option
var SoraCamDevicesVideosExportCmdTo int64

// SoraCamDevicesVideosExportCmdBody holds contents of request body to be sent
var SoraCamDevicesVideosExportCmdBody string

func InitSoraCamDevicesVideosExportCmd() {
	SoraCamDevicesVideosExportCmd.Flags().StringVar(&SoraCamDevicesVideosExportCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesVideosExportCmd.Flags().Int64Var(&SoraCamDevicesVideosExportCmdFrom, "from", 0, TRAPI("Start time for exporting (unix time in milliseconds)."))

	SoraCamDevicesVideosExportCmd.Flags().Int64Var(&SoraCamDevicesVideosExportCmdTo, "to", 0, TRAPI("End time for exporting (unix time in milliseconds).- The maximum time for a single API call to export is 900 seconds (15 minutes). Make sure the difference between 'from' and 'to' does not exceed 900 seconds."))

	SoraCamDevicesVideosExportCmd.Flags().StringVar(&SoraCamDevicesVideosExportCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesVideosExportCmd.RunE = SoraCamDevicesVideosExportCmdRunE

	SoraCamDevicesVideosCmd.AddCommand(SoraCamDevicesVideosExportCmd)
}

// SoraCamDevicesVideosExportCmd defines 'export' subcommand
var SoraCamDevicesVideosExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/sora_cam/devices/{device_id}/videos/exports:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/videos/exports:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "exportSoraCamDeviceRecordedVideo"),
}

func SoraCamDevicesVideosExportCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesVideosExportCmdParams(ac)
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

func collectSoraCamDevicesVideosExportCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesVideosExportCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesVideosExportCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("from", "from", "body", parsedBody, SoraCamDevicesVideosExportCmdFrom)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("to", "to", "body", parsedBody, SoraCamDevicesVideosExportCmdTo)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesVideosExportCmd("/sora_cam/devices/{device_id}/videos/exports"),
		query:       buildQueryForSoraCamDevicesVideosExportCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesVideosExportCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesVideosExportCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesVideosExportCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesVideosExportCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesVideosExportCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesVideosExportCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesVideosExportCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesVideosExportCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesVideosExportCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if SoraCamDevicesVideosExportCmd.Flags().Lookup("from").Changed {
		result["from"] = SoraCamDevicesVideosExportCmdFrom
	}

	if SoraCamDevicesVideosExportCmd.Flags().Lookup("to").Changed {
		result["to"] = SoraCamDevicesVideosExportCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
