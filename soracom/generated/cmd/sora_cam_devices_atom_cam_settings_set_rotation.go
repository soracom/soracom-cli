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

// SoraCamDevicesAtomCamSettingsSetRotationCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomCamSettingsSetRotationCmdDeviceId string

// SoraCamDevicesAtomCamSettingsSetRotationCmdState holds value of 'state' option
var SoraCamDevicesAtomCamSettingsSetRotationCmdState int64

// SoraCamDevicesAtomCamSettingsSetRotationCmdBody holds contents of request body to be sent
var SoraCamDevicesAtomCamSettingsSetRotationCmdBody string

func InitSoraCamDevicesAtomCamSettingsSetRotationCmd() {
	SoraCamDevicesAtomCamSettingsSetRotationCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetRotationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesAtomCamSettingsSetRotationCmd.Flags().Int64Var(&SoraCamDevicesAtomCamSettingsSetRotationCmdState, "state", 0, TRAPI("Settings for rotating the image by 180 degrees. Set to '180' when installing the SoraCam compatible camera device upside down.- '0': OFF (no rotation).- '180': ON (180-degree rotation)."))

	SoraCamDevicesAtomCamSettingsSetRotationCmd.Flags().StringVar(&SoraCamDevicesAtomCamSettingsSetRotationCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesAtomCamSettingsSetRotationCmd.RunE = SoraCamDevicesAtomCamSettingsSetRotationCmdRunE

	SoraCamDevicesAtomCamSettingsCmd.AddCommand(SoraCamDevicesAtomCamSettingsSetRotationCmd)
}

// SoraCamDevicesAtomCamSettingsSetRotationCmd defines 'set-rotation' subcommand
var SoraCamDevicesAtomCamSettingsSetRotationCmd = &cobra.Command{
	Use:   "set-rotation",
	Short: TRAPI("/sora_cam/devices/{device_id}/atom_cam/settings/rotation:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atom_cam/settings/rotation:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "setSoraCamDeviceAtomCamSettingsRotation"),
}

func SoraCamDevicesAtomCamSettingsSetRotationCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamSettingsSetRotationCmdParams(ac)
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

func collectSoraCamDevicesAtomCamSettingsSetRotationCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesAtomCamSettingsSetRotationCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomCamSettingsSetRotationCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("state", "state", "body", parsedBody, SoraCamDevicesAtomCamSettingsSetRotationCmdState)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesAtomCamSettingsSetRotationCmd("/sora_cam/devices/{device_id}/atom_cam/settings/rotation"),
		query:       buildQueryForSoraCamDevicesAtomCamSettingsSetRotationCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamSettingsSetRotationCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomCamSettingsSetRotationCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomCamSettingsSetRotationCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesAtomCamSettingsSetRotationCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesAtomCamSettingsSetRotationCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesAtomCamSettingsSetRotationCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesAtomCamSettingsSetRotationCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesAtomCamSettingsSetRotationCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesAtomCamSettingsSetRotationCmdBody)
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

	if SoraCamDevicesAtomCamSettingsSetRotationCmd.Flags().Lookup("state").Changed {
		result["state"] = SoraCamDevicesAtomCamSettingsSetRotationCmdState
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
