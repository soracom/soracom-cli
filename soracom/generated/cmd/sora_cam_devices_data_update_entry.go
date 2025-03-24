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

// SoraCamDevicesDataUpdateEntryCmdContentType holds value of 'content-type' option
var SoraCamDevicesDataUpdateEntryCmdContentType string

// SoraCamDevicesDataUpdateEntryCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesDataUpdateEntryCmdDeviceId string

// SoraCamDevicesDataUpdateEntryCmdTime holds value of 'time' option
var SoraCamDevicesDataUpdateEntryCmdTime int64

// SoraCamDevicesDataUpdateEntryCmdBody holds contents of request body to be sent
var SoraCamDevicesDataUpdateEntryCmdBody string

func InitSoraCamDevicesDataUpdateEntryCmd() {
	SoraCamDevicesDataUpdateEntryCmd.Flags().StringVar(&SoraCamDevicesDataUpdateEntryCmdContentType, "content-type", "", TRAPI("The Content-Type of the data to be saved.- For JSON data, specify 'application/json'.- For text data, specify 'text/plain'.- For binary data, specify 'application/octet-stream'."))

	SoraCamDevicesDataUpdateEntryCmd.Flags().StringVar(&SoraCamDevicesDataUpdateEntryCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesDataUpdateEntryCmd.Flags().Int64Var(&SoraCamDevicesDataUpdateEntryCmdTime, "time", 0, TRAPI("The timestamp of the target data entry (UNIX time in milliseconds)."))

	SoraCamDevicesDataUpdateEntryCmd.Flags().StringVar(&SoraCamDevicesDataUpdateEntryCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesDataUpdateEntryCmd.RunE = SoraCamDevicesDataUpdateEntryCmdRunE

	SoraCamDevicesDataCmd.AddCommand(SoraCamDevicesDataUpdateEntryCmd)
}

// SoraCamDevicesDataUpdateEntryCmd defines 'update-entry' subcommand
var SoraCamDevicesDataUpdateEntryCmd = &cobra.Command{
	Use:   "update-entry",
	Short: TRAPI("/sora_cam/devices/{device_id}/data/{time}:put:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/data/{time}:put:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "updateSoraCamDeviceDataEntry"),
}

func SoraCamDevicesDataUpdateEntryCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesDataUpdateEntryCmdParams(ac)
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

func collectSoraCamDevicesDataUpdateEntryCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesDataUpdateEntryCmd()
	if err != nil {
		return nil, err
	}
	contentType := SoraCamDevicesDataUpdateEntryCmdContentType

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesDataUpdateEntryCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("time", "time", "path", parsedBody, SoraCamDevicesDataUpdateEntryCmdTime)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSoraCamDevicesDataUpdateEntryCmd("/sora_cam/devices/{device_id}/data/{time}"),
		query:       buildQueryForSoraCamDevicesDataUpdateEntryCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesDataUpdateEntryCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesDataUpdateEntryCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	path = strReplace(path, "{"+"time"+"}", url.PathEscape(sprintf("%d", SoraCamDevicesDataUpdateEntryCmdTime)), -1)

	return path
}

func buildQueryForSoraCamDevicesDataUpdateEntryCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesDataUpdateEntryCmd() (string, error) {
	var b []byte
	var err error

	if SoraCamDevicesDataUpdateEntryCmdBody != "" {
		if strings.HasPrefix(SoraCamDevicesDataUpdateEntryCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesDataUpdateEntryCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesDataUpdateEntryCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesDataUpdateEntryCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
