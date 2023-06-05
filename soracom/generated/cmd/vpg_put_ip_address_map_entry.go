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

// VpgPutIpAddressMapEntryCmdIpAddress holds value of 'ipAddress' option
var VpgPutIpAddressMapEntryCmdIpAddress string

// VpgPutIpAddressMapEntryCmdKey holds value of 'key' option
var VpgPutIpAddressMapEntryCmdKey string

// VpgPutIpAddressMapEntryCmdVpgId holds value of 'vpg_id' option
var VpgPutIpAddressMapEntryCmdVpgId string

// VpgPutIpAddressMapEntryCmdBody holds contents of request body to be sent
var VpgPutIpAddressMapEntryCmdBody string

func InitVpgPutIpAddressMapEntryCmd() {
	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdIpAddress, "ip-address", "", TRAPI(""))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdKey, "key", "", TRAPI(""))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgPutIpAddressMapEntryCmd.Flags().StringVar(&VpgPutIpAddressMapEntryCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgPutIpAddressMapEntryCmd.RunE = VpgPutIpAddressMapEntryCmdRunE

	VpgCmd.AddCommand(VpgPutIpAddressMapEntryCmd)
}

// VpgPutIpAddressMapEntryCmd defines 'put-ip-address-map-entry' subcommand
var VpgPutIpAddressMapEntryCmd = &cobra.Command{
	Use:   "put-ip-address-map-entry",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/ip_address_map:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/ip_address_map:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "putVirtualPrivateGatewayIpAddressMapEntry"),
}

func VpgPutIpAddressMapEntryCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgPutIpAddressMapEntryCmdParams(ac)
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

func collectVpgPutIpAddressMapEntryCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgPutIpAddressMapEntryCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("ipAddress", "ip-address", "body", parsedBody, VpgPutIpAddressMapEntryCmdIpAddress)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("key", "key", "body", parsedBody, VpgPutIpAddressMapEntryCmdKey)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgPutIpAddressMapEntryCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgPutIpAddressMapEntryCmd("/virtual_private_gateways/{vpg_id}/ip_address_map"),
		query:       buildQueryForVpgPutIpAddressMapEntryCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgPutIpAddressMapEntryCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgPutIpAddressMapEntryCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgPutIpAddressMapEntryCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgPutIpAddressMapEntryCmd() (string, error) {
	var result map[string]interface{}

	if VpgPutIpAddressMapEntryCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgPutIpAddressMapEntryCmdBody, "@") {
			fname := strings.TrimPrefix(VpgPutIpAddressMapEntryCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgPutIpAddressMapEntryCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgPutIpAddressMapEntryCmdBody)
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
