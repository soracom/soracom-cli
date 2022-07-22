// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsReportLocalInfoCmdSimId holds value of 'sim_id' option
var SimsReportLocalInfoCmdSimId string

func init() {
	SimsReportLocalInfoCmd.Flags().StringVar(&SimsReportLocalInfoCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsReportLocalInfoCmd)
}

// SimsReportLocalInfoCmd defines 'report-local-info' subcommand
var SimsReportLocalInfoCmd = &cobra.Command{
	Use:   "report-local-info",
	Short: TRAPI("/sims/{sim_id}/report_local_info:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/report_local_info:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "reportSimLocalInfo"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectSimsReportLocalInfoCmdParams(ac)
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
	},
}

func collectSimsReportLocalInfoCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsReportLocalInfoCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsReportLocalInfoCmd("/sims/{sim_id}/report_local_info"),
		query:  buildQueryForSimsReportLocalInfoCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsReportLocalInfoCmd(path string) string {

	escapedSimId := url.PathEscape(SimsReportLocalInfoCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsReportLocalInfoCmd() url.Values {
	result := url.Values{}

	return result
}
