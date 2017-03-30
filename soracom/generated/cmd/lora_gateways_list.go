package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraGatewaysListCmdLastEvaluatedKey string

// LoraGatewaysListCmdTagName holds value of 'tag_name' option
var LoraGatewaysListCmdTagName string

// LoraGatewaysListCmdTagValue holds value of 'tag_value' option
var LoraGatewaysListCmdTagValue string

// LoraGatewaysListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var LoraGatewaysListCmdTagValueMatchMode string

// LoraGatewaysListCmdLimit holds value of 'limit' option
var LoraGatewaysListCmdLimit int64

func init() {
	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdLastEvaluatedKey, "last-evaluated-key", "", TR("lora_gateways.list.parameters.last_evaluated_key.description"))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagName, "tag-name", "", TR("lora_gateways.list.parameters.tag_name.description"))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValue, "tag-value", "", TR("lora_gateways.list.parameters.tag_value.description"))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValueMatchMode, "tag-value-match-mode", "", TR("lora_gateways.list.parameters.tag_value_match_mode.description"))

	LoraGatewaysListCmd.Flags().Int64Var(&LoraGatewaysListCmdLimit, "limit", 0, TR("lora_gateways.list.parameters.limit.description"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysListCmd)
}

// LoraGatewaysListCmd defines 'list' subcommand
var LoraGatewaysListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("lora_gateways.list.summary"),
	Long:  TR(`lora_gateways.list.description`),
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

		param, err := collectLoraGatewaysListCmdParams()
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

func collectLoraGatewaysListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraGatewaysListCmd("/lora_gateways"),
		query:  buildQueryForLoraGatewaysListCmd(),
	}, nil
}

func buildPathForLoraGatewaysListCmd(path string) string {

	return path
}

func buildQueryForLoraGatewaysListCmd() string {
	result := []string{}

	if LoraGatewaysListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LoraGatewaysListCmdLastEvaluatedKey))
	}

	if LoraGatewaysListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", LoraGatewaysListCmdTagName))
	}

	if LoraGatewaysListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", LoraGatewaysListCmdTagValue))
	}

	if LoraGatewaysListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", LoraGatewaysListCmdTagValueMatchMode))
	}

	if LoraGatewaysListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LoraGatewaysListCmdLimit))
	}

	return strings.Join(result, "&")
}
