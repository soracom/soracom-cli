// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgUnsetInspectionCmdVpgId holds value of 'vpg_id' option
var VpgUnsetInspectionCmdVpgId string

func InitVpgUnsetInspectionCmd() {
	VpgUnsetInspectionCmd.Flags().StringVar(&VpgUnsetInspectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgUnsetInspectionCmd.RunE = VpgUnsetInspectionCmdRunE

	VpgCmd.AddCommand(VpgUnsetInspectionCmd)
}

// VpgUnsetInspectionCmd defines 'unset-inspection' subcommand
var VpgUnsetInspectionCmd = &cobra.Command{
	Use:   "unset-inspection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/unset_inspection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/unset_inspection:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "unsetInspectionConfiguration"),
}

func VpgUnsetInspectionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgUnsetInspectionCmdParams(ac)
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

func collectVpgUnsetInspectionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgUnsetInspectionCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgUnsetInspectionCmd("/virtual_private_gateways/{vpg_id}/junction/unset_inspection"),
		query:  buildQueryForVpgUnsetInspectionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgUnsetInspectionCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgUnsetInspectionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgUnsetInspectionCmd() url.Values {
	result := url.Values{}

	return result
}
