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
	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The device ID of the last device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	LoraGatewaysListCmd.Flags().StringVar(&LoraGatewaysListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	LoraGatewaysListCmd.Flags().Int64Var(&LoraGatewaysListCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa devices to retrieve."))

	LoraGatewaysCmd.AddCommand(LoraGatewaysListCmd)
}

// LoraGatewaysListCmd defines 'list' subcommand
var LoraGatewaysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lora_gateways:get:summary"),
	Long:  TRAPI(`/lora_gateways:get:description`),
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

		param, err := collectLoraGatewaysListCmdParams(ac)
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

func collectLoraGatewaysListCmdParams(ac *apiClient) (*apiParams, error) {

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
