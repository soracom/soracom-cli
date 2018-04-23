package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgUnsetInspectionCmdId holds value of 'id' option
var VpgUnsetInspectionCmdId string

func init() {
	VpgUnsetInspectionCmd.Flags().StringVar(&VpgUnsetInspectionCmdId, "id", "", TRAPI("VPG ID"))

	VpgCmd.AddCommand(VpgUnsetInspectionCmd)
}

// VpgUnsetInspectionCmd defines 'unset-inspection' subcommand
var VpgUnsetInspectionCmd = &cobra.Command{
	Use:   "unset-inspection",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/unset_inspection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/unset_inspection:post:description`),
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

		param, err := collectVpgUnsetInspectionCmdParams(ac)
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

func collectVpgUnsetInspectionCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgUnsetInspectionCmd("/virtual_private_gateways/{id}/junction/unset_inspection"),
		query:  buildQueryForVpgUnsetInspectionCmd(),
	}, nil
}

func buildPathForVpgUnsetInspectionCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgUnsetInspectionCmdId, -1)

	return path
}

func buildQueryForVpgUnsetInspectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
