package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsRevokePermissionCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsRevokePermissionCmdNsId string

// LoraNetworkSetsRevokePermissionCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsRevokePermissionCmdOperatorId string

// LoraNetworkSetsRevokePermissionCmdBody holds contents of request body to be sent
var LoraNetworkSetsRevokePermissionCmdBody string

func init() {
	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsRevokePermissionCmd.Flags().StringVar(&LoraNetworkSetsRevokePermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsRevokePermissionCmd)
}

// LoraNetworkSetsRevokePermissionCmd defines 'revoke-permission' subcommand
var LoraNetworkSetsRevokePermissionCmd = &cobra.Command{
	Use:   "revoke-permission",
	Short: TRAPI("/lora_network_sets/{ns_id}/revoke_permission:post:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/revoke_permission:post:description`),
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

		param, err := collectLoraNetworkSetsRevokePermissionCmdParams(ac)
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

func collectLoraNetworkSetsRevokePermissionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLoraNetworkSetsRevokePermissionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsRevokePermissionCmd("/lora_network_sets/{ns_id}/revoke_permission"),
		query:       buildQueryForLoraNetworkSetsRevokePermissionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraNetworkSetsRevokePermissionCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsRevokePermissionCmdNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsRevokePermissionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraNetworkSetsRevokePermissionCmd() (string, error) {
	if LoraNetworkSetsRevokePermissionCmdBody != "" {
		if strings.HasPrefix(LoraNetworkSetsRevokePermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsRevokePermissionCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraNetworkSetsRevokePermissionCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraNetworkSetsRevokePermissionCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LoraNetworkSetsRevokePermissionCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsRevokePermissionCmdOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
