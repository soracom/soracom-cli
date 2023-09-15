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

// LoraDevicesPutTagsCmdDeviceId holds value of 'device_id' option
var LoraDevicesPutTagsCmdDeviceId string

// LoraDevicesPutTagsCmdBody holds contents of request body to be sent
var LoraDevicesPutTagsCmdBody string

func InitLoraDevicesPutTagsCmd() {
	LoraDevicesPutTagsCmd.Flags().StringVar(&LoraDevicesPutTagsCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRaWAN device."))

	LoraDevicesPutTagsCmd.Flags().StringVar(&LoraDevicesPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraDevicesPutTagsCmd.RunE = LoraDevicesPutTagsCmdRunE

	LoraDevicesCmd.AddCommand(LoraDevicesPutTagsCmd)
}

// LoraDevicesPutTagsCmd defines 'put-tags' subcommand
var LoraDevicesPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/lora_devices/{device_id}/tags:put:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/tags:put:description`) + "\n\n" + createLinkToAPIReference("LoraDevice", "putLoraDeviceTags"),
}

func LoraDevicesPutTagsCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectLoraDevicesPutTagsCmdParams(ac)
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
}

func collectLoraDevicesPutTagsCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraDevicesPutTagsCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, LoraDevicesPutTagsCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLoraDevicesPutTagsCmd("/lora_devices/{device_id}/tags"),
		query:       buildQueryForLoraDevicesPutTagsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesPutTagsCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesPutTagsCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesPutTagsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraDevicesPutTagsCmd() (string, error) {
	var b []byte
	var err error

	if LoraDevicesPutTagsCmdBody != "" {
		if strings.HasPrefix(LoraDevicesPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesPutTagsCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraDevicesPutTagsCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraDevicesPutTagsCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
