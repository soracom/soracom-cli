package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesDeleteObjectModelCmdModelId holds value of 'model_id' option
var DevicesDeleteObjectModelCmdModelId string

func init() {
	DevicesDeleteObjectModelCmd.Flags().StringVar(&DevicesDeleteObjectModelCmdModelId, "model-id", "", TRAPI("Device object model ID"))

	DevicesCmd.AddCommand(DevicesDeleteObjectModelCmd)
}

// DevicesDeleteObjectModelCmd defines 'delete-object-model' subcommand
var DevicesDeleteObjectModelCmd = &cobra.Command{
	Use:   "delete-object-model",
	Short: TRAPI("/device_object_models/{model_id}:delete:summary"),
	Long:  TRAPI(`/device_object_models/{model_id}:delete:description`),
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

		param, err := collectDevicesDeleteObjectModelCmdParams(ac)
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

func collectDevicesDeleteObjectModelCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteObjectModelCmd("/device_object_models/{model_id}"),
		query:  buildQueryForDevicesDeleteObjectModelCmd(),
	}, nil
}

func buildPathForDevicesDeleteObjectModelCmd(path string) string {

	path = strings.Replace(path, "{"+"model_id"+"}", DevicesDeleteObjectModelCmdModelId, -1)

	return path
}

func buildQueryForDevicesDeleteObjectModelCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
