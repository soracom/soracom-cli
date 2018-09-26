package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsGetCmdGroupId holds value of 'group_id' option
var GroupsGetCmdGroupId string

func init() {
	GroupsGetCmd.Flags().StringVar(&GroupsGetCmdGroupId, "group-id", "", TRAPI("Target group ID."))

	GroupsCmd.AddCommand(GroupsGetCmd)
}

// GroupsGetCmd defines 'get' subcommand
var GroupsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/groups/{group_id}:get:summary"),
	Long:  TRAPI(`/groups/{group_id}:get:description`),
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

		param, err := collectGroupsGetCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectGroupsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGroupsGetCmd("/groups/{group_id}"),
		query:  buildQueryForGroupsGetCmd(),
	}, nil
}

func buildPathForGroupsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsGetCmdGroupId, -1)

	return path
}

func buildQueryForGroupsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
