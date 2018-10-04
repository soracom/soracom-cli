package cmd

import (
	"encoding/json"

	"io/ioutil"

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

func init() {
	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdDescription, "description", "", TRAPI(""))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdPermission, "permission", "", TRAPI(""))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdRoleId, "role-id", "", TRAPI("role_id"))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	RolesCmd.AddCommand(RolesCreateCmd)
}

// RolesCreateCmd defines 'create' subcommand
var RolesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators/{operator_id}/roles/{role_id}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles/{role_id}:post:description`),
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

		param, err := collectRolesCreateCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectRolesCreateCmdParams(ac *apiClient) (*apiParams, error) {

	if RolesCreateCmdOperatorId == "" {
		RolesCreateCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForRolesCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForRolesCreateCmd("/operators/{operator_id}/roles/{role_id}"),
		query:       buildQueryForRolesCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForRolesCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesCreateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"role_id"+"}", RolesCreateCmdRoleId, -1)

	return path
}

func buildQueryForRolesCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForRolesCreateCmd() (string, error) {
	var result map[string]interface{}

	if RolesCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(RolesCreateCmdBody, "@") {
			fname := strings.TrimPrefix(RolesCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if RolesCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
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
