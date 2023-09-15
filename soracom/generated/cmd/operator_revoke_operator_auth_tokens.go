// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorRevokeOperatorAuthTokensCmdOperatorId holds value of 'operator_id' option
var OperatorRevokeOperatorAuthTokensCmdOperatorId string

func InitOperatorRevokeOperatorAuthTokensCmd() {
	OperatorRevokeOperatorAuthTokensCmd.Flags().StringVar(&OperatorRevokeOperatorAuthTokensCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	OperatorRevokeOperatorAuthTokensCmd.RunE = OperatorRevokeOperatorAuthTokensCmdRunE

	OperatorCmd.AddCommand(OperatorRevokeOperatorAuthTokensCmd)
}

// OperatorRevokeOperatorAuthTokensCmd defines 'revoke-operator-auth-tokens' subcommand
var OperatorRevokeOperatorAuthTokensCmd = &cobra.Command{
	Use:   "revoke-operator-auth-tokens",
	Short: TRAPI("/operators/{operator_id}/tokens:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/tokens:delete:description`) + "\n\n" + createLinkToAPIReference("Operator", "revokeOperatorAuthTokens"),
}

func OperatorRevokeOperatorAuthTokensCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorRevokeOperatorAuthTokensCmdParams(ac)
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

func collectOperatorRevokeOperatorAuthTokensCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorRevokeOperatorAuthTokensCmdOperatorId == "" {
		OperatorRevokeOperatorAuthTokensCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForOperatorRevokeOperatorAuthTokensCmd("/operators/{operator_id}/tokens"),
		query:  buildQueryForOperatorRevokeOperatorAuthTokensCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorRevokeOperatorAuthTokensCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorRevokeOperatorAuthTokensCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorRevokeOperatorAuthTokensCmd() url.Values {
	result := url.Values{}

	return result
}
