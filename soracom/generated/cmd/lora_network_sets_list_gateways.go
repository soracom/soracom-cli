package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsListGatewaysCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraNetworkSetsListGatewaysCmdLastEvaluatedKey string

// LoraNetworkSetsListGatewaysCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsListGatewaysCmdNsId string

// LoraNetworkSetsListGatewaysCmdLimit holds value of 'limit' option
var LoraNetworkSetsListGatewaysCmdLimit int64

func init() {
	LoraNetworkSetsListGatewaysCmd.Flags().StringVar(&LoraNetworkSetsListGatewaysCmdLastEvaluatedKey, "last-evaluated-key", "", TR("lora_network_sets.list_gateways.parameters.last_evaluated_key.description"))

	LoraNetworkSetsListGatewaysCmd.Flags().StringVar(&LoraNetworkSetsListGatewaysCmdNsId, "ns-id", "", TR("lora_network_sets.list_gateways.parameters.ns_id.description"))

	LoraNetworkSetsListGatewaysCmd.Flags().Int64Var(&LoraNetworkSetsListGatewaysCmdLimit, "limit", 0, TR("lora_network_sets.list_gateways.parameters.limit.description"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsListGatewaysCmd)
}

// LoraNetworkSetsListGatewaysCmd defines 'list-gateways' subcommand
var LoraNetworkSetsListGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: TR("lora_network_sets.list_gateways.summary"),
	Long:  TR(`lora_network_sets.list_gateways.description`),
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

		param, err := collectLoraNetworkSetsListGatewaysCmdParams()
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

func collectLoraNetworkSetsListGatewaysCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsListGatewaysCmd("/lora_network_sets/{ns_id}/gateways"),
		query:  buildQueryForLoraNetworkSetsListGatewaysCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsListGatewaysCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsListGatewaysCmdNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsListGatewaysCmd() string {
	result := []string{}

	if LoraNetworkSetsListGatewaysCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LoraNetworkSetsListGatewaysCmdLastEvaluatedKey))
	}

	if LoraNetworkSetsListGatewaysCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LoraNetworkSetsListGatewaysCmdLimit))
	}

	return strings.Join(result, "&")
}
