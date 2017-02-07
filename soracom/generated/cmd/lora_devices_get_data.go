package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesGetDataCmdDeviceId holds value of 'device_id' option
var LoraDevicesGetDataCmdDeviceId string

// LoraDevicesGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraDevicesGetDataCmdLastEvaluatedKey string

// LoraDevicesGetDataCmdSort holds value of 'sort' option
var LoraDevicesGetDataCmdSort string

// LoraDevicesGetDataCmdFrom holds value of 'from' option
var LoraDevicesGetDataCmdFrom int64

// LoraDevicesGetDataCmdLimit holds value of 'limit' option
var LoraDevicesGetDataCmdLimit int64

// LoraDevicesGetDataCmdTo holds value of 'to' option
var LoraDevicesGetDataCmdTo int64

func init() {
	LoraDevicesGetDataCmd.Flags().StringVar(&LoraDevicesGetDataCmdDeviceId, "device-id", "", TR("lora_devices.get_data_from_lora_device.get.parameters.device_id.description"))

	LoraDevicesGetDataCmd.Flags().StringVar(&LoraDevicesGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TR("lora_devices.get_data_from_lora_device.get.parameters.last_evaluated_key.description"))

	LoraDevicesGetDataCmd.Flags().StringVar(&LoraDevicesGetDataCmdSort, "sort", "", TR("lora_devices.get_data_from_lora_device.get.parameters.sort.description"))

	LoraDevicesGetDataCmd.Flags().Int64Var(&LoraDevicesGetDataCmdFrom, "from", 0, TR("lora_devices.get_data_from_lora_device.get.parameters.from.description"))

	LoraDevicesGetDataCmd.Flags().Int64Var(&LoraDevicesGetDataCmdLimit, "limit", 0, TR("lora_devices.get_data_from_lora_device.get.parameters.limit.description"))

	LoraDevicesGetDataCmd.Flags().Int64Var(&LoraDevicesGetDataCmdTo, "to", 0, TR("lora_devices.get_data_from_lora_device.get.parameters.to.description"))

	LoraDevicesCmd.AddCommand(LoraDevicesGetDataCmd)
}

// LoraDevicesGetDataCmd defines 'get-data' subcommand
var LoraDevicesGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TR("lora_devices.get_data_from_lora_device.get.summary"),
	Long:  TR(`lora_devices.get_data_from_lora_device.get.description`),
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

		param, err := collectLoraDevicesGetDataCmdParams()
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

func collectLoraDevicesGetDataCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraDevicesGetDataCmd("/lora_devices/{device_id}/data"),
		query:  buildQueryForLoraDevicesGetDataCmd(),
	}, nil
}

func buildPathForLoraDevicesGetDataCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesGetDataCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesGetDataCmd() string {
	result := []string{}

	if LoraDevicesGetDataCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LoraDevicesGetDataCmdLastEvaluatedKey))
	}

	if LoraDevicesGetDataCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", LoraDevicesGetDataCmdSort))
	}

	if LoraDevicesGetDataCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", LoraDevicesGetDataCmdFrom))
	}

	if LoraDevicesGetDataCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LoraDevicesGetDataCmdLimit))
	}

	if LoraDevicesGetDataCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", LoraDevicesGetDataCmdTo))
	}

	return strings.Join(result, "&")
}
