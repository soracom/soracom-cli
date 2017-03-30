package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysSetNetworkSetCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysSetNetworkSetCmdGatewayId string

// LoraGatewaysSetNetworkSetCmdNetworkSetId holds value of 'networkSetId' option
var LoraGatewaysSetNetworkSetCmdNetworkSetId string

// LoraGatewaysSetNetworkSetCmdBody holds contents of request body to be sent
var LoraGatewaysSetNetworkSetCmdBody string

func init() {
	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdGatewayId, "gateway-id", "", TR("lora_gateways.set_network_set.parameters.gateway_id.description"))

	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdNetworkSetId, "network-set-id", "", TR(""))

	LoraGatewaysSetNetworkSetCmd.Flags().StringVar(&LoraGatewaysSetNetworkSetCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysSetNetworkSetCmd)
}

// LoraGatewaysSetNetworkSetCmd defines 'set-network-set' subcommand
var LoraGatewaysSetNetworkSetCmd = &cobra.Command{
	Use:   "set-network-set",
	Short: TR("lora_gateways.set_network_set.summary"),
	Long:  TR(`lora_gateways.set_network_set.description`),
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

		param, err := collectLoraGatewaysSetNetworkSetCmdParams()
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

func collectLoraGatewaysSetNetworkSetCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraGatewaysSetNetworkSetCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraGatewaysSetNetworkSetCmd("/lora_gateways/{gateway_id}/set_network_set"),
		query:       buildQueryForLoraGatewaysSetNetworkSetCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraGatewaysSetNetworkSetCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysSetNetworkSetCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysSetNetworkSetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraGatewaysSetNetworkSetCmd() (string, error) {
	if LoraGatewaysSetNetworkSetCmdBody != "" {
		if strings.HasPrefix(LoraGatewaysSetNetworkSetCmdBody, "@") {
			fname := strings.TrimPrefix(LoraGatewaysSetNetworkSetCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraGatewaysSetNetworkSetCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraGatewaysSetNetworkSetCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LoraGatewaysSetNetworkSetCmdNetworkSetId != "" {
		result["networkSetId"] = LoraGatewaysSetNetworkSetCmdNetworkSetId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
