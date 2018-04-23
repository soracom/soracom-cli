package cmd

import (
	"os"
	"strings"

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

func init() {
	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The time stamp of the last event retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next event onward."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdFrom, "from", 0, TRAPI("Start time for the events search range (unixtime)."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdLimit, "limit", 0, TRAPI("Maximum number of events to retrieve."))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdTo, "to", 0, TRAPI("End time for the events search range (unixtime)."))

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

func collectSubscribersSessionEventsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersSessionEventsCmd("/subscribers/{imsi}/events/sessions"),
		query:  buildQueryForSubscribersSessionEventsCmd(),
	}, nil
}

func buildPathForSubscribersSessionEventsCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSessionEventsCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSessionEventsCmd() string {
	result := []string{}

	if SubscribersSessionEventsCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", SubscribersSessionEventsCmdLastEvaluatedKey))
	}

	if SubscribersSessionEventsCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", SubscribersSessionEventsCmdFrom))
	}

	if SubscribersSessionEventsCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", SubscribersSessionEventsCmdLimit))
	}

	if SubscribersSessionEventsCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", SubscribersSessionEventsCmdTo))
	}

	return strings.Join(result, "&")
}
