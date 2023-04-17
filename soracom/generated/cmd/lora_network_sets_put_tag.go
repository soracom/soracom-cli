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

// LoraNetworkSetsPutTagCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsPutTagCmdNsId string

// LoraNetworkSetsPutTagCmdBody holds contents of request body to be sent
var LoraNetworkSetsPutTagCmdBody string

func init() {
	LoraNetworkSetsPutTagCmd.Flags().StringVar(&LoraNetworkSetsPutTagCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsPutTagCmd.Flags().StringVar(&LoraNetworkSetsPutTagCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsPutTagCmd)
}

// LoraNetworkSetsPutTagCmd defines 'put-tag' subcommand
var LoraNetworkSetsPutTagCmd = &cobra.Command{
	Use:   "put-tag",
	Short: TRAPI("/lora_network_sets/{ns_id}/tags:put:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/tags:put:description`) + "\n\n" + createLinkToAPIReference("LoraNetworkSet", "putLoraNetworkSetTags"),
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

		param, err := collectLoraNetworkSetsPutTagCmdParams(ac)
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

func collectLoraNetworkSetsPutTagCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraNetworkSetsPutTagCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("ns_id", "ns-id", "path", parsedBody, LoraNetworkSetsPutTagCmdNsId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLoraNetworkSetsPutTagCmd("/lora_network_sets/{ns_id}/tags"),
		query:       buildQueryForLoraNetworkSetsPutTagCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsPutTagCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsPutTagCmdNsId)

	path = strReplace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsPutTagCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraNetworkSetsPutTagCmd() (string, error) {
	var b []byte
	var err error

	if LoraNetworkSetsPutTagCmdBody != "" {
		if strings.HasPrefix(LoraNetworkSetsPutTagCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsPutTagCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraNetworkSetsPutTagCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsPutTagCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
