package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

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

func collectDevicesUpdateObjectModelCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesUpdateObjectModelCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesUpdateObjectModelCmd("/device_object_models/{model_id}"),
		query:       buildQueryForDevicesUpdateObjectModelCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesUpdateObjectModelCmd(path string) string {

	path = strings.Replace(path, "{"+"model_id"+"}", DevicesUpdateObjectModelCmdModelId, -1)

	return path
}

func buildQueryForDevicesUpdateObjectModelCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesUpdateObjectModelCmd() (string, error) {
	if DevicesUpdateObjectModelCmdBody != "" {
		if strings.HasPrefix(DevicesUpdateObjectModelCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesUpdateObjectModelCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if DevicesUpdateObjectModelCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return DevicesUpdateObjectModelCmdBody, nil
		}
	}

	result := map[string]interface{}{}

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
