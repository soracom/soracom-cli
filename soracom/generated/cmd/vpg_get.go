// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// VpgGetCmdVpgId holds value of 'vpg_id' option
var VpgGetCmdVpgId string

func init() {
	VpgGetCmd.Flags().StringVar(&VpgGetCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))
	VpgCmd.AddCommand(VpgGetCmd)
}

// VpgGetCmd defines 'get' subcommand
var VpgGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}:get:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectVpgGetCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectVpgGetCmdParams(ac *apiClient) (*apiParams, error) {
	if VpgGetCmdVpgId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "vpg-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForVpgGetCmd("/virtual_private_gateways/{vpg_id}"),
		query:  buildQueryForVpgGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgGetCmd(path string) string {

	escapedVpgId := url.PathEscape(VpgGetCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgGetCmd() url.Values {
	result := url.Values{}

	return result
}
