package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersEnableTerminationCmdImsi holds value of 'imsi' option
var SubscribersEnableTerminationCmdImsi string

func init() {
	SubscribersEnableTerminationCmd.Flags().StringVar(&SubscribersEnableTerminationCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersEnableTerminationCmd)
}

// SubscribersEnableTerminationCmd defines 'enable-termination' subcommand
var SubscribersEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/subscribers/{imsi}/enable_termination:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/enable_termination:post:description`),
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

		param, err := collectSubscribersEnableTerminationCmdParams()
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

func collectSubscribersEnableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersEnableTerminationCmd("/subscribers/{imsi}/enable_termination"),
		query:  buildQueryForSubscribersEnableTerminationCmd(),
	}, nil
}

func buildPathForSubscribersEnableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersEnableTerminationCmdImsi, -1)

	return path
}

func buildQueryForSubscribersEnableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
