package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysEnableTerminationCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysEnableTerminationCmdGatewayId string

func init() {
	LoraGatewaysEnableTerminationCmd.Flags().StringVar(&LoraGatewaysEnableTerminationCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRa gateway."))

	LoraGatewaysCmd.AddCommand(LoraGatewaysEnableTerminationCmd)
}

// LoraGatewaysEnableTerminationCmd defines 'enable-termination' subcommand
var LoraGatewaysEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/lora_gateways/{gateway_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/enable_termination:post:description`),
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

		param, err := collectLoraGatewaysEnableTerminationCmdParams(ac)
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

func collectLoraGatewaysEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysEnableTerminationCmd("/lora_gateways/{gateway_id}/enable_termination"),
		query:  buildQueryForLoraGatewaysEnableTerminationCmd(),
	}, nil
}

func buildPathForLoraGatewaysEnableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysEnableTerminationCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysEnableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
