// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// DevicesPutDeviceTagsCmdDeviceId holds value of 'device_id' option
var DevicesPutDeviceTagsCmdDeviceId string

// DevicesPutDeviceTagsCmdBody holds contents of request body to be sent
var DevicesPutDeviceTagsCmdBody string

func init() {
	DevicesPutDeviceTagsCmd.Flags().StringVar(&DevicesPutDeviceTagsCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesPutDeviceTagsCmd.Flags().StringVar(&DevicesPutDeviceTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DevicesCmd.AddCommand(DevicesPutDeviceTagsCmd)
}

// DevicesPutDeviceTagsCmd defines 'put-device-tags' subcommand
var DevicesPutDeviceTagsCmd = &cobra.Command{
	Use:   "put-device-tags",
	Short: TRAPI("/devices/{device_id}/tags:put:summary"),
	Long:  TRAPI(`/devices/{device_id}/tags:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectDevicesPutDeviceTagsCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesPutDeviceTagsCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForDevicesPutDeviceTagsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if DevicesPutDeviceTagsCmdDeviceId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForDevicesPutDeviceTagsCmd("/devices/{device_id}/tags"),
		query:       buildQueryForDevicesPutDeviceTagsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesPutDeviceTagsCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesPutDeviceTagsCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForDevicesPutDeviceTagsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesPutDeviceTagsCmd() (string, error) {
	var b []byte
	var err error

	if DevicesPutDeviceTagsCmdBody != "" {
		if strings.HasPrefix(DevicesPutDeviceTagsCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesPutDeviceTagsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesPutDeviceTagsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesPutDeviceTagsCmdBody)
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
