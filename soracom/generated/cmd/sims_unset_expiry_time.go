// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsUnsetExpiryTimeCmdSimId holds value of 'sim_id' option
var SimsUnsetExpiryTimeCmdSimId string

func init() {
	SimsUnsetExpiryTimeCmd.Flags().StringVar(&SimsUnsetExpiryTimeCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsUnsetExpiryTimeCmd)
}

// SimsUnsetExpiryTimeCmd defines 'unset-expiry-time' subcommand
var SimsUnsetExpiryTimeCmd = &cobra.Command{
	Use:   "unset-expiry-time",
	Short: TRAPI("/sims/{sim_id}/unset_expiry_time:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/unset_expiry_time:post:description`),
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

		param, err := collectSimsUnsetExpiryTimeCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSimsUnsetExpiryTimeCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsUnsetExpiryTimeCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsUnsetExpiryTimeCmd("/sims/{sim_id}/unset_expiry_time"),
		query:  buildQueryForSimsUnsetExpiryTimeCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsUnsetExpiryTimeCmd(path string) string {

	escapedSimId := url.PathEscape(SimsUnsetExpiryTimeCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsUnsetExpiryTimeCmd() url.Values {
	result := url.Values{}

	return result
}
