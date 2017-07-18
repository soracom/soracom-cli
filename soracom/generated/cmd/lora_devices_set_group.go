package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesSetGroupCmdDeviceId holds value of 'device_id' option
var LoraDevicesSetGroupCmdDeviceId string

// LoraDevicesSetGroupCmdGroupId holds value of 'groupId' option
var LoraDevicesSetGroupCmdGroupId string

// LoraDevicesSetGroupCmdOperatorId holds value of 'operatorId' option
var LoraDevicesSetGroupCmdOperatorId string

// LoraDevicesSetGroupCmdCreatedTime holds value of 'createdTime' option
var LoraDevicesSetGroupCmdCreatedTime int64

// LoraDevicesSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var LoraDevicesSetGroupCmdLastModifiedTime int64

// LoraDevicesSetGroupCmdBody holds contents of request body to be sent
var LoraDevicesSetGroupCmdBody string

func init() {
	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().Int64Var(&LoraDevicesSetGroupCmdCreatedTime, "created-time", 0, TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().Int64Var(&LoraDevicesSetGroupCmdLastModifiedTime, "last-modified-time", 0, TRAPI(""))

	LoraDevicesSetGroupCmd.Flags().StringVar(&LoraDevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraDevicesCmd.AddCommand(LoraDevicesSetGroupCmd)
}

// LoraDevicesSetGroupCmd defines 'set-group' subcommand
var LoraDevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/lora_devices/{device_id}/set_group:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/set_group:post:description`),
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

		param, err := collectLoraDevicesSetGroupCmdParams()
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

func collectLoraDevicesSetGroupCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraDevicesSetGroupCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesSetGroupCmd("/lora_devices/{device_id}/set_group"),
		query:       buildQueryForLoraDevicesSetGroupCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraDevicesSetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesSetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesSetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraDevicesSetGroupCmd() (string, error) {
	if LoraDevicesSetGroupCmdBody != "" {
		if strings.HasPrefix(LoraDevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesSetGroupCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraDevicesSetGroupCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraDevicesSetGroupCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LoraDevicesSetGroupCmdGroupId != "" {
		result["groupId"] = LoraDevicesSetGroupCmdGroupId
	}

	if LoraDevicesSetGroupCmdOperatorId != "" {
		result["operatorId"] = LoraDevicesSetGroupCmdOperatorId
	}

	if LoraDevicesSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = LoraDevicesSetGroupCmdCreatedTime
	}

	if LoraDevicesSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = LoraDevicesSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
