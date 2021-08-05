// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// LagoonMigrationGetInfoCmdPlan holds value of 'plan' option
var LagoonMigrationGetInfoCmdPlan string

// LagoonMigrationGetInfoCmdBody holds contents of request body to be sent
var LagoonMigrationGetInfoCmdBody string

func init() {
	LagoonMigrationGetInfoCmd.Flags().StringVar(&LagoonMigrationGetInfoCmdPlan, "plan", "", TRAPI(""))

	LagoonMigrationGetInfoCmd.Flags().StringVar(&LagoonMigrationGetInfoCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonMigrationCmd.AddCommand(LagoonMigrationGetInfoCmd)
}

// LagoonMigrationGetInfoCmd defines 'get-info' subcommand
var LagoonMigrationGetInfoCmd = &cobra.Command{
	Use:   "get-info",
	Short: TRAPI("/lagoon/migration:get:summary"),
	Long:  TRAPI(`/lagoon/migration:get:description`),
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

		param, err := collectLagoonMigrationGetInfoCmdParams(ac)
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

func collectLagoonMigrationGetInfoCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForLagoonMigrationGetInfoCmd()
	if err != nil {
		return nil, err
	}
	contentType := ""

	return &apiParams{
		method:      "GET",
		path:        buildPathForLagoonMigrationGetInfoCmd("/lagoon/migration"),
		query:       buildQueryForLagoonMigrationGetInfoCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonMigrationGetInfoCmd(path string) string {

	return path
}

func buildQueryForLagoonMigrationGetInfoCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonMigrationGetInfoCmd() (string, error) {
	var result map[string]interface{}

	if LagoonMigrationGetInfoCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonMigrationGetInfoCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonMigrationGetInfoCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonMigrationGetInfoCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonMigrationGetInfoCmdBody)
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

	if LagoonMigrationGetInfoCmdPlan != "" {
		result["plan"] = LagoonMigrationGetInfoCmdPlan
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
