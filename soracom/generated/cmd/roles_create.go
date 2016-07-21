package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

var RolesCreateCmdDescription string

var RolesCreateCmdOperatorId string

var RolesCreateCmdPermission string

var RolesCreateCmdRoleId string

var RolesCreateCmdBody string

func init() {
	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdDescription, "description", "", TR(""))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdOperatorId, "operator-id", "", TR("operator_id"))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdPermission, "permission", "", TR(""))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdRoleId, "role-id", "", TR("role_id"))

	RolesCreateCmd.Flags().StringVar(&RolesCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	RolesCmd.AddCommand(RolesCreateCmd)
}

var RolesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("roles.create_role.post.summary"),
	Long:  TR(`roles.create_role.post.description`),
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

		param, err := collectRolesCreateCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectRolesCreateCmdParams() (*apiParams, error) {

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
	if RolesCreateCmdBody != "" {
		if strings.HasPrefix(RolesCreateCmdBody, "@") {
			fname := strings.TrimPrefix(RolesCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if RolesCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return RolesCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

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
