// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgSetFixedPublicIpAddressesCmdVpgId holds value of 'vpg_id' option
var VpgSetFixedPublicIpAddressesCmdVpgId string

func InitVpgSetFixedPublicIpAddressesCmd() {
	VpgSetFixedPublicIpAddressesCmd.Flags().StringVar(&VpgSetFixedPublicIpAddressesCmdVpgId, "vpg-id", "", TRAPI(""))

	VpgSetFixedPublicIpAddressesCmd.RunE = VpgSetFixedPublicIpAddressesCmdRunE

	VpgCmd.AddCommand(VpgSetFixedPublicIpAddressesCmd)
}

// VpgSetFixedPublicIpAddressesCmd defines 'set-fixed-public-ip-addresses' subcommand
var VpgSetFixedPublicIpAddressesCmd = &cobra.Command{
	Use:   "set-fixed-public-ip-addresses",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "assignFixedPublicIpAddresses"),
}

func VpgSetFixedPublicIpAddressesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgSetFixedPublicIpAddressesCmdParams(ac)
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

func collectVpgSetFixedPublicIpAddressesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgSetFixedPublicIpAddressesCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgSetFixedPublicIpAddressesCmd("/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses"),
		query:  buildQueryForVpgSetFixedPublicIpAddressesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgSetFixedPublicIpAddressesCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgSetFixedPublicIpAddressesCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgSetFixedPublicIpAddressesCmd() url.Values {
	result := url.Values{}

	return result
}