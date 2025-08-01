// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// AnalysisQueriesStartCmdSql holds value of 'sql' option
var AnalysisQueriesStartCmdSql string

// AnalysisQueriesStartCmdFrom holds value of 'from' option
var AnalysisQueriesStartCmdFrom int64

// AnalysisQueriesStartCmdTo holds value of 'to' option
var AnalysisQueriesStartCmdTo int64

// AnalysisQueriesStartCmdBody holds contents of request body to be sent
var AnalysisQueriesStartCmdBody string

func InitAnalysisQueriesStartCmd() {
	AnalysisQueriesStartCmd.Flags().StringVar(&AnalysisQueriesStartCmdSql, "sql", "", TRAPI("Database query (SQL)."))

	AnalysisQueriesStartCmd.Flags().Int64Var(&AnalysisQueriesStartCmdFrom, "from", 0, TRAPI("Start of the period to apply the database query (UNIX time (seconds))."))

	AnalysisQueriesStartCmd.Flags().Int64Var(&AnalysisQueriesStartCmdTo, "to", 0, TRAPI("End of the period to apply the database query (UNIX time (seconds))."))

	AnalysisQueriesStartCmd.Flags().StringVar(&AnalysisQueriesStartCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	AnalysisQueriesStartCmd.RunE = AnalysisQueriesStartCmdRunE

	AnalysisQueriesCmd.AddCommand(AnalysisQueriesStartCmd)
}

// AnalysisQueriesStartCmd defines 'start' subcommand
var AnalysisQueriesStartCmd = &cobra.Command{
	Use:   "start",
	Short: TRAPI("/analysis/queries:post:summary"),
	Long:  TRAPI(`/analysis/queries:post:description`) + "\n\n" + createLinkToAPIReference("Analysis", "startAnalysisQueries"),
}

func AnalysisQueriesStartCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectAnalysisQueriesStartCmdParams(ac)
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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectAnalysisQueriesStartCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForAnalysisQueriesStartCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForAnalysisQueriesStartCmd("/analysis/queries"),
		query:       buildQueryForAnalysisQueriesStartCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAnalysisQueriesStartCmd(path string) string {

	return path
}

func buildQueryForAnalysisQueriesStartCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForAnalysisQueriesStartCmd() (string, error) {
	var result map[string]interface{}

	if AnalysisQueriesStartCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(AnalysisQueriesStartCmdBody, "@") {
			fname := strings.TrimPrefix(AnalysisQueriesStartCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if AnalysisQueriesStartCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(AnalysisQueriesStartCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if AnalysisQueriesStartCmdSql != "" {
		result["sql"] = AnalysisQueriesStartCmdSql
	}

	if AnalysisQueriesStartCmd.Flags().Lookup("from").Changed {
		result["from"] = AnalysisQueriesStartCmdFrom
	}

	if AnalysisQueriesStartCmd.Flags().Lookup("to").Changed {
		result["to"] = AnalysisQueriesStartCmdTo
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
