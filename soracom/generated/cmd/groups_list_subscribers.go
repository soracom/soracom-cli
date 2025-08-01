// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GroupsListSubscribersCmdGroupId holds value of 'group_id' option
var GroupsListSubscribersCmdGroupId string

// GroupsListSubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GroupsListSubscribersCmdLastEvaluatedKey string

// GroupsListSubscribersCmdLimit holds value of 'limit' option
var GroupsListSubscribersCmdLimit int64

// GroupsListSubscribersCmdPaginate indicates to do pagination or not
var GroupsListSubscribersCmdPaginate bool

// GroupsListSubscribersCmdOutputJSONL indicates to output with jsonl format
var GroupsListSubscribersCmdOutputJSONL bool

func InitGroupsListSubscribersCmd() {
	GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdGroupId, "group-id", "", TRAPI("ID of the target Group."))

	GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the previous page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	GroupsListSubscribersCmd.Flags().Int64Var(&GroupsListSubscribersCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	GroupsListSubscribersCmd.Flags().BoolVar(&GroupsListSubscribersCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	GroupsListSubscribersCmd.Flags().BoolVar(&GroupsListSubscribersCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	GroupsListSubscribersCmd.RunE = GroupsListSubscribersCmdRunE

	GroupsCmd.AddCommand(GroupsListSubscribersCmd)
}

// GroupsListSubscribersCmd defines 'list-subscribers' subcommand
var GroupsListSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: TRAPI("/groups/{group_id}/subscribers:get:summary"),
	Long:  TRAPI(`/groups/{group_id}/subscribers:get:description`) + "\n\n" + createLinkToAPIReference("Group", "listSubscribersInGroup"),
}

func GroupsListSubscribersCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectGroupsListSubscribersCmdParams(ac)
	if err != nil {
		return err
	}

	body, err := ac.callAPI(param)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	if body == "" {
		return nil
	}

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		if GroupsListSubscribersCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectGroupsListSubscribersCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("group_id", "group-id", "path", parsedBody, GroupsListSubscribersCmdGroupId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsListSubscribersCmd("/groups/{group_id}/subscribers"),
		query:  buildQueryForGroupsListSubscribersCmd(),

		doPagination:                      GroupsListSubscribersCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsListSubscribersCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsListSubscribersCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	return path
}

func buildQueryForGroupsListSubscribersCmd() url.Values {
	result := url.Values{}

	if GroupsListSubscribersCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", GroupsListSubscribersCmdLastEvaluatedKey)
	}

	if GroupsListSubscribersCmdLimit != 0 {
		result.Add("limit", sprintf("%d", GroupsListSubscribersCmdLimit))
	}

	return result
}
