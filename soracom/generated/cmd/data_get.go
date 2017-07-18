package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DataGetCmdImsi holds value of 'imsi' option
var DataGetCmdImsi string

// DataGetCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DataGetCmdLastEvaluatedKey string

// DataGetCmdSort holds value of 'sort' option
var DataGetCmdSort string

// DataGetCmdFrom holds value of 'from' option
var DataGetCmdFrom int64

// DataGetCmdLimit holds value of 'limit' option
var DataGetCmdLimit int64

// DataGetCmdTo holds value of 'to' option
var DataGetCmdTo int64

func init() {
	DataGetCmd.Flags().StringVar(&DataGetCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber that generated data entries."))

	DataGetCmd.Flags().StringVar(&DataGetCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	DataGetCmd.Flags().StringVar(&DataGetCmdSort, "sort", "", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	DataGetCmd.Flags().Int64Var(&DataGetCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	DataCmd.AddCommand(DataGetCmd)
}

// DataGetCmd defines 'get' subcommand
var DataGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/subscribers/{imsi}/data:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/data:get:description`),
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

		param, err := collectDataGetCmdParams()
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

func collectDataGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDataGetCmd("/subscribers/{imsi}/data"),
		query:  buildQueryForDataGetCmd(),
	}, nil
}

func buildPathForDataGetCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", DataGetCmdImsi, -1)

	return path
}

func buildQueryForDataGetCmd() string {
	result := []string{}

	if DataGetCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", DataGetCmdLastEvaluatedKey))
	}

	if DataGetCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", DataGetCmdSort))
	}

	if DataGetCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", DataGetCmdFrom))
	}

	if DataGetCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", DataGetCmdLimit))
	}

	if DataGetCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", DataGetCmdTo))
	}

	return strings.Join(result, "&")
}
