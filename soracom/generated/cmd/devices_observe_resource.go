// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// DevicesObserveResourceCmdDeviceId holds value of 'device_id' option
var DevicesObserveResourceCmdDeviceId string

// DevicesObserveResourceCmdInstance holds value of 'instance' option
var DevicesObserveResourceCmdInstance string

// DevicesObserveResourceCmdObject holds value of 'object' option
var DevicesObserveResourceCmdObject string

// DevicesObserveResourceCmdResource holds value of 'resource' option
var DevicesObserveResourceCmdResource string

// DevicesObserveResourceCmdModel holds value of 'model' option
var DevicesObserveResourceCmdModel bool

func init() {
	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdObject, "object", "", TRAPI("Object ID"))

	DevicesObserveResourceCmd.Flags().StringVar(&DevicesObserveResourceCmdResource, "resource", "", TRAPI("Resource ID"))

	DevicesObserveResourceCmd.Flags().BoolVar(&DevicesObserveResourceCmdModel, "model", false, TRAPI("Whether or not to add model information"))
	DevicesCmd.AddCommand(DevicesObserveResourceCmd)
}

// DevicesObserveResourceCmd defines 'observe-resource' subcommand
var DevicesObserveResourceCmd = &cobra.Command{
	Use:   "observe-resource",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/{resource}/observe:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/{resource}/observe:post:description`),
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

		param, err := collectDevicesObserveResourceCmdParams(ac)
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

func collectDevicesObserveResourceCmdParams(ac *apiClient) (*apiParams, error) {
	if DevicesObserveResourceCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	if DevicesObserveResourceCmdInstance == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "instance")
	}

	if DevicesObserveResourceCmdObject == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "object")
	}

	if DevicesObserveResourceCmdResource == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "resource")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesObserveResourceCmd("/devices/{device_id}/{object}/{instance}/{resource}/observe"),
		query:  buildQueryForDevicesObserveResourceCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesObserveResourceCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesObserveResourceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesObserveResourceCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesObserveResourceCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	escapedResource := url.PathEscape(DevicesObserveResourceCmdResource)

	path = strReplace(path, "{"+"resource"+"}", escapedResource, -1)

	return path
}

func buildQueryForDevicesObserveResourceCmd() url.Values {
	result := url.Values{}

	if DevicesObserveResourceCmdModel != false {
		result.Add("model", sprintf("%t", DevicesObserveResourceCmdModel))
	}

	return result
}
