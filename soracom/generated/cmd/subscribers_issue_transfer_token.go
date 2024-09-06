// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail holds value of 'transferDestinationOperatorEmail' option
var SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail string

// SubscribersIssueTransferTokenCmdTransferDestinationOperatorId holds value of 'transferDestinationOperatorId' option
var SubscribersIssueTransferTokenCmdTransferDestinationOperatorId string

// SubscribersIssueTransferTokenCmdTransferImsi holds multiple values of 'transferImsi' option
var SubscribersIssueTransferTokenCmdTransferImsi []string

// SubscribersIssueTransferTokenCmdBody holds contents of request body to be sent
var SubscribersIssueTransferTokenCmdBody string

func InitSubscribersIssueTransferTokenCmd() {
	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail, "transfer-destination-operator-email", "", TRAPI("Primary email address of the destination operator. Please confirm with the owner of the destination operator."))

	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorId, "transfer-destination-operator-id", "", TRAPI("Operator ID of the destination operator. Please confirm with the owner of the destination operator."))

	SubscribersIssueTransferTokenCmd.Flags().StringSliceVar(&SubscribersIssueTransferTokenCmdTransferImsi, "transfer-imsi", []string{}, TRAPI("IMSI of the SIM to be transferred. The IMSI can be obtained from the [Sim:listSims API](#!/Sim/listSims).If there is a possibility of canceling SIMs one by one, please specify only one IMSI."))

	SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersIssueTransferTokenCmd.RunE = SubscribersIssueTransferTokenCmdRunE

	SubscribersCmd.AddCommand(SubscribersIssueTransferTokenCmd)
}

// SubscribersIssueTransferTokenCmd defines 'issue-transfer-token' subcommand
var SubscribersIssueTransferTokenCmd = &cobra.Command{
	Use:   "issue-transfer-token",
	Short: TRAPI("/subscribers/transfer_token/issue:post:summary"),
	Long:  TRAPI(`/subscribers/transfer_token/issue:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "issueSubscriberTransferToken"),
}

func SubscribersIssueTransferTokenCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersIssueTransferTokenCmdParams(ac)
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

func collectSubscribersIssueTransferTokenCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSubscribersIssueTransferTokenCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("transferDestinationOperatorEmail", "transfer-destination-operator-email", "body", parsedBody, SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("transferDestinationOperatorId", "transfer-destination-operator-id", "body", parsedBody, SubscribersIssueTransferTokenCmdTransferDestinationOperatorId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringSliceParameterIsSupplied("transferImsi", "transfer-imsi", "body", parsedBody, SubscribersIssueTransferTokenCmdTransferImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersIssueTransferTokenCmd("/subscribers/transfer_token/issue"),
		query:       buildQueryForSubscribersIssueTransferTokenCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersIssueTransferTokenCmd(path string) string {

	return path
}

func buildQueryForSubscribersIssueTransferTokenCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersIssueTransferTokenCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersIssueTransferTokenCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersIssueTransferTokenCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersIssueTransferTokenCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SubscribersIssueTransferTokenCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersIssueTransferTokenCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail != "" {
		result["transferDestinationOperatorEmail"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail
	}

	if SubscribersIssueTransferTokenCmdTransferDestinationOperatorId != "" {
		result["transferDestinationOperatorId"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorId
	}

	if len(SubscribersIssueTransferTokenCmdTransferImsi) != 0 {
		result["transferImsi"] = SubscribersIssueTransferTokenCmdTransferImsi
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
