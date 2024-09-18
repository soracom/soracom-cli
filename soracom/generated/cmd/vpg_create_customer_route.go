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

// VpgCreateCustomerRouteCmdDestinationCidr holds value of 'destinationCidr' option
var VpgCreateCustomerRouteCmdDestinationCidr string

// VpgCreateCustomerRouteCmdTarget holds value of 'target' option
var VpgCreateCustomerRouteCmdTarget string

// VpgCreateCustomerRouteCmdVpgId holds value of 'vpg_id' option
var VpgCreateCustomerRouteCmdVpgId string

// VpgCreateCustomerRouteCmdIgnoreDeviceSubnetCidrRangeOverlap holds value of 'ignoreDeviceSubnetCidrRangeOverlap' option
var VpgCreateCustomerRouteCmdIgnoreDeviceSubnetCidrRangeOverlap bool

// VpgCreateCustomerRouteCmdBody holds contents of request body to be sent
var VpgCreateCustomerRouteCmdBody string

func InitVpgCreateCustomerRouteCmd() {
	VpgCreateCustomerRouteCmd.Flags().StringVar(&VpgCreateCustomerRouteCmdDestinationCidr, "destination-cidr", "", TRAPI("Destination CIDR block"))

	VpgCreateCustomerRouteCmd.Flags().StringVar(&VpgCreateCustomerRouteCmdTarget, "target", "", TRAPI("Transit Gateway ID"))

	VpgCreateCustomerRouteCmd.Flags().StringVar(&VpgCreateCustomerRouteCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCreateCustomerRouteCmd.Flags().BoolVar(&VpgCreateCustomerRouteCmdIgnoreDeviceSubnetCidrRangeOverlap, "ignore-device-subnet-cidr-range-overlap", false, TRAPI(""))

	VpgCreateCustomerRouteCmd.Flags().StringVar(&VpgCreateCustomerRouteCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCreateCustomerRouteCmd.RunE = VpgCreateCustomerRouteCmdRunE

	VpgCmd.AddCommand(VpgCreateCustomerRouteCmd)
}

// VpgCreateCustomerRouteCmd defines 'create-customer-route' subcommand
var VpgCreateCustomerRouteCmd = &cobra.Command{
	Use:   "create-customer-route",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/customer_routes:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/customer_routes:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "createCustomerRoute"),
}

func VpgCreateCustomerRouteCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgCreateCustomerRouteCmdParams(ac)
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

func collectVpgCreateCustomerRouteCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreateCustomerRouteCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("destinationCidr", "destination-cidr", "body", parsedBody, VpgCreateCustomerRouteCmdDestinationCidr)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("target", "target", "body", parsedBody, VpgCreateCustomerRouteCmdTarget)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCreateCustomerRouteCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateCustomerRouteCmd("/virtual_private_gateways/{vpg_id}/customer_routes"),
		query:       buildQueryForVpgCreateCustomerRouteCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreateCustomerRouteCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCreateCustomerRouteCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCreateCustomerRouteCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreateCustomerRouteCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreateCustomerRouteCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreateCustomerRouteCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateCustomerRouteCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgCreateCustomerRouteCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreateCustomerRouteCmdBody)
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

	if VpgCreateCustomerRouteCmdDestinationCidr != "" {
		result["destinationCidr"] = VpgCreateCustomerRouteCmdDestinationCidr
	}

	if VpgCreateCustomerRouteCmdTarget != "" {
		result["target"] = VpgCreateCustomerRouteCmdTarget
	}

	if VpgCreateCustomerRouteCmd.Flags().Lookup("ignore-device-subnet-cidr-range-overlap").Changed {
		result["ignoreDeviceSubnetCidrRangeOverlap"] = VpgCreateCustomerRouteCmdIgnoreDeviceSubnetCidrRangeOverlap
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
