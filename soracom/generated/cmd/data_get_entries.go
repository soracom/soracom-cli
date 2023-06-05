// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

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

// DataGetEntriesCmdPaginate indicates to do pagination or not
var DataGetEntriesCmdPaginate bool

// DataGetEntriesCmdOutputJSONL indicates to output with jsonl format
var DataGetEntriesCmdOutputJSONL bool

func InitDataGetEntriesCmd() {
	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of 'time' in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdResourceId, "resource-id", "", TRAPI("ID of data source resource"))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdResourceType, "resource-type", "", TRAPI("Type of data source resource"))

	DataGetEntriesCmd.Flags().StringVar(&DataGetEntriesCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataGetEntriesCmd.Flags().Int64Var(&DataGetEntriesCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	DataGetEntriesCmd.Flags().BoolVar(&DataGetEntriesCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	DataGetEntriesCmd.Flags().BoolVar(&DataGetEntriesCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	DataGetEntriesCmd.RunE = DataGetEntriesCmdRunE

	DataCmd.AddCommand(DataGetEntriesCmd)
}

// DataGetEntriesCmd defines 'get-entries' subcommand
var DataGetEntriesCmd = &cobra.Command{
	Use:   "get-entries",
	Short: TRAPI("/data/{resource_type}/{resource_id}:get:summary"),
	Long:  TRAPI(`/data/{resource_type}/{resource_id}:get:description`) + "\n\n" + createLinkToAPIReference("DataEntry", "getDataEntries"),
}

func DataGetEntriesCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
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

	param, err := collectDataGetEntriesCmdParams(ac)
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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		if DataGetEntriesCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectDataGetEntriesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error

	err = checkIfRequiredStringParameterIsSupplied("resource_id", "resource-id", "path", parsedBody, DataGetEntriesCmdResourceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("resource_type", "resource-type", "path", parsedBody, DataGetEntriesCmdResourceType)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetEntriesCmd("/data/{resource_type}/{resource_id}"),
		query:  buildQueryForDataGetEntriesCmd(),

		doPagination:                      DataGetEntriesCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDataGetEntriesCmd(path string) string {

	escapedResourceId := url.PathEscape(DataGetEntriesCmdResourceId)

	path = strReplace(path, "{"+"resource_id"+"}", escapedResourceId, -1)

	escapedResourceType := url.PathEscape(DataGetEntriesCmdResourceType)

	path = strReplace(path, "{"+"resource_type"+"}", escapedResourceType, -1)

	return path
}

func buildQueryForDataGetEntriesCmd() url.Values {
	result := url.Values{}

	if DataGetEntriesCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", DataGetEntriesCmdLastEvaluatedKey)
	}

	if DataGetEntriesCmdSort != "desc" {
		result.Add("sort", DataGetEntriesCmdSort)
	}

	if DataGetEntriesCmdFrom != 0 {
		result.Add("from", sprintf("%d", DataGetEntriesCmdFrom))
	}

	if DataGetEntriesCmdLimit != 0 {
		result.Add("limit", sprintf("%d", DataGetEntriesCmdLimit))
	}

	if DataGetEntriesCmdTo != 0 {
		result.Add("to", sprintf("%d", DataGetEntriesCmdTo))
	}

	return result
}
