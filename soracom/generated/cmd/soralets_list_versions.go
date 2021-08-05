// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraletsListVersionsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SoraletsListVersionsCmdLastEvaluatedKey string

// SoraletsListVersionsCmdSoraletId holds value of 'soralet_id' option
var SoraletsListVersionsCmdSoraletId string

// SoraletsListVersionsCmdSort holds value of 'sort' option
var SoraletsListVersionsCmdSort string

// SoraletsListVersionsCmdLimit holds value of 'limit' option
var SoraletsListVersionsCmdLimit int64

// SoraletsListVersionsCmdPaginate indicates to do pagination or not
var SoraletsListVersionsCmdPaginate bool

func init() {
	SoraletsListVersionsCmd.Flags().StringVar(&SoraletsListVersionsCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The identifier of the last version retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next version onward."))

	SoraletsListVersionsCmd.Flags().StringVar(&SoraletsListVersionsCmdSoraletId, "soralet-id", "", TRAPI("The identifier of Soralet."))

	SoraletsListVersionsCmd.Flags().StringVar(&SoraletsListVersionsCmdSort, "sort", "desc", TRAPI("Sort order"))

	SoraletsListVersionsCmd.Flags().Int64Var(&SoraletsListVersionsCmdLimit, "limit", 0, TRAPI("The maximum number of items in a response."))

	SoraletsListVersionsCmd.Flags().BoolVar(&SoraletsListVersionsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	SoraletsCmd.AddCommand(SoraletsListVersionsCmd)
}

// SoraletsListVersionsCmd defines 'list-versions' subcommand
var SoraletsListVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: TRAPI("/soralets/{soralet_id}/versions:get:summary"),
	Long:  TRAPI(`/soralets/{soralet_id}/versions:get:description`),
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

		param, err := collectSoraletsListVersionsCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSoraletsListVersionsCmdParams(ac *apiClient) (*apiParams, error) {

	if SoraletsListVersionsCmdSoraletId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "soralet-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraletsListVersionsCmd("/soralets/{soralet_id}/versions"),
		query:  buildQueryForSoraletsListVersionsCmd(),

		doPagination:                      SoraletsListVersionsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraletsListVersionsCmd(path string) string {

	escapedSoraletId := url.PathEscape(SoraletsListVersionsCmdSoraletId)

	path = strReplace(path, "{"+"soralet_id"+"}", escapedSoraletId, -1)

	return path
}

func buildQueryForSoraletsListVersionsCmd() url.Values {
	result := url.Values{}

	if SoraletsListVersionsCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", SoraletsListVersionsCmdLastEvaluatedKey)
	}

	if SoraletsListVersionsCmdSort != "desc" {
		result.Add("sort", SoraletsListVersionsCmdSort)
	}

	if SoraletsListVersionsCmdLimit != 0 {
		result.Add("limit", sprintf("%d", SoraletsListVersionsCmdLimit))
	}

	return result
}
