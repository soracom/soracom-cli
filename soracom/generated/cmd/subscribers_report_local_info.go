// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersReportLocalInfoCmdImsi holds value of 'imsi' option
var SubscribersReportLocalInfoCmdImsi string

func InitSubscribersReportLocalInfoCmd() {
	SubscribersReportLocalInfoCmd.Flags().StringVar(&SubscribersReportLocalInfoCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersReportLocalInfoCmd.RunE = SubscribersReportLocalInfoCmdRunE

	SubscribersCmd.AddCommand(SubscribersReportLocalInfoCmd)
}

// SubscribersReportLocalInfoCmd defines 'report-local-info' subcommand
var SubscribersReportLocalInfoCmd = &cobra.Command{
	Use:   "report-local-info",
	Short: TRAPI("/subscribers/{imsi}/report_local_info:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/report_local_info:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "reportLocalInfo"),
}

func SubscribersReportLocalInfoCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersReportLocalInfoCmdParams(ac)
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

func collectSubscribersReportLocalInfoCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersReportLocalInfoCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersReportLocalInfoCmd("/subscribers/{imsi}/report_local_info"),
		query:  buildQueryForSubscribersReportLocalInfoCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersReportLocalInfoCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersReportLocalInfoCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersReportLocalInfoCmd() url.Values {
	result := url.Values{}

	return result
}
