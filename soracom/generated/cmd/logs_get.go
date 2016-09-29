package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LogsGetCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var LogsGetCmdLastEvaluatedKey string

// LogsGetCmdResourceId holds value of 'resource_id' option
var LogsGetCmdResourceId string

// LogsGetCmdResourceType holds value of 'resource_type' option
var LogsGetCmdResourceType string

// LogsGetCmdService holds value of 'service' option
var LogsGetCmdService string

// LogsGetCmdFrom holds value of 'from' option
var LogsGetCmdFrom int64

// LogsGetCmdLimit holds value of 'limit' option
var LogsGetCmdLimit int64

// LogsGetCmdTo holds value of 'to' option
var LogsGetCmdTo int64

func init() {
	LogsGetCmd.Flags().StringVar(&LogsGetCmdLastEvaluatedKey, "last-evaluated-key", "", TR("logs.get_logs.get.parameters.last_evaluated_key.description"))

	LogsGetCmd.Flags().StringVar(&LogsGetCmdResourceId, "resource-id", "", TR("logs.get_logs.get.parameters.resource_id.description"))

	LogsGetCmd.Flags().StringVar(&LogsGetCmdResourceType, "resource-type", "", TR("logs.get_logs.get.parameters.resource_type.description"))

	LogsGetCmd.Flags().StringVar(&LogsGetCmdService, "service", "", TR("logs.get_logs.get.parameters.service.description"))

	LogsGetCmd.Flags().Int64Var(&LogsGetCmdFrom, "from", 0, TR("logs.get_logs.get.parameters.from.description"))

	LogsGetCmd.Flags().Int64Var(&LogsGetCmdLimit, "limit", 0, TR("logs.get_logs.get.parameters.limit.description"))

	LogsGetCmd.Flags().Int64Var(&LogsGetCmdTo, "to", 0, TR("logs.get_logs.get.parameters.to.description"))

	LogsCmd.AddCommand(LogsGetCmd)
}

// LogsGetCmd defines 'get' subcommand
var LogsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("logs.get_logs.get.summary"),
	Long:  TR(`logs.get_logs.get.description`),
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

		param, err := collectLogsGetCmdParams()
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

func collectLogsGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLogsGetCmd("/logs"),
		query:  buildQueryForLogsGetCmd(),
	}, nil
}

func buildPathForLogsGetCmd(path string) string {

	return path
}

func buildQueryForLogsGetCmd() string {
	result := []string{}

	if LogsGetCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", LogsGetCmdLastEvaluatedKey))
	}

	if LogsGetCmdResourceId != "" {
		result = append(result, sprintf("%s=%s", "resource_id", LogsGetCmdResourceId))
	}

	if LogsGetCmdResourceType != "" {
		result = append(result, sprintf("%s=%s", "resource_type", LogsGetCmdResourceType))
	}

	if LogsGetCmdService != "" {
		result = append(result, sprintf("%s=%s", "service", LogsGetCmdService))
	}

	if LogsGetCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", LogsGetCmdFrom))
	}

	if LogsGetCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", LogsGetCmdLimit))
	}

	if LogsGetCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", LogsGetCmdTo))
	}

	return strings.Join(result, "&")
}
