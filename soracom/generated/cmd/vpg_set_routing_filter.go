// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// VpgSetRoutingFilterCmdVpgId holds value of 'vpg_id' option
var VpgSetRoutingFilterCmdVpgId string

// VpgSetRoutingFilterCmdBody holds contents of request body to be sent
var VpgSetRoutingFilterCmdBody string

func init() {
	VpgSetRoutingFilterCmd.Flags().StringVar(&VpgSetRoutingFilterCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgSetRoutingFilterCmd.Flags().StringVar(&VpgSetRoutingFilterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VpgCmd.AddCommand(VpgSetRoutingFilterCmd)
}

// VpgSetRoutingFilterCmd defines 'set-routing-filter' subcommand
var VpgSetRoutingFilterCmd = &cobra.Command{
	Use:   "set-routing-filter",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/set_routing_filter:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/set_routing_filter:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "setRoutingFilter"),
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

		param, err := collectVpgSetRoutingFilterCmdParams(ac)
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

func collectVpgSetRoutingFilterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgSetRoutingFilterCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgSetRoutingFilterCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetRoutingFilterCmd("/virtual_private_gateways/{vpg_id}/set_routing_filter"),
		query:       buildQueryForVpgSetRoutingFilterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgSetRoutingFilterCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgSetRoutingFilterCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgSetRoutingFilterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgSetRoutingFilterCmd() (string, error) {
	var b []byte
	var err error

	if VpgSetRoutingFilterCmdBody != "" {
		if strings.HasPrefix(VpgSetRoutingFilterCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetRoutingFilterCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgSetRoutingFilterCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgSetRoutingFilterCmdBody)
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
