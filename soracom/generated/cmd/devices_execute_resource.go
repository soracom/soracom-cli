// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// DevicesExecuteResourceCmdDeviceId holds value of 'device_id' option
var DevicesExecuteResourceCmdDeviceId string

// DevicesExecuteResourceCmdInstance holds value of 'instance' option
var DevicesExecuteResourceCmdInstance string

// DevicesExecuteResourceCmdObject holds value of 'object' option
var DevicesExecuteResourceCmdObject string

// DevicesExecuteResourceCmdResource holds value of 'resource' option
var DevicesExecuteResourceCmdResource string

// DevicesExecuteResourceCmdBody holds contents of request body to be sent
var DevicesExecuteResourceCmdBody string

func init() {
	DevicesExecuteResourceCmd.Flags().StringVar(&DevicesExecuteResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesExecuteResourceCmd.Flags().StringVar(&DevicesExecuteResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesExecuteResourceCmd.Flags().StringVar(&DevicesExecuteResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesExecuteResourceCmd.Flags().StringVar(&DevicesExecuteResourceCmdResource, "resource", "", TRAPI("Resource ID"))

	DevicesExecuteResourceCmd.Flags().StringVar(&DevicesExecuteResourceCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DevicesCmd.AddCommand(DevicesExecuteResourceCmd)
}

// DevicesExecuteResourceCmd defines 'execute-resource' subcommand
var DevicesExecuteResourceCmd = &cobra.Command{
	Use:   "execute-resource",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/{resource}/execute:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}/execute:post:description`) + "\n\n" + createLinkToAPIReference("Device", "executeDeviceResource"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectDevicesExecuteResourceCmdParams(ac)
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
	},
}

func collectDevicesExecuteResourceCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForDevicesExecuteResourceCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, DevicesExecuteResourceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("instance", "instance", "path", parsedBody, DevicesExecuteResourceCmdInstance)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("object", "object", "path", parsedBody, DevicesExecuteResourceCmdObject)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("resource", "resource", "path", parsedBody, DevicesExecuteResourceCmdResource)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesExecuteResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}/execute"),
		query:       buildQueryForDevicesExecuteResourceCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesExecuteResourceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesExecuteResourceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesExecuteResourceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesExecuteResourceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	escapedResource := url.PathEscape(DevicesExecuteResourceCmdResource)

	path = strReplace(path, "{"+"resource"+"}", escapedResource, -1)

	return path
}

func buildQueryForDevicesExecuteResourceCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesExecuteResourceCmd() (string, error) {
	var result map[string]interface{}

	if DevicesExecuteResourceCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesExecuteResourceCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesExecuteResourceCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesExecuteResourceCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesExecuteResourceCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
