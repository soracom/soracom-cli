// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDeleteTransitGatewayVpcAttachmentCmdCustomerTgwVpcAttachmentId holds value of 'customer_tgw_vpc_attachment_id' option
var VpgDeleteTransitGatewayVpcAttachmentCmdCustomerTgwVpcAttachmentId string

// VpgDeleteTransitGatewayVpcAttachmentCmdVpgId holds value of 'vpg_id' option
var VpgDeleteTransitGatewayVpcAttachmentCmdVpgId string

func InitVpgDeleteTransitGatewayVpcAttachmentCmd() {
	VpgDeleteTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgDeleteTransitGatewayVpcAttachmentCmdCustomerTgwVpcAttachmentId, "customer-tgw-vpc-attachment-id", "", TRAPI("Transit gateway VPC attachment ID"))

	VpgDeleteTransitGatewayVpcAttachmentCmd.Flags().StringVar(&VpgDeleteTransitGatewayVpcAttachmentCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgDeleteTransitGatewayVpcAttachmentCmd.RunE = VpgDeleteTransitGatewayVpcAttachmentCmdRunE

	VpgCmd.AddCommand(VpgDeleteTransitGatewayVpcAttachmentCmd)
}

// VpgDeleteTransitGatewayVpcAttachmentCmd defines 'delete-transit-gateway-vpc-attachment' subcommand
var VpgDeleteTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "delete-transit-gateway-vpc-attachment",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments/{customer_tgw_vpc_attachment_id}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments/{customer_tgw_vpc_attachment_id}:delete:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "deleteTransitGatewayVpcAttachment"),
}

func VpgDeleteTransitGatewayVpcAttachmentCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgDeleteTransitGatewayVpcAttachmentCmdParams(ac)
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

func collectVpgDeleteTransitGatewayVpcAttachmentCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("customer_tgw_vpc_attachment_id", "customer-tgw-vpc-attachment-id", "path", parsedBody, VpgDeleteTransitGatewayVpcAttachmentCmdCustomerTgwVpcAttachmentId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgDeleteTransitGatewayVpcAttachmentCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteTransitGatewayVpcAttachmentCmd("/virtual_private_gateways/{vpg_id}/transit_gateway_vpc_attachments/{customer_tgw_vpc_attachment_id}"),
		query:  buildQueryForVpgDeleteTransitGatewayVpcAttachmentCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDeleteTransitGatewayVpcAttachmentCmd(path string) string {

	escapedCustomerTgwVpcAttachmentId := url.PathEscape(VpgDeleteTransitGatewayVpcAttachmentCmdCustomerTgwVpcAttachmentId)

	path = strReplace(path, "{"+"customer_tgw_vpc_attachment_id"+"}", escapedCustomerTgwVpcAttachmentId, -1)

	escapedVpgId := url.PathEscape(VpgDeleteTransitGatewayVpcAttachmentCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDeleteTransitGatewayVpcAttachmentCmd() url.Values {
	result := url.Values{}

	return result
}
