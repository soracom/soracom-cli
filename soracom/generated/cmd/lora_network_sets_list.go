package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraNetworkSetsListCmdLastEvaluatedKey string

// LoraNetworkSetsListCmdTagName holds value of 'tag_name' option
var LoraNetworkSetsListCmdTagName string

// LoraNetworkSetsListCmdTagValue holds value of 'tag_value' option
var LoraNetworkSetsListCmdTagValue string

// LoraNetworkSetsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var LoraNetworkSetsListCmdTagValueMatchMode string

// LoraNetworkSetsListCmdLimit holds value of 'limit' option
var LoraNetworkSetsListCmdLimit int64

func init() {
	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdLastEvaluatedKey, "last-evaluated-key", "", TR("lora_network_sets.list.parameters.last_evaluated_key.description"))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagName, "tag-name", "", TR("lora_network_sets.list.parameters.tag_name.description"))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValue, "tag-value", "", TR("lora_network_sets.list.parameters.tag_value.description"))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValueMatchMode, "tag-value-match-mode", "", TR("lora_network_sets.list.parameters.tag_value_match_mode.description"))

	LoraNetworkSetsListCmd.Flags().Int64Var(&LoraNetworkSetsListCmdLimit, "limit", 0, TR("lora_network_sets.list.parameters.limit.description"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsListCmd)
}

// LoraNetworkSetsListCmd defines 'list' subcommand
var LoraNetworkSetsListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("lora_network_sets.list.summary"),
	Long:  TR(`lora_network_sets.list.description`),
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

		param, err := collectLoraNetworkSetsListCmdParams()
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

func collectLoraNetworkSetsListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsListCmd("/lora_network_sets"),
		query:  buildQueryForLoraNetworkSetsListCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsListCmd(path string) string {

	return path
}

func buildQueryForLoraNetworkSetsListCmd() string {
	result := []string{}

	if LoraNetworkSetsListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LoraNetworkSetsListCmdLastEvaluatedKey))
	}

	if LoraNetworkSetsListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", LoraNetworkSetsListCmdTagName))
	}

	if LoraNetworkSetsListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", LoraNetworkSetsListCmdTagValue))
	}

	if LoraNetworkSetsListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", LoraNetworkSetsListCmdTagValueMatchMode))
	}

	if LoraNetworkSetsListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LoraNetworkSetsListCmdLimit))
	}

	return strings.Join(result, "&")
}
