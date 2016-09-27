package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgGetCmdVpgId holds value of 'vpg_id' option
var VpgGetCmdVpgId string

func init() {
	VpgGetCmd.Flags().StringVar(&VpgGetCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.get_virtual_private_gateway.get.parameters.vpg_id.description"))

	VpgCmd.AddCommand(VpgGetCmd)
}

// VpgGetCmd defines 'get' subcommand
var VpgGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("virtual_private_gateway.get_virtual_private_gateway.get.summary"),
	Long:  TR(`virtual_private_gateway.get_virtual_private_gateway.get.description`),
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

		param, err := collectVpgGetCmdParams()
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

func collectVpgGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgGetCmd("/virtual_private_gateways/{vpg_id}"),
		query:  buildQueryForVpgGetCmd(),
	}, nil
}

func buildPathForVpgGetCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgGetCmdVpgId, -1)

	return path
}

func buildQueryForVpgGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
