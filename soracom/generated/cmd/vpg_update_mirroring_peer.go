package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgUpdateMirroringPeerCmdId holds value of 'id' option
var VpgUpdateMirroringPeerCmdId string

// VpgUpdateMirroringPeerCmdIpaddr holds value of 'ipaddr' option
var VpgUpdateMirroringPeerCmdIpaddr string

// VpgUpdateMirroringPeerCmdBody holds contents of request body to be sent
var VpgUpdateMirroringPeerCmdBody string

func init() {
	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdId, "id", "", TRAPI("VPG ID"))

	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdIpaddr, "ipaddr", "", TRAPI("Mirroring peer IP address"))

	VpgUpdateMirroringPeerCmd.Flags().StringVar(&VpgUpdateMirroringPeerCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgUpdateMirroringPeerCmd)
}

// VpgUpdateMirroringPeerCmd defines 'update-mirroring-peer' subcommand
var VpgUpdateMirroringPeerCmd = &cobra.Command{
	Use:   "update-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}:put:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}:put:description`),
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

		param, err := collectVpgUpdateMirroringPeerCmdParams()
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

func collectVpgUpdateMirroringPeerCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgUpdateMirroringPeerCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForVpgUpdateMirroringPeerCmd("/virtual_private_gateways/{id}/junction/mirroring/peers/{ipaddr}"),
		query:       buildQueryForVpgUpdateMirroringPeerCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgUpdateMirroringPeerCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgUpdateMirroringPeerCmdId, -1)

	path = strings.Replace(path, "{"+"ipaddr"+"}", VpgUpdateMirroringPeerCmdIpaddr, -1)

	return path
}

func buildQueryForVpgUpdateMirroringPeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgUpdateMirroringPeerCmd() (string, error) {
	if VpgUpdateMirroringPeerCmdBody != "" {
		if strings.HasPrefix(VpgUpdateMirroringPeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgUpdateMirroringPeerCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgUpdateMirroringPeerCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgUpdateMirroringPeerCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
