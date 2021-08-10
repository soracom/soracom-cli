// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// CredentialsDeleteCmdCredentialsId holds value of 'credentials_id' option
var CredentialsDeleteCmdCredentialsId string

func init() {
	CredentialsDeleteCmd.Flags().StringVar(&CredentialsDeleteCmdCredentialsId, "credentials-id", "", TRAPI("Credentials ID"))
	CredentialsCmd.AddCommand(CredentialsDeleteCmd)
}

// CredentialsDeleteCmd defines 'delete' subcommand
var CredentialsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/credentials/{credentials_id}:delete:summary"),
	Long:  TRAPI(`/credentials/{credentials_id}:delete:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectCredentialsDeleteCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectCredentialsDeleteCmdParams(ac *apiClient) (*apiParams, error) {
	if CredentialsDeleteCmdCredentialsId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "credentials-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForCredentialsDeleteCmd("/credentials/{credentials_id}"),
		query:  buildQueryForCredentialsDeleteCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCredentialsDeleteCmd(path string) string {

	escapedCredentialsId := url.PathEscape(CredentialsDeleteCmdCredentialsId)

	path = strReplace(path, "{"+"credentials_id"+"}", escapedCredentialsId, -1)

	return path
}

func buildQueryForCredentialsDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
