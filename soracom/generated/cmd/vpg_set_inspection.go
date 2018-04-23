package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgSetInspectionCmdId holds value of 'id' option
var VpgSetInspectionCmdId string

// VpgSetInspectionCmdEnabled holds value of 'enabled' option
var VpgSetInspectionCmdEnabled bool

// VpgSetInspectionCmdBody holds contents of request body to be sent
var VpgSetInspectionCmdBody string

func init() {
	VpgSetInspectionCmd.Flags().StringVar(&VpgSetInspectionCmdId, "id", "", TRAPI("VPG ID"))

	VpgSetInspectionCmd.Flags().BoolVar(&VpgSetInspectionCmdEnabled, "enabled", false, TRAPI(""))

	VpgSetInspectionCmd.Flags().StringVar(&VpgSetInspectionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgSetInspectionCmd)
}

// VpgSetInspectionCmd defines 'set-inspection' subcommand
var VpgSetInspectionCmd = &cobra.Command{
	Use:   "set-inspection",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/set_inspection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/set_inspection:post:description`),
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

		param, err := collectVpgSetInspectionCmdParams(ac)
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

func collectVpgSetInspectionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgSetInspectionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetInspectionCmd("/virtual_private_gateways/{id}/junction/set_inspection"),
		query:       buildQueryForVpgSetInspectionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgSetInspectionCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgSetInspectionCmdId, -1)

	return path
}

func buildQueryForVpgSetInspectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgSetInspectionCmd() (string, error) {
	if VpgSetInspectionCmdBody != "" {
		if strings.HasPrefix(VpgSetInspectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetInspectionCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgSetInspectionCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgSetInspectionCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgSetInspectionCmdEnabled != false {
		result["enabled"] = VpgSetInspectionCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
