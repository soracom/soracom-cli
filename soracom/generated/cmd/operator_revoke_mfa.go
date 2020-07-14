// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorRevokeMfaCmdOperatorId holds value of 'operator_id' option
var OperatorRevokeMfaCmdOperatorId string

func init() {
	OperatorRevokeMfaCmd.Flags().StringVar(&OperatorRevokeMfaCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorRevokeMfaCmd)
}

// OperatorRevokeMfaCmd defines 'revoke-mfa' subcommand
var OperatorRevokeMfaCmd = &cobra.Command{
	Use:   "revoke-mfa",
	Short: TRAPI("/operators/{operator_id}/mfa:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/mfa:delete:description`),
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

		param, err := collectOperatorRevokeMfaCmdParams(ac)
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

func collectOperatorRevokeMfaCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorRevokeMfaCmdOperatorId == "" {
		OperatorRevokeMfaCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForOperatorRevokeMfaCmd("/operators/{operator_id}/mfa"),
		query:  buildQueryForOperatorRevokeMfaCmd(),
	}, nil
}

func buildPathForOperatorRevokeMfaCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorRevokeMfaCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorRevokeMfaCmd() url.Values {
	result := url.Values{}

	return result
}
