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

// RolesCreateCmdDescription holds value of 'description' option
var RolesCreateCmdDescription string

// RolesCreateCmdOperatorId holds value of 'operator_id' option
var RolesCreateCmdOperatorId string

// RolesCreateCmdPermission holds value of 'permission' option
var RolesCreateCmdPermission string

// RolesCreateCmdRoleId holds value of 'role_id' option
var RolesCreateCmdRoleId string

// RolesCreateCmdBody holds contents of request body to be sent
var RolesCreateCmdBody string

func InitRolesCreateCmd() {
	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdDescription, "description", "", TRAPI("The description of the operator managed role."))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdPermission, "permission", "", TRAPI("Permission as JSON for the operator managed role."))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdRoleId, "role-id", "", TRAPI("Role ID."))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	RolesCreateCmd.RunE = RolesCreateCmdRunE

	RolesCmd.AddCommand(RolesCreateCmd)
}

// RolesCreateCmd defines 'create' subcommand
var RolesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:post:description`) + "\n\n" + createLinkToAPIReference("Role", "createRole"),
}

func RolesCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectRolesCreateCmdParams(ac)
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

func collectRolesCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if RolesCreateCmdOperatorId == "" {
		RolesCreateCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForRolesCreateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("permission", "permission", "body", parsedBody, RolesCreateCmdPermission)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("role_id", "role-id", "path", parsedBody, RolesCreateCmdRoleId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForRolesCreateCmd("/operators/{operator_id}/roles/{role_id}"),
		query:       buildQueryForRolesCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForRolesCreateCmd(path string) string {

	escapedOperatorId := url.PathEscape(RolesCreateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedRoleId := url.PathEscape(RolesCreateCmdRoleId)

	path = strReplace(path, "{"+"role_id"+"}", escapedRoleId, -1)

	return path
}

func buildQueryForRolesCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForRolesCreateCmd() (string, error) {
	var result map[string]interface{}

	if RolesCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(RolesCreateCmdBody, "@") {
			fname := strings.TrimPrefix(RolesCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if RolesCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(RolesCreateCmdBody)
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

	if RolesCreateCmdDescription != "" {
		result["description"] = RolesCreateCmdDescription
	}

	if RolesCreateCmdPermission != "" {
		result["permission"] = RolesCreateCmdPermission
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
