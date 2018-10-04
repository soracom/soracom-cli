package cmd

import (
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersPutTagsCmdImsi holds value of 'imsi' option
var SubscribersPutTagsCmdImsi string

// SubscribersPutTagsCmdBody holds contents of request body to be sent
var SubscribersPutTagsCmdBody string

func init() {
	SubscribersPutTagsCmd.Flags().StringVar(&SubscribersPutTagsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersPutTagsCmd.Flags().StringVar(&SubscribersPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersPutTagsCmd)
}

// SubscribersPutTagsCmd defines 'put-tags' subcommand
var SubscribersPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/subscribers/{imsi}/tags:put:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/tags:put:description`),
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

		param, err := collectSubscribersPutTagsCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectSubscribersPutTagsCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersPutTagsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSubscribersPutTagsCmd("/subscribers/{imsi}/tags"),
		query:       buildQueryForSubscribersPutTagsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersPutTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersPutTagsCmdImsi, -1)

	return path
}

func buildQueryForSubscribersPutTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersPutTagsCmd() (string, error) {
	var b []byte
	var err error

	if SubscribersPutTagsCmdBody != "" {
		if strings.HasPrefix(SubscribersPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersPutTagsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersPutTagsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersPutTagsCmdBody)
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
