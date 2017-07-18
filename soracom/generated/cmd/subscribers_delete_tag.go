package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersDeleteTagCmdImsi holds value of 'imsi' option
var SubscribersDeleteTagCmdImsi string

// SubscribersDeleteTagCmdTagName holds value of 'tag_name' option
var SubscribersDeleteTagCmdTagName string

func init() {
	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	SubscribersCmd.AddCommand(SubscribersDeleteTagCmd)
}

// SubscribersDeleteTagCmd defines 'delete-tag' subcommand
var SubscribersDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/subscribers/{imsi}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/tags/{tag_name}:delete:description`),
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

		param, err := collectSubscribersDeleteTagCmdParams()
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

func collectSubscribersDeleteTagCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSubscribersDeleteTagCmd("/subscribers/{imsi}/tags/{tag_name}"),
		query:  buildQueryForSubscribersDeleteTagCmd(),
	}, nil
}

func buildPathForSubscribersDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersDeleteTagCmdImsi, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", SubscribersDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForSubscribersDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
