// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SimsGetDataCmdLastEvaluatedKey string

// SimsGetDataCmdSimId holds value of 'sim_id' option
var SimsGetDataCmdSimId string

// SimsGetDataCmdSort holds value of 'sort' option
var SimsGetDataCmdSort string

// SimsGetDataCmdFrom holds value of 'from' option
var SimsGetDataCmdFrom int64

// SimsGetDataCmdLimit holds value of 'limit' option
var SimsGetDataCmdLimit int64

// SimsGetDataCmdTo holds value of 'to' option
var SimsGetDataCmdTo int64

func init() {
	SimsGetDataCmd.Flags().StringVar(&SimsGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	SimsGetDataCmd.Flags().StringVar(&SimsGetDataCmdSimId, "sim-id", "", TRAPI("Sim Id of the target SIM."))

	SimsGetDataCmd.Flags().StringVar(&SimsGetDataCmdSort, "sort", "desc", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	SimsGetDataCmd.Flags().Int64Var(&SimsGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	SimsGetDataCmd.Flags().Int64Var(&SimsGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	SimsGetDataCmd.Flags().Int64Var(&SimsGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))
	SimsCmd.AddCommand(SimsGetDataCmd)
}

// SimsGetDataCmd defines 'get-data' subcommand
var SimsGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/sims/{sim_id}/data:get:summary"),
	Long:  TRAPI(`/sims/{sim_id}/data:get:description`),
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

		param, err := collectSimsGetDataCmdParams(ac)
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
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSimsGetDataCmdParams(ac *apiClient) (*apiParams, error) {

	if SimsGetDataCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsGetDataCmd("/sims/{sim_id}/data"),
		query:  buildQueryForSimsGetDataCmd(),
	}, nil
}

func buildPathForSimsGetDataCmd(path string) string {

	escapedSimId := url.PathEscape(SimsGetDataCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsGetDataCmd() url.Values {
	result := url.Values{}

	if SimsGetDataCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SimsGetDataCmdLastEvaluatedKey)
	}

	if SimsGetDataCmdSort != "desc" {
		result.Add("sort", SimsGetDataCmdSort)
	}

	if SimsGetDataCmdFrom != 0 {
		result.Add("from", sprintf("%d", SimsGetDataCmdFrom))
	}

	if SimsGetDataCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SimsGetDataCmdLimit))
	}

	if SimsGetDataCmdTo != 0 {
		result.Add("to", sprintf("%d", SimsGetDataCmdTo))
	}

	return result
}
