// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesRegisterCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesRegisterCmdDeviceId string

// SigfoxDevicesRegisterCmdRegistrationSecret holds value of 'registrationSecret' option
var SigfoxDevicesRegisterCmdRegistrationSecret string

// SigfoxDevicesRegisterCmdBody holds contents of request body to be sent
var SigfoxDevicesRegisterCmdBody string

func init() {
	SigfoxDevicesRegisterCmd.Flags().StringVar(&SigfoxDevicesRegisterCmdDeviceId, "device-id", "", TRAPI("Device ID of the target sigfox device to register"))

	SigfoxDevicesRegisterCmd.Flags().StringVar(&SigfoxDevicesRegisterCmdRegistrationSecret, "registration-secret", "", TRAPI(""))

	SigfoxDevicesRegisterCmd.Flags().StringVar(&SigfoxDevicesRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SigfoxDevicesCmd.AddCommand(SigfoxDevicesRegisterCmd)
}

// SigfoxDevicesRegisterCmd defines 'register' subcommand
var SigfoxDevicesRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/sigfox_devices/{device_id}/register:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/register:post:description`) + "\n\n" + createLinkToAPIReference("SigfoxDevice", "registerSigfoxDevice"),
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

		param, err := collectSigfoxDevicesRegisterCmdParams(ac)
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

func collectSigfoxDevicesRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSigfoxDevicesRegisterCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SigfoxDevicesRegisterCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSigfoxDevicesRegisterCmd("/sigfox_devices/{device_id}/register"),
		query:       buildQueryForSigfoxDevicesRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSigfoxDevicesRegisterCmd(path string) string {

	escapedDeviceId := url.PathEscape(SigfoxDevicesRegisterCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSigfoxDevicesRegisterCmd() (string, error) {
	var result map[string]interface{}

	if SigfoxDevicesRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SigfoxDevicesRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(SigfoxDevicesRegisterCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SigfoxDevicesRegisterCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SigfoxDevicesRegisterCmdBody)
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

	if SigfoxDevicesRegisterCmdRegistrationSecret != "" {
		result["registrationSecret"] = SigfoxDevicesRegisterCmdRegistrationSecret
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
