// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BatchGroupsJobsGetCmdBatchGroupId holds value of 'batch_group_id' option
var BatchGroupsJobsGetCmdBatchGroupId string

// BatchGroupsJobsGetCmdJobId holds value of 'job_id' option
var BatchGroupsJobsGetCmdJobId string

func InitBatchGroupsJobsGetCmd() {
	BatchGroupsJobsGetCmd.Flags().StringVar(&BatchGroupsJobsGetCmdBatchGroupId, "batch-group-id", "", TRAPI("Batch group ID."))

	BatchGroupsJobsGetCmd.Flags().StringVar(&BatchGroupsJobsGetCmdJobId, "job-id", "", TRAPI("Batch job ID."))

	BatchGroupsJobsGetCmd.RunE = BatchGroupsJobsGetCmdRunE

	BatchGroupsJobsCmd.AddCommand(BatchGroupsJobsGetCmd)
}

// BatchGroupsJobsGetCmd defines 'get' subcommand
var BatchGroupsJobsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/batch_groups/{batch_group_id}/jobs/{job_id}:get:summary"),
	Long:  TRAPI(`/batch_groups/{batch_group_id}/jobs/{job_id}:get:description`) + "\n\n" + createLinkToAPIReference("Batch", "getBatchJob"),
}

func BatchGroupsJobsGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectBatchGroupsJobsGetCmdParams(ac)
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
}

func collectBatchGroupsJobsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("batch_group_id", "batch-group-id", "path", parsedBody, BatchGroupsJobsGetCmdBatchGroupId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("job_id", "job-id", "path", parsedBody, BatchGroupsJobsGetCmdJobId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForBatchGroupsJobsGetCmd("/batch_groups/{batch_group_id}/jobs/{job_id}"),
		query:  buildQueryForBatchGroupsJobsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForBatchGroupsJobsGetCmd(path string) string {

	escapedBatchGroupId := url.PathEscape(BatchGroupsJobsGetCmdBatchGroupId)

	path = strReplace(path, "{"+"batch_group_id"+"}", escapedBatchGroupId, -1)

	escapedJobId := url.PathEscape(BatchGroupsJobsGetCmdJobId)

	path = strReplace(path, "{"+"job_id"+"}", escapedJobId, -1)

	return path
}

func buildQueryForBatchGroupsJobsGetCmd() url.Values {
	result := url.Values{}

	return result
}
