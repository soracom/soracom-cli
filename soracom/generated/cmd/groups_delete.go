// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// GroupsDeleteCmdGroupId holds value of 'group_id' option
var GroupsDeleteCmdGroupId string

func init() {
	GroupsDeleteCmd.Flags().StringVar(&GroupsDeleteCmdGroupId, "group-id", "", TRAPI("Target group ID."))
	GroupsCmd.AddCommand(GroupsDeleteCmd)
}

// GroupsDeleteCmd defines 'delete' subcommand
var GroupsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/groups/{group_id}:delete:summary"),
	Long:  TRAPI(`/groups/{group_id}:delete:description`),
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

		param, err := collectGroupsDeleteCmdParams(ac)
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

func collectGroupsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if GroupsDeleteCmdGroupId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "group-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteCmd("/groups/{group_id}"),
		query:  buildQueryForGroupsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsDeleteCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsDeleteCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	return path
}

func buildQueryForGroupsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
