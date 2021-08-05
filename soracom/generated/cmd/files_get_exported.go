// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesGetExportedCmdExportedFileId holds value of 'exported_file_id' option
var FilesGetExportedCmdExportedFileId string

func init() {
	FilesGetExportedCmd.Flags().StringVar(&FilesGetExportedCmdExportedFileId, "exported-file-id", "", TRAPI("file export id"))
	FilesCmd.AddCommand(FilesGetExportedCmd)
}

// FilesGetExportedCmd defines 'get-exported' subcommand
var FilesGetExportedCmd = &cobra.Command{
	Use:   "get-exported",
	Short: TRAPI("/files/exported/{exported_file_id}:get:summary"),
	Long:  TRAPI(`/files/exported/{exported_file_id}:get:description`),
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

		param, err := collectFilesGetExportedCmdParams(ac)
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

func collectFilesGetExportedCmdParams(ac *apiClient) (*apiParams, error) {
	if FilesGetExportedCmdExportedFileId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "exported-file-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForFilesGetExportedCmd("/files/exported/{exported_file_id}"),
		query:  buildQueryForFilesGetExportedCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForFilesGetExportedCmd(path string) string {

	escapedExportedFileId := url.PathEscape(FilesGetExportedCmdExportedFileId)

	path = strReplace(path, "{"+"exported_file_id"+"}", escapedExportedFileId, -1)

	return path
}

func buildQueryForFilesGetExportedCmd() url.Values {
	result := url.Values{}

	return result
}
