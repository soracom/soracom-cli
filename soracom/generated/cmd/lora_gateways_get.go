package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysGetCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysGetCmdGatewayId string

func init() {
	LoraGatewaysGetCmd.Flags().StringVar(&LoraGatewaysGetCmdGatewayId, "gateway-id", "", TRAPI("Gateway ID of the target LoRa gateway."))

	LoraGatewaysCmd.AddCommand(LoraGatewaysGetCmd)
}

// LoraGatewaysGetCmd defines 'get' subcommand
var LoraGatewaysGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/lora_gateways/{gateway_id}:get:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}:get:description`),
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

		param, err := collectLoraGatewaysGetCmdParams(ac)
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

func collectLoraGatewaysGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraGatewaysGetCmd("/lora_gateways/{gateway_id}"),
		query:  buildQueryForLoraGatewaysGetCmd(),
	}, nil
}

func buildPathForLoraGatewaysGetCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysGetCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
