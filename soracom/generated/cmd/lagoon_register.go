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

// LagoonRegisterCmdPlan holds value of 'plan' option
var LagoonRegisterCmdPlan string

// LagoonRegisterCmdUserPassword holds value of 'userPassword' option
var LagoonRegisterCmdUserPassword string

// LagoonRegisterCmdBody holds contents of request body to be sent
var LagoonRegisterCmdBody string

func init() {
	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdPlan, "plan", "", TRAPI(""))

	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdUserPassword, "user-password", "", TRAPI("This password is used by the initial user's login."))

	LagoonRegisterCmd.Flags().StringVar(&LagoonRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonCmd.AddCommand(LagoonRegisterCmd)
}

// LagoonRegisterCmd defines 'register' subcommand
var LagoonRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/lagoon/register:post:summary"),
	Long:  TRAPI(`/lagoon/register:post:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "registerLagoon"),
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

		param, err := collectLagoonRegisterCmdParams(ac)
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

func collectLagoonRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonRegisterCmd()
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
		path:        buildPathForLagoonRegisterCmd("/lagoon/register"),
		query:       buildQueryForLagoonRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonRegisterCmd(path string) string {

	return path
}

func buildQueryForLagoonRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonRegisterCmd() (string, error) {
	var result map[string]interface{}

	if LagoonRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonRegisterCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonRegisterCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonRegisterCmdBody)
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
