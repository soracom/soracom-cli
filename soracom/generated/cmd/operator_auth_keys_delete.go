package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorAuthKeysDeleteCmdAuthKeyId holds value of 'auth_key_id' option
var OperatorAuthKeysDeleteCmdAuthKeyId string

// OperatorAuthKeysDeleteCmdOperatorId holds value of 'operator_id' option
var OperatorAuthKeysDeleteCmdOperatorId string

func init() {
	OperatorAuthKeysDeleteCmd.Flags().StringVar(&OperatorAuthKeysDeleteCmdAuthKeyId, "auth-key-id", "", TRAPI("auth_key_id"))

	OperatorAuthKeysDeleteCmd.Flags().StringVar(&OperatorAuthKeysDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysDeleteCmd)
}

// OperatorAuthKeysDeleteCmd defines 'delete' subcommand
var OperatorAuthKeysDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/auth_keys/{auth_key_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/auth_keys/{auth_key_id}:delete:description`),
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

		param, err := collectOperatorAuthKeysDeleteCmdParams()
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

func collectOperatorAuthKeysDeleteCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForOperatorAuthKeysDeleteCmd("/operators/{operator_id}/auth_keys/{auth_key_id}"),
		query:  buildQueryForOperatorAuthKeysDeleteCmd(),
	}, nil
}

func buildPathForOperatorAuthKeysDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"auth_key_id"+"}", OperatorAuthKeysDeleteCmdAuthKeyId, -1)

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorAuthKeysDeleteCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorAuthKeysDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
