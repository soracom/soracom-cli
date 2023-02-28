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

// LagoonDashboardsUpdatePermissionsCmdDashboardId holds value of 'dashboard_id' option
var LagoonDashboardsUpdatePermissionsCmdDashboardId int64

// LagoonDashboardsUpdatePermissionsCmdClassic holds value of 'classic' option
var LagoonDashboardsUpdatePermissionsCmdClassic bool

// LagoonDashboardsUpdatePermissionsCmdBody holds contents of request body to be sent
var LagoonDashboardsUpdatePermissionsCmdBody string

func init() {
	LagoonDashboardsUpdatePermissionsCmd.Flags().Int64Var(&LagoonDashboardsUpdatePermissionsCmdDashboardId, "dashboard-id", 0, TRAPI("dashboard_id"))

	LagoonDashboardsUpdatePermissionsCmd.Flags().BoolVar(&LagoonDashboardsUpdatePermissionsCmdClassic, "classic", false, TRAPI("If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled."))

	LagoonDashboardsUpdatePermissionsCmd.Flags().StringVar(&LagoonDashboardsUpdatePermissionsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonDashboardsCmd.AddCommand(LagoonDashboardsUpdatePermissionsCmd)
}

// LagoonDashboardsUpdatePermissionsCmd defines 'update-permissions' subcommand
var LagoonDashboardsUpdatePermissionsCmd = &cobra.Command{
	Use:   "update-permissions",
	Short: TRAPI("/lagoon/dashboards/{dashboard_id}/permissions:put:summary"),
	Long:  TRAPI(`/lagoon/dashboards/{dashboard_id}/permissions:put:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "updateLagoonDashboardPermissions"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectLagoonDashboardsUpdatePermissionsCmdParams(ac)
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

func collectLagoonDashboardsUpdatePermissionsCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonDashboardsUpdatePermissionsCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("dashboard_id", "dashboard-id", "path", parsedBody, LagoonDashboardsUpdatePermissionsCmdDashboardId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonDashboardsUpdatePermissionsCmd("/lagoon/dashboards/{dashboard_id}/permissions"),
		query:       buildQueryForLagoonDashboardsUpdatePermissionsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonDashboardsUpdatePermissionsCmd(path string) string {

	path = strReplace(path, "{"+"dashboard_id"+"}", url.PathEscape(sprintf("%d", LagoonDashboardsUpdatePermissionsCmdDashboardId)), -1)

	return path
}

func buildQueryForLagoonDashboardsUpdatePermissionsCmd() url.Values {
	result := url.Values{}

	if LagoonDashboardsUpdatePermissionsCmdClassic != false {
		result.Add("classic", sprintf("%t", LagoonDashboardsUpdatePermissionsCmdClassic))
	}

	return result
}

func buildBodyForLagoonDashboardsUpdatePermissionsCmd() (string, error) {
	var result map[string]interface{}

	if LagoonDashboardsUpdatePermissionsCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonDashboardsUpdatePermissionsCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonDashboardsUpdatePermissionsCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonDashboardsUpdatePermissionsCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonDashboardsUpdatePermissionsCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
