package cmd

import (
	"os"
	"strings"

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

func init() {
	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdImsi, "imsi", "", TR("subscribers.get_data_from_subscriber.get.parameters.imsi.description"))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TR("subscribers.get_data_from_subscriber.get.parameters.last_evaluated_key.description"))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdSort, "sort", "", TR("subscribers.get_data_from_subscriber.get.parameters.sort.description"))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdFrom, "from", 0, TR("subscribers.get_data_from_subscriber.get.parameters.from.description"))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdLimit, "limit", 0, TR("subscribers.get_data_from_subscriber.get.parameters.limit.description"))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdTo, "to", 0, TR("subscribers.get_data_from_subscriber.get.parameters.to.description"))

	SubscribersCmd.AddCommand(SubscribersGetDataCmd)
}

// SubscribersGetDataCmd defines 'get-data' subcommand
var SubscribersGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TR("subscribers.get_data_from_subscriber.get.summary"),
	Long:  TR(`subscribers.get_data_from_subscriber.get.description`),
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

		param, err := collectSubscribersGetDataCmdParams()
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

func collectSubscribersGetDataCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersGetDataCmd("/subscribers/{imsi}/data"),
		query:  buildQueryForSubscribersGetDataCmd(),
	}, nil
}

func buildPathForSubscribersGetDataCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersGetDataCmdImsi, -1)

	return path
}

func buildQueryForSubscribersGetDataCmd() string {
	result := []string{}

	if SubscribersGetDataCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", SubscribersGetDataCmdLastEvaluatedKey))
	}

	if SubscribersGetDataCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", SubscribersGetDataCmdSort))
	}

	if SubscribersGetDataCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", SubscribersGetDataCmdFrom))
	}

	if SubscribersGetDataCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", SubscribersGetDataCmdLimit))
	}

	if SubscribersGetDataCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", SubscribersGetDataCmdTo))
	}

	return strings.Join(result, "&")
}
