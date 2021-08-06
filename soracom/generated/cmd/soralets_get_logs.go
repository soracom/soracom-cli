// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraletsGetLogsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraletsGetLogsCmdLastEvaluatedKey string

// SoraletsGetLogsCmdSoraletId holds value of 'soralet_id' option
var SoraletsGetLogsCmdSoraletId string

// SoraletsGetLogsCmdSort holds value of 'sort' option
var SoraletsGetLogsCmdSort string

// SoraletsGetLogsCmdLimit holds value of 'limit' option
var SoraletsGetLogsCmdLimit int64

// SoraletsGetLogsCmdPaginate indicates to do pagination or not
var SoraletsGetLogsCmdPaginate bool

func init() {
	SoraletsGetLogsCmd.Flags().StringVar(&SoraletsGetLogsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The identifier of the last log message retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next log message onward."))

	SoraletsGetLogsCmd.Flags().StringVar(&SoraletsGetLogsCmdSoraletId, "soralet-id", "", TRAPI("The identifier of Soralet."))

	SoraletsGetLogsCmd.Flags().StringVar(&SoraletsGetLogsCmdSort, "sort", "desc", TRAPI("Sort order"))

	SoraletsGetLogsCmd.Flags().Int64Var(&SoraletsGetLogsCmdLimit, "limit", 0, TRAPI("The maximum number of items in a response."))

	SoraletsGetLogsCmd.Flags().BoolVar(&SoraletsGetLogsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	SoraletsCmd.AddCommand(SoraletsGetLogsCmd)
}

// SoraletsGetLogsCmd defines 'get-logs' subcommand
var SoraletsGetLogsCmd = &cobra.Command{
	Use:   "get-logs",
	Short: TRAPI("/soralets/{soralet_id}/logs:get:summary"),
	Long:  TRAPI(`/soralets/{soralet_id}/logs:get:description`),
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

		param, err := collectSoraletsGetLogsCmdParams(ac)
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

func collectSoraletsGetLogsCmdParams(ac *apiClient) (*apiParams, error) {

	if SoraletsGetLogsCmdSoraletId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "soralet-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraletsGetLogsCmd("/soralets/{soralet_id}/logs"),
		query:  buildQueryForSoraletsGetLogsCmd(),

		doPagination:                      SoraletsGetLogsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraletsGetLogsCmd(path string) string {

	escapedSoraletId := url.PathEscape(SoraletsGetLogsCmdSoraletId)

	path = strReplace(path, "{"+"soralet_id"+"}", escapedSoraletId, -1)

	return path
}

func buildQueryForSoraletsGetLogsCmd() url.Values {
	result := url.Values{}

	if SoraletsGetLogsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraletsGetLogsCmdLastEvaluatedKey)
	}

	if SoraletsGetLogsCmdSort != "desc" {
		result.Add("sort", SoraletsGetLogsCmdSort)
	}

	if SoraletsGetLogsCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SoraletsGetLogsCmdLimit))
	}

	return result
}
