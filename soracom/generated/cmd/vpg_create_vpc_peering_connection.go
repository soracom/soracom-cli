package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock holds value of 'destinationCidrBlock' option
var VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock string

// VpgCreateVpcPeeringConnectionCmdPeerOwnerId holds value of 'peerOwnerId' option
var VpgCreateVpcPeeringConnectionCmdPeerOwnerId string

// VpgCreateVpcPeeringConnectionCmdPeerVpcId holds value of 'peerVpcId' option
var VpgCreateVpcPeeringConnectionCmdPeerVpcId string

// VpgCreateVpcPeeringConnectionCmdVpgId holds value of 'vpg_id' option
var VpgCreateVpcPeeringConnectionCmdVpgId string

// VpgCreateVpcPeeringConnectionCmdBody holds contents of request body to be sent
var VpgCreateVpcPeeringConnectionCmdBody string

func init() {
	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock, "destination-cidr-block", "", TR(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerOwnerId, "peer-owner-id", "", TR(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerVpcId, "peer-vpc-id", "", TR(""))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.create_vpc_peering_connection.post.parameters.vpg_id.description"))

	VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgCreateVpcPeeringConnectionCmd)
}

// VpgCreateVpcPeeringConnectionCmd defines 'create-vpc-peering-connection' subcommand
var VpgCreateVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "create-vpc-peering-connection",
	Short: TR("virtual_private_gateway.create_vpc_peering_connection.post.summary"),
	Long:  TR(`virtual_private_gateway.create_vpc_peering_connection.post.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		param, err := collectVpgCreateVpcPeeringConnectionCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectVpgCreateVpcPeeringConnectionCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgCreateVpcPeeringConnectionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateVpcPeeringConnectionCmd("/virtual_private_gateways/{vpg_id}/vpc_peering_connections"),
		query:       buildQueryForVpgCreateVpcPeeringConnectionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgCreateVpcPeeringConnectionCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgCreateVpcPeeringConnectionCmdVpgId, -1)

	return path
}

func buildQueryForVpgCreateVpcPeeringConnectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgCreateVpcPeeringConnectionCmd() (string, error) {
	if VpgCreateVpcPeeringConnectionCmdBody != "" {
		if strings.HasPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgCreateVpcPeeringConnectionCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgCreateVpcPeeringConnectionCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock != "" {
		result["destinationCidrBlock"] = VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock
	}

	if VpgCreateVpcPeeringConnectionCmdPeerOwnerId != "" {
		result["peerOwnerId"] = VpgCreateVpcPeeringConnectionCmdPeerOwnerId
	}

	if VpgCreateVpcPeeringConnectionCmdPeerVpcId != "" {
		result["peerVpcId"] = VpgCreateVpcPeeringConnectionCmdPeerVpcId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
