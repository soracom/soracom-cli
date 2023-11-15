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

// VpgPutSimBasedRoutingRoutesCmdVpgId holds value of 'vpg_id' option
var VpgPutSimBasedRoutingRoutesCmdVpgId string

// VpgPutSimBasedRoutingRoutesCmdBody holds contents of request body to be sent
var VpgPutSimBasedRoutingRoutesCmdBody string

func InitVpgPutSimBasedRoutingRoutesCmd() {
	VpgPutSimBasedRoutingRoutesCmd.Flags().StringVar(&VpgPutSimBasedRoutingRoutesCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgPutSimBasedRoutingRoutesCmd.Flags().StringVar(&VpgPutSimBasedRoutingRoutesCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgPutSimBasedRoutingRoutesCmd.RunE = VpgPutSimBasedRoutingRoutesCmdRunE

	VpgCmd.AddCommand(VpgPutSimBasedRoutingRoutesCmd)
}

// VpgPutSimBasedRoutingRoutesCmd defines 'put-sim-based-routing-routes' subcommand
var VpgPutSimBasedRoutingRoutesCmd = &cobra.Command{
	Use:   "put-sim-based-routing-routes",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/routes:put:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/routes:put:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "putSimBasedRoutingRoutes"),
}

func VpgPutSimBasedRoutingRoutesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgPutSimBasedRoutingRoutesCmdParams(ac)
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

func collectVpgPutSimBasedRoutingRoutesCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForVpgPutSimBasedRoutingRoutesCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgPutSimBasedRoutingRoutesCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForVpgPutSimBasedRoutingRoutesCmd("/virtual_private_gateways/{vpg_id}/gate/routing/static/sims/routes"),
		query:       buildQueryForVpgPutSimBasedRoutingRoutesCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgPutSimBasedRoutingRoutesCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgPutSimBasedRoutingRoutesCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgPutSimBasedRoutingRoutesCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVpgPutSimBasedRoutingRoutesCmd() (string, error) {
	var result map[string]interface{}

	if VpgPutSimBasedRoutingRoutesCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgPutSimBasedRoutingRoutesCmdBody, "@") {
			fname := strings.TrimPrefix(VpgPutSimBasedRoutingRoutesCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if VpgPutSimBasedRoutingRoutesCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgPutSimBasedRoutingRoutesCmdBody)
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