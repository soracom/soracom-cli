// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LagoonDashboardsInitPermissionsCmdDashboardId holds value of 'dashboard_id' option
var LagoonDashboardsInitPermissionsCmdDashboardId int64

// LagoonDashboardsInitPermissionsCmdClassic holds value of 'classic' option
var LagoonDashboardsInitPermissionsCmdClassic bool

func init() {
	LagoonDashboardsInitPermissionsCmd.Flags().Int64Var(&LagoonDashboardsInitPermissionsCmdDashboardId, "dashboard-id", 0, TRAPI("dashboard_id"))

	LagoonDashboardsInitPermissionsCmd.Flags().BoolVar(&LagoonDashboardsInitPermissionsCmdClassic, "classic", false, TRAPI("If the value is true, a request will be issued to Lagoon Classic.  This is only valid if both Lagoon and Lagoon Classic are enabled."))
	LagoonDashboardsCmd.AddCommand(LagoonDashboardsInitPermissionsCmd)
}

// LagoonDashboardsInitPermissionsCmd defines 'init-permissions' subcommand
var LagoonDashboardsInitPermissionsCmd = &cobra.Command{
	Use:   "init-permissions",
	Short: TRAPI("/lagoon/dashboards/{dashboard_id}/permissions/init:post:summary"),
	Long:  TRAPI(`/lagoon/dashboards/{dashboard_id}/permissions/init:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonDashboardsInitPermissionsCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLagoonDashboardsInitPermissionsCmdParams(ac *apiClient) (*apiParams, error) {
	if LagoonDashboardsInitPermissionsCmdDashboardId == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "dashboard-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLagoonDashboardsInitPermissionsCmd("/lagoon/dashboards/{dashboard_id}/permissions/init"),
		query:  buildQueryForLagoonDashboardsInitPermissionsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonDashboardsInitPermissionsCmd(path string) string {

	path = strReplace(path, "{"+"dashboard_id"+"}", url.PathEscape(sprintf("%d", LagoonDashboardsInitPermissionsCmdDashboardId)), -1)

	return path
}

func buildQueryForLagoonDashboardsInitPermissionsCmd() url.Values {
	result := url.Values{}

	if LagoonDashboardsInitPermissionsCmdClassic != false {
		result.Add("classic", sprintf("%t", LagoonDashboardsInitPermissionsCmdClassic))
	}

	return result
}
