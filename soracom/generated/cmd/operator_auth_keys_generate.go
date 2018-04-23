package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorAuthKeysGenerateCmdOperatorId holds value of 'operator_id' option
var OperatorAuthKeysGenerateCmdOperatorId string

func init() {
	OperatorAuthKeysGenerateCmd.Flags().StringVar(&OperatorAuthKeysGenerateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysGenerateCmd)
}

// OperatorAuthKeysGenerateCmd defines 'generate' subcommand
var OperatorAuthKeysGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: TRAPI("/operators/{operator_id}/auth_keys:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/auth_keys:post:description`),
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

		param, err := collectOperatorAuthKeysGenerateCmdParams(ac)
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

func collectOperatorAuthKeysGenerateCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorAuthKeysGenerateCmdOperatorId == "" {
		OperatorAuthKeysGenerateCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorAuthKeysGenerateCmd("/operators/{operator_id}/auth_keys"),
		query:  buildQueryForOperatorAuthKeysGenerateCmd(),
	}, nil
}

func buildPathForOperatorAuthKeysGenerateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorAuthKeysGenerateCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorAuthKeysGenerateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
