package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesPutResourceCmdDeviceId holds value of 'deviceId' option
var DevicesPutResourceCmdDeviceId string

// DevicesPutResourceCmdInstance holds value of 'instance' option
var DevicesPutResourceCmdInstance string

// DevicesPutResourceCmdObject holds value of 'object' option
var DevicesPutResourceCmdObject string

// DevicesPutResourceCmdResource holds value of 'resource' option
var DevicesPutResourceCmdResource string

// DevicesPutResourceCmdBody holds contents of request body to be sent
var DevicesPutResourceCmdBody string

func init() {
	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdResource, "resource", "", TRAPI("Resource ID"))

	DevicesPutResourceCmd.Flags().StringVar(&DevicesPutResourceCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCmd.AddCommand(DevicesPutResourceCmd)
}

// DevicesPutResourceCmd defines 'put-resource' subcommand
var DevicesPutResourceCmd = &cobra.Command{
	Use:   "put-resource",
	Short: TRAPI("/devices/{deviceId}/{object}/{instance}/{resource}:put:summary"),
	Long:  TRAPI(`/devices/{deviceId}/{object}/{instance}/{resource}:put:description`),
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

		param, err := collectDevicesPutResourceCmdParams()
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

func collectDevicesPutResourceCmdParams() (*apiParams, error) {

	body, err := buildBodyForDevicesPutResourceCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForDevicesPutResourceCmd("/devices/{deviceId}/{object}/{instance}/{resource}"),
		query:       buildQueryForDevicesPutResourceCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesPutResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesPutResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesPutResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesPutResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesPutResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesPutResourceCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesPutResourceCmd() (string, error) {
	if DevicesPutResourceCmdBody != "" {
		if strings.HasPrefix(DevicesPutResourceCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesPutResourceCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if DevicesPutResourceCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return DevicesPutResourceCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
