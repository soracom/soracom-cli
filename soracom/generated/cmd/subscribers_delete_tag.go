// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersDeleteTagCmdImsi holds value of 'imsi' option
var SubscribersDeleteTagCmdImsi string

// SubscribersDeleteTagCmdTagName holds value of 'tag_name' option
var SubscribersDeleteTagCmdTagName string

func InitSubscribersDeleteTagCmd() {
	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	SubscribersDeleteTagCmd.RunE = SubscribersDeleteTagCmdRunE

	SubscribersCmd.AddCommand(SubscribersDeleteTagCmd)
}

// SubscribersDeleteTagCmd defines 'delete-tag' subcommand
var SubscribersDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/subscribers/{imsi}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/tags/{tag_name}:delete:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "deleteSubscriberTag"),
}

func SubscribersDeleteTagCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSubscribersDeleteTagCmdParams(ac)
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

func collectSubscribersDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SubscribersDeleteTagCmdImsi)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("tag_name", "tag-name", "path", parsedBody, SubscribersDeleteTagCmdTagName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSubscribersDeleteTagCmd("/subscribers/{imsi}/tags/{tag_name}"),
		query:  buildQueryForSubscribersDeleteTagCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersDeleteTagCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersDeleteTagCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	escapedTagName := url.PathEscape(SubscribersDeleteTagCmdTagName)

	path = strReplace(path, "{"+"tag_name"+"}", escapedTagName, -1)

	return path
}

func buildQueryForSubscribersDeleteTagCmd() url.Values {
	result := url.Values{}

	return result
}
