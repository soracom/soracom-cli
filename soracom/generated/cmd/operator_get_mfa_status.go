package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorGetMfaStatusCmdOperatorId holds value of 'operator_id' option
var OperatorGetMfaStatusCmdOperatorId string

func init() {
	OperatorGetMfaStatusCmd.Flags().StringVar(&OperatorGetMfaStatusCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorCmd.AddCommand(OperatorGetMfaStatusCmd)
}

// OperatorGetMfaStatusCmd defines 'get-mfa-status' subcommand
var OperatorGetMfaStatusCmd = &cobra.Command{
	Use:   "get-mfa-status",
	Short: TRAPI("/operators/{operator_id}/mfa:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/mfa:get:description`),
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

		param, err := collectOperatorGetMfaStatusCmdParams(ac)
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

func collectOperatorGetMfaStatusCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorGetMfaStatusCmdOperatorId == "" {
		OperatorGetMfaStatusCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorGetMfaStatusCmd("/operators/{operator_id}/mfa"),
		query:  buildQueryForOperatorGetMfaStatusCmd(),
	}, nil
}

func buildPathForOperatorGetMfaStatusCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorGetMfaStatusCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorGetMfaStatusCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
