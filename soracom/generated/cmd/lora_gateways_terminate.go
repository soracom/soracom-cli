// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraGatewaysTerminateCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysTerminateCmdGatewayId string

func init() {
	LoraGatewaysTerminateCmd.Flags().StringVar(&LoraGatewaysTerminateCmdGatewayId, "gateway-id", "", TRAPI("Device ID of the target LoRa gateway."))
	LoraGatewaysCmd.AddCommand(LoraGatewaysTerminateCmd)
}

// LoraGatewaysTerminateCmd defines 'terminate' subcommand
var LoraGatewaysTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/lora_gateways/{gateway_id}/terminate:post:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/terminate:post:description`) + "\n\n" + createLinkToAPIReference("LoraGateway", "terminateLoraGateway"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLoraGatewaysTerminateCmdParams(ac)
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
	},
}

func collectLoraGatewaysTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("gateway_id", "gateway-id", "path", parsedBody, LoraGatewaysTerminateCmdGatewayId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysTerminateCmd("/lora_gateways/{gateway_id}/terminate"),
		query:  buildQueryForLoraGatewaysTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysTerminateCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysTerminateCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysTerminateCmd() url.Values {
	result := url.Values{}

	return result
}
