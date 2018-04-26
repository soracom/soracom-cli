package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgUnsetRedirectionCmdVpgId holds value of 'vpg_id' option
var VpgUnsetRedirectionCmdVpgId string

func init() {
	VpgUnsetRedirectionCmd.Flags().StringVar(&VpgUnsetRedirectionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgCmd.AddCommand(VpgUnsetRedirectionCmd)
}

// VpgUnsetRedirectionCmd defines 'unset-redirection' subcommand
var VpgUnsetRedirectionCmd = &cobra.Command{
	Use:   "unset-redirection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/junction/unset_redirection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/junction/unset_redirection:post:description`),
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

		param, err := collectVpgUnsetRedirectionCmdParams(ac)
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

func collectVpgUnsetRedirectionCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgUnsetRedirectionCmd("/virtual_private_gateways/{vpg_id}/junction/unset_redirection"),
		query:  buildQueryForVpgUnsetRedirectionCmd(),
	}, nil
}

func buildPathForVpgUnsetRedirectionCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgUnsetRedirectionCmdVpgId, -1)

	return path
}

func buildQueryForVpgUnsetRedirectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
