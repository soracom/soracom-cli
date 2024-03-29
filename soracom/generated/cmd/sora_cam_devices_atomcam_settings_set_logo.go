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

// SoraCamDevicesAtomcamSettingsSetLogoCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesAtomcamSettingsSetLogoCmdDeviceId string

// SoraCamDevicesAtomcamSettingsSetLogoCmdState holds value of 'state' option
var SoraCamDevicesAtomcamSettingsSetLogoCmdState string

// SoraCamDevicesAtomcamSettingsSetLogoCmdBody holds contents of request body to be sent
var SoraCamDevicesAtomcamSettingsSetLogoCmdBody string

func InitSoraCamDevicesAtomcamSettingsSetLogoCmd() {
	SoraCamDevicesAtomcamSettingsSetLogoCmd.Flags().StringVar(&SoraCamDevicesAtomcamSettingsSetLogoCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesAtomcamSettingsSetLogoCmd.Flags().StringVar(&SoraCamDevicesAtomcamSettingsSetLogoCmdState, "state", "", TRAPI("Display settings for the ATOM tech Inc. logo shown on the bottom left of the captured image.- 'on'- 'off'"))

	SoraCamDevicesAtomcamSettingsSetLogoCmd.Flags().StringVar(&SoraCamDevicesAtomcamSettingsSetLogoCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SoraCamDevicesAtomcamSettingsSetLogoCmd.RunE = SoraCamDevicesAtomcamSettingsSetLogoCmdRunE

	SoraCamDevicesAtomcamSettingsCmd.AddCommand(SoraCamDevicesAtomcamSettingsSetLogoCmd)
}

// SoraCamDevicesAtomcamSettingsSetLogoCmd defines 'set-logo' subcommand
var SoraCamDevicesAtomcamSettingsSetLogoCmd = &cobra.Command{
	Use:   "set-logo",
	Short: TRAPI("/sora_cam/devices/{device_id}/atomcam/settings/logo:post:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/atomcam/settings/logo:post:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "setSoraCamDeviceAtomCamSettingsLogo"),
}

func SoraCamDevicesAtomcamSettingsSetLogoCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomcamSettingsSetLogoCmdParams(ac)
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

func collectSoraCamDevicesAtomcamSettingsSetLogoCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesAtomcamSettingsSetLogoCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesAtomcamSettingsSetLogoCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraCamDevicesAtomcamSettingsSetLogoCmd("/sora_cam/devices/{device_id}/atomcam/settings/logo"),
		query:       buildQueryForSoraCamDevicesAtomcamSettingsSetLogoCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomcamSettingsSetLogoCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesAtomcamSettingsSetLogoCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesAtomcamSettingsSetLogoCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesAtomcamSettingsSetLogoCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesAtomcamSettingsSetLogoCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesAtomcamSettingsSetLogoCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesAtomcamSettingsSetLogoCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesAtomcamSettingsSetLogoCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesAtomcamSettingsSetLogoCmdBody)
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

	if SoraCamDevicesAtomcamSettingsSetLogoCmdState != "" {
		result["state"] = SoraCamDevicesAtomcamSettingsSetLogoCmdState
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
