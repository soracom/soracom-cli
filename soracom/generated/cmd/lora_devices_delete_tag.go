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
	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdDeviceId, "device-id", "", TRAPI("device ID of the target LoRa device."))

	LoraDevicesDeleteTagCmd.Flags().StringVar(&LoraDevicesDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	LoraDevicesCmd.AddCommand(LoraDevicesDeleteTagCmd)
}

// LoraDevicesDeleteTagCmd defines 'delete-tag' subcommand
var LoraDevicesDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/lora_devices/{device_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/tags/{tag_name}:delete:description`),
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

		param, err := collectLoraDevicesDeleteTagCmdParams(ac)
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

func collectLoraDevicesDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {

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
