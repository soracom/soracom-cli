package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgDeleteMirroringPeerCmdIpaddr holds value of 'ipaddr' option
var VpgDeleteMirroringPeerCmdIpaddr string

// VpgDeleteMirroringPeerCmdVpgId holds value of 'vpg_id' option
var VpgDeleteMirroringPeerCmdVpgId string

func init() {
	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdIpaddr, "ipaddr", "", TRAPI("IP address of mirroring peer"))

	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgCmd.AddCommand(VpgDeleteMirroringPeerCmd)
}

// VpgDeleteMirroringPeerCmd defines 'delete-mirroring-peer' subcommand
var VpgDeleteMirroringPeerCmd = &cobra.Command{
	Use:   "delete-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:delete:description`),
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

		param, err := collectVpgDeleteMirroringPeerCmdParams(ac)
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

func collectVpgDeleteMirroringPeerCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteMirroringPeerCmd("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}"),
		query:  buildQueryForVpgDeleteMirroringPeerCmd(),
	}, nil
}

func buildPathForVpgDeleteMirroringPeerCmd(path string) string {

	path = strings.Replace(path, "{"+"ipaddr"+"}", VpgDeleteMirroringPeerCmdIpaddr, -1)

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgDeleteMirroringPeerCmdVpgId, -1)

	return path
}

func buildQueryForVpgDeleteMirroringPeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
