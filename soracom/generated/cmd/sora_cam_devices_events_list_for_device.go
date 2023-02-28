// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesEventsListForDeviceCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesEventsListForDeviceCmdDeviceId string

// SoraCamDevicesEventsListForDeviceCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraCamDevicesEventsListForDeviceCmdLastEvaluatedKey string

// SoraCamDevicesEventsListForDeviceCmdSort holds value of 'sort' option
var SoraCamDevicesEventsListForDeviceCmdSort string

// SoraCamDevicesEventsListForDeviceCmdFrom holds value of 'from' option
var SoraCamDevicesEventsListForDeviceCmdFrom int64

// SoraCamDevicesEventsListForDeviceCmdLimit holds value of 'limit' option
var SoraCamDevicesEventsListForDeviceCmdLimit int64

// SoraCamDevicesEventsListForDeviceCmdTo holds value of 'to' option
var SoraCamDevicesEventsListForDeviceCmdTo int64

// SoraCamDevicesEventsListForDeviceCmdPaginate indicates to do pagination or not
var SoraCamDevicesEventsListForDeviceCmdPaginate bool

// SoraCamDevicesEventsListForDeviceCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesEventsListForDeviceCmdOutputJSONL bool

func init() {
	SoraCamDevicesEventsListForDeviceCmd.Flags().StringVar(&SoraCamDevicesEventsListForDeviceCmdDeviceId, "device-id", "", TRAPI("ID of the target SoraCam device."))

	SoraCamDevicesEventsListForDeviceCmd.Flags().StringVar(&SoraCamDevicesEventsListForDeviceCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("Value of the x-soracom-next-key header in the response to the last listSoraCamDeviceEventsForDevice request. By specifying this parameter, you can continue to retrieve the list from the last request."))

	SoraCamDevicesEventsListForDeviceCmd.Flags().StringVar(&SoraCamDevicesEventsListForDeviceCmdSort, "sort", "desc", TRAPI("Sort order of the events.- `desc`: Descending order (latest data entry first)- `asc`: Ascending order (oldest data entry first)"))

	SoraCamDevicesEventsListForDeviceCmd.Flags().Int64Var(&SoraCamDevicesEventsListForDeviceCmdFrom, "from", 0, TRAPI("Start time of the events to be searched (unix time in milliseconds). If not specified, `from` is set to the oldest event time."))

	SoraCamDevicesEventsListForDeviceCmd.Flags().Int64Var(&SoraCamDevicesEventsListForDeviceCmdLimit, "limit", 10, TRAPI("Maximum number of items to retrieve in one request. Note that the response may contain fewer items than the specified limit."))

	SoraCamDevicesEventsListForDeviceCmd.Flags().Int64Var(&SoraCamDevicesEventsListForDeviceCmdTo, "to", 0, TRAPI("End time of the events to be searched (unix time in milliseconds). If not specified, `to` is set to the current time."))

	SoraCamDevicesEventsListForDeviceCmd.Flags().BoolVar(&SoraCamDevicesEventsListForDeviceCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	SoraCamDevicesEventsListForDeviceCmd.Flags().BoolVar(&SoraCamDevicesEventsListForDeviceCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	SoraCamDevicesEventsCmd.AddCommand(SoraCamDevicesEventsListForDeviceCmd)
}

// SoraCamDevicesEventsListForDeviceCmd defines 'list-for-device' subcommand
var SoraCamDevicesEventsListForDeviceCmd = &cobra.Command{
	Use:   "list-for-device",
	Short: TRAPI("/sora_cam/devices/{device_id}/events:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/events:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceEventsForDevice"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectSoraCamDevicesEventsListForDeviceCmdParams(ac)
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
			if SoraCamDevicesEventsListForDeviceCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSoraCamDevicesEventsListForDeviceCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesEventsListForDeviceCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesEventsListForDeviceCmd("/sora_cam/devices/{device_id}/events"),
		query:  buildQueryForSoraCamDevicesEventsListForDeviceCmd(),

		doPagination:                      SoraCamDevicesEventsListForDeviceCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesEventsListForDeviceCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesEventsListForDeviceCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesEventsListForDeviceCmd() url.Values {
	result := url.Values{}

	if SoraCamDevicesEventsListForDeviceCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraCamDevicesEventsListForDeviceCmdLastEvaluatedKey)
	}

	if SoraCamDevicesEventsListForDeviceCmdSort != "desc" {
		result.Add("sort", SoraCamDevicesEventsListForDeviceCmdSort)
	}

	if SoraCamDevicesEventsListForDeviceCmdFrom != 0 {
		result.Add("from", sprintf("%d", SoraCamDevicesEventsListForDeviceCmdFrom))
	}

	if SoraCamDevicesEventsListForDeviceCmdLimit != 10 {
		result.Add("limit", sprintf("%d", SoraCamDevicesEventsListForDeviceCmdLimit))
	}

	if SoraCamDevicesEventsListForDeviceCmdTo != 0 {
		result.Add("to", sprintf("%d", SoraCamDevicesEventsListForDeviceCmdTo))
	}

	return result
}
