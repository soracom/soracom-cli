package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var SubscribersSessionEventsCmdImsi string

var SubscribersSessionEventsCmdLastEvaluatedKey string

var SubscribersSessionEventsCmdFrom int64

var SubscribersSessionEventsCmdLimit int64

var SubscribersSessionEventsCmdTo int64

func init() {
	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdImsi, "imsi", "", TR("subscribers.list_session_events.get.parameters.imsi.description"))

	SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", TR("subscribers.list_session_events.get.parameters.last_evaluated_key.description"))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdFrom, "from", 0, TR("subscribers.list_session_events.get.parameters.from.description"))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdLimit, "limit", 0, TR("subscribers.list_session_events.get.parameters.limit.description"))

	SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdTo, "to", 0, TR("subscribers.list_session_events.get.parameters.to.description"))

	SubscribersCmd.AddCommand(SubscribersSessionEventsCmd)
}

var SubscribersSessionEventsCmd = &cobra.Command{
	Use:   "session-events",
	Short: TR("subscribers.list_session_events.get.summary"),
	Long:  TR(`subscribers.list_session_events.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectSubscribersSessionEventsCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectSubscribersSessionEventsCmdParams() (*apiParams, error) {

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
