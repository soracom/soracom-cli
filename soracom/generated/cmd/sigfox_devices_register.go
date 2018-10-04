package cmd

import (
	"encoding/json"

	"io/ioutil"

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
	Long:  TRAPI(`/sigfox_devices/{device_id}/register:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectSigfoxDevicesRegisterCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSigfoxDevicesRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSigfoxDevicesRegisterCmd("/sigfox_devices/{device_id}/register"),
		query:       buildQueryForSigfoxDevicesRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSigfoxDevicesRegisterCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesRegisterCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
