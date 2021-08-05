// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

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

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDataGetEntryCmdParams(ac *apiClient) (*apiParams, error) {
	if DataGetEntryCmdResourceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "resource-id")
	}

	if DataGetEntryCmdResourceType == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "resource-type")
	}

	if DataGetEntryCmdTime == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "time")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetEntryCmd("/data/{resource_type}/{resource_id}/{time}"),
		query:  buildQueryForDataGetEntryCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDataGetEntryCmd(path string) string {

	escapedResourceId := url.PathEscape(DataGetEntryCmdResourceId)

	path = strReplace(path, "{"+"resource_id"+"}", escapedResourceId, -1)

	escapedResourceType := url.PathEscape(DataGetEntryCmdResourceType)

	path = strReplace(path, "{"+"resource_type"+"}", escapedResourceType, -1)

	path = strReplace(path, "{"+"time"+"}", url.PathEscape(sprintf("%d", DataGetEntryCmdTime)), -1)

	return path
}

func buildQueryForDataGetEntryCmd() url.Values {
	result := url.Values{}

	return result
}
