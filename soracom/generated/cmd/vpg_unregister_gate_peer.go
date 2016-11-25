package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgUnregisterGatePeerCmdOuterIpAddress holds value of 'outer_ip_address' option
var VpgUnregisterGatePeerCmdOuterIpAddress string

// VpgUnregisterGatePeerCmdVpgId holds value of 'vpg_id' option
var VpgUnregisterGatePeerCmdVpgId string

func init() {
	VpgUnregisterGatePeerCmd.Flags().StringVar(&VpgUnregisterGatePeerCmdOuterIpAddress, "outer-ip-address", "", TR("virtual_private_gateway.unregister_virtual_private_gateway_peer.delete.parameters.host_id.description"))

	VpgUnregisterGatePeerCmd.Flags().StringVar(&VpgUnregisterGatePeerCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.unregister_virtual_private_gateway_peer.delete.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgUnregisterGatePeerCmd)
}

// VpgUnregisterGatePeerCmd defines 'unregister-gate-peer' subcommand
var VpgUnregisterGatePeerCmd = &cobra.Command{
	Use:   "unregister-gate-peer",
	Short: TR("virtual_private_gateway.unregister_virtual_private_gateway_peer.delete.summary"),
	Long:  TR(`virtual_private_gateway.unregister_virtual_private_gateway_peer.delete.description`),
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

		param, err := collectVpgUnregisterGatePeerCmdParams()
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

func collectVpgUnregisterGatePeerCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgUnregisterGatePeerCmd("/virtual_private_gateways/{vpg_id}/gate/peers/{outer_ip_address}"),
		query:  buildQueryForVpgUnregisterGatePeerCmd(),
	}, nil
}

func buildPathForVpgUnregisterGatePeerCmd(path string) string {

	path = strings.Replace(path, "{"+"outer_ip_address"+"}", VpgUnregisterGatePeerCmdOuterIpAddress, -1)

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgUnregisterGatePeerCmdVpgId, -1)

	return path
}

func buildQueryForVpgUnregisterGatePeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
