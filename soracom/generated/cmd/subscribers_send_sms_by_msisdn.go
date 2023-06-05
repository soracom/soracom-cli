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

// SubscribersSendSmsByMsisdnCmdMsisdn holds value of 'msisdn' option
var SubscribersSendSmsByMsisdnCmdMsisdn string

// SubscribersSendSmsByMsisdnCmdPayload holds value of 'payload' option
var SubscribersSendSmsByMsisdnCmdPayload string

// SubscribersSendSmsByMsisdnCmdEncodingType holds value of 'encodingType' option
var SubscribersSendSmsByMsisdnCmdEncodingType int64

// SubscribersSendSmsByMsisdnCmdBody holds contents of request body to be sent
var SubscribersSendSmsByMsisdnCmdBody string

func InitSubscribersSendSmsByMsisdnCmd() {
	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdMsisdn, "msisdn", "", TRAPI("MSISDN of the target subscriber."))

	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdPayload, "payload", "", TRAPI(""))

	SubscribersSendSmsByMsisdnCmd.Flags().Int64Var(&SubscribersSendSmsByMsisdnCmdEncodingType, "encoding-type", 2, TRAPI("Encoding type of the message body. Default is '2' ('DCS_UCS2').- '1': Send in GSM 7-bit that only supports standard alphabet. Kanji, Cyrillic, and Arabic characters cannot be sent. Maximum 160 characters (maximum 140 bytes).    Example: '{\"encodingType\": 1, \"payload\": \"Test message.\"}'- '2': Send in UCS-2, which supports Kanji, Cyrillic, Arabic, etc. Maximum 70 characters.    Example: '{\"encodingType\": 2, \"payload\": \"テストメッセージ\"}'"))

	SubscribersSendSmsByMsisdnCmd.Flags().StringVar(&SubscribersSendSmsByMsisdnCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersSendSmsByMsisdnCmd.RunE = SubscribersSendSmsByMsisdnCmdRunE

	SubscribersCmd.AddCommand(SubscribersSendSmsByMsisdnCmd)
}

// SubscribersSendSmsByMsisdnCmd defines 'send-sms-by-msisdn' subcommand
var SubscribersSendSmsByMsisdnCmd = &cobra.Command{
	Use:   "send-sms-by-msisdn",
	Short: TRAPI("/subscribers/msisdn/{msisdn}/send_sms:post:summary"),
	Long:  TRAPI(`/subscribers/msisdn/{msisdn}/send_sms:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "sendSmsByMsisdn"),
}

func SubscribersSendSmsByMsisdnCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := authHelper(ac, cmd, args)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSubscribersSendSmsByMsisdnCmdParams(ac)
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

func collectSubscribersSendSmsByMsisdnCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSubscribersSendSmsByMsisdnCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("msisdn", "msisdn", "path", parsedBody, SubscribersSendSmsByMsisdnCmdMsisdn)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSendSmsByMsisdnCmd("/subscribers/msisdn/{msisdn}/send_sms"),
		query:       buildQueryForSubscribersSendSmsByMsisdnCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersSendSmsByMsisdnCmd(path string) string {

	escapedMsisdn := url.PathEscape(SubscribersSendSmsByMsisdnCmdMsisdn)

	path = strReplace(path, "{"+"msisdn"+"}", escapedMsisdn, -1)

	return path
}

func buildQueryForSubscribersSendSmsByMsisdnCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersSendSmsByMsisdnCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSendSmsByMsisdnCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSendSmsByMsisdnCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSendSmsByMsisdnCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SubscribersSendSmsByMsisdnCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSendSmsByMsisdnCmdBody)
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

	if SubscribersSendSmsByMsisdnCmdPayload != "" {
		result["payload"] = SubscribersSendSmsByMsisdnCmdPayload
	}

	if SubscribersSendSmsByMsisdnCmd.Flags().Lookup("encoding-type").Changed {
		result["encodingType"] = SubscribersSendSmsByMsisdnCmdEncodingType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
