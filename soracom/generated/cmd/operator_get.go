// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorGetCmdOperatorId holds value of 'operator_id' option
var OperatorGetCmdOperatorId string

func init() {
	OperatorGetCmd.Flags().StringVar(&OperatorGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorGetCmd)
}

// OperatorGetCmd defines 'get' subcommand
var OperatorGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}:get:description`),
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

		param, err := collectOperatorGetCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectOperatorGetCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorGetCmdOperatorId == "" {
		OperatorGetCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorGetCmd("/operators/{operator_id}"),
		query:  buildQueryForOperatorGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorGetCmd() url.Values {
	result := url.Values{}

	return result
}
