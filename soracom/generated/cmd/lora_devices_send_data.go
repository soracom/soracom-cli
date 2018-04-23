package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesSendDataCmdData holds value of 'data' option
var LoraDevicesSendDataCmdData string

// LoraDevicesSendDataCmdDeviceId holds value of 'device_id' option
var LoraDevicesSendDataCmdDeviceId string

// LoraDevicesSendDataCmdBody holds contents of request body to be sent
var LoraDevicesSendDataCmdBody string

func init() {
	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdData, "data", "", TRAPI(""))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdDeviceId, "device-id", "", TRAPI("ID of the recipient device."))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraDevicesCmd.AddCommand(LoraDevicesSendDataCmd)
}

// LoraDevicesSendDataCmd defines 'send-data' subcommand
var LoraDevicesSendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: TRAPI("/lora_devices/{device_id}/data:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/data:post:description`),
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

		param, err := collectLoraDevicesSendDataCmdParams(ac)
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

func collectLoraDevicesSendDataCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLoraDevicesSendDataCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesSendDataCmd("/lora_devices/{device_id}/data"),
		query:       buildQueryForLoraDevicesSendDataCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraDevicesSendDataCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesSendDataCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesSendDataCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraDevicesSendDataCmd() (string, error) {
	if LoraDevicesSendDataCmdBody != "" {
		if strings.HasPrefix(LoraDevicesSendDataCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesSendDataCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraDevicesSendDataCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraDevicesSendDataCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LoraDevicesSendDataCmdData != "" {
		result["data"] = LoraDevicesSendDataCmdData
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
