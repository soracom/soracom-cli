package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail holds value of 'transferDestinationOperatorEmail' option
var SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail string

// SubscribersIssueTransferTokenCmdTransferDestinationOperatorId holds value of 'transferDestinationOperatorId' option
var SubscribersIssueTransferTokenCmdTransferDestinationOperatorId string

// SubscribersIssueTransferTokenCmdBody holds contents of request body to be sent
var SubscribersIssueTransferTokenCmdBody string

func init() {
	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail, "transfer-destination-operator-email", "", TRAPI(""))

	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorId, "transfer-destination-operator-id", "", TRAPI(""))

	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersIssueTransferTokenCmd)
}

// SubscribersIssueTransferTokenCmd defines 'issue-transfer-token' subcommand
var SubscribersIssueTransferTokenCmd = &cobra.Command{
	Use:   "issue-transfer-token",
	Short: TRAPI("/subscribers/transfer_token/issue:post:summary"),
	Long:  TRAPI(`/subscribers/transfer_token/issue:post:description`),
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

		param, err := collectSubscribersIssueTransferTokenCmdParams(ac)
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

func collectSubscribersIssueTransferTokenCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersIssueTransferTokenCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersIssueTransferTokenCmd("/subscribers/transfer_token/issue"),
		query:       buildQueryForSubscribersIssueTransferTokenCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersIssueTransferTokenCmd(path string) string {

	return path
}

func buildQueryForSubscribersIssueTransferTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersIssueTransferTokenCmd() (string, error) {
	if SubscribersIssueTransferTokenCmdBody != "" {
		if strings.HasPrefix(SubscribersIssueTransferTokenCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersIssueTransferTokenCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if SubscribersIssueTransferTokenCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return SubscribersIssueTransferTokenCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail != "" {
		result["transferDestinationOperatorEmail"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail
	}

	if SubscribersIssueTransferTokenCmdTransferDestinationOperatorId != "" {
		result["transferDestinationOperatorId"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
