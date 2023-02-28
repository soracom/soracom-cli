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

// VpgCreateCmdDeviceSubnetCidrRange holds value of 'deviceSubnetCidrRange' option
var VpgCreateCmdDeviceSubnetCidrRange string

// VpgCreateCmdType holds value of 'type' option
var VpgCreateCmdType int64

// VpgCreateCmdUseInternetGateway holds value of 'useInternetGateway' option
var VpgCreateCmdUseInternetGateway bool

// VpgCreateCmdBody holds contents of request body to be sent
var VpgCreateCmdBody string

func init() {
	VpgCreateCmd.Flags().StringVar(&VpgCreateCmdDeviceSubnetCidrRange, "device-subnet-cidr-range", "10.128.0.0/9", TRAPI(""))

	VpgCreateCmd.Flags().Int64Var(&VpgCreateCmdType, "type", 0, TRAPI("VPG Type.- `14` : Type-E- `15` : Type-F"))

	VpgCreateCmd.Flags().BoolVar(&VpgCreateCmdUseInternetGateway, "use-internet-gateway", true, TRAPI(""))

	VpgCreateCmd.Flags().StringVar(&VpgCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VpgCmd.AddCommand(VpgCreateCmd)
}

// VpgCreateCmd defines 'create' subcommand
var VpgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/virtual_private_gateways:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "createVirtualPrivateGateway"),
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

		param, err := collectVpgCreateCmdParams(ac)
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

func collectVpgCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreateCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("type", "type", "body", parsedBody, VpgCreateCmdType)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateCmd("/virtual_private_gateways"),
		query:       buildQueryForVpgCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreateCmd(path string) string {

	return path
}

func buildQueryForVpgCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreateCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreateCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreateCmdBody)
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

	if VpgCreateCmdDeviceSubnetCidrRange != "10.128.0.0/9" {
		result["deviceSubnetCidrRange"] = VpgCreateCmdDeviceSubnetCidrRange
	}

	result["type"] = VpgCreateCmdType

	if VpgCreateCmdUseInternetGateway != true {

		result["useInternetGateway"] = VpgCreateCmdUseInternetGateway

	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
