package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgOpenGateCmdVpgId holds value of 'vpg_id' option
var VpgOpenGateCmdVpgId string

func init() {
	VpgOpenGateCmd.Flags().StringVar(&VpgOpenGateCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCmd.AddCommand(VpgOpenGateCmd)
}

// VpgOpenGateCmd defines 'open-gate' subcommand
var VpgOpenGateCmd = &cobra.Command{
	Use:   "open-gate",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/open:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/open:post:description`),
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

		param, err := collectVpgOpenGateCmdParams()
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
