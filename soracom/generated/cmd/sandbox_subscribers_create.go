package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	SandboxSubscribersCmd.AddCommand(SandboxSubscribersCreateCmd)
}

// SandboxSubscribersCreateCmd defines 'create' subcommand
var SandboxSubscribersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sandbox/subscribers/create:post:summary"),
	Long:  TRAPI(`/sandbox/subscribers/create:post:description`),
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

		param, err := collectSandboxSubscribersCreateCmdParams(ac)
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

func collectSandboxSubscribersCreateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSandboxSubscribersCreateCmd("/sandbox/subscribers/create"),
		query:  buildQueryForSandboxSubscribersCreateCmd(),
	}, nil
}

func buildPathForSandboxSubscribersCreateCmd(path string) string {

	return path
}

func buildQueryForSandboxSubscribersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
