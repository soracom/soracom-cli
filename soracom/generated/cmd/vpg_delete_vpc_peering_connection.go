package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgDeleteVpcPeeringConnectionCmdPcxId holds value of 'pcx_id' option
var VpgDeleteVpcPeeringConnectionCmdPcxId string

// VpgDeleteVpcPeeringConnectionCmdVpgId holds value of 'vpg_id' option
var VpgDeleteVpcPeeringConnectionCmdVpgId string

func init() {
	VpgDeleteVpcPeeringConnectionCmd.Flags().StringVar(&VpgDeleteVpcPeeringConnectionCmdPcxId, "pcx-id", "", TR("virtual_private_gateway.delete_vpc_peering_connection.delete.parameters.pcx_id.description"))

	VpgDeleteVpcPeeringConnectionCmd.Flags().StringVar(&VpgDeleteVpcPeeringConnectionCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.delete_vpc_peering_connection.delete.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgDeleteVpcPeeringConnectionCmd)
}

// VpgDeleteVpcPeeringConnectionCmd defines 'delete-vpc-peering-connection' subcommand
var VpgDeleteVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-peering-connection",
	Short: TR("virtual_private_gateway.delete_vpc_peering_connection.delete.summary"),
	Long:  TR(`virtual_private_gateway.delete_vpc_peering_connection.delete.description`),
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

		param, err := collectVpgDeleteVpcPeeringConnectionCmdParams()
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

func collectVpgDeleteVpcPeeringConnectionCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteVpcPeeringConnectionCmd("/virtual_private_gateways/{vpg_id}/vpc_peering_connections/{pcx_id}"),
		query:  buildQueryForVpgDeleteVpcPeeringConnectionCmd(),
	}, nil
}

func buildPathForVpgDeleteVpcPeeringConnectionCmd(path string) string {

	path = strings.Replace(path, "{"+"pcx_id"+"}", VpgDeleteVpcPeeringConnectionCmdPcxId, -1)

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgDeleteVpcPeeringConnectionCmdVpgId, -1)

	return path
}

func buildQueryForVpgDeleteVpcPeeringConnectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
