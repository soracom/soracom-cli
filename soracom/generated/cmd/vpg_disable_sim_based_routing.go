// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDisableSimBasedRoutingCmdVpgId holds value of 'vpg_id' option
var VpgDisableSimBasedRoutingCmdVpgId string

func InitVpgDisableSimBasedRoutingCmd() {
	VpgDisableSimBasedRoutingCmd.Flags().StringVar(&VpgDisableSimBasedRoutingCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgDisableSimBasedRoutingCmd.RunE = VpgDisableSimBasedRoutingCmdRunE

	VpgCmd.AddCommand(VpgDisableSimBasedRoutingCmd)
}

// VpgDisableSimBasedRoutingCmd defines 'disable-sim-based-routing' subcommand
var VpgDisableSimBasedRoutingCmd = &cobra.Command{
	Use:   "disable-sim-based-routing",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/disable:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/disable:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "disableSimBasedRouging"),
}

func VpgDisableSimBasedRoutingCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgDisableSimBasedRoutingCmdParams(ac)
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

func collectVpgDisableSimBasedRoutingCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgDisableSimBasedRoutingCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgDisableSimBasedRoutingCmd("/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/disable"),
		query:  buildQueryForVpgDisableSimBasedRoutingCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDisableSimBasedRoutingCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgDisableSimBasedRoutingCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDisableSimBasedRoutingCmd() url.Values {
	result := url.Values{}

	return result
}
