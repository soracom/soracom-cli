package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorGetCmdOperatorId holds value of 'operator_id' option
var OperatorGetCmdOperatorId string

func init() {
	OperatorGetCmd.Flags().StringVar(&OperatorGetCmdOperatorId, "operator-id", "", TR("operator ID"))

	OperatorCmd.AddCommand(OperatorGetCmd)
}

// OperatorGetCmd defines 'get' subcommand
var OperatorGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("operator.get_operator.get.summary"),
	Long:  TR(`operator.get_operator.get.description`),
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

		param, err := collectOperatorGetCmdParams()
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

func collectOperatorGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOperatorGetCmd("/operators/{operator_id}"),
		query:  buildQueryForOperatorGetCmd(),
	}, nil
}

func buildPathForOperatorGetCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorGetCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
