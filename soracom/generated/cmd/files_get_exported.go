package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// FilesGetExportedCmdExportedFileId holds value of 'exported_file_id' option
var FilesGetExportedCmdExportedFileId string

func init() {
	FilesGetExportedCmd.Flags().StringVar(&FilesGetExportedCmdExportedFileId, "exported-file-id", "", TR("file export id"))

	FilesCmd.AddCommand(FilesGetExportedCmd)
}

// FilesGetExportedCmd defines 'get-exported' subcommand
var FilesGetExportedCmd = &cobra.Command{
	Use:   "get-exported",
	Short: TR("files.get_exported.get.summary"),
	Long:  TR(`files.get_exported.get.description`),
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

		param, err := collectFilesGetExportedCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectFilesGetExportedCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForFilesGetExportedCmd("/files/exported/{exported_file_id}"),
		query:  buildQueryForFilesGetExportedCmd(),
	}, nil
}

func buildPathForFilesGetExportedCmd(path string) string {

	path = strings.Replace(path, "{"+"exported_file_id"+"}", FilesGetExportedCmdExportedFileId, -1)

	return path
}

func buildQueryForFilesGetExportedCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
