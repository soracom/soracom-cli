package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var SubscribersGetCmdImsi string

func init() {
	SubscribersGetCmd.Flags().StringVar(&SubscribersGetCmdImsi, "imsi", "", TR("subscribers.get_subscriber.get.parameters.imsi.description"))

	SubscribersCmd.AddCommand(SubscribersGetCmd)
}

var SubscribersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("subscribers.get_subscriber.get.summary"),
	Long:  TR(`subscribers.get_subscriber.get.description`),
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

		param, err := collectSubscribersGetCmdParams()
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

func collectSubscribersGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersGetCmd("/subscribers/{imsi}"),
		query:  buildQueryForSubscribersGetCmd(),
	}, nil
}

func buildPathForSubscribersGetCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersGetCmdImsi, -1)

	return path
}

func buildQueryForSubscribersGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
