package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgOpenGateCmdVpgId holds value of 'vpg_id' option
var VpgOpenGateCmdVpgId string

func init() {
	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.open_gate.post.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgOpenGateCmd)
}

// VpgOpenGateCmd defines 'open-gate' subcommand
var VpgOpenGateCmd = &cobra.Command{
	Use:   "open-gate",
	Short: TR("virtual_private_gateway.open_gate.post.summary"),
	Long:  TR(`virtual_private_gateway.open_gate.post.description`),
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

		param, err := collectVpgOpenGateCmdParams()
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

func collectVpgOpenGateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgOpenGateCmd("/virtual_private_gateways/{vpg_id}/gate/open"),
		query:  buildQueryForVpgOpenGateCmd(),
	}, nil
}

func buildPathForVpgOpenGateCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgOpenGateCmdVpgId, -1)

	return path
}

func buildQueryForVpgOpenGateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
