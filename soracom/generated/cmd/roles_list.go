// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// RolesListCmdOperatorId holds value of 'operator_id' option
var RolesListCmdOperatorId string

// RolesListCmdOwner holds value of 'owner' option
var RolesListCmdOwner string

// RolesListCmdOutputJSONL indicates to output with jsonl format
var RolesListCmdOutputJSONL bool

func InitRolesListCmd() {
	RolesListCmd.Flags().StringVar(&RolesListCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	RolesListCmd.Flags().StringVar(&RolesListCmdOwner, "owner", "", TRAPI("Filters by the type of the role. Specify one of the following:- 'operator': Operator managed role (default)- 'soracom': Soracom managed role- 'all': All roles"))

	RolesListCmd.Flags().BoolVar(&RolesListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	RolesListCmd.RunE = RolesListCmdRunE

	RolesCmd.AddCommand(RolesListCmd)
}

// RolesListCmd defines 'list' subcommand
var RolesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/roles:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles:get:description`) + "\n\n" + createLinkToAPIReference("Role", "listRoles"),
}

func RolesListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectRolesListCmdParams(ac)
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
		if RolesListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectRolesListCmdParams(ac *apiClient) (*apiParams, error) {
	if RolesListCmdOperatorId == "" {
		RolesListCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListCmd("/operators/{operator_id}/roles"),
		query:  buildQueryForRolesListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesListCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForRolesListCmd() url.Values {
	result := url.Values{}

	if RolesListCmdOwner != "" {
		result.Add("owner", RolesListCmdOwner)
	}

	return result
}
