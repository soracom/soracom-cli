// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// ResourceSummariesGetCmdResourceSummaryType holds value of 'resource_summary_type' option
var ResourceSummariesGetCmdResourceSummaryType string

func InitResourceSummariesGetCmd() {
	ResourceSummariesGetCmd.Flags().StringVar(&ResourceSummariesGetCmdResourceSummaryType, "resource-summary-type", "", TRAPI("The type of the resource summary.- 'simsPerStatus': The number of IoT SIMs per status"))

	ResourceSummariesGetCmd.RunE = ResourceSummariesGetCmdRunE

	ResourceSummariesCmd.AddCommand(ResourceSummariesGetCmd)
}

// ResourceSummariesGetCmd defines 'get' subcommand
var ResourceSummariesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/resource_summaries/{resource_summary_type}:get:summary"),
	Long:  TRAPI(`/resource_summaries/{resource_summary_type}:get:description`) + "\n\n" + createLinkToAPIReference("ResourceSummary", "getResourceSummary"),
}

func ResourceSummariesGetCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectResourceSummariesGetCmdParams(ac)
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
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectResourceSummariesGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("resource_summary_type", "resource-summary-type", "path", parsedBody, ResourceSummariesGetCmdResourceSummaryType)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForResourceSummariesGetCmd("/resource_summaries/{resource_summary_type}"),
		query:  buildQueryForResourceSummariesGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForResourceSummariesGetCmd(path string) string {

	escapedResourceSummaryType := url.PathEscape(ResourceSummariesGetCmdResourceSummaryType)

	path = strReplace(path, "{"+"resource_summary_type"+"}", escapedResourceSummaryType, -1)

	return path
}

func buildQueryForResourceSummariesGetCmd() url.Values {
	result := url.Values{}

	return result
}
