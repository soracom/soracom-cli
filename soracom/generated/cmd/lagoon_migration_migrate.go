// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/soracom/soracom-cli/generators/lib"

	"github.com/spf13/cobra"
)

// LagoonMigrationMigrateCmdDashboardIds holds multiple values of 'dashboardIds' option
var LagoonMigrationMigrateCmdDashboardIds []string

// LagoonMigrationMigrateCmdBody holds contents of request body to be sent
var LagoonMigrationMigrateCmdBody string

func init() {
	LagoonMigrationMigrateCmd.Flags().StringSliceVar(&LagoonMigrationMigrateCmdDashboardIds, "dashboard-ids", []string{}, TRAPI("A list of dashboard IDs to migrate"))

	LagoonMigrationMigrateCmd.Flags().StringVar(&LagoonMigrationMigrateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonMigrationCmd.AddCommand(LagoonMigrationMigrateCmd)
}

// LagoonMigrationMigrateCmd defines 'migrate' subcommand
var LagoonMigrationMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: TRAPI("/lagoon/migration:post:summary"),
	Long:  TRAPI(`/lagoon/migration:post:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "migrateLagoon"),
	RunE: func(cmd *cobra.Command, args []string) error {
		lib.WarnfStderr(TRCLI("cli.deprecated-api") + "\n")

		if len(args) > 0 {
			return fmt.Errorf("unexpected arguments passed => %v", args)
		}

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

		param, err := collectLagoonMigrationMigrateCmdParams(ac)
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
	},
}

func collectLagoonMigrationMigrateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonMigrationMigrateCmd()
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
		path:        buildPathForLagoonMigrationMigrateCmd("/lagoon/migration"),
		query:       buildQueryForLagoonMigrationMigrateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonMigrationMigrateCmd(path string) string {

	return path
}

func buildQueryForLagoonMigrationMigrateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonMigrationMigrateCmd() (string, error) {
	var result map[string]interface{}

	if LagoonMigrationMigrateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonMigrationMigrateCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonMigrationMigrateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonMigrationMigrateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonMigrationMigrateCmdBody)
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

	if len(LagoonMigrationMigrateCmdDashboardIds) != 0 {
		result["dashboardIds"] = LagoonMigrationMigrateCmdDashboardIds
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
