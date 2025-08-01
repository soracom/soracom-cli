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

// VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId holds value of 'customerAwsAccountId' option
var VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId string

// VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId holds value of 'customerVpcId' option
var VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId string

// VpgCreateTransitGatewayVpcAttachmentCmdName holds value of 'name' option
var VpgCreateTransitGatewayVpcAttachmentCmdName string

// VpgCreateTransitGatewayVpcAttachmentCmdVpgId holds value of 'vpg_id' option
var VpgCreateTransitGatewayVpcAttachmentCmdVpgId string

// VpgCreateTransitGatewayVpcAttachmentCmdBody holds contents of request body to be sent
var VpgCreateTransitGatewayVpcAttachmentCmdBody string

func InitVpgCreateTransitGatewayVpcAttachmentCmd() {
	VpgCreateTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId, "customer-aws-account-id", "", TRAPI("AWS account ID of the VPC to be attached"))

	VpgCreateTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId, "customer-vpc-id", "", TRAPI("AWS VPC ID of the VPC to be attached"))

	VpgCreateTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgCreateTransitGatewayVpcAttachmentCmdName, "name", "", TRAPI("A name used to identify the VPC attachment"))

	VpgCreateTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgCreateTransitGatewayVpcAttachmentCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgCreateTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgCreateTransitGatewayVpcAttachmentCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCreateTransitGatewayVpcAttachmentCmd.RunE = VpgCreateTransitGatewayVpcAttachmentCmdRunE

	VpgCmd.AddCommand(VpgCreateTransitGatewayVpcAttachmentCmd)
}

// VpgCreateTransitGatewayVpcAttachmentCmd defines 'create-transit-gateway-vpc-attachment' subcommand
var VpgCreateTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "create-transit-gateway-vpc-attachment",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "createTransitGatewayVpcAttachment"),
}

func VpgCreateTransitGatewayVpcAttachmentCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectVpgCreateTransitGatewayVpcAttachmentCmdParams(ac)
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

func collectVpgCreateTransitGatewayVpcAttachmentCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgCreateTransitGatewayVpcAttachmentCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("customerAwsAccountId", "customer-aws-account-id", "body", parsedBody, VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("customerVpcId", "customer-vpc-id", "body", parsedBody, VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgCreateTransitGatewayVpcAttachmentCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateTransitGatewayVpcAttachmentCmd("/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments"),
		query:       buildQueryForVpgCreateTransitGatewayVpcAttachmentCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgCreateTransitGatewayVpcAttachmentCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgCreateTransitGatewayVpcAttachmentCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgCreateTransitGatewayVpcAttachmentCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgCreateTransitGatewayVpcAttachmentCmd() (string, error) {
	var result map[string]interface{}

	if VpgCreateTransitGatewayVpcAttachmentCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgCreateTransitGatewayVpcAttachmentCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateTransitGatewayVpcAttachmentCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgCreateTransitGatewayVpcAttachmentCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgCreateTransitGatewayVpcAttachmentCmdBody)
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

	if VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId != "" {
		result["customerAwsAccountId"] = VpgCreateTransitGatewayVpcAttachmentCmdCustomerAwsAccountId
	}

	if VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId != "" {
		result["customerVpcId"] = VpgCreateTransitGatewayVpcAttachmentCmdCustomerVpcId
	}

	if VpgCreateTransitGatewayVpcAttachmentCmdName != "" {
		result["name"] = VpgCreateTransitGatewayVpcAttachmentCmdName
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
