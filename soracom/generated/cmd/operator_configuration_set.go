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

// OperatorConfigurationSetCmdNamespace holds value of 'namespace' option
var OperatorConfigurationSetCmdNamespace string

// OperatorConfigurationSetCmdOperatorId holds value of 'operator_id' option
var OperatorConfigurationSetCmdOperatorId string

// OperatorConfigurationSetCmdExpectedVersion holds value of 'expectedVersion' option
var OperatorConfigurationSetCmdExpectedVersion float64

// OperatorConfigurationSetCmdBody holds contents of request body to be sent
var OperatorConfigurationSetCmdBody string

func InitOperatorConfigurationSetCmd() {
	OperatorConfigurationSetCmd.Flags().StringVar(&OperatorConfigurationSetCmdNamespace, "namespace", "", TRAPI("Namespace of operator configuration."))

	OperatorConfigurationSetCmd.Flags().StringVar(&OperatorConfigurationSetCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	OperatorConfigurationSetCmd.Flags().Float64Var(&OperatorConfigurationSetCmdExpectedVersion, "expected-version", 0, TRAPI("This property is used to avoid update confliction. To use it, retrieve the version by getOperatorConfigurationNamespace API, then specify the version here. If the expectedVersion is not specified, the API does not check an update confliction."))

	OperatorConfigurationSetCmd.Flags().StringVar(&OperatorConfigurationSetCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorConfigurationSetCmd.RunE = OperatorConfigurationSetCmdRunE

	OperatorConfigurationCmd.AddCommand(OperatorConfigurationSetCmd)
}

// OperatorConfigurationSetCmd defines 'set' subcommand
var OperatorConfigurationSetCmd = &cobra.Command{
	Use:   "set",
	Short: TRAPI("/operators/{operator_id}/configuration/{namespace}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/configuration/{namespace}:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "setOperatorConfigurationNamespace"),
}

func OperatorConfigurationSetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorConfigurationSetCmdParams(ac)
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

func collectOperatorConfigurationSetCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if OperatorConfigurationSetCmdOperatorId == "" {
		OperatorConfigurationSetCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForOperatorConfigurationSetCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("namespace", "namespace", "path", parsedBody, OperatorConfigurationSetCmdNamespace)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorConfigurationSetCmd("/operators/{operator_id}/configuration/{namespace}"),
		query:       buildQueryForOperatorConfigurationSetCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorConfigurationSetCmd(path string) string {

	escapedNamespace := url.PathEscape(OperatorConfigurationSetCmdNamespace)

	path = strReplace(path, "{"+"namespace"+"}", escapedNamespace, -1)

	escapedOperatorId := url.PathEscape(OperatorConfigurationSetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorConfigurationSetCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorConfigurationSetCmd() (string, error) {
	var result map[string]interface{}

	if OperatorConfigurationSetCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorConfigurationSetCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorConfigurationSetCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if OperatorConfigurationSetCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorConfigurationSetCmdBody)
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

	if OperatorConfigurationSetCmd.Flags().Lookup("expected-version").Changed {
		result["expectedVersion"] = OperatorConfigurationSetCmdExpectedVersion
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
