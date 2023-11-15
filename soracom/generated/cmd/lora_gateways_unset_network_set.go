// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraGatewaysUnsetNetworkSetCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysUnsetNetworkSetCmdGatewayId string

func InitLoraGatewaysUnsetNetworkSetCmd() {
	LoraGatewaysUnsetNetworkSetCmd.Flags().StringVar(&LoraGatewaysUnsetNetworkSetCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRaWAN gateway."))

	LoraGatewaysUnsetNetworkSetCmd.RunE = LoraGatewaysUnsetNetworkSetCmdRunE

	LoraGatewaysCmd.AddCommand(LoraGatewaysUnsetNetworkSetCmd)
}

// LoraGatewaysUnsetNetworkSetCmd defines 'unset-network-set' subcommand
var LoraGatewaysUnsetNetworkSetCmd = &cobra.Command{
	Use:   "unset-network-set",
	Short: TRAPI("/lora_gateways/{gateway_id}/unset_network_set:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/unset_network_set:post:description`) + "\n\n" + createLinkToAPIReference("LoraGateway", "unsetLoraNetworkSet"),
}

func LoraGatewaysUnsetNetworkSetCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectLoraGatewaysUnsetNetworkSetCmdParams(ac)
	if err != nil {
		return err
	}

	body, err := ac.callAPI(param)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	if body == "" {
		return nil
	}

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectLoraGatewaysUnsetNetworkSetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("gateway_id", "gateway-id", "path", parsedBody, LoraGatewaysUnsetNetworkSetCmdGatewayId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysUnsetNetworkSetCmd("/lora_gateways/{gateway_id}/unset_network_set"),
		query:  buildQueryForLoraGatewaysUnsetNetworkSetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysUnsetNetworkSetCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysUnsetNetworkSetCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysUnsetNetworkSetCmd() url.Values {
	result := url.Values{}

	return result
}
