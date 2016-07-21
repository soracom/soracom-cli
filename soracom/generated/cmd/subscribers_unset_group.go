package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var SubscribersUnsetGroupCmdImsi string

func init() {
	SubscribersUnsetGroupCmd.Flags().StringVar(&SubscribersUnsetGroupCmdImsi, "imsi", "", TR("subscribers.unset_group.post.parameters.imsi.description"))

	SubscribersCmd.AddCommand(SubscribersUnsetGroupCmd)
}

var SubscribersUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TR("subscribers.unset_group.post.summary"),
	Long:  TR(`subscribers.unset_group.post.description`),
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

		param, err := collectSubscribersUnsetGroupCmdParams()
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

func collectSubscribersUnsetGroupCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersUnsetGroupCmd("/subscribers/{imsi}/unset_group"),
		query:  buildQueryForSubscribersUnsetGroupCmd(),
	}, nil
}

func buildPathForSubscribersUnsetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersUnsetGroupCmdImsi, -1)

	return path
}

func buildQueryForSubscribersUnsetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
