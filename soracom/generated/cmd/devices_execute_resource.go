package cmd

import (
	"encoding/json"

	"io/ioutil"

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
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}/execute:post:description`),
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

		param, err := collectDevicesExecuteResourceCmdParams(ac)
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

func collectDevicesExecuteResourceCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesExecuteResourceCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesExecuteResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}/execute"),
		query:       buildQueryForDevicesExecuteResourceCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesExecuteResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesExecuteResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesExecuteResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesExecuteResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesExecuteResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesExecuteResourceCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
