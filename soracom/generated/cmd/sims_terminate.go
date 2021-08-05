// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsTerminateCmdSimId holds value of 'sim_id' option
var SimsTerminateCmdSimId string

func init() {
	SimsTerminateCmd.Flags().StringVar(&SimsTerminateCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsTerminateCmd)
}

// SimsTerminateCmd defines 'terminate' subcommand
var SimsTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/sims/{sim_id}/terminate:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/terminate:post:description`),
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

		param, err := collectSimsTerminateCmdParams(ac)
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

func collectSimsTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsTerminateCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsTerminateCmd("/sims/{sim_id}/terminate"),
		query:  buildQueryForSimsTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsTerminateCmd(path string) string {

	escapedSimId := url.PathEscape(SimsTerminateCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsTerminateCmd() url.Values {
	result := url.Values{}

	return result
}
