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

// LoraNetworkSetsAddPermissionCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsAddPermissionCmdNsId string

// LoraNetworkSetsAddPermissionCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsAddPermissionCmdOperatorId string

// LoraNetworkSetsAddPermissionCmdBody holds contents of request body to be sent
var LoraNetworkSetsAddPermissionCmdBody string

func init() {
	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsAddPermissionCmd.Flags().StringVar(&LoraNetworkSetsAddPermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsAddPermissionCmd)
}

// LoraNetworkSetsAddPermissionCmd defines 'add-permission' subcommand
var LoraNetworkSetsAddPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: TRAPI("/lora_network_sets/{ns_id}/add_permission:post:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/add_permission:post:description`) + "\n\n" + createLinkToAPIReference("LoraNetworkSet", "addPermissionToLoraNetworkSet"),
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

		param, err := collectLoraNetworkSetsAddPermissionCmdParams(ac)
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

func collectLoraNetworkSetsAddPermissionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraNetworkSetsAddPermissionCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("ns_id", "ns-id", "path", parsedBody, LoraNetworkSetsAddPermissionCmdNsId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsAddPermissionCmd("/lora_network_sets/{ns_id}/add_permission"),
		query:       buildQueryForLoraNetworkSetsAddPermissionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsAddPermissionCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsAddPermissionCmdNsId)

	path = strReplace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsAddPermissionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraNetworkSetsAddPermissionCmd() (string, error) {
	var result map[string]interface{}

	if LoraNetworkSetsAddPermissionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraNetworkSetsAddPermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsAddPermissionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraNetworkSetsAddPermissionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsAddPermissionCmdBody)
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

	if LoraNetworkSetsAddPermissionCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsAddPermissionCmdOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
