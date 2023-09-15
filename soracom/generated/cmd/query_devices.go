// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// QueryDevicesCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QueryDevicesCmdLastEvaluatedKey string

// QueryDevicesCmdSearchType holds value of 'search_type' option
var QueryDevicesCmdSearchType string

// QueryDevicesCmdDeviceId holds multiple values of 'deviceId' option
var QueryDevicesCmdDeviceId []string

// QueryDevicesCmdGroup holds multiple values of 'group' option
var QueryDevicesCmdGroup []string

// QueryDevicesCmdImei holds multiple values of 'imei' option
var QueryDevicesCmdImei []string

// QueryDevicesCmdImsi holds multiple values of 'imsi' option
var QueryDevicesCmdImsi []string

// QueryDevicesCmdName holds multiple values of 'name' option
var QueryDevicesCmdName []string

// QueryDevicesCmdTag holds multiple values of 'tag' option
var QueryDevicesCmdTag []string

// QueryDevicesCmdLimit holds value of 'limit' option
var QueryDevicesCmdLimit int64

// QueryDevicesCmdPaginate indicates to do pagination or not
var QueryDevicesCmdPaginate bool

// QueryDevicesCmdOutputJSONL indicates to output with jsonl format
var QueryDevicesCmdOutputJSONL bool

func InitQueryDevicesCmd() {
	QueryDevicesCmd.Flags().StringVar(&QueryDevicesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The SORACOM Inventory device ID of the last Inventory device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next Inventory device onward."))

	QueryDevicesCmd.Flags().StringVar(&QueryDevicesCmdSearchType, "search-type", "and", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdDeviceId, "device-id", []string{}, TRAPI("SORACOM Inventory device ID to search"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdGroup, "group", []string{}, TRAPI("Group name to search"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdImei, "imei", []string{}, TRAPI("IMEI of the device that was used on bootstrapping"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdImsi, "imsi", []string{}, TRAPI("IMSI of the device that was used on bootstrapping"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdName, "name", []string{}, TRAPI("Name to search"))

	QueryDevicesCmd.Flags().StringSliceVar(&QueryDevicesCmdTag, "tag", []string{}, TRAPI("String of tag values to search"))

	QueryDevicesCmd.Flags().Int64Var(&QueryDevicesCmdLimit, "limit", 10, TRAPI("The maximum number of item to retrieve"))

	QueryDevicesCmd.Flags().BoolVar(&QueryDevicesCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	QueryDevicesCmd.Flags().BoolVar(&QueryDevicesCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	QueryDevicesCmd.RunE = QueryDevicesCmdRunE

	QueryCmd.AddCommand(QueryDevicesCmd)
}

// QueryDevicesCmd defines 'devices' subcommand
var QueryDevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: TRAPI("/query/devices:get:summary"),
	Long:  TRAPI(`/query/devices:get:description`) + "\n\n" + createLinkToAPIReference("Query", "searchDevices"),
}

func QueryDevicesCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectQueryDevicesCmdParams(ac)
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
		if QueryDevicesCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectQueryDevicesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQueryDevicesCmd("/query/devices"),
		query:  buildQueryForQueryDevicesCmd(),

		doPagination:                      QueryDevicesCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForQueryDevicesCmd(path string) string {

	return path
}

func buildQueryForQueryDevicesCmd() url.Values {
	result := url.Values{}

	if QueryDevicesCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", QueryDevicesCmdLastEvaluatedKey)
	}

	if QueryDevicesCmdSearchType != "and" {
		result.Add("search_type", QueryDevicesCmdSearchType)
	}

	for _, s := range QueryDevicesCmdDeviceId {
		if s != "" {
			result.Add("deviceId", s)
		}
	}

	for _, s := range QueryDevicesCmdGroup {
		if s != "" {
			result.Add("group", s)
		}
	}

	for _, s := range QueryDevicesCmdImei {
		if s != "" {
			result.Add("imei", s)
		}
	}

	for _, s := range QueryDevicesCmdImsi {
		if s != "" {
			result.Add("imsi", s)
		}
	}

	for _, s := range QueryDevicesCmdName {
		if s != "" {
			result.Add("name", s)
		}
	}

	for _, s := range QueryDevicesCmdTag {
		if s != "" {
			result.Add("tag", s)
		}
	}

	if QueryDevicesCmdLimit != 10 {
		result.Add("limit", sprintf("%d", QueryDevicesCmdLimit))
	}

	return result
}
