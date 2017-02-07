package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesDeleteTagCmdDeviceId holds value of 'device_id' option
var LoraDevicesDeleteTagCmdDeviceId string

// LoraDevicesDeleteTagCmdTagName holds value of 'tag_name' option
var LoraDevicesDeleteTagCmdTagName string

func init() {
	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdDeviceId, "device-id", "", TR("lora_devices.delete_lora_device_tag.delete.parameters.device_id.description"))

	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdTagName, "tag-name", "", TR("lora_devices.delete_lora_device_tag.delete.parameters.tag_name.description"))

	LoraDevicesCmd.AddCommand(LoraDevicesDeleteTagCmd)
}

// LoraDevicesDeleteTagCmd defines 'delete-tag' subcommand
var LoraDevicesDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TR("lora_devices.delete_lora_device_tag.delete.summary"),
	Long:  TR(`lora_devices.delete_lora_device_tag.delete.description`),
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

		param, err := collectLoraDevicesDeleteTagCmdParams()
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

func collectLoraDevicesDeleteTagCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraDevicesDeleteTagCmd("/lora_devices/{device_id}/tags/{tag_name}"),
		query:  buildQueryForLoraDevicesDeleteTagCmd(),
	}, nil
}

func buildPathForLoraDevicesDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesDeleteTagCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", LoraDevicesDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForLoraDevicesDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
