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

// SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdDeviceId string

// SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState holds value of 'state' option
var SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState string

// SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody holds contents of request body to be sent
var SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody string

func InitSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd() {
	SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState, "state", "", TRAPI("Display settings for motion tagging.- 'on'- 'off'"))

	SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd.RunE = SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd)
}

// SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd defines 'set-motion-tagging' subcommand
var SoraCamDevicesAtomCamSettingsSetMotionTaggingCmd = &cobra.Command{
	Use:   "set-motion-tagging",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/motion_tagging:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/motion_tagging:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "setSoraCamDeviceAtomCamSettingsMotionTagging"),
}

func SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsSetMotionTaggingCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsSetMotionTaggingCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("state", "state", "body", parsedBody, SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd("/sora_cam/devices/{device_id}/atom_cam/settings/motion_tagging"),
		query:       buildQueryForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdBody)
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

	if SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState != "" {
		result["state"] = SoraCamDevicesAtomCamSettingsSetMotionTaggingCmdState
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
