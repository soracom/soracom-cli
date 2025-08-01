// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// AuditLogsNapterGetCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var AuditLogsNapterGetCmdLastEvaluatedKey string

// AuditLogsNapterGetCmdResourceId holds value of 'resource_id' option
var AuditLogsNapterGetCmdResourceId string

// AuditLogsNapterGetCmdResourceType holds value of 'resource_type' option
var AuditLogsNapterGetCmdResourceType string

// AuditLogsNapterGetCmdFrom holds value of 'from' option
var AuditLogsNapterGetCmdFrom int64

// AuditLogsNapterGetCmdLimit holds value of 'limit' option
var AuditLogsNapterGetCmdLimit int64

// AuditLogsNapterGetCmdTo holds value of 'to' option
var AuditLogsNapterGetCmdTo int64

// AuditLogsNapterGetCmdPaginate indicates to do pagination or not
var AuditLogsNapterGetCmdPaginate bool

// AuditLogsNapterGetCmdOutputJSONL indicates to output with jsonl format
var AuditLogsNapterGetCmdOutputJSONL bool

func InitAuditLogsNapterGetCmd() {
	AuditLogsNapterGetCmd.Flags().StringVar(&AuditLogsNapterGetCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of 'time' in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	AuditLogsNapterGetCmd.Flags().StringVar(&AuditLogsNapterGetCmdResourceId, "resource-id", "", TRAPI("Identity of the target resource to query log entries."))

	AuditLogsNapterGetCmd.Flags().StringVar(&AuditLogsNapterGetCmdResourceType, "resource-type", "", TRAPI("Type of the target resource to query log entries."))

	AuditLogsNapterGetCmd.Flags().Int64Var(&AuditLogsNapterGetCmdFrom, "from", 0, TRAPI("Start time for the log search range (unixtime milliseconds)."))

	AuditLogsNapterGetCmd.Flags().Int64Var(&AuditLogsNapterGetCmdLimit, "limit", 0, TRAPI("Maximum number of log entries to retrieve (value range is 1 to 1000)."))

	AuditLogsNapterGetCmd.Flags().Int64Var(&AuditLogsNapterGetCmdTo, "to", 0, TRAPI("End time for the log search range (unixtime milliseconds)."))

	AuditLogsNapterGetCmd.Flags().BoolVar(&AuditLogsNapterGetCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	AuditLogsNapterGetCmd.Flags().BoolVar(&AuditLogsNapterGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	AuditLogsNapterGetCmd.RunE = AuditLogsNapterGetCmdRunE

	AuditLogsNapterCmd.AddCommand(AuditLogsNapterGetCmd)
}

// AuditLogsNapterGetCmd defines 'get' subcommand
var AuditLogsNapterGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/audit_logs/napter:get:summary"),
	Long:  TRAPI(`/audit_logs/napter:get:description`) + "\n\n" + createLinkToAPIReference("AuditLog", "getNapterAuditLogs"),
}

func AuditLogsNapterGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectAuditLogsNapterGetCmdParams(ac)
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
		if AuditLogsNapterGetCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectAuditLogsNapterGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForAuditLogsNapterGetCmd("/audit_logs/napter"),
		query:  buildQueryForAuditLogsNapterGetCmd(),

		doPagination:                      AuditLogsNapterGetCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAuditLogsNapterGetCmd(path string) string {

	return path
}

func buildQueryForAuditLogsNapterGetCmd() url.Values {
	result := url.Values{}

	if AuditLogsNapterGetCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", AuditLogsNapterGetCmdLastEvaluatedKey)
	}

	if AuditLogsNapterGetCmdResourceId != "" {
		result.Add("resource_id", AuditLogsNapterGetCmdResourceId)
	}

	if AuditLogsNapterGetCmdResourceType != "" {
		result.Add("resource_type", AuditLogsNapterGetCmdResourceType)
	}

	if AuditLogsNapterGetCmdFrom != 0 {
		result.Add("from", sprintf("%d", AuditLogsNapterGetCmdFrom))
	}

	if AuditLogsNapterGetCmdLimit != 0 {
		result.Add("limit", sprintf("%d", AuditLogsNapterGetCmdLimit))
	}

	if AuditLogsNapterGetCmdTo != 0 {
		result.Add("to", sprintf("%d", AuditLogsNapterGetCmdTo))
	}

	return result
}
