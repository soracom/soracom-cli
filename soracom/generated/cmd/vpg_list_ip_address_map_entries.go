package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var VpgListIpAddressMapEntriesCmdVpgId string

func init() {
	VpgListIpAddressMapEntriesCmd.Flags().StringVar(&VpgListIpAddressMapEntriesCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.list_virtual_private_gateway_ip_address_map_entries.get.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgListIpAddressMapEntriesCmd)
}

var VpgListIpAddressMapEntriesCmd = &cobra.Command{
	Use:   "list-ip-address-map-entries",
	Short: TR("virtual_private_gateway.list_virtual_private_gateway_ip_address_map_entries.get.summary"),
	Long:  TR(`virtual_private_gateway.list_virtual_private_gateway_ip_address_map_entries.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectVpgListIpAddressMapEntriesCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectVpgListIpAddressMapEntriesCmdParams() (*apiParams, error) {

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
