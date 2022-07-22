// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgEnableGatePrivacySeparatorCmdVpgId holds value of 'vpg_id' option
var VpgEnableGatePrivacySeparatorCmdVpgId string

func init() {
	VpgEnableGatePrivacySeparatorCmd.Flags().StringVar(&VpgEnableGatePrivacySeparatorCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))
	VpgCmd.AddCommand(VpgEnableGatePrivacySeparatorCmd)
}

// VpgEnableGatePrivacySeparatorCmd defines 'enable-gate-privacy-separator' subcommand
var VpgEnableGatePrivacySeparatorCmd = &cobra.Command{
	Use:   "enable-gate-privacy-separator",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/enable_privacy_separator:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/enable_privacy_separator:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "enableGatePrivacySeparator"),
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

		param, err := collectVpgEnableGatePrivacySeparatorCmdParams(ac)
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

func collectVpgEnableGatePrivacySeparatorCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgEnableGatePrivacySeparatorCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgEnableGatePrivacySeparatorCmd("/virtual_private_gateways/{vpg_id}/gate/enable_privacy_separator"),
		query:  buildQueryForVpgEnableGatePrivacySeparatorCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgEnableGatePrivacySeparatorCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgEnableGatePrivacySeparatorCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgEnableGatePrivacySeparatorCmd() url.Values {
	result := url.Values{}

	return result
}
