package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgCreateCmdDeviceSubnetCidrRange holds value of 'deviceSubnetCidrRange' option
var VpgCreateCmdDeviceSubnetCidrRange string

// VpgCreateCmdPrimaryServiceName holds value of 'primaryServiceName' option
var VpgCreateCmdPrimaryServiceName string

// VpgCreateCmdUseInternetGateway holds value of 'useInternetGateway' option
var VpgCreateCmdUseInternetGateway bool

// VpgCreateCmdBody holds contents of request body to be sent
var VpgCreateCmdBody string

func init() {
	VpgCreateCmd.Flags().StringVar(&VpgCreateCmdDeviceSubnetCidrRange, "device-subnet-cidr-range", "", TRAPI(""))

	VpgCreateCmd.Flags().StringVar(&VpgCreateCmdPrimaryServiceName, "primary-service-name", "", TRAPI(""))

	VpgCreateCmd.Flags().BoolVar(&VpgCreateCmdUseInternetGateway, "use-internet-gateway", false, TRAPI(""))

	VpgCreateCmd.Flags().StringVar(&VpgCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgCreateCmd)
}

// VpgCreateCmd defines 'create' subcommand
var VpgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/virtual_private_gateways:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways:post:description`),
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

		param, err := collectVpgCreateCmdParams()
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

func collectVpgCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateCmd("/virtual_private_gateways"),
		query:       buildQueryForVpgCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgCreateCmd(path string) string {

	return path
}

func buildQueryForVpgCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgCreateCmd() (string, error) {
	if VpgCreateCmdBody != "" {
		if strings.HasPrefix(VpgCreateCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgCreateCmdDeviceSubnetCidrRange != "" {
		result["deviceSubnetCidrRange"] = VpgCreateCmdDeviceSubnetCidrRange
	}

	if VpgCreateCmdPrimaryServiceName != "" {
		result["primaryServiceName"] = VpgCreateCmdPrimaryServiceName
	}

	if VpgCreateCmdUseInternetGateway != false {
		result["useInternetGateway"] = VpgCreateCmdUseInternetGateway
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
