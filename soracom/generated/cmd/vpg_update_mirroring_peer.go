package cmd

import (
	"encoding/json"
	"io/ioutil"

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
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}:put:description`),
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

		param, err := collectVpgUpdateMirroringPeerCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectVpgUpdateMirroringPeerCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgUpdateMirroringPeerCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForVpgUpdateMirroringPeerCmd("/virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr}"),
		query:       buildQueryForVpgUpdateMirroringPeerCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgUpdateMirroringPeerCmd(path string) string {

	path = strings.Replace(path, "{"+"ipaddr"+"}", VpgUpdateMirroringPeerCmdIpaddr, -1)

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgUpdateMirroringPeerCmdVpgId, -1)

	return path
}

func buildQueryForVpgUpdateMirroringPeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgUpdateMirroringPeerCmd() (string, error) {
	var result map[string]interface{}

	if VpgUpdateMirroringPeerCmdBody != "" {
		var b []byte
		var err error

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
