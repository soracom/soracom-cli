package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorGetSupportTokenCmdOperatorId holds value of 'operator_id' option
var OperatorGetSupportTokenCmdOperatorId string

func init() {
	OperatorGetSupportTokenCmd.Flags().StringVar(&OperatorGetSupportTokenCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorCmd.AddCommand(OperatorGetSupportTokenCmd)
}

// OperatorGetSupportTokenCmd defines 'get-support-token' subcommand
var OperatorGetSupportTokenCmd = &cobra.Command{
	Use:   "get-support-token",
	Short: TRAPI("/operators/{operator_id}/support/token:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/support/token:post:description`),
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

		param, err := collectOperatorGetSupportTokenCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectOperatorGetSupportTokenCmdParams(ac *apiClient) (*apiParams, error) {

	if OperatorGetSupportTokenCmdOperatorId == "" {
		OperatorGetSupportTokenCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorGetSupportTokenCmd("/operators/{operator_id}/support/token"),
		query:  buildQueryForOperatorGetSupportTokenCmd(),
	}, nil
}

func buildPathForOperatorGetSupportTokenCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorGetSupportTokenCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorGetSupportTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
