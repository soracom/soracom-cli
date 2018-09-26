package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgListGatePeersCmdVpgId holds value of 'vpg_id' option
var VpgListGatePeersCmdVpgId string

func init() {
	VpgListGatePeersCmd.Flags().StringVar(&VpgListGatePeersCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCmd.AddCommand(VpgListGatePeersCmd)
}

// VpgListGatePeersCmd defines 'list-gate-peers' subcommand
var VpgListGatePeersCmd = &cobra.Command{
	Use:   "list-gate-peers",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/peers:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/peers:get:description`),
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

		param, err := collectVpgListGatePeersCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectVpgListGatePeersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgListGatePeersCmd("/virtual_private_gateways/{vpg_id}/gate/peers"),
		query:  buildQueryForVpgListGatePeersCmd(),
	}, nil
}

func buildPathForVpgListGatePeersCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgListGatePeersCmdVpgId, -1)

	return path
}

func buildQueryForVpgListGatePeersCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
