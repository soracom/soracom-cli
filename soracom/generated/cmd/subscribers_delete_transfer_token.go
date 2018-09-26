package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersDeleteTransferTokenCmdToken holds value of 'token' option
var SubscribersDeleteTransferTokenCmdToken string

func init() {
	SubscribersDeleteTransferTokenCmd.Flags().StringVar(&SubscribersDeleteTransferTokenCmdToken, "token", "", TRAPI("token"))

	SubscribersCmd.AddCommand(SubscribersDeleteTransferTokenCmd)
}

// SubscribersDeleteTransferTokenCmd defines 'delete-transfer-token' subcommand
var SubscribersDeleteTransferTokenCmd = &cobra.Command{
	Use:   "delete-transfer-token",
	Short: TRAPI("/subscribers/transfer_token/{token}:delete:summary"),
	Long:  TRAPI(`/subscribers/transfer_token/{token}:delete:description`),
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

		param, err := collectSubscribersDeleteTransferTokenCmdParams(ac)
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

func collectSubscribersDeleteTransferTokenCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSubscribersDeleteTransferTokenCmd("/subscribers/transfer_token/{token}"),
		query:  buildQueryForSubscribersDeleteTransferTokenCmd(),
	}, nil
}

func buildPathForSubscribersDeleteTransferTokenCmd(path string) string {

	path = strings.Replace(path, "{"+"token"+"}", SubscribersDeleteTransferTokenCmdToken, -1)

	return path
}

func buildQueryForSubscribersDeleteTransferTokenCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
