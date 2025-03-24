// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// CredentialsDeleteCmdCredentialsId holds value of 'credentials_id' option
var CredentialsDeleteCmdCredentialsId string

func InitCredentialsDeleteCmd() {
	CredentialsDeleteCmd.Flags().StringVar(&CredentialsDeleteCmdCredentialsId, "credentials-id", "", TRAPI("Credential set ID."))

	CredentialsDeleteCmd.RunE = CredentialsDeleteCmdRunE

	CredentialsCmd.AddCommand(CredentialsDeleteCmd)
}

// CredentialsDeleteCmd defines 'delete' subcommand
var CredentialsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/credentials/{credentials_id}:delete:summary"),
	Long:  TRAPI(`/credentials/{credentials_id}:delete:description`) + "\n\n" + createLinkToAPIReference("Credential", "deleteCredential"),
}

func CredentialsDeleteCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectCredentialsDeleteCmdParams(ac)
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

func collectCredentialsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("credentials_id", "credentials-id", "path", parsedBody, CredentialsDeleteCmdCredentialsId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForCredentialsDeleteCmd("/credentials/{credentials_id}"),
		query:  buildQueryForCredentialsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCredentialsDeleteCmd(path string) string {

	escapedCredentialsId := url.PathEscape(CredentialsDeleteCmdCredentialsId)

	path = strReplace(path, "{"+"credentials_id"+"}", escapedCredentialsId, -1)

	return path
}

func buildQueryForCredentialsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
