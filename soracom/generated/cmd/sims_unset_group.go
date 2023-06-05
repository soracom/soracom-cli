// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsUnsetGroupCmdSimId holds value of 'sim_id' option
var SimsUnsetGroupCmdSimId string

func InitSimsUnsetGroupCmd() {
	SimsUnsetGroupCmd.Flags().StringVar(&SimsUnsetGroupCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsUnsetGroupCmd.RunE = SimsUnsetGroupCmdRunE

	SimsCmd.AddCommand(SimsUnsetGroupCmd)
}

// SimsUnsetGroupCmd defines 'unset-group' subcommand
var SimsUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/sims/{sim_id}/unset_group:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/unset_group:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "unsetSimGroup"),
}

func SimsUnsetGroupCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimsUnsetGroupCmdParams(ac)
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

func collectSimsUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsUnsetGroupCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsUnsetGroupCmd("/sims/{sim_id}/unset_group"),
		query:  buildQueryForSimsUnsetGroupCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsUnsetGroupCmd(path string) string {

	escapedSimId := url.PathEscape(SimsUnsetGroupCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsUnsetGroupCmd() url.Values {
	result := url.Values{}

	return result
}
