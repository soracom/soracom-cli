package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DataGetEntryCmdResourceId holds value of 'resource_id' option
var DataGetEntryCmdResourceId string

// DataGetEntryCmdResourceType holds value of 'resource_type' option
var DataGetEntryCmdResourceType string

// DataGetEntryCmdTime holds value of 'time' option
var DataGetEntryCmdTime int64

func init() {
	DataGetEntryCmd.Flags().StringVar(&DataGetEntryCmdResourceId, "resource-id", "", TRAPI("ID of data source resource"))

	DataGetEntryCmd.Flags().StringVar(&DataGetEntryCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataGetEntryCmd.Flags().Int64Var(&DataGetEntryCmdTime, "time", 0, TRAPI("Timestamp of the target data entry to get (unixtime in milliseconds)."))

	DataCmd.AddCommand(DataGetEntryCmd)
}

// DataGetEntryCmd defines 'get-entry' subcommand
var DataGetEntryCmd = &cobra.Command{
	Use:   "get-entry",
	Short: TRAPI("/data/{resource_type}/{resource_id}/{time}:get:summary"),
	Long:  TRAPI(`/data/{resource_type}/{resource_id}/{time}:get:description`),
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

		param, err := collectDataGetEntryCmdParams(ac)
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

func collectDataGetEntryCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetEntryCmd("/data/{resource_type}/{resource_id}/{time}"),
		query:  buildQueryForDataGetEntryCmd(),
	}, nil
}

func buildPathForDataGetEntryCmd(path string) string {

	path = strings.Replace(path, "{"+"resource_id"+"}", DataGetEntryCmdResourceId, -1)

	path = strings.Replace(path, "{"+"resource_type"+"}", DataGetEntryCmdResourceType, -1)

	return path
}

func buildQueryForDataGetEntryCmd() string {
	result := []string{}

	if DataGetEntryCmdTime != 0 {
		result = append(result, sprintf("%s=%d", "time", DataGetEntryCmdTime))
	}

	return strings.Join(result, "&")
}
