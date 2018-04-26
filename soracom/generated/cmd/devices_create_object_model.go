package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesCreateObjectModelCmdCreatedTime holds value of 'createdTime' option
var DevicesCreateObjectModelCmdCreatedTime string

// DevicesCreateObjectModelCmdFormat holds value of 'format' option
var DevicesCreateObjectModelCmdFormat string

// DevicesCreateObjectModelCmdLastModifiedTime holds value of 'lastModifiedTime' option
var DevicesCreateObjectModelCmdLastModifiedTime string

// DevicesCreateObjectModelCmdObjectId holds value of 'objectId' option
var DevicesCreateObjectModelCmdObjectId string

// DevicesCreateObjectModelCmdObjectName holds value of 'objectName' option
var DevicesCreateObjectModelCmdObjectName string

// DevicesCreateObjectModelCmdOperatorId holds value of 'operatorId' option
var DevicesCreateObjectModelCmdOperatorId string

// DevicesCreateObjectModelCmdScope holds value of 'scope' option
var DevicesCreateObjectModelCmdScope string

// DevicesCreateObjectModelCmdBody holds contents of request body to be sent
var DevicesCreateObjectModelCmdBody string

func init() {
	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdCreatedTime, "created-time", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdFormat, "format", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdObjectId, "object-id", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdObjectName, "object-name", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdOperatorId, "operator-id", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdScope, "scope", "", TRAPI(""))

	DevicesCreateObjectModelCmd.Flags().StringVar(&DevicesCreateObjectModelCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCmd.AddCommand(DevicesCreateObjectModelCmd)
}

// DevicesCreateObjectModelCmd defines 'create-object-model' subcommand
var DevicesCreateObjectModelCmd = &cobra.Command{
	Use:   "create-object-model",
	Short: TRAPI("/device_object_models:post:summary"),
	Long:  TRAPI(`/device_object_models:post:description`),
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

		param, err := collectDevicesCreateObjectModelCmdParams(ac)
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

func collectDevicesCreateObjectModelCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesCreateObjectModelCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesCreateObjectModelCmd("/device_object_models"),
		query:       buildQueryForDevicesCreateObjectModelCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesCreateObjectModelCmd(path string) string {

	return path
}

func buildQueryForDevicesCreateObjectModelCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesCreateObjectModelCmd() (string, error) {
	if DevicesCreateObjectModelCmdBody != "" {
		if strings.HasPrefix(DevicesCreateObjectModelCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesCreateObjectModelCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if DevicesCreateObjectModelCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return DevicesCreateObjectModelCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if DevicesCreateObjectModelCmdCreatedTime != "" {
		result["createdTime"] = DevicesCreateObjectModelCmdCreatedTime
	}

	if DevicesCreateObjectModelCmdFormat != "" {
		result["format"] = DevicesCreateObjectModelCmdFormat
	}

	if DevicesCreateObjectModelCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = DevicesCreateObjectModelCmdLastModifiedTime
	}

	if DevicesCreateObjectModelCmdObjectId != "" {
		result["objectId"] = DevicesCreateObjectModelCmdObjectId
	}

	if DevicesCreateObjectModelCmdObjectName != "" {
		result["objectName"] = DevicesCreateObjectModelCmdObjectName
	}

	if DevicesCreateObjectModelCmdOperatorId != "" {
		result["operatorId"] = DevicesCreateObjectModelCmdOperatorId
	}

	if DevicesCreateObjectModelCmdScope != "" {
		result["scope"] = DevicesCreateObjectModelCmdScope
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
