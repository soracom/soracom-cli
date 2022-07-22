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

// VpgUpdateMirroringPeerCmdIpaddr holds value of 'ipaddr' option
var VpgUpdateMirroringPeerCmdIpaddr string

// VpgUpdateMirroringPeerCmdVpgId holds value of 'vpg_id' option
var VpgUpdateMirroringPeerCmdVpgId string

// VpgUpdateMirroringPeerCmdBody holds contents of request body to be sent
var VpgUpdateMirroringPeerCmdBody string

func init() {
	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdIpaddr, "ipaddr", "", TRAPI("Mirroring peer IP address"))

	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VpgCmd.AddCommand(VpgUpdateMirroringPeerCmd)
}

// VpgUpdateMirroringPeerCmd defines 'update-mirroring-peer' subcommand
var VpgUpdateMirroringPeerCmd = &cobra.Command{
	Use:   "update-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:put:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:put:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "updateMirroringPeer"),
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

		param, err := collectVpgUpdateMirroringPeerCmdParams(ac)
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

func collectVpgUpdateMirroringPeerCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgUpdateMirroringPeerCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("ipaddr", "ipaddr", "path", parsedBody, VpgUpdateMirroringPeerCmdIpaddr)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgUpdateMirroringPeerCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForVpgUpdateMirroringPeerCmd("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}"),
		query:       buildQueryForVpgUpdateMirroringPeerCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgUpdateMirroringPeerCmd(path string) string {

	escapedIpaddr := url.PathEscape(VpgUpdateMirroringPeerCmdIpaddr)

	path = strReplace(path, "{"+"ipaddr"+"}", escapedIpaddr, -1)

	escapedVpgId := url.PathEscape(VpgUpdateMirroringPeerCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgUpdateMirroringPeerCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgUpdateMirroringPeerCmd() (string, error) {
	var b []byte
	var err error

	if VpgUpdateMirroringPeerCmdBody != "" {
		if strings.HasPrefix(VpgUpdateMirroringPeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgUpdateMirroringPeerCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgUpdateMirroringPeerCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgUpdateMirroringPeerCmdBody)
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
