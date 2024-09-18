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

// VpgUpdateTagsCmdVpgId holds value of 'vpg_id' option
var VpgUpdateTagsCmdVpgId string

// VpgUpdateTagsCmdBody holds contents of request body to be sent
var VpgUpdateTagsCmdBody string

func InitVpgUpdateTagsCmd() {
	VpgUpdateTagsCmd.Flags().StringVar(&VpgUpdateTagsCmdVpgId, "vpg-id", "", TRAPI(""))

	VpgUpdateTagsCmd.Flags().StringVar(&VpgUpdateTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgUpdateTagsCmd.RunE = VpgUpdateTagsCmdRunE

	VpgCmd.AddCommand(VpgUpdateTagsCmd)
}

// VpgUpdateTagsCmd defines 'update-tags' subcommand
var VpgUpdateTagsCmd = &cobra.Command{
	Use:   "update-tags",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/tags:put:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/tags:put:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "updateVirtualPrivateGatewayTags"),
}

func VpgUpdateTagsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgUpdateTagsCmdParams(ac)
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

func collectVpgUpdateTagsCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgUpdateTagsCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgUpdateTagsCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForVpgUpdateTagsCmd("/virtual_private_gateways/{vpg_id}/tags"),
		query:       buildQueryForVpgUpdateTagsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgUpdateTagsCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgUpdateTagsCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgUpdateTagsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgUpdateTagsCmd() (string, error) {
	var b []byte
	var err error

	if VpgUpdateTagsCmdBody != "" {
		if strings.HasPrefix(VpgUpdateTagsCmdBody, "@") {
			fname := strings.TrimPrefix(VpgUpdateTagsCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgUpdateTagsCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgUpdateTagsCmdBody)
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