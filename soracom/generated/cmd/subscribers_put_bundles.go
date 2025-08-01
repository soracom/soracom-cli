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

// SubscribersPutBundlesCmdImsi holds value of 'imsi' option
var SubscribersPutBundlesCmdImsi string

// SubscribersPutBundlesCmdBody holds contents of request body to be sent
var SubscribersPutBundlesCmdBody string

func InitSubscribersPutBundlesCmd() {
	SubscribersPutBundlesCmd.Flags().StringVar(&SubscribersPutBundlesCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber. The IMSI can be obtained from the [Sim:listSims API](#!/Sim/listSims)."))

	SubscribersPutBundlesCmd.Flags().StringVar(&SubscribersPutBundlesCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersPutBundlesCmd.RunE = SubscribersPutBundlesCmdRunE

	SubscribersCmd.AddCommand(SubscribersPutBundlesCmd)
}

// SubscribersPutBundlesCmd defines 'put-bundles' subcommand
var SubscribersPutBundlesCmd = &cobra.Command{
	Use:   "put-bundles",
	Short: TRAPI("/subscribers/{imsi}/bundles:put:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/bundles:put:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "putBundles"),
}

func SubscribersPutBundlesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersPutBundlesCmdParams(ac)
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

func collectSubscribersPutBundlesCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSubscribersPutBundlesCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersPutBundlesCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSubscribersPutBundlesCmd("/subscribers/{imsi}/bundles"),
		query:       buildQueryForSubscribersPutBundlesCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersPutBundlesCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersPutBundlesCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersPutBundlesCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersPutBundlesCmd() (string, error) {
	var b []byte
	var err error

	if SubscribersPutBundlesCmdBody != "" {
		if strings.HasPrefix(SubscribersPutBundlesCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersPutBundlesCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SubscribersPutBundlesCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersPutBundlesCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
