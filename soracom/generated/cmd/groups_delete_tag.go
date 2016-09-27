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
	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdGroupId, "group-id", "", TR("groups.delete_group_tag.delete.parameters.group_id.description"))

	GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdTagName, "tag-name", "", TR("groups.delete_group_tag.delete.parameters.tag_name.description"))

	GroupsCmd.AddCommand(GroupsDeleteTagCmd)
}

// GroupsDeleteTagCmd defines 'delete-tag' subcommand
var GroupsDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TR("groups.delete_group_tag.delete.summary"),
	Long:  TR(`groups.delete_group_tag.delete.description`),
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

		param, err := collectGroupsDeleteTagCmdParams()
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

func collectGroupsDeleteTagCmdParams() (*apiParams, error) {

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
