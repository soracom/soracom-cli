package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsDeleteConfigCmdGroupId holds value of 'group_id' option
var GroupsDeleteConfigCmdGroupId string

// GroupsDeleteConfigCmdName holds value of 'name' option
var GroupsDeleteConfigCmdName string

// GroupsDeleteConfigCmdNamespace holds value of 'namespace' option
var GroupsDeleteConfigCmdNamespace string

func init() {
	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdGroupId, "group-id", "", TR("groups.delete_configuration_parameter.delete.parameters.group_id.description"))

	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdName, "name", "", TR("groups.delete_configuration_parameter.delete.parameters.name.description"))

	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdNamespace, "namespace", "", TR("groups.delete_configuration_parameter.delete.parameters.namespace.description"))

	GroupsCmd.AddCommand(GroupsDeleteConfigCmd)
}

// GroupsDeleteConfigCmd defines 'delete-config' subcommand
var GroupsDeleteConfigCmd = &cobra.Command{
	Use:   "delete-config",
	Short: TR("groups.delete_configuration_parameter.delete.summary"),
	Long:  TR(`groups.delete_configuration_parameter.delete.description`),
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

		param, err := collectGroupsDeleteConfigCmdParams()
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

func collectGroupsDeleteConfigCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteConfigCmd("/groups/{group_id}/configuration/{namespace}/{name}"),
		query:  buildQueryForGroupsDeleteConfigCmd(),
	}, nil
}

func buildPathForGroupsDeleteConfigCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsDeleteConfigCmdGroupId, -1)

	path = strings.Replace(path, "{"+"name"+"}", GroupsDeleteConfigCmdName, -1)

	path = strings.Replace(path, "{"+"namespace"+"}", GroupsDeleteConfigCmdNamespace, -1)

	return path
}

func buildQueryForGroupsDeleteConfigCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
