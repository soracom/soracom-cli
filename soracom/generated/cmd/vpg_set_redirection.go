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

// VpgSetRedirectionCmdDescription holds value of 'description' option
var VpgSetRedirectionCmdDescription string

// VpgSetRedirectionCmdGateway holds value of 'gateway' option
var VpgSetRedirectionCmdGateway string

// VpgSetRedirectionCmdVpgId holds value of 'vpg_id' option
var VpgSetRedirectionCmdVpgId string

// VpgSetRedirectionCmdEnabled holds value of 'enabled' option
var VpgSetRedirectionCmdEnabled bool

// VpgSetRedirectionCmdBody holds contents of request body to be sent
var VpgSetRedirectionCmdBody string

func InitVpgSetRedirectionCmd() {
	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdDescription, "description", "", TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdGateway, "gateway", "", TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgSetRedirectionCmd.Flags().BoolVar(&VpgSetRedirectionCmdEnabled, "enabled", false, TRAPI(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgSetRedirectionCmd.RunE = VpgSetRedirectionCmdRunE

	VpgCmd.AddCommand(VpgSetRedirectionCmd)
}

// VpgSetRedirectionCmd defines 'set-redirection' subcommand
var VpgSetRedirectionCmd = &cobra.Command{
	Use:   "set-redirection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/set_redirection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/set_redirection:post:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "setRedirectionConfiguration"),
}

func VpgSetRedirectionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgSetRedirectionCmdParams(ac)
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

func collectVpgSetRedirectionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgSetRedirectionCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgSetRedirectionCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetRedirectionCmd("/virtual_private_gateways/{vpg_id}/junction/set_redirection"),
		query:       buildQueryForVpgSetRedirectionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgSetRedirectionCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgSetRedirectionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgSetRedirectionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgSetRedirectionCmd() (string, error) {
	var result map[string]interface{}

	if VpgSetRedirectionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgSetRedirectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetRedirectionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgSetRedirectionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgSetRedirectionCmdBody)
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

	if VpgSetRedirectionCmdDescription != "" {
		result["description"] = VpgSetRedirectionCmdDescription
	}

	if VpgSetRedirectionCmdGateway != "" {
		result["gateway"] = VpgSetRedirectionCmdGateway
	}

	if VpgSetRedirectionCmd.Flags().Lookup("enabled").Changed {
		result["enabled"] = VpgSetRedirectionCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
