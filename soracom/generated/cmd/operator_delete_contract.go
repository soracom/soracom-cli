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

// OperatorDeleteContractCmdContractName holds value of 'contract_name' option
var OperatorDeleteContractCmdContractName string

// OperatorDeleteContractCmdOperatorId holds value of 'operator_id' option
var OperatorDeleteContractCmdOperatorId string

func init() {
	OperatorDeleteContractCmd.Flags().StringVar(&OperatorDeleteContractCmdContractName, "contract-name", "", TRAPI("contract_name"))

	OperatorDeleteContractCmd.Flags().StringVar(&OperatorDeleteContractCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorDeleteContractCmd)
}

// OperatorDeleteContractCmd defines 'delete-contract' subcommand
var OperatorDeleteContractCmd = &cobra.Command{
	Use:   "delete-contract",
	Short: TRAPI("/operators/{operator_id}/contracts/{contract_name}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/contracts/{contract_name}:delete:description`),
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

		param, err := collectOperatorDeleteContractCmdParams(ac)
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

func collectOperatorDeleteContractCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorDeleteContractCmdOperatorId == "" {
		OperatorDeleteContractCmdOperatorId = ac.OperatorID
	}

	if OperatorDeleteContractCmdContractName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "contract-name")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForOperatorDeleteContractCmd("/operators/{operator_id}/contracts/{contract_name}"),
		query:  buildQueryForOperatorDeleteContractCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorDeleteContractCmd(path string) string {

	escapedContractName := url.PathEscape(OperatorDeleteContractCmdContractName)

	path = strReplace(path, "{"+"contract_name"+"}", escapedContractName, -1)

	escapedOperatorId := url.PathEscape(OperatorDeleteContractCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorDeleteContractCmd() url.Values {
	result := url.Values{}

	return result
}
