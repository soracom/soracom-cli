package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSetToStandbyCmdImsi holds value of 'imsi' option
var SubscribersSetToStandbyCmdImsi string

func init() {
	SubscribersSetToStandbyCmd.Flags().StringVar(&SubscribersSetToStandbyCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersCmd.AddCommand(SubscribersSetToStandbyCmd)
}

// SubscribersSetToStandbyCmd defines 'set-to-standby' subcommand
var SubscribersSetToStandbyCmd = &cobra.Command{
	Use:   "set-to-standby",
	Short: TRAPI("/subscribers/{imsi}/set_to_standby:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/set_to_standby:post:description`),
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

		param, err := collectSubscribersSetToStandbyCmdParams(ac)
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

func collectSubscribersSetToStandbyCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersSetToStandbyCmd("/subscribers/{imsi}/set_to_standby"),
		query:  buildQueryForSubscribersSetToStandbyCmd(),
	}, nil
}

func buildPathForSubscribersSetToStandbyCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSetToStandbyCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSetToStandbyCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
