// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var FilesListCmdLastEvaluatedKey string

// FilesListCmdLimit holds value of 'limit' option
var FilesListCmdLimit string

// FilesListCmdPath holds value of 'path' option
var FilesListCmdPath string

// FilesListCmdScope holds value of 'scope' option
var FilesListCmdScope string

// FilesListCmdPaginate indicates to do pagination or not
var FilesListCmdPaginate bool

func init() {
	FilesListCmd.Flags().StringVar(&FilesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The filename  of the last file entry retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next file entry onward."))

	FilesListCmd.Flags().StringVar(&FilesListCmdLimit, "limit", "", TRAPI("Num of entries"))

	FilesListCmd.Flags().StringVar(&FilesListCmdPath, "path", "/", TRAPI("Target path"))

	FilesListCmd.Flags().StringVar(&FilesListCmdScope, "scope", "private", TRAPI("Scope of the request"))

	FilesListCmd.Flags().BoolVar(&FilesListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	FilesCmd.AddCommand(FilesListCmd)
}

// FilesListCmd defines 'list' subcommand
var FilesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/files/{scope}/{path}/:get:summary"),
	Long:  TRAPI(`/files/{scope}/{path}/:get:description`),
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

		param, err := collectFilesListCmdParams(ac)
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

func collectFilesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForFilesListCmd("/files/{scope}/{path}/"),
		query:  buildQueryForFilesListCmd(),

		doPagination:                      FilesListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForFilesListCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesListCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesListCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesListCmd() url.Values {
	result := url.Values{}

	if FilesListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", FilesListCmdLastEvaluatedKey)
	}

	if FilesListCmdLimit != "" {
		result.Add("limit", FilesListCmdLimit)
	}

	return result
}
