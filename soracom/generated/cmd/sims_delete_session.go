// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsDeleteSessionCmdSimId holds value of 'sim_id' option
var SimsDeleteSessionCmdSimId string

func init() {
	SimsDeleteSessionCmd.Flags().StringVar(&SimsDeleteSessionCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsDeleteSessionCmd)
}

// SimsDeleteSessionCmd defines 'delete-session' subcommand
var SimsDeleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: TRAPI("/sims/{sim_id}/delete_session:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/delete_session:post:description`),
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

		param, err := collectSimsDeleteSessionCmdParams(ac)
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

func collectSimsDeleteSessionCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsDeleteSessionCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsDeleteSessionCmd("/sims/{sim_id}/delete_session"),
		query:  buildQueryForSimsDeleteSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsDeleteSessionCmd(path string) string {

	escapedSimId := url.PathEscape(SimsDeleteSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsDeleteSessionCmd() url.Values {
	result := url.Values{}

	return result
}
