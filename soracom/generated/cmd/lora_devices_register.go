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

// LoraDevicesRegisterCmdDeviceId holds value of 'device_id' option
var LoraDevicesRegisterCmdDeviceId string

// LoraDevicesRegisterCmdGroupId holds value of 'groupId' option
var LoraDevicesRegisterCmdGroupId string

// LoraDevicesRegisterCmdRegistrationSecret holds value of 'registrationSecret' option
var LoraDevicesRegisterCmdRegistrationSecret string

// LoraDevicesRegisterCmdBody holds contents of request body to be sent
var LoraDevicesRegisterCmdBody string

func init() {
	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdGroupId, "group-id", "", TRAPI(""))

	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdRegistrationSecret, "registration-secret", "", TRAPI(""))

	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraDevicesCmd.AddCommand(LoraDevicesRegisterCmd)
}

// LoraDevicesRegisterCmd defines 'register' subcommand
var LoraDevicesRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/lora_devices/{device_id}/register:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/register:post:description`) + "\n\n" + createLinkToAPIReference("LoraDevice", "registerLoraDevice"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLoraDevicesRegisterCmdParams(ac)
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
	},
}

func collectLoraDevicesRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraDevicesRegisterCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, LoraDevicesRegisterCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesRegisterCmd("/lora_devices/{device_id}/register"),
		query:       buildQueryForLoraDevicesRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesRegisterCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesRegisterCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraDevicesRegisterCmd() (string, error) {
	var result map[string]interface{}

	if LoraDevicesRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraDevicesRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesRegisterCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraDevicesRegisterCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraDevicesRegisterCmdBody)
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

	if LoraDevicesRegisterCmdGroupId != "" {
		result["groupId"] = LoraDevicesRegisterCmdGroupId
	}

	if LoraDevicesRegisterCmdRegistrationSecret != "" {
		result["registrationSecret"] = LoraDevicesRegisterCmdRegistrationSecret
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
