package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysUnsetNetworkSetCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysUnsetNetworkSetCmdGatewayId string

func init() {
	LoraGatewaysUnsetNetworkSetCmd.Flags().StringVar(&LoraGatewaysUnsetNetworkSetCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRa gateway."))

	LoraGatewaysCmd.AddCommand(LoraGatewaysUnsetNetworkSetCmd)
}

// LoraGatewaysUnsetNetworkSetCmd defines 'unset-network-set' subcommand
var LoraGatewaysUnsetNetworkSetCmd = &cobra.Command{
	Use:   "unset-network-set",
	Short: TRAPI("/lora_gateways/{gateway_id}/unset_network_set:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/unset_network_set:post:description`),
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

		param, err := collectLoraGatewaysUnsetNetworkSetCmdParams(ac)
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

func collectLoraGatewaysUnsetNetworkSetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysUnsetNetworkSetCmd("/lora_gateways/{gateway_id}/unset_network_set"),
		query:  buildQueryForLoraGatewaysUnsetNetworkSetCmd(),
	}, nil
}

func buildPathForLoraGatewaysUnsetNetworkSetCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysUnsetNetworkSetCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysUnsetNetworkSetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
