package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgCloseGateCmdVpgId holds value of 'vpg_id' option
var VpgCloseGateCmdVpgId string

func init() {
	VpgCloseGateCmd.Flags().StringVar(&VpgCloseGateCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.close_gate.post.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgCloseGateCmd)
}

// VpgCloseGateCmd defines 'close-gate' subcommand
var VpgCloseGateCmd = &cobra.Command{
	Use:   "close-gate",
	Short: TR("virtual_private_gateway.close_gate.post.summary"),
	Long:  TR(`virtual_private_gateway.close_gate.post.description`),
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

		param, err := collectVpgCloseGateCmdParams()
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

func collectVpgCloseGateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgCloseGateCmd("/virtual_private_gateways/{vpg_id}/gate/close"),
		query:  buildQueryForVpgCloseGateCmd(),
	}, nil
}

func buildPathForVpgCloseGateCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgCloseGateCmdVpgId, -1)

	return path
}

func buildQueryForVpgCloseGateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
