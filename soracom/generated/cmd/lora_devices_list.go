package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LoraDevicesListCmdLastEvaluatedKey string

// LoraDevicesListCmdTagName holds value of 'tag_name' option
var LoraDevicesListCmdTagName string

// LoraDevicesListCmdTagValue holds value of 'tag_value' option
var LoraDevicesListCmdTagValue string

// LoraDevicesListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var LoraDevicesListCmdTagValueMatchMode string

// LoraDevicesListCmdLimit holds value of 'limit' option
var LoraDevicesListCmdLimit int64

func init() {
	LoraDevicesListCmd.Flags().StringVar(&LoraDevicesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The device ID of the last device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next device onward."))

	LoraDevicesListCmd.Flags().StringVar(&LoraDevicesListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	LoraDevicesListCmd.Flags().StringVar(&LoraDevicesListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	LoraDevicesListCmd.Flags().StringVar(&LoraDevicesListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	LoraDevicesListCmd.Flags().Int64Var(&LoraDevicesListCmdLimit, "limit", 0, TRAPI("Maximum number of LoRa devices to retrieve."))

	LoraDevicesCmd.AddCommand(LoraDevicesListCmd)
}

// LoraDevicesListCmd defines 'list' subcommand
var LoraDevicesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/lora_devices:get:summary"),
	Long:  TRAPI(`/lora_devices:get:description`),
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

		param, err := collectLoraDevicesListCmdParams(ac)
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

func collectLoraDevicesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraDevicesListCmd("/lora_devices"),
		query:  buildQueryForLoraDevicesListCmd(),
	}, nil
}

func buildPathForLoraDevicesListCmd(path string) string {

	return path
}

func buildQueryForLoraDevicesListCmd() string {
	result := []string{}

	if LoraDevicesListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LoraDevicesListCmdLastEvaluatedKey))
	}

	if LoraDevicesListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", LoraDevicesListCmdTagName))
	}

	if LoraDevicesListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", LoraDevicesListCmdTagValue))
	}

	if LoraDevicesListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", LoraDevicesListCmdTagValueMatchMode))
	}

	if LoraDevicesListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LoraDevicesListCmdLimit))
	}

	return strings.Join(result, "&")
}
