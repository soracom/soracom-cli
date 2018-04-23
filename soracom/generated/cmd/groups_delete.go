package cmd

import (
	"os"
	"strings"

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

func collectGroupsDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteCmd("/groups/{group_id}"),
		query:  buildQueryForGroupsDeleteCmd(),
	}, nil
}

func buildPathForGroupsDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsDeleteCmdGroupId, -1)

	return path
}

func buildQueryForGroupsDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
