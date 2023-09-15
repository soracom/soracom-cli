// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// StatsFunkGetCmdImsi holds value of 'imsi' option
var StatsFunkGetCmdImsi string

// StatsFunkGetCmdPeriod holds value of 'period' option
var StatsFunkGetCmdPeriod string

// StatsFunkGetCmdFrom holds value of 'from' option
var StatsFunkGetCmdFrom int64

// StatsFunkGetCmdTo holds value of 'to' option
var StatsFunkGetCmdTo int64

// StatsFunkGetCmdOutputJSONL indicates to output with jsonl format
var StatsFunkGetCmdOutputJSONL bool

func InitStatsFunkGetCmd() {
	StatsFunkGetCmd.Flags().StringVar(&StatsFunkGetCmdImsi, "imsi", "", TRAPI("imsi"))

	StatsFunkGetCmd.Flags().StringVar(&StatsFunkGetCmdPeriod, "period", "", TRAPI("Unit of aggregation. minutes outputs the usage report at the finest granularity. However, while the device is connected to the Soracom platform, the amount of usage will be recorded at approximately 5-minute intervals."))

	StatsFunkGetCmd.Flags().Int64Var(&StatsFunkGetCmdFrom, "from", 0, TRAPI("Start time for the aggregate data (UNIX time in seconds)."))

	StatsFunkGetCmd.Flags().Int64Var(&StatsFunkGetCmdTo, "to", 0, TRAPI("End time for the aggregate data (UNIX time in seconds)."))

	StatsFunkGetCmd.Flags().BoolVar(&StatsFunkGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	StatsFunkGetCmd.RunE = StatsFunkGetCmdRunE

	StatsFunkCmd.AddCommand(StatsFunkGetCmd)
}

// StatsFunkGetCmd defines 'get' subcommand
var StatsFunkGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/funk/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/stats/funk/subscribers/{imsi}:get:description`) + "\n\n" + createLinkToAPIReference("Stats", "getFunkStats"),
}

func StatsFunkGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectStatsFunkGetCmdParams(ac)
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
		if StatsFunkGetCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectStatsFunkGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, StatsFunkGetCmdImsi)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("period", "period", "query", parsedBody, StatsFunkGetCmdPeriod)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("from", "from", "query", parsedBody, StatsFunkGetCmdFrom)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("to", "to", "query", parsedBody, StatsFunkGetCmdTo)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsFunkGetCmd("/stats/funk/subscribers/{imsi}"),
		query:  buildQueryForStatsFunkGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForStatsFunkGetCmd(path string) string {

	escapedImsi := url.PathEscape(StatsFunkGetCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForStatsFunkGetCmd() url.Values {
	result := url.Values{}

	if StatsFunkGetCmdPeriod != "" {
		result.Add("period", StatsFunkGetCmdPeriod)
	}

	if StatsFunkGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", StatsFunkGetCmdFrom))
	}

	if StatsFunkGetCmdTo != 0 {
		result.Add("to", sprintf("%d", StatsFunkGetCmdTo))
	}

	return result
}
