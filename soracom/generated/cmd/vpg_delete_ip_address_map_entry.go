// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDeleteIpAddressMapEntryCmdKey holds value of 'key' option
var VpgDeleteIpAddressMapEntryCmdKey string

// VpgDeleteIpAddressMapEntryCmdVpgId holds value of 'vpg_id' option
var VpgDeleteIpAddressMapEntryCmdVpgId string

func init() {
	VpgDeleteIpAddressMapEntryCmd.Flags().StringVar(&VpgDeleteIpAddressMapEntryCmdKey, "key", "", TRAPI("Target key to remove."))

	VpgDeleteIpAddressMapEntryCmd.Flags().StringVar(&VpgDeleteIpAddressMapEntryCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))
	VpgCmd.AddCommand(VpgDeleteIpAddressMapEntryCmd)
}

// VpgDeleteIpAddressMapEntryCmd defines 'delete-ip-address-map-entry' subcommand
var VpgDeleteIpAddressMapEntryCmd = &cobra.Command{
	Use:   "delete-ip-address-map-entry",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/ip_address_map/{key}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/ip_address_map/{key}:delete:description`),
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

		param, err := collectVpgDeleteIpAddressMapEntryCmdParams(ac)
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

func collectVpgDeleteIpAddressMapEntryCmdParams(ac *apiClient) (*apiParams, error) {
	if VpgDeleteIpAddressMapEntryCmdKey == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "key")
	}

	if VpgDeleteIpAddressMapEntryCmdVpgId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "vpg-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteIpAddressMapEntryCmd("/virtual_private_gateways/{vpg_id}/ip_address_map/{key}"),
		query:  buildQueryForVpgDeleteIpAddressMapEntryCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDeleteIpAddressMapEntryCmd(path string) string {

	escapedKey := url.PathEscape(VpgDeleteIpAddressMapEntryCmdKey)

	path = strReplace(path, "{"+"key"+"}", escapedKey, -1)

	escapedVpgId := url.PathEscape(VpgDeleteIpAddressMapEntryCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDeleteIpAddressMapEntryCmd() url.Values {
	result := url.Values{}

	return result
}
