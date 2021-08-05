// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesGetMetadataCmdPath holds value of 'path' option
var FilesGetMetadataCmdPath string

// FilesGetMetadataCmdScope holds value of 'scope' option
var FilesGetMetadataCmdScope string

func init() {
	FilesGetMetadataCmd.Flags().StringVar(&FilesGetMetadataCmdPath, "path", "", TRAPI("Target path"))

	FilesGetMetadataCmd.Flags().StringVar(&FilesGetMetadataCmdScope, "scope", "private", TRAPI("Scope of the request"))
	FilesCmd.AddCommand(FilesGetMetadataCmd)
}

// FilesGetMetadataCmd defines 'get-metadata' subcommand
var FilesGetMetadataCmd = &cobra.Command{
	Use:   "get-metadata",
	Short: TRAPI("/files/{scope}/{path}:head:summary"),
	Long:  TRAPI(`/files/{scope}/{path}:head:description`),
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

		param, err := collectFilesGetMetadataCmdParams(ac)
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

func collectFilesGetMetadataCmdParams(ac *apiClient) (*apiParams, error) {
	if FilesGetMetadataCmdPath == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "path")
	}

	return &apiParams{
		method: "HEAD",
		path:   buildPathForFilesGetMetadataCmd("/files/{scope}/{path}"),
		query:  buildQueryForFilesGetMetadataCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForFilesGetMetadataCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesGetMetadataCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesGetMetadataCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesGetMetadataCmd() url.Values {
	result := url.Values{}

	return result
}
