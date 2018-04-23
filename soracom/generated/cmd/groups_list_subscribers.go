package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsListSubscribersCmdGroupId holds value of 'group_id' option
var GroupsListSubscribersCmdGroupId string

// GroupsListSubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GroupsListSubscribersCmdLastEvaluatedKey string

// GroupsListSubscribersCmdLimit holds value of 'limit' option
var GroupsListSubscribersCmdLimit int64

func init() {
	GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdGroupId, "group-id", "", TRAPI("Target group ID."))

	GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	GroupsListSubscribersCmd.Flags().Int64Var(&GroupsListSubscribersCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	GroupsCmd.AddCommand(GroupsListSubscribersCmd)
}

// GroupsListSubscribersCmd defines 'list-subscribers' subcommand
var GroupsListSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: TRAPI("/groups/{group_id}/subscribers:get:summary"),
	Long:  TRAPI(`/groups/{group_id}/subscribers:get:description`),
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

		param, err := collectGroupsListSubscribersCmdParams(ac)
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

func collectGroupsListSubscribersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsListSubscribersCmd("/groups/{group_id}/subscribers"),
		query:  buildQueryForGroupsListSubscribersCmd(),
	}, nil
}

func buildPathForGroupsListSubscribersCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsListSubscribersCmdGroupId, -1)

	return path
}

func buildQueryForGroupsListSubscribersCmd() string {
	result := []string{}

	if GroupsListSubscribersCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", GroupsListSubscribersCmdLastEvaluatedKey))
	}

	if GroupsListSubscribersCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", GroupsListSubscribersCmdLimit))
	}

	return strings.Join(result, "&")
}
