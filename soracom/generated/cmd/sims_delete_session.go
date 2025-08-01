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

func InitSimsDeleteSessionCmd() {
	SimsDeleteSessionCmd.Flags().StringVar(&SimsDeleteSessionCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM. The SIM ID can be obtained from the [Sim:listSims API](#!/Sim/listSims)."))

	SimsDeleteSessionCmd.RunE = SimsDeleteSessionCmdRunE

	SimsCmd.AddCommand(SimsDeleteSessionCmd)
}

// SimsDeleteSessionCmd defines 'delete-session' subcommand
var SimsDeleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: TRAPI("/sims/{sim_id}/delete_session:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/delete_session:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "deleteSimSession"),
}

func SimsDeleteSessionCmdRunE(cmd *cobra.Command, args []string) error {

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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSimsDeleteSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsDeleteSessionCmdSimId)
	if err != nil {
		return nil, err
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
