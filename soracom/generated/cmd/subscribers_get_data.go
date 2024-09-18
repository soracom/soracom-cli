// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersGetDataCmdImsi holds value of 'imsi' option
var SubscribersGetDataCmdImsi string

// SubscribersGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SubscribersGetDataCmdLastEvaluatedKey string

// SubscribersGetDataCmdSort holds value of 'sort' option
var SubscribersGetDataCmdSort string

// SubscribersGetDataCmdFrom holds value of 'from' option
var SubscribersGetDataCmdFrom int64

// SubscribersGetDataCmdLimit holds value of 'limit' option
var SubscribersGetDataCmdLimit int64

// SubscribersGetDataCmdTo holds value of 'to' option
var SubscribersGetDataCmdTo int64

// SubscribersGetDataCmdOutputJSONL indicates to output with jsonl format
var SubscribersGetDataCmdOutputJSONL bool

func InitSubscribersGetDataCmd() {
	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber that generated data entries."))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Key in the last data entry retrieved in the previous page. The key for data entries in this API is '${UNIX time in milliseconds}_${IMSI}'. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (UNIX time in milliseconds)."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve (value range is 1 to 1000). The default is '10'."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (UNIX time in milliseconds)."))

	SubscribersGetDataCmd.Flags().BoolVar(&SubscribersGetDataCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SubscribersGetDataCmd.RunE = SubscribersGetDataCmdRunE

	SubscribersCmd.AddCommand(SubscribersGetDataCmd)
}

// SubscribersGetDataCmd defines 'get-data' subcommand
var SubscribersGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/subscribers/{imsi}/data:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/data:get:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "getDataFromSubscriber"),
}

func SubscribersGetDataCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersGetDataCmdParams(ac)
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
		if SubscribersGetDataCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSubscribersGetDataCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersGetDataCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersGetDataCmd("/subscribers/{imsi}/data"),
		query:  buildQueryForSubscribersGetDataCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersGetDataCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersGetDataCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersGetDataCmd() url.Values {
	result := url.Values{}

	if SubscribersGetDataCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SubscribersGetDataCmdLastEvaluatedKey)
	}

	if SubscribersGetDataCmdSort != "desc" {
		result.Add("sort", SubscribersGetDataCmdSort)
	}

	if SubscribersGetDataCmdFrom != 0 {
		result.Add("from", sprintf("%d", SubscribersGetDataCmdFrom))
	}

	if SubscribersGetDataCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SubscribersGetDataCmdLimit))
	}

	if SubscribersGetDataCmdTo != 0 {
		result.Add("to", sprintf("%d", SubscribersGetDataCmdTo))
	}

	return result
}
