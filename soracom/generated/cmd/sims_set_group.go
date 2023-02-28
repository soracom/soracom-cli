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

// SimsSetGroupCmdGroupId holds value of 'groupId' option
var SimsSetGroupCmdGroupId string

// SimsSetGroupCmdSimId holds value of 'sim_id' option
var SimsSetGroupCmdSimId string

// SimsSetGroupCmdBody holds contents of request body to be sent
var SimsSetGroupCmdBody string

func init() {
	SimsSetGroupCmd.Flags().StringVar(&SimsSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	SimsSetGroupCmd.Flags().StringVar(&SimsSetGroupCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsSetGroupCmd.Flags().StringVar(&SimsSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsSetGroupCmd)
}

// SimsSetGroupCmd defines 'set-group' subcommand
var SimsSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/sims/{sim_id}/set_group:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/set_group:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "setSimGroup"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectSimsSetGroupCmdParams(ac)
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
	},
}

func collectSimsSetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsSetGroupCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsSetGroupCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsSetGroupCmd("/sims/{sim_id}/set_group"),
		query:       buildQueryForSimsSetGroupCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsSetGroupCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSetGroupCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSetGroupCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsSetGroupCmd() (string, error) {
	var result map[string]interface{}

	if SimsSetGroupCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(SimsSetGroupCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SimsSetGroupCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsSetGroupCmdBody)
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

	if SimsSetGroupCmdGroupId != "" {
		result["groupId"] = SimsSetGroupCmdGroupId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
