// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgReleaseFixedPublicIpAddressesCmdVpgId holds value of 'vpg_id' option
var VpgReleaseFixedPublicIpAddressesCmdVpgId string

func InitVpgReleaseFixedPublicIpAddressesCmd() {
	VpgReleaseFixedPublicIpAddressesCmd.Flags().StringVar(&VpgReleaseFixedPublicIpAddressesCmdVpgId, "vpg-id", "", TRAPI(""))

	VpgReleaseFixedPublicIpAddressesCmd.RunE = VpgReleaseFixedPublicIpAddressesCmdRunE

	VpgCmd.AddCommand(VpgReleaseFixedPublicIpAddressesCmd)
}

// VpgReleaseFixedPublicIpAddressesCmd defines 'release-fixed-public-ip-addresses' subcommand
var VpgReleaseFixedPublicIpAddressesCmd = &cobra.Command{
	Use:   "release-fixed-public-ip-addresses",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses:delete:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "releaseFixedPublicIpAddresses"),
}

func VpgReleaseFixedPublicIpAddressesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgReleaseFixedPublicIpAddressesCmdParams(ac)
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

func collectVpgReleaseFixedPublicIpAddressesCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgReleaseFixedPublicIpAddressesCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgReleaseFixedPublicIpAddressesCmd("/virtual_private_gateways/{vpg_id}/fixed_public_ip_addresses"),
		query:  buildQueryForVpgReleaseFixedPublicIpAddressesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgReleaseFixedPublicIpAddressesCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgReleaseFixedPublicIpAddressesCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgReleaseFixedPublicIpAddressesCmd() url.Values {
	result := url.Values{}

	return result
}
