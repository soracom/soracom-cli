package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonRegisterCmdPlan holds value of 'plan' option
var LagoonRegisterCmdPlan string

// LagoonRegisterCmdUserPassword holds value of 'userPassword' option
var LagoonRegisterCmdUserPassword string

// LagoonRegisterCmdBody holds contents of request body to be sent
var LagoonRegisterCmdBody string

func init() {
	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdPlan, "plan", "", TRAPI(""))

	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdUserPassword, "user-password", "", TRAPI(""))

	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonRegisterCmd)
}

// LagoonRegisterCmd defines 'register' subcommand
var LagoonRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/lagoon/register:post:summary"),
	Long:  TRAPI(`/lagoon/register:post:description`),
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

		param, err := collectLagoonRegisterCmdParams(ac)
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLagoonRegisterCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLagoonRegisterCmd("/lagoon/register"),
		query:       buildQueryForLagoonRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonRegisterCmd(path string) string {

	return path
}

func buildQueryForLagoonRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonRegisterCmd() (string, error) {
	if LagoonRegisterCmdBody != "" {
		if strings.HasPrefix(LagoonRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonRegisterCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LagoonRegisterCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LagoonRegisterCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LagoonRegisterCmdPlan != "" {
		result["plan"] = LagoonRegisterCmdPlan
	}

	if LagoonRegisterCmdUserPassword != "" {
		result["userPassword"] = LagoonRegisterCmdUserPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
