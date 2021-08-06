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

// SubscribersSessionEventsCmdPaginate indicates to do pagination or not
var SubscribersSessionEventsCmdPaginate bool

func init() {
	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The time stamp of the last event retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next event onward."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdFrom, "from", 0, TRAPI("Start time for the events search range (unixtime)."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdLimit, "limit", 0, TRAPI("Maximum number of events to retrieve."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdTo, "to", 0, TRAPI("End time for the events search range (unixtime)."))

	SubscribersSessionEventsCmd.Flags().BoolVar(&SubscribersSessionEventsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	SubscribersCmd.AddCommand(SubscribersSessionEventsCmd)
}

// SubscribersSessionEventsCmd defines 'session-events' subcommand
var SubscribersSessionEventsCmd = &cobra.Command{
	Use:   "session-events",
	Short: TRAPI("/subscribers/{imsi}/events/sessions:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/events/sessions:get:description`),
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersSessionEventsCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersSessionEventsCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersSessionEventsCmd("/subscribers/{imsi}/events/sessions"),
		query:  buildQueryForSubscribersSessionEventsCmd(),

		doPagination:                      SubscribersSessionEventsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

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
