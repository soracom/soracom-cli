package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgListIpAddressMapEntriesCmdVpgId holds value of 'vpg_id' option
var VpgListIpAddressMapEntriesCmdVpgId string

func init() {
	VpgListIpAddressMapEntriesCmd.Flags().StringVar(&VpgListIpAddressMapEntriesCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCmd.AddCommand(VpgListIpAddressMapEntriesCmd)
}

// VpgListIpAddressMapEntriesCmd defines 'list-ip-address-map-entries' subcommand
var VpgListIpAddressMapEntriesCmd = &cobra.Command{
	Use:   "list-ip-address-map-entries",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/ip_address_map:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/ip_address_map:get:description`),
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

		param, err := collectVpgListIpAddressMapEntriesCmdParams(ac)
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

func collectVpgListIpAddressMapEntriesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgListIpAddressMapEntriesCmd("/virtual_private_gateways/{vpg_id}/ip_address_map"),
		query:  buildQueryForVpgListIpAddressMapEntriesCmd(),
	}, nil
}

func buildPathForVpgListIpAddressMapEntriesCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgListIpAddressMapEntriesCmdVpgId, -1)

	return path
}

func buildQueryForVpgListIpAddressMapEntriesCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
