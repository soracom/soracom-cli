// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersSessionEventsCmdImsi holds value of 'imsi' option
var SubscribersSessionEventsCmdImsi string

// SubscribersSessionEventsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SubscribersSessionEventsCmdLastEvaluatedKey string

// SubscribersSessionEventsCmdFrom holds value of 'from' option
var SubscribersSessionEventsCmdFrom int64

// SubscribersSessionEventsCmdLimit holds value of 'limit' option
var SubscribersSessionEventsCmdLimit int64

// SubscribersSessionEventsCmdTo holds value of 'to' option
var SubscribersSessionEventsCmdTo int64

// SubscribersSessionEventsCmdOutputJSONL indicates to output with jsonl format
var SubscribersSessionEventsCmdOutputJSONL bool

func InitSubscribersSessionEventsCmd() {
	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The time stamp of the last event retrieved on the previous page. By specifying this parameter, you can continue to retrieve the list from the next event onward."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdFrom, "from", 0, TRAPI("Start time for the events search range (UNIX time in milliseconds)."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdLimit, "limit", 0, TRAPI("Maximum number of events to retrieve."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdTo, "to", 0, TRAPI("End time for the events search range (UNIX time in milliseconds)."))

	SubscribersSessionEventsCmd.Flags().BoolVar(&SubscribersSessionEventsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SubscribersSessionEventsCmd.RunE = SubscribersSessionEventsCmdRunE

	SubscribersCmd.AddCommand(SubscribersSessionEventsCmd)
}

// SubscribersSessionEventsCmd defines 'session-events' subcommand
var SubscribersSessionEventsCmd = &cobra.Command{
	Use:   "session-events",
	Short: TRAPI("/subscribers/{imsi}/events/sessions:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/events/sessions:get:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "listSessionEvents"),
}

func SubscribersSessionEventsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersSessionEventsCmdParams(ac)
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
		if SubscribersSessionEventsCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSubscribersSessionEventsCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersSessionEventsCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersSessionEventsCmd("/subscribers/{imsi}/events/sessions"),
		query:  buildQueryForSubscribersSessionEventsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersSessionEventsCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersSessionEventsCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersSessionEventsCmd() url.Values {
	result := url.Values{}

	if SubscribersSessionEventsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SubscribersSessionEventsCmdLastEvaluatedKey)
	}

	if SubscribersSessionEventsCmdFrom != 0 {
		result.Add("from", sprintf("%d", SubscribersSessionEventsCmdFrom))
	}

	if SubscribersSessionEventsCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SubscribersSessionEventsCmdLimit))
	}

	if SubscribersSessionEventsCmdTo != 0 {
		result.Add("to", sprintf("%d", SubscribersSessionEventsCmdTo))
	}

	return result
}
