package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgPutIpAddressMapEntryCmdIpAddress holds value of 'ipAddress' option
var VpgPutIpAddressMapEntryCmdIpAddress string

// VpgPutIpAddressMapEntryCmdKey holds value of 'key' option
var VpgPutIpAddressMapEntryCmdKey string

// VpgPutIpAddressMapEntryCmdVpgId holds value of 'vpg_id' option
var VpgPutIpAddressMapEntryCmdVpgId string

// VpgPutIpAddressMapEntryCmdBody holds contents of request body to be sent
var VpgPutIpAddressMapEntryCmdBody string

func init() {
	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdIpAddress, "ip-address", "", TRAPI(""))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdKey, "key", "", TRAPI(""))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgPutIpAddressMapEntryCmd)
}

// VpgPutIpAddressMapEntryCmd defines 'put-ip-address-map-entry' subcommand
var VpgPutIpAddressMapEntryCmd = &cobra.Command{
	Use:   "put-ip-address-map-entry",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/ip_address_map:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/ip_address_map:post:description`),
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

		param, err := collectVpgPutIpAddressMapEntryCmdParams()
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

func collectVpgPutIpAddressMapEntryCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgPutIpAddressMapEntryCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgPutIpAddressMapEntryCmd("/virtual_private_gateways/{vpg_id}/ip_address_map"),
		query:       buildQueryForVpgPutIpAddressMapEntryCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgPutIpAddressMapEntryCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgPutIpAddressMapEntryCmdVpgId, -1)

	return path
}

func buildQueryForVpgPutIpAddressMapEntryCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgPutIpAddressMapEntryCmd() (string, error) {
	if VpgPutIpAddressMapEntryCmdBody != "" {
		if strings.HasPrefix(VpgPutIpAddressMapEntryCmdBody, "@") {
			fname := strings.TrimPrefix(VpgPutIpAddressMapEntryCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgPutIpAddressMapEntryCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgPutIpAddressMapEntryCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgPutIpAddressMapEntryCmdIpAddress != "" {
		result["ipAddress"] = VpgPutIpAddressMapEntryCmdIpAddress
	}

	if VpgPutIpAddressMapEntryCmdKey != "" {
		result["key"] = VpgPutIpAddressMapEntryCmdKey
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
