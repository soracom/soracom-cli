// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsSuspendCmdSimId holds value of 'sim_id' option
var SimsSuspendCmdSimId string

func InitSimsSuspendCmd() {
	SimsSuspendCmd.Flags().StringVar(&SimsSuspendCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM. The SIM ID can be obtained from the [Sim:listSims API](#!/Sim/listSims)."))

	SimsSuspendCmd.RunE = SimsSuspendCmdRunE

	SimsCmd.AddCommand(SimsSuspendCmd)
}

// SimsSuspendCmd defines 'suspend' subcommand
var SimsSuspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: TRAPI("/sims/{sim_id}/suspend:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/suspend:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "suspendSim"),
}

func SimsSuspendCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimsSuspendCmdParams(ac)
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

func collectSimsSuspendCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsSuspendCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsSuspendCmd("/sims/{sim_id}/suspend"),
		query:  buildQueryForSimsSuspendCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsSuspendCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSuspendCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSuspendCmd() url.Values {
	result := url.Values{}

	return result
}
