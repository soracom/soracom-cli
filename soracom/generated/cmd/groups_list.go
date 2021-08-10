// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// GroupsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GroupsListCmdLastEvaluatedKey string

// GroupsListCmdTagName holds value of 'tag_name' option
var GroupsListCmdTagName string

// GroupsListCmdTagValue holds value of 'tag_value' option
var GroupsListCmdTagValue string

// GroupsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var GroupsListCmdTagValueMatchMode string

// GroupsListCmdLimit holds value of 'limit' option
var GroupsListCmdLimit int64

// GroupsListCmdPaginate indicates to do pagination or not
var GroupsListCmdPaginate bool

func init() {
	GroupsListCmd.Flags().StringVar(&GroupsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The last Group ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next group onward."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagName, "tag-name", "", TRAPI("Tag name of the group. Filters through all groups that exactly match the tag name. When tag_name is specified, tag_value is required."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValue, "tag-value", "", TRAPI("Tag value of the groups."))

	GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode."))

	GroupsListCmd.Flags().Int64Var(&GroupsListCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	GroupsListCmd.Flags().BoolVar(&GroupsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	GroupsCmd.AddCommand(GroupsListCmd)
}

// GroupsListCmd defines 'list' subcommand
var GroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/groups:get:summary"),
	Long:  TRAPI(`/groups:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
		}

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

		param, err := collectGroupsListCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectGroupsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsListCmd("/groups"),
		query:  buildQueryForGroupsListCmd(),

		doPagination:                      GroupsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsListCmd(path string) string {

	return path
}

func buildQueryForGroupsListCmd() url.Values {
	result := url.Values{}

	if GroupsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", GroupsListCmdLastEvaluatedKey)
	}

	if GroupsListCmdTagName != "" {
		result.Add("tag_name", GroupsListCmdTagName)
	}

	if GroupsListCmdTagValue != "" {
		result.Add("tag_value", GroupsListCmdTagValue)
	}

	if GroupsListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", GroupsListCmdTagValueMatchMode)
	}

	if GroupsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", GroupsListCmdLimit))
	}

	return result
}
