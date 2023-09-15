// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// GroupsCreateCmdBody holds contents of request body to be sent
var GroupsCreateCmdBody string

func InitGroupsCreateCmd() {
	GroupsCreateCmd.Flags().StringVar(&GroupsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	GroupsCreateCmd.RunE = GroupsCreateCmdRunE

	GroupsCmd.AddCommand(GroupsCreateCmd)
}

// GroupsCreateCmd defines 'create' subcommand
var GroupsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/groups:post:summary"),
	Long:  TRAPI(`/groups:post:description`) + "\n\n" + createLinkToAPIReference("Group", "createGroup"),
}

func GroupsCreateCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
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

	param, err := collectGroupsCreateCmdParams(ac)
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
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectGroupsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForGroupsCreateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGroupsCreateCmd("/groups"),
		query:       buildQueryForGroupsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsCreateCmd(path string) string {

	return path
}

func buildQueryForGroupsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForGroupsCreateCmd() (string, error) {
	var result map[string]interface{}

	if GroupsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(GroupsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(GroupsCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if GroupsCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(GroupsCreateCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
