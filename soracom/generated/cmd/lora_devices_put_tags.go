package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesPutTagsCmdDeviceId holds value of 'device_id' option
var LoraDevicesPutTagsCmdDeviceId string

// LoraDevicesPutTagsCmdBody holds contents of request body to be sent
var LoraDevicesPutTagsCmdBody string

func init() {
	LoraDevicesPutTagsCmd.Flags().StringVar(&LoraDevicesPutTagsCmdDeviceId, "device-id", "", TR("lora_devices.put_lora_device_tags.put.parameters.device_id.description"))

	LoraDevicesPutTagsCmd.Flags().StringVar(&LoraDevicesPutTagsCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	LoraDevicesCmd.AddCommand(LoraDevicesPutTagsCmd)
}

// LoraDevicesPutTagsCmd defines 'put-tags' subcommand
var LoraDevicesPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TR("lora_devices.put_lora_device_tags.put.summary"),
	Long:  TR(`lora_devices.put_lora_device_tags.put.description`),
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

		param, err := collectLoraDevicesPutTagsCmdParams()
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

func collectLoraDevicesPutTagsCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraDevicesPutTagsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLoraDevicesPutTagsCmd("/lora_devices/{device_id}/tags"),
		query:       buildQueryForLoraDevicesPutTagsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraDevicesPutTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesPutTagsCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesPutTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraDevicesPutTagsCmd() (string, error) {
	if LoraDevicesPutTagsCmdBody != "" {
		if strings.HasPrefix(LoraDevicesPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesPutTagsCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraDevicesPutTagsCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraDevicesPutTagsCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
