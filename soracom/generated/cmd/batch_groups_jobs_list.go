// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BatchGroupsJobsListCmdBatchGroupId holds value of 'batch_group_id' option
var BatchGroupsJobsListCmdBatchGroupId string

// BatchGroupsJobsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var BatchGroupsJobsListCmdLastEvaluatedKey string

// BatchGroupsJobsListCmdLimit holds value of 'limit' option
var BatchGroupsJobsListCmdLimit int64

// BatchGroupsJobsListCmdPaginate indicates to do pagination or not
var BatchGroupsJobsListCmdPaginate bool

// BatchGroupsJobsListCmdOutputJSONL indicates to output with jsonl format
var BatchGroupsJobsListCmdOutputJSONL bool

func InitBatchGroupsJobsListCmd() {
	BatchGroupsJobsListCmd.Flags().StringVar(&BatchGroupsJobsListCmdBatchGroupId, "batch-group-id", "", TRAPI("Batch group ID."))

	BatchGroupsJobsListCmd.Flags().StringVar(&BatchGroupsJobsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of the 'x-soracom-next-key' header returned in the response to the last request. Specify this to retrieve the next page of batch jobs."))

	BatchGroupsJobsListCmd.Flags().Int64Var(&BatchGroupsJobsListCmdLimit, "limit", 0, TRAPI("Maximum number of batch jobs to retrieve. The number of batch jobs returned may be less than the specified value."))

	BatchGroupsJobsListCmd.Flags().BoolVar(&BatchGroupsJobsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	BatchGroupsJobsListCmd.Flags().BoolVar(&BatchGroupsJobsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	BatchGroupsJobsListCmd.RunE = BatchGroupsJobsListCmdRunE

	BatchGroupsJobsCmd.AddCommand(BatchGroupsJobsListCmd)
}

// BatchGroupsJobsListCmd defines 'list' subcommand
var BatchGroupsJobsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/batch_groups/{batch_group_id}/jobs:get:summary"),
	Long:  TRAPI(`/batch_groups/{batch_group_id}/jobs:get:description`) + "\n\n" + createLinkToAPIReference("Batch", "listBatchJobs"),
}

func BatchGroupsJobsListCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectBatchGroupsJobsListCmdParams(ac)
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
		if BatchGroupsJobsListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectBatchGroupsJobsListCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("batch_group_id", "batch-group-id", "path", parsedBody, BatchGroupsJobsListCmdBatchGroupId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForBatchGroupsJobsListCmd("/batch_groups/{batch_group_id}/jobs"),
		query:  buildQueryForBatchGroupsJobsListCmd(),

		doPagination:                      BatchGroupsJobsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBatchGroupsJobsListCmd(path string) string {

	escapedBatchGroupId := url.PathEscape(BatchGroupsJobsListCmdBatchGroupId)

	path = strReplace(path, "{"+"batch_group_id"+"}", escapedBatchGroupId, -1)

	return path
}

func buildQueryForBatchGroupsJobsListCmd() url.Values {
	result := url.Values{}

	if BatchGroupsJobsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", BatchGroupsJobsListCmdLastEvaluatedKey)
	}

	if BatchGroupsJobsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", BatchGroupsJobsListCmdLimit))
	}

	return result
}
