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

// DevicesGetObjectModelCmdModelId holds value of 'model_id' option
var DevicesGetObjectModelCmdModelId string

func init() {
	DevicesGetObjectModelCmd.Flags().StringVar(&DevicesGetObjectModelCmdModelId, "model-id", "", TRAPI("Device object model ID"))
	DevicesCmd.AddCommand(DevicesGetObjectModelCmd)
}

// DevicesGetObjectModelCmd defines 'get-object-model' subcommand
var DevicesGetObjectModelCmd = &cobra.Command{
	Use:   "get-object-model",
	Short: TRAPI("/device_object_models/{model_id}:get:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}:get:description`),
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

		param, err := collectDevicesGetObjectModelCmdParams(ac)
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

func collectDevicesGetObjectModelCmdParams(ac *apiClient) (*apiParams, error) {
	if DevicesGetObjectModelCmdModelId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "model-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesGetObjectModelCmd("/device_object_models/{model_id}"),
		query:  buildQueryForDevicesGetObjectModelCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesGetObjectModelCmd(path string) string {

	escapedModelId := url.PathEscape(DevicesGetObjectModelCmdModelId)

	path = strReplace(path, "{"+"model_id"+"}", escapedModelId, -1)

	return path
}

func buildQueryForDevicesGetObjectModelCmd() url.Values {
	result := url.Values{}

	return result
}
