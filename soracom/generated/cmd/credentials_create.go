package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CredentialsCreateCmdCredentialsId holds value of 'credentials_id' option
var CredentialsCreateCmdCredentialsId string

// CredentialsCreateCmdDescription holds value of 'description' option
var CredentialsCreateCmdDescription string

// CredentialsCreateCmdType holds value of 'type' option
var CredentialsCreateCmdType string

// CredentialsCreateCmdBody holds contents of request body to be sent
var CredentialsCreateCmdBody string

func init() {
	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdCredentialsId, "credentials-id", "", TRAPI("credentials_id"))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdDescription, "description", "", TRAPI(""))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdType, "type", "", TRAPI(""))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	CredentialsCmd.AddCommand(CredentialsCreateCmd)
}

// CredentialsCreateCmd defines 'create' subcommand
var CredentialsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/credentials/{credentials_id}:post:summary"),
	Long:  TRAPI(`/credentials/{credentials_id}:post:description`),
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

		param, err := collectCredentialsCreateCmdParams(ac)
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

func collectCredentialsCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForCredentialsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForCredentialsCreateCmd("/credentials/{credentials_id}"),
		query:       buildQueryForCredentialsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForCredentialsCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"credentials_id"+"}", CredentialsCreateCmdCredentialsId, -1)

	return path
}

func buildQueryForCredentialsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForCredentialsCreateCmd() (string, error) {
	var result map[string]interface{}

	if CredentialsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(CredentialsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(CredentialsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if CredentialsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(CredentialsCreateCmdBody)
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

	if CredentialsCreateCmdDescription != "" {
		result["description"] = CredentialsCreateCmdDescription
	}

	if CredentialsCreateCmdType != "" {
		result["type"] = CredentialsCreateCmdType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
