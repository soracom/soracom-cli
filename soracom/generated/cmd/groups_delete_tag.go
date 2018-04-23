package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsDeleteTagCmdGroupId holds value of 'group_id' option
var GroupsDeleteTagCmdGroupId string

// GroupsDeleteTagCmdTagName holds value of 'tag_name' option
var GroupsDeleteTagCmdTagName string

func init() {
	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdGroupId, "group-id", "", TRAPI("Target group ID."))

	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	GroupsCmd.AddCommand(GroupsDeleteTagCmd)
}

// GroupsDeleteTagCmd defines 'delete-tag' subcommand
var GroupsDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/groups/{group_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/groups/{group_id}/tags/{tag_name}:delete:description`),
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

		param, err := collectGroupsDeleteTagCmdParams(ac)
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

func collectGroupsDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteTagCmd("/groups/{group_id}/tags/{tag_name}"),
		query:  buildQueryForGroupsDeleteTagCmd(),
	}, nil
}

func buildPathForGroupsDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsDeleteTagCmdGroupId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", GroupsDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForGroupsDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
