package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var GroupsListCmdLastEvaluatedKey string

var GroupsListCmdTagName string

var GroupsListCmdTagValue string

var GroupsListCmdTagValueMatchMode string

var GroupsListCmdLimit int64

func init() {
	GroupsListCmd.Flags().StringVar(&GroupsListCmdLastEvaluatedKey, "last-evaluated-key", "", TR("groups.list_groups.get.parameters.last_evaluated_key.description"))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagName, "tag-name", "", TR("groups.list_groups.get.parameters.tag_name.description"))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValue, "tag-value", "", TR("groups.list_groups.get.parameters.tag_value.description"))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValueMatchMode, "tag-value-match-mode", "", TR("groups.list_groups.get.parameters.tag_value_match_mode.description"))

	GroupsListCmd.Flags().Int64Var(&GroupsListCmdLimit, "limit", 0, TR("groups.list_groups.get.parameters.limit.description"))

	GroupsCmd.AddCommand(GroupsListCmd)
}

var GroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("groups.list_groups.get.summary"),
	Long:  TR(`groups.list_groups.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectGroupsListCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectGroupsListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsListCmd("/groups"),
		query:  buildQueryForGroupsListCmd(),
	}, nil
}

func buildPathForGroupsListCmd(path string) string {

	return path
}

func buildQueryForGroupsListCmd() string {
	result := []string{}

	if GroupsListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", GroupsListCmdLastEvaluatedKey))
	}

	if GroupsListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", GroupsListCmdTagName))
	}

	if GroupsListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", GroupsListCmdTagValue))
	}

	if GroupsListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", GroupsListCmdTagValueMatchMode))
	}

	if GroupsListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", GroupsListCmdLimit))
	}

	return strings.Join(result, "&")
}
