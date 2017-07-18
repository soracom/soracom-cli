package cmd

import (
	"os"
	"strings"

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

		param, err := collectVpgDeleteIpAddressMapEntryCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectVpgDeleteIpAddressMapEntryCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteIpAddressMapEntryCmd("/virtual_private_gateways/{vpg_id}/ip_address_map/{key}"),
		query:  buildQueryForVpgDeleteIpAddressMapEntryCmd(),
	}, nil
}

func buildPathForVpgDeleteIpAddressMapEntryCmd(path string) string {

	path = strings.Replace(path, "{"+"key"+"}", VpgDeleteIpAddressMapEntryCmdKey, -1)

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgDeleteIpAddressMapEntryCmdVpgId, -1)

	return path
}

func buildQueryForVpgDeleteIpAddressMapEntryCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
