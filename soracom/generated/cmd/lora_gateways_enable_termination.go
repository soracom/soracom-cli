// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
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

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectLoraGatewaysEnableTerminationCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLoraGatewaysEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if LoraGatewaysEnableTerminationCmdGatewayId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "gateway-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraGatewaysEnableTerminationCmd("/lora_gateways/{gateway_id}/enable_termination"),
		query:  buildQueryForLoraGatewaysEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysEnableTerminationCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysEnableTerminationCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
