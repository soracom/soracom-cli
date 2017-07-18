package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsPutTagsCmdGroupId holds value of 'group_id' option
var GroupsPutTagsCmdGroupId string

// GroupsPutTagsCmdBody holds contents of request body to be sent
var GroupsPutTagsCmdBody string

func init() {
	GroupsPutTagsCmd.Flags().StringVar(&GroupsPutTagsCmdGroupId, "group-id", "", TRAPI("Target group ID."))

	GroupsPutTagsCmd.Flags().StringVar(&GroupsPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	GroupsCmd.AddCommand(GroupsPutTagsCmd)
}

// GroupsPutTagsCmd defines 'put-tags' subcommand
var GroupsPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/groups/{group_id}/tags:put:summary"),
	Long:  TRAPI(`/groups/{group_id}/tags:put:description`),
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

		param, err := collectGroupsPutTagsCmdParams()
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

func collectGroupsPutTagsCmdParams() (*apiParams, error) {

	body, err := buildBodyForGroupsPutTagsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForGroupsPutTagsCmd("/groups/{group_id}/tags"),
		query:       buildQueryForGroupsPutTagsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForGroupsPutTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"group_id"+"}", GroupsPutTagsCmdGroupId, -1)

	return path
}

func buildQueryForGroupsPutTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForGroupsPutTagsCmd() (string, error) {
	if GroupsPutTagsCmdBody != "" {
		if strings.HasPrefix(GroupsPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(GroupsPutTagsCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if GroupsPutTagsCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return GroupsPutTagsCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
