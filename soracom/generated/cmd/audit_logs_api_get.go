// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// AuditLogsApiGetCmdApiKind holds value of 'api_kind' option
var AuditLogsApiGetCmdApiKind string

// AuditLogsApiGetCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var AuditLogsApiGetCmdLastEvaluatedKey string

// AuditLogsApiGetCmdFromEpochMs holds value of 'from_epoch_ms' option
var AuditLogsApiGetCmdFromEpochMs int64

// AuditLogsApiGetCmdLimit holds value of 'limit' option
var AuditLogsApiGetCmdLimit int64

// AuditLogsApiGetCmdToEpochMs holds value of 'to_epoch_ms' option
var AuditLogsApiGetCmdToEpochMs int64

// AuditLogsApiGetCmdPaginate indicates to do pagination or not
var AuditLogsApiGetCmdPaginate bool

// AuditLogsApiGetCmdOutputJSONL indicates to output with jsonl format
var AuditLogsApiGetCmdOutputJSONL bool

func init() {
	AuditLogsApiGetCmd.Flags().StringVar(&AuditLogsApiGetCmdApiKind, "api-kind", "", TRAPI("Filter item for audit log retrieval by API kind (e.g. `/v1/auth`)."))

	AuditLogsApiGetCmd.Flags().StringVar(&AuditLogsApiGetCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `requestedTimeEpochMs` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	AuditLogsApiGetCmd.Flags().Int64Var(&AuditLogsApiGetCmdFromEpochMs, "from-epoch-ms", 0, TRAPI("Start time for the log search range (unixtime milliseconds)."))

	AuditLogsApiGetCmd.Flags().Int64Var(&AuditLogsApiGetCmdLimit, "limit", 0, TRAPI("Maximum number of log entries to retrieve."))

	AuditLogsApiGetCmd.Flags().Int64Var(&AuditLogsApiGetCmdToEpochMs, "to-epoch-ms", 0, TRAPI("End time for the log search range (unixtime milliseconds)."))

	AuditLogsApiGetCmd.Flags().BoolVar(&AuditLogsApiGetCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	AuditLogsApiGetCmd.Flags().BoolVar(&AuditLogsApiGetCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	AuditLogsApiCmd.AddCommand(AuditLogsApiGetCmd)
}

// AuditLogsApiGetCmd defines 'get' subcommand
var AuditLogsApiGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/audit_logs/api:get:summary"),
	Long:  TRAPI(`/audit_logs/api:get:description`) + "\n\n" + createLinkToAPIReference("AuditLog", "getApiAuditLogs"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectAuditLogsApiGetCmdParams(ac)
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
			if AuditLogsApiGetCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectAuditLogsApiGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForAuditLogsApiGetCmd("/audit_logs/api"),
		query:  buildQueryForAuditLogsApiGetCmd(),

		doPagination:                      AuditLogsApiGetCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAuditLogsApiGetCmd(path string) string {

	return path
}

func buildQueryForAuditLogsApiGetCmd() url.Values {
	result := url.Values{}

	if AuditLogsApiGetCmdApiKind != "" {
		result.Add("api_kind", AuditLogsApiGetCmdApiKind)
	}

	if AuditLogsApiGetCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", AuditLogsApiGetCmdLastEvaluatedKey)
	}

	if AuditLogsApiGetCmdFromEpochMs != 0 {
		result.Add("from_epoch_ms", sprintf("%d", AuditLogsApiGetCmdFromEpochMs))
	}

	if AuditLogsApiGetCmdLimit != 0 {
		result.Add("limit", sprintf("%d", AuditLogsApiGetCmdLimit))
	}

	if AuditLogsApiGetCmdToEpochMs != 0 {
		result.Add("to_epoch_ms", sprintf("%d", AuditLogsApiGetCmdToEpochMs))
	}

	return result
}
