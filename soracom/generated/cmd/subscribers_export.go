// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersExportCmdExportMode holds value of 'export_mode' option
var SubscribersExportCmdExportMode string

func InitSubscribersExportCmd() {
	SubscribersExportCmd.Flags().StringVar(&SubscribersExportCmdExportMode, "export-mode", "sync", TRAPI("export mode (async, sync)"))

	SubscribersExportCmd.RunE = SubscribersExportCmdRunE

	SubscribersCmd.AddCommand(SubscribersExportCmd)
}

// SubscribersExportCmd defines 'export' subcommand
var SubscribersExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/subscribers/export:post:summary"),
	Long:  TRAPI(`/subscribers/export:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "exportSubscribers"),
}

func SubscribersExportCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersExportCmdParams(ac)
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
}

func collectSubscribersExportCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersExportCmd("/subscribers/export"),
		query:  buildQueryForSubscribersExportCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersExportCmd(path string) string {

	return path
}

func buildQueryForSubscribersExportCmd() url.Values {
	result := url.Values{}

	if SubscribersExportCmdExportMode != "sync" {
		result.Add("export_mode", SubscribersExportCmdExportMode)
	}

	return result
}
