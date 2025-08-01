// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var FilesListCmdLastEvaluatedKey string

// FilesListCmdPath holds value of 'path' option
var FilesListCmdPath string

// FilesListCmdScope holds value of 'scope' option
var FilesListCmdScope string

// FilesListCmdLimit holds value of 'limit' option
var FilesListCmdLimit int64

// FilesListCmdPaginate indicates to do pagination or not
var FilesListCmdPaginate bool

// FilesListCmdOutputJSONL indicates to output with jsonl format
var FilesListCmdOutputJSONL bool

func InitFilesListCmd() {
	FilesListCmd.Flags().StringVar(&FilesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The filename of the last file entry retrieved on the previous page. By specifying this parameter, you can continue to retrieve the list from the next file entry onward."))

	FilesListCmd.Flags().StringVar(&FilesListCmdPath, "path", "/", TRAPI("Target path."))

	FilesListCmd.Flags().StringVar(&FilesListCmdScope, "scope", "private", TRAPI("Scope of the request. Specify 'private' to handle files uploaded to Harvest Files."))

	FilesListCmd.Flags().Int64Var(&FilesListCmdLimit, "limit", 10, TRAPI("Maximum number of file entries to be returned."))

	FilesListCmd.Flags().BoolVar(&FilesListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	FilesListCmd.Flags().BoolVar(&FilesListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	FilesListCmd.RunE = FilesListCmdRunE

	FilesCmd.AddCommand(FilesListCmd)
}

// FilesListCmd defines 'list' subcommand
var FilesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/files/{scope}/{path}/:get:summary"),
	Long:  TRAPI(`/files/{scope}/{path}/:get:description`) + "\n\n" + createLinkToAPIReference("FileEntry", "listFiles"),
}

func FilesListCmdRunE(cmd *cobra.Command, args []string) error {

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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		if FilesListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
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

	if FilesListCmdLimit != 10 {
		result.Add("limit", sprintf("%d", FilesListCmdLimit))
	}

	return result
}
