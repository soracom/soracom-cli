package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysTerminateCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysTerminateCmdGatewayId string

func init() {
	LoraGatewaysTerminateCmd.Flags().StringVar(&LoraGatewaysTerminateCmdGatewayId, "gateway-id", "", TR("lora_gateways.terminate.parameters.gateway_id.description"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysTerminateCmd)
}

// LoraGatewaysTerminateCmd defines 'terminate' subcommand
var LoraGatewaysTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TR("lora_gateways.terminate.summary"),
	Long:  TR(`lora_gateways.terminate.description`),
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

		param, err := collectLoraGatewaysTerminateCmdParams()
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

func collectLoraGatewaysTerminateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysTerminateCmd("/lora_gateways/{gateway_id}/terminate"),
		query:  buildQueryForLoraGatewaysTerminateCmd(),
	}, nil
}

func buildPathForLoraGatewaysTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysTerminateCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
