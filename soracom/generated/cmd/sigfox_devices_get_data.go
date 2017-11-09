package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesGetDataCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesGetDataCmdDeviceId string

// SigfoxDevicesGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SigfoxDevicesGetDataCmdLastEvaluatedKey string

// SigfoxDevicesGetDataCmdSort holds value of 'sort' option
var SigfoxDevicesGetDataCmdSort string

// SigfoxDevicesGetDataCmdFrom holds value of 'from' option
var SigfoxDevicesGetDataCmdFrom int64

// SigfoxDevicesGetDataCmdLimit holds value of 'limit' option
var SigfoxDevicesGetDataCmdLimit int64

// SigfoxDevicesGetDataCmdTo holds value of 'to' option
var SigfoxDevicesGetDataCmdTo int64

func init() {
	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdDeviceId, "device-id", "", TRAPI("Device ID of the target subscriber that generated data entries."))

	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	SigfoxDevicesGetDataCmd.Flags().StringVar(&SigfoxDevicesGetDataCmdSort, "sort", "", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	SigfoxDevicesGetDataCmd.Flags().Int64Var(&SigfoxDevicesGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesGetDataCmd)
}

// SigfoxDevicesGetDataCmd defines 'get-data' subcommand
var SigfoxDevicesGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/sigfox_devices/{device_id}/data:get:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/data:get:description`),
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

		param, err := collectSigfoxDevicesGetDataCmdParams()
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

func collectSigfoxDevicesGetDataCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSigfoxDevicesGetDataCmd("/sigfox_devices/{device_id}/data"),
		query:  buildQueryForSigfoxDevicesGetDataCmd(),
	}, nil
}

func buildPathForSigfoxDevicesGetDataCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesGetDataCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesGetDataCmd() string {
	result := []string{}

	if SigfoxDevicesGetDataCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", SigfoxDevicesGetDataCmdLastEvaluatedKey))
	}

	if SigfoxDevicesGetDataCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", SigfoxDevicesGetDataCmdSort))
	}

	if SigfoxDevicesGetDataCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", SigfoxDevicesGetDataCmdFrom))
	}

	if SigfoxDevicesGetDataCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", SigfoxDevicesGetDataCmdLimit))
	}

	if SigfoxDevicesGetDataCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", SigfoxDevicesGetDataCmdTo))
	}

	return strings.Join(result, "&")
}
