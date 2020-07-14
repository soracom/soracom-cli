// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesGetCmdPath holds value of 'path' option
var FilesGetCmdPath string

// FilesGetCmdScope holds value of 'scope' option
var FilesGetCmdScope string

func init() {
	FilesGetCmd.Flags().StringVar(&FilesGetCmdPath, "path", "", TRAPI("Target path"))

	FilesGetCmd.Flags().StringVar(&FilesGetCmdScope, "scope", "private", TRAPI("Scope of the request"))
	FilesCmd.AddCommand(FilesGetCmd)
}

// FilesGetCmd defines 'get' subcommand
var FilesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/files/{scope}/{path}:get:summary"),
	Long:  TRAPI(`/files/{scope}/{path}:get:description`),
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

		param, err := collectFilesGetCmdParams(ac)
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
		rawOutput = true

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectFilesGetCmdParams(ac *apiClient) (*apiParams, error) {
	if FilesGetCmdPath == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "path")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForFilesGetCmd("/files/{scope}/{path}"),
		query:  buildQueryForFilesGetCmd(),
	}, nil
}

func buildPathForFilesGetCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesGetCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesGetCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesGetCmd() url.Values {
	result := url.Values{}

	return result
}
