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

// SandboxSubscribersCreateCmdSubscription holds value of 'subscription' option
var SandboxSubscribersCreateCmdSubscription string

// SandboxSubscribersCreateCmdBundles holds multiple values of 'bundles' option
var SandboxSubscribersCreateCmdBundles []string

// SandboxSubscribersCreateCmdBody holds contents of request body to be sent
var SandboxSubscribersCreateCmdBody string

func InitSandboxSubscribersCreateCmd() {
	SandboxSubscribersCreateCmd.Flags().StringVar(&SandboxSubscribersCreateCmdSubscription, "subscription", "", TRAPI("Subscription. Specify one of:"))

	SandboxSubscribersCreateCmd.Flags().StringSliceVar(&SandboxSubscribersCreateCmdBundles, "bundles", []string{}, TRAPI("Bundle. If necessary, specify one of:"))

	SandboxSubscribersCreateCmd.Flags().StringVar(&SandboxSubscribersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxSubscribersCreateCmd.RunE = SandboxSubscribersCreateCmdRunE

	SandboxSubscribersCmd.AddCommand(SandboxSubscribersCreateCmd)
}

// SandboxSubscribersCreateCmd defines 'create' subcommand
var SandboxSubscribersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sandbox/subscribers/create:post:summary"),
	Long:  TRAPI(`/sandbox/subscribers/create:post:description`) + "\n\n" + createLinkToAPIReference("Subscriber", "sandboxCreateSubscriber"),
}

func SandboxSubscribersCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSandboxSubscribersCreateCmdParams(ac)
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

func collectSandboxSubscribersCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSandboxSubscribersCreateCmd()
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
		path:        buildPathForSandboxSubscribersCreateCmd("/sandbox/subscribers/create"),
		query:       buildQueryForSandboxSubscribersCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxSubscribersCreateCmd(path string) string {

	return path
}

func buildQueryForSandboxSubscribersCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxSubscribersCreateCmd() (string, error) {
	var result map[string]interface{}

	if SandboxSubscribersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxSubscribersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxSubscribersCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SandboxSubscribersCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxSubscribersCreateCmdBody)
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

	if SandboxSubscribersCreateCmdSubscription != "" {
		result["subscription"] = SandboxSubscribersCreateCmdSubscription
	}

	if len(SandboxSubscribersCreateCmdBundles) != 0 {
		result["bundles"] = SandboxSubscribersCreateCmdBundles
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
