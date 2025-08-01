// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersDeleteTransferTokenCmdToken holds value of 'token' option
var SubscribersDeleteTransferTokenCmdToken string

func InitSubscribersDeleteTransferTokenCmd() {
	SubscribersDeleteTransferTokenCmd.Flags().StringVar(&SubscribersDeleteTransferTokenCmdToken, "token", "", TRAPI("Transfer token. The transfer token can be obtained from [Subscriber:issueSubscriberTransferToken](#!/Subscriber/issueSubscriberTransferToken)."))

	SubscribersDeleteTransferTokenCmd.RunE = SubscribersDeleteTransferTokenCmdRunE

	SubscribersCmd.AddCommand(SubscribersDeleteTransferTokenCmd)
}

// SubscribersDeleteTransferTokenCmd defines 'delete-transfer-token' subcommand
var SubscribersDeleteTransferTokenCmd = &cobra.Command{
	Use:   "delete-transfer-token",
	Short: TRAPI("/subscribers/transfer_token/{token}:delete:summary"),
	Long:  TRAPI(`/subscribers/transfer_token/{token}:delete:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "deleteSubscriberTransferToken"),
}

func SubscribersDeleteTransferTokenCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
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

	param, err := collectSubscribersDeleteTransferTokenCmdParams(ac)
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

func collectSubscribersDeleteTransferTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("token", "token", "path", parsedBody, SubscribersDeleteTransferTokenCmdToken)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSubscribersDeleteTransferTokenCmd("/subscribers/transfer_token/{token}"),
		query:  buildQueryForSubscribersDeleteTransferTokenCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersDeleteTransferTokenCmd(path string) string {

	escapedToken := url.PathEscape(SubscribersDeleteTransferTokenCmdToken)

	path = strReplace(path, "{"+"token"+"}", escapedToken, -1)

	return path
}

func buildQueryForSubscribersDeleteTransferTokenCmd() url.Values {
	result := url.Values{}

	return result
}
