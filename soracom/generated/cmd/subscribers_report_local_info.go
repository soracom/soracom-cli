package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersReportLocalInfoCmdImsi holds value of 'imsi' option
var SubscribersReportLocalInfoCmdImsi string

func init() {
	SubscribersReportLocalInfoCmd.Flags().StringVar(&SubscribersReportLocalInfoCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersReportLocalInfoCmd)
}

// SubscribersReportLocalInfoCmd defines 'report-local-info' subcommand
var SubscribersReportLocalInfoCmd = &cobra.Command{
	Use:   "report-local-info",
	Short: TRAPI("/subscribers/{imsi}/report_local_info:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/report_local_info:post:description`),
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

		param, err := collectSubscribersReportLocalInfoCmdParams()
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

func collectSubscribersReportLocalInfoCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersReportLocalInfoCmd("/subscribers/{imsi}/report_local_info"),
		query:  buildQueryForSubscribersReportLocalInfoCmd(),
	}, nil
}

func buildPathForSubscribersReportLocalInfoCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersReportLocalInfoCmdImsi, -1)

	return path
}

func buildQueryForSubscribersReportLocalInfoCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
