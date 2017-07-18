package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersActivateCmdImsi holds value of 'imsi' option
var SubscribersActivateCmdImsi string

func init() {
	SubscribersActivateCmd.Flags().StringVar(&SubscribersActivateCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersActivateCmd)
}

// SubscribersActivateCmd defines 'activate' subcommand
var SubscribersActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: TRAPI("/subscribers/{imsi}/activate:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/activate:post:description`),
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

		param, err := collectSubscribersActivateCmdParams()
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

func collectSubscribersActivateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersActivateCmd("/subscribers/{imsi}/activate"),
		query:  buildQueryForSubscribersActivateCmd(),
	}, nil
}

func buildPathForSubscribersActivateCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersActivateCmdImsi, -1)

	return path
}

func buildQueryForSubscribersActivateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
