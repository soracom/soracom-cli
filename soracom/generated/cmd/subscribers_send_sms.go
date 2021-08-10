// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// SubscribersSendSmsCmdImsi holds value of 'imsi' option
var SubscribersSendSmsCmdImsi string

// SubscribersSendSmsCmdPayload holds value of 'payload' option
var SubscribersSendSmsCmdPayload string

// SubscribersSendSmsCmdEncodingType holds value of 'encodingType' option
var SubscribersSendSmsCmdEncodingType int64

// SubscribersSendSmsCmdBody holds contents of request body to be sent
var SubscribersSendSmsCmdBody string

func init() {
	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdPayload, "payload", "", TRAPI(""))

	SubscribersSendSmsCmd.Flags().Int64Var(&SubscribersSendSmsCmdEncodingType, "encoding-type", 2, TRAPI(""))

	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SubscribersCmd.AddCommand(SubscribersSendSmsCmd)
}

// SubscribersSendSmsCmd defines 'send-sms' subcommand
var SubscribersSendSmsCmd = &cobra.Command{
	Use:   "send-sms",
	Short: TRAPI("/subscribers/{imsi}/send_sms:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/send_sms:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectSubscribersSendSmsCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersSendSmsCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSubscribersSendSmsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SubscribersSendSmsCmdImsi == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSendSmsCmd("/subscribers/{imsi}/send_sms"),
		query:       buildQueryForSubscribersSendSmsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersSendSmsCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersSendSmsCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersSendSmsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersSendSmsCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSendSmsCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSendSmsCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSendSmsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersSendSmsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSendSmsCmdBody)
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

	if SubscribersSendSmsCmdPayload != "" {
		result["payload"] = SubscribersSendSmsCmdPayload
	}

	if SubscribersSendSmsCmdEncodingType != 2 {
		result["encodingType"] = SubscribersSendSmsCmdEncodingType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
