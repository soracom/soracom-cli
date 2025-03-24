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

// SoraCamDevicesAtomCamSettingsSetTimestampCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsSetTimestampCmdDeviceId string

// SoraCamDevicesAtomCamSettingsSetTimestampCmdState holds value of 'state' option
var SoraCamDevicesAtomCamSettingsSetTimestampCmdState string

// SoraCamDevicesAtomCamSettingsSetTimestampCmdBody holds contents of request body to be sent
var SoraCamDevicesAtomCamSettingsSetTimestampCmdBody string

func InitSoraCamDevicesAtomCamSettingsSetTimestampCmd() {
	SoraCamDevicesAtomCamSettingsSetTimestampCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetTimestampCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsSetTimestampCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetTimestampCmdState, "state", "", TRAPI("Display settings for the timestamp shown on the bottom right of the captured image.- 'on'- 'off'"))

	SoraCamDevicesAtomCamSettingsSetTimestampCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetTimestampCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesAtomCamSettingsSetTimestampCmd.RunE = SoraCamDevicesAtomCamSettingsSetTimestampCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsSetTimestampCmd)
}

// SoraCamDevicesAtomCamSettingsSetTimestampCmd defines 'set-timestamp' subcommand
var SoraCamDevicesAtomCamSettingsSetTimestampCmd = &cobra.Command{
	Use:   "set-timestamp",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/timestamp:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/timestamp:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "setSoraCamDeviceAtomCamSettingsTimestamp"),
}

func SoraCamDevicesAtomCamSettingsSetTimestampCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsSetTimestampCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsSetTimestampCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesAtomCamSettingsSetTimestampCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsSetTimestampCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("state", "state", "body", parsedBody, SoraCamDevicesAtomCamSettingsSetTimestampCmdState)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesAtomCamSettingsSetTimestampCmd("/sora_cam/devices/{device_id}/atom_cam/settings/timestamp"),
		query:       buildQueryForSoraCamDevicesAtomCamSettingsSetTimestampCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsSetTimestampCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsSetTimestampCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsSetTimestampCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesAtomCamSettingsSetTimestampCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesAtomCamSettingsSetTimestampCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesAtomCamSettingsSetTimestampCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesAtomCamSettingsSetTimestampCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesAtomCamSettingsSetTimestampCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesAtomCamSettingsSetTimestampCmdBody)
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

	if SoraCamDevicesAtomCamSettingsSetTimestampCmdState != "" {
		result["state"] = SoraCamDevicesAtomCamSettingsSetTimestampCmdState
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
