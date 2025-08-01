// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// RolesListVersionsCmdOperatorId holds value of 'operator_id' option
var RolesListVersionsCmdOperatorId string

// RolesListVersionsCmdRoleId holds value of 'role_id' option
var RolesListVersionsCmdRoleId string

// RolesListVersionsCmdOutputJSONL indicates to output with jsonl format
var RolesListVersionsCmdOutputJSONL bool

func InitRolesListVersionsCmd() {
	RolesListVersionsCmd.Flags().StringVar(&RolesListVersionsCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	RolesListVersionsCmd.Flags().StringVar(&RolesListVersionsCmdRoleId, "role-id", "", TRAPI("SRN (Soracom Resource Name) of the Soracom managed role. The format is 'srn:soracom:::Role:[roleName]'."))

	RolesListVersionsCmd.Flags().BoolVar(&RolesListVersionsCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	RolesListVersionsCmd.RunE = RolesListVersionsCmdRunE

	RolesCmd.AddCommand(RolesListVersionsCmd)
}

// RolesListVersionsCmd defines 'list-versions' subcommand
var RolesListVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}/versions:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}/versions:get:description`) + "\n\n" + createLinkToAPIReference("Role", "listRoleVersions"),
}

func RolesListVersionsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectRolesListVersionsCmdParams(ac)
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
		if RolesListVersionsCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectRolesListVersionsCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if RolesListVersionsCmdOperatorId == "" {
		RolesListVersionsCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("role_id", "role-id", "path", parsedBody, RolesListVersionsCmdRoleId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListVersionsCmd("/operators/{operator_id}/roles/{role_id}/versions"),
		query:  buildQueryForRolesListVersionsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesListVersionsCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesListVersionsCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(RolesListVersionsCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	return path
}

func buildQueryForRolesListVersionsCmd() url.Values {
	result := url.Values{}

	return result
}
