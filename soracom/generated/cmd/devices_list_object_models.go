package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesListObjectModelsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DevicesListObjectModelsCmdLastEvaluatedKey string

// DevicesListObjectModelsCmdLimit holds value of 'limit' option
var DevicesListObjectModelsCmdLimit int64

func init() {
	DevicesListObjectModelsCmd.Flags().StringVar(&DevicesListObjectModelsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("ID of the last device object model in the previous page"))

	DevicesListObjectModelsCmd.Flags().Int64Var(&DevicesListObjectModelsCmdLimit, "limit", 0, TRAPI("Max number of device object models in a response"))

	DevicesCmd.AddCommand(DevicesListObjectModelsCmd)
}

// DevicesListObjectModelsCmd defines 'list-object-models' subcommand
var DevicesListObjectModelsCmd = &cobra.Command{
	Use:   "list-object-models",
	Short: TRAPI("/device_object_models:get:summary"),
	Long:  TRAPI(`/device_object_models:get:description`),
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

		param, err := collectDevicesListObjectModelsCmdParams(ac)
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

func collectDevicesListObjectModelsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesListObjectModelsCmd("/device_object_models"),
		query:  buildQueryForDevicesListObjectModelsCmd(),
	}, nil
}

func buildPathForDevicesListObjectModelsCmd(path string) string {

	return path
}

func buildQueryForDevicesListObjectModelsCmd() string {
	result := []string{}

	if DevicesListObjectModelsCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", DevicesListObjectModelsCmdLastEvaluatedKey))
	}

	if DevicesListObjectModelsCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", DevicesListObjectModelsCmdLimit))
	}

	return strings.Join(result, "&")
}
