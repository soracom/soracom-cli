package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesRegisterCmdDeviceId holds value of 'device_id' option
var LoraDevicesRegisterCmdDeviceId string

// LoraDevicesRegisterCmdBody holds contents of request body to be sent
var LoraDevicesRegisterCmdBody string

func init() {
	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesRegisterCmd.Flags().StringVar(&LoraDevicesRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraDevicesCmd.AddCommand(LoraDevicesRegisterCmd)
}

// LoraDevicesRegisterCmd defines 'register' subcommand
var LoraDevicesRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/lora_devices/{device_id}/register:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/register:post:description`),
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

		param, err := collectLoraDevicesRegisterCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLoraDevicesRegisterCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraDevicesRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesRegisterCmd("/lora_devices/{device_id}/register"),
		query:       buildQueryForLoraDevicesRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraDevicesRegisterCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesRegisterCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraDevicesRegisterCmd() (string, error) {
	if LoraDevicesRegisterCmdBody != "" {
		if strings.HasPrefix(LoraDevicesRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesRegisterCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraDevicesRegisterCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraDevicesRegisterCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
