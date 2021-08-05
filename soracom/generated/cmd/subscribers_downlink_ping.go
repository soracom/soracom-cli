// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SubscribersDownlinkPingCmdImsi holds value of 'imsi' option
var SubscribersDownlinkPingCmdImsi string

// SubscribersDownlinkPingCmdNumberOfPingRequests holds value of 'numberOfPingRequests' option
var SubscribersDownlinkPingCmdNumberOfPingRequests int64

// SubscribersDownlinkPingCmdTimeoutSeconds holds value of 'timeoutSeconds' option
var SubscribersDownlinkPingCmdTimeoutSeconds int64

// SubscribersDownlinkPingCmdBody holds contents of request body to be sent
var SubscribersDownlinkPingCmdBody string

func init() {
	SubscribersDownlinkPingCmd.Flags().StringVar(&SubscribersDownlinkPingCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersDownlinkPingCmd.Flags().Int64Var(&SubscribersDownlinkPingCmdNumberOfPingRequests, "number-of-ping-requests", 1, TRAPI(""))

	SubscribersDownlinkPingCmd.Flags().Int64Var(&SubscribersDownlinkPingCmdTimeoutSeconds, "timeout-seconds", 1, TRAPI(""))

	SubscribersDownlinkPingCmd.Flags().StringVar(&SubscribersDownlinkPingCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SubscribersCmd.AddCommand(SubscribersDownlinkPingCmd)
}

// SubscribersDownlinkPingCmd defines 'downlink-ping' subcommand
var SubscribersDownlinkPingCmd = &cobra.Command{
	Use:   "downlink-ping",
	Short: TRAPI("/subscribers/{imsi}/downlink/ping:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/downlink/ping:post:description`),
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

		param, err := collectSubscribersDownlinkPingCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSubscribersDownlinkPingCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSubscribersDownlinkPingCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SubscribersDownlinkPingCmdImsi == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersDownlinkPingCmd("/subscribers/{imsi}/downlink/ping"),
		query:       buildQueryForSubscribersDownlinkPingCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersDownlinkPingCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersDownlinkPingCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersDownlinkPingCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersDownlinkPingCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersDownlinkPingCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersDownlinkPingCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersDownlinkPingCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersDownlinkPingCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersDownlinkPingCmdBody)
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

	if SubscribersDownlinkPingCmdNumberOfPingRequests != 1 {
		result["numberOfPingRequests"] = SubscribersDownlinkPingCmdNumberOfPingRequests
	}

	if SubscribersDownlinkPingCmdTimeoutSeconds != 1 {
		result["timeoutSeconds"] = SubscribersDownlinkPingCmdTimeoutSeconds
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
