package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersGetCmdImsi holds value of 'imsi' option
var SubscribersGetCmdImsi string

func init() {
	SubscribersGetCmd.Flags().StringVar(&SubscribersGetCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersGetCmd)
}

// SubscribersGetCmd defines 'get' subcommand
var SubscribersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}:get:description`),
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

		param, err := collectSubscribersGetCmdParams(ac)
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

func collectSubscribersGetCmdParams(ac *apiClient) (*apiParams, error) {

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
