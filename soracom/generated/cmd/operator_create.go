package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

var OperatorCreateCmdEmail string

var OperatorCreateCmdPassword string

var OperatorCreateCmdBody string

func init() {
	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdEmail, "email", "", TR(""))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdPassword, "password", "", TR(""))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorCreateCmd)
}

var OperatorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("operator.create_operator.post.summary"),
	Long:  TR(`operator.create_operator.post.description`),
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

		param, err := collectOperatorCreateCmdParams()
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

func collectOperatorCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForOperatorCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorCreateCmd("/operators"),
		query:       buildQueryForOperatorCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorCreateCmd(path string) string {

	return path
}

func buildQueryForOperatorCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorCreateCmd() (string, error) {
	if OperatorCreateCmdBody != "" {
		if strings.HasPrefix(OperatorCreateCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if OperatorCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return OperatorCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if OperatorCreateCmdEmail != "" {
		result["email"] = OperatorCreateCmdEmail
	}

	if OperatorCreateCmdPassword != "" {
		result["password"] = OperatorCreateCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
