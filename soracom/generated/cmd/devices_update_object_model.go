// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

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

// DevicesUpdateObjectModelCmdCreatedTime holds value of 'createdTime' option
var DevicesUpdateObjectModelCmdCreatedTime string

// DevicesUpdateObjectModelCmdFormat holds value of 'format' option
var DevicesUpdateObjectModelCmdFormat string

// DevicesUpdateObjectModelCmdLastModifiedTime holds value of 'lastModifiedTime' option
var DevicesUpdateObjectModelCmdLastModifiedTime string

// DevicesUpdateObjectModelCmdModelId holds value of 'model_id' option
var DevicesUpdateObjectModelCmdModelId string

// DevicesUpdateObjectModelCmdObjectId holds value of 'objectId' option
var DevicesUpdateObjectModelCmdObjectId string

// DevicesUpdateObjectModelCmdObjectName holds value of 'objectName' option
var DevicesUpdateObjectModelCmdObjectName string

// DevicesUpdateObjectModelCmdOperatorId holds value of 'operatorId' option
var DevicesUpdateObjectModelCmdOperatorId string

// DevicesUpdateObjectModelCmdScope holds value of 'scope' option
var DevicesUpdateObjectModelCmdScope string

// DevicesUpdateObjectModelCmdBody holds contents of request body to be sent
var DevicesUpdateObjectModelCmdBody string

func init() {
	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdCreatedTime, "created-time", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdFormat, "format", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdModelId, "model-id", "", TRAPI("Device object model ID"))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdObjectId, "object-id", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdObjectName, "object-name", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdOperatorId, "operator-id", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdScope, "scope", "", TRAPI(""))

	DevicesUpdateObjectModelCmd.Flags().StringVar(&DevicesUpdateObjectModelCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	DevicesCmd.AddCommand(DevicesUpdateObjectModelCmd)
}

// DevicesUpdateObjectModelCmd defines 'update-object-model' subcommand
var DevicesUpdateObjectModelCmd = &cobra.Command{
	Use:   "update-object-model",
	Short: TRAPI("/device_object_models/{model_id}:post:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}:post:description`),
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

		param, err := collectDevicesUpdateObjectModelCmdParams(ac)
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

func collectDevicesUpdateObjectModelCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForDevicesUpdateObjectModelCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if DevicesUpdateObjectModelCmdModelId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "model-id")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesUpdateObjectModelCmd("/device_object_models/{model_id}"),
		query:       buildQueryForDevicesUpdateObjectModelCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesUpdateObjectModelCmd(path string) string {

	escapedModelId := url.PathEscape(DevicesUpdateObjectModelCmdModelId)

	path = strReplace(path, "{"+"model_id"+"}", escapedModelId, -1)

	return path
}

func buildQueryForDevicesUpdateObjectModelCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForDevicesUpdateObjectModelCmd() (string, error) {
	var result map[string]interface{}

	if DevicesUpdateObjectModelCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(DevicesUpdateObjectModelCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesUpdateObjectModelCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if DevicesUpdateObjectModelCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(DevicesUpdateObjectModelCmdBody)
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

	if DevicesUpdateObjectModelCmdCreatedTime != "" {
		result["createdTime"] = DevicesUpdateObjectModelCmdCreatedTime
	}

	if DevicesUpdateObjectModelCmdFormat != "" {
		result["format"] = DevicesUpdateObjectModelCmdFormat
	}

	if DevicesUpdateObjectModelCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = DevicesUpdateObjectModelCmdLastModifiedTime
	}

	if DevicesUpdateObjectModelCmdObjectId != "" {
		result["objectId"] = DevicesUpdateObjectModelCmdObjectId
	}

	if DevicesUpdateObjectModelCmdObjectName != "" {
		result["objectName"] = DevicesUpdateObjectModelCmdObjectName
	}

	if DevicesUpdateObjectModelCmdOperatorId != "" {
		result["operatorId"] = DevicesUpdateObjectModelCmdOperatorId
	}

	if DevicesUpdateObjectModelCmdScope != "" {
		result["scope"] = DevicesUpdateObjectModelCmdScope
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
