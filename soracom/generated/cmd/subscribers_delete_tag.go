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
	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdImsi, "imsi", "", TR("subscribers.delete_subscriber_tag.delete.parameters.imsi.description"))

	SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdTagName, "tag-name", "", TR("subscribers.delete_subscriber_tag.delete.parameters.tag_name.description"))

	SubscribersCmd.AddCommand(SubscribersDeleteTagCmd)
}

// SubscribersDeleteTagCmd defines 'delete-tag' subcommand
var SubscribersDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TR("subscribers.delete_subscriber_tag.delete.summary"),
	Long:  TR(`subscribers.delete_subscriber_tag.delete.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
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

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
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
