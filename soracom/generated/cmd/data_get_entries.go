package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DataGetEntriesCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DataGetEntriesCmdLastEvaluatedKey string

// DataGetEntriesCmdResourceId holds value of 'resource_id' option
var DataGetEntriesCmdResourceId string

// DataGetEntriesCmdResourceType holds value of 'resource_type' option
var DataGetEntriesCmdResourceType string

// DataGetEntriesCmdSort holds value of 'sort' option
var DataGetEntriesCmdSort string

// DataGetEntriesCmdFrom holds value of 'from' option
var DataGetEntriesCmdFrom int64

// DataGetEntriesCmdLimit holds value of 'limit' option
var DataGetEntriesCmdLimit int64

// DataGetEntriesCmdTo holds value of 'to' option
var DataGetEntriesCmdTo int64

func init() {
	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdResourceId, "resource-id", "", TRAPI("ID of data source resource"))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdSort, "sort", "", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	DataCmd.AddCommand(DataGetEntriesCmd)
}

// DataGetEntriesCmd defines 'get-entries' subcommand
var DataGetEntriesCmd = &cobra.Command{
	Use:   "get-entries",
	Short: TRAPI("/data/{resource_type}/{resource_id}:get:summary"),
	Long:  TRAPI(`/data/{resource_type}/{resource_id}:get:description`),
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

		param, err := collectDataGetEntriesCmdParams(ac)
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

func collectDataGetEntriesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetEntriesCmd("/data/{resource_type}/{resource_id}"),
		query:  buildQueryForDataGetEntriesCmd(),
	}, nil
}

func buildPathForDataGetEntriesCmd(path string) string {

	path = strings.Replace(path, "{"+"resource_id"+"}", DataGetEntriesCmdResourceId, -1)

	path = strings.Replace(path, "{"+"resource_type"+"}", DataGetEntriesCmdResourceType, -1)

	return path
}

func buildQueryForDataGetEntriesCmd() string {
	result := []string{}

	if DataGetEntriesCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", DataGetEntriesCmdLastEvaluatedKey))
	}

	if DataGetEntriesCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", DataGetEntriesCmdSort))
	}

	if DataGetEntriesCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", DataGetEntriesCmdFrom))
	}

	if DataGetEntriesCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", DataGetEntriesCmdLimit))
	}

	if DataGetEntriesCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", DataGetEntriesCmdTo))
	}

	return strings.Join(result, "&")
}
