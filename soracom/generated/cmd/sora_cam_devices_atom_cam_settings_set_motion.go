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

// SoraCamDevicesAtomCamSettingsSetMotionCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsSetMotionCmdDeviceId string

// SoraCamDevicesAtomCamSettingsSetMotionCmdState holds value of 'state' option
var SoraCamDevicesAtomCamSettingsSetMotionCmdState string

// SoraCamDevicesAtomCamSettingsSetMotionCmdBody holds contents of request body to be sent
var SoraCamDevicesAtomCamSettingsSetMotionCmdBody string

func InitSoraCamDevicesAtomCamSettingsSetMotionCmd() {
	SoraCamDevicesAtomCamSettingsSetMotionCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesAtomCamSettingsSetMotionCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionCmdState, "state", "", TRAPI("Settings for motion detection.- 'on'- 'off'"))

	SoraCamDevicesAtomCamSettingsSetMotionCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesAtomCamSettingsSetMotionCmd.RunE = SoraCamDevicesAtomCamSettingsSetMotionCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsSetMotionCmd)
}

// SoraCamDevicesAtomCamSettingsSetMotionCmd defines 'set-motion' subcommand
var SoraCamDevicesAtomCamSettingsSetMotionCmd = &cobra.Command{
	Use:   "set-motion",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/motion:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/motion:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "setSoraCamDeviceAtomCamSettingsMotion"),
}

func SoraCamDevicesAtomCamSettingsSetMotionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsSetMotionCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsSetMotionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesAtomCamSettingsSetMotionCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsSetMotionCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("state", "state", "body", parsedBody, SoraCamDevicesAtomCamSettingsSetMotionCmdState)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesAtomCamSettingsSetMotionCmd("/sora_cam/devices/{device_id}/atom_cam/settings/motion"),
		query:       buildQueryForSoraCamDevicesAtomCamSettingsSetMotionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsSetMotionCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsSetMotionCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsSetMotionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesAtomCamSettingsSetMotionCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesAtomCamSettingsSetMotionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesAtomCamSettingsSetMotionCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesAtomCamSettingsSetMotionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesAtomCamSettingsSetMotionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesAtomCamSettingsSetMotionCmdBody)
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

	if SoraCamDevicesAtomCamSettingsSetMotionCmdState != "" {
		result["state"] = SoraCamDevicesAtomCamSettingsSetMotionCmdState
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}