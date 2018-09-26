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
	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID of the last network set retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	LoraNetworkSetsListCmd.Flags().StringVar(&LoraNetworkSetsListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	LoraNetworkSetsListCmd.Flags().Int64Var(&LoraNetworkSetsListCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa devices to retrieve."))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsListCmd)
}

// LoraNetworkSetsListCmd defines 'list' subcommand
var LoraNetworkSetsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lora_network_sets:get:summary"),
	Long:  TRAPI(`/lora_network_sets:get:description`),
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

		param, err := collectLoraNetworkSetsListCmdParams(ac)
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

func collectLoraNetworkSetsListCmdParams(ac *apiClient) (*apiParams, error) {

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
