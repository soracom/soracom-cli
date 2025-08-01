// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesEventsListCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesEventsListCmdDeviceId string

// SoraCamDevicesEventsListCmdSort holds value of 'sort' option
var SoraCamDevicesEventsListCmdSort string

// SoraCamDevicesEventsListCmdFrom holds value of 'from' option
var SoraCamDevicesEventsListCmdFrom int64

// SoraCamDevicesEventsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesEventsListCmdLastEvaluatedKey int64

// SoraCamDevicesEventsListCmdLimit holds value of 'limit' option
var SoraCamDevicesEventsListCmdLimit int64

// SoraCamDevicesEventsListCmdTo holds value of 'to' option
var SoraCamDevicesEventsListCmdTo int64

// SoraCamDevicesEventsListCmdPaginate indicates to do pagination or not
var SoraCamDevicesEventsListCmdPaginate bool

// SoraCamDevicesEventsListCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesEventsListCmdOutputJSONL bool

func InitSoraCamDevicesEventsListCmd() {
	SoraCamDevicesEventsListCmd.Flags().StringVar(&SoraCamDevicesEventsListCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device. If not specified, retrieves event history for all SoraCam compatible camera devices owned by the operator."))

	SoraCamDevicesEventsListCmd.Flags().StringVar(&SoraCamDevicesEventsListCmdSort, "sort", "desc", TRAPI("Event list sort order.- 'desc': Descending order (newest events detected by SoraCam compatible camera device first).- 'asc': Ascending order (oldest events detected by SoraCam compatible camera device first)."))

	SoraCamDevicesEventsListCmd.Flags().Int64Var(&SoraCamDevicesEventsListCmdFrom, "from", 0, TRAPI("Start time of the events to retrieve (unix time in milliseconds). If not specified, 'from' is set to the oldest event time."))

	SoraCamDevicesEventsListCmd.Flags().Int64Var(&SoraCamDevicesEventsListCmdLastEvaluatedKey, "last-evaluated-key", 0, TRAPI("Value of the 'x-soracom-next-key' header in the previous response. Use this to continue pagination."))

	SoraCamDevicesEventsListCmd.Flags().Int64Var(&SoraCamDevicesEventsListCmdLimit, "limit", 10, TRAPI("Maximum number of events to retrieve in one request (value range is 1 to 1000). Note that the response may contain fewer events than the specified limit."))

	SoraCamDevicesEventsListCmd.Flags().Int64Var(&SoraCamDevicesEventsListCmdTo, "to", 0, TRAPI("End time of the events to retrieve (unix time in milliseconds). If not specified, 'to' is set to the current time."))

	SoraCamDevicesEventsListCmd.Flags().BoolVar(&SoraCamDevicesEventsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SoraCamDevicesEventsListCmd.Flags().BoolVar(&SoraCamDevicesEventsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SoraCamDevicesEventsListCmd.RunE = SoraCamDevicesEventsListCmdRunE

	SoraCamDevicesEventsCmd.AddCommand(SoraCamDevicesEventsListCmd)
}

// SoraCamDevicesEventsListCmd defines 'list' subcommand
var SoraCamDevicesEventsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/sora_cam/devices/events:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/events:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceEvents"),
}

func SoraCamDevicesEventsListCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectSoraCamDevicesEventsListCmdParams(ac)
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
		if SoraCamDevicesEventsListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamDevicesEventsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesEventsListCmd("/sora_cam/devices/events"),
		query:  buildQueryForSoraCamDevicesEventsListCmd(),

		doPagination:                      SoraCamDevicesEventsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesEventsListCmd(path string) string {

	return path
}

func buildQueryForSoraCamDevicesEventsListCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesEventsListCmdDeviceId != "" {
		result.Add("device_id", SoraCamDevicesEventsListCmdDeviceId)
	}

	if SoraCamDevicesEventsListCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesEventsListCmdSort)
	}

	if SoraCamDevicesEventsListCmdFrom != 0 {
		result.Add("from", sprintf("%d", SoraCamDevicesEventsListCmdFrom))
	}

	if SoraCamDevicesEventsListCmdLastEvaluatedKey != 0 {
		result.Add("last_evaluated_key", sprintf("%d", SoraCamDevicesEventsListCmdLastEvaluatedKey))
	}

	if SoraCamDevicesEventsListCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesEventsListCmdLimit))
	}

	if SoraCamDevicesEventsListCmdTo != 0 {
		result.Add("to", sprintf("%d", SoraCamDevicesEventsListCmdTo))
	}

	return result
}
