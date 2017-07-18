package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgDeleteMirroringPeerCmdId holds value of 'id' option
var VpgDeleteMirroringPeerCmdId string

// VpgDeleteMirroringPeerCmdIpaddr holds value of 'ipaddr' option
var VpgDeleteMirroringPeerCmdIpaddr string

func init() {
	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdId, "id", "", TRAPI("VPG ID"))

	VpgDeleteMirroringPeerCmd.Flags().StringVar(&VpgDeleteMirroringPeerCmdIpaddr, "ipaddr", "", TRAPI("IP address of mirroring peer"))

	VpgCmd.AddCommand(VpgDeleteMirroringPeerCmd)
}

// VpgDeleteMirroringPeerCmd defines 'delete-mirroring-peer' subcommand
var VpgDeleteMirroringPeerCmd = &cobra.Command{
	Use:   "delete-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}:delete:description`),
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

		param, err := collectVpgDeleteMirroringPeerCmdParams()
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

func collectVpgDeleteMirroringPeerCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteMirroringPeerCmd("/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}"),
		query:  buildQueryForVpgDeleteMirroringPeerCmd(),
	}, nil
}

func buildPathForVpgDeleteMirroringPeerCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgDeleteMirroringPeerCmdId, -1)

	path = strings.Replace(path, "{"+"ipaddr"+"}", VpgDeleteMirroringPeerCmdIpaddr, -1)

	return path
}

func buildQueryForVpgDeleteMirroringPeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
