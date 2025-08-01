// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorDeleteContractCmdContractName holds value of 'contract_name' option
var OperatorDeleteContractCmdContractName string

// OperatorDeleteContractCmdOperatorId holds value of 'operator_id' option
var OperatorDeleteContractCmdOperatorId string

func InitOperatorDeleteContractCmd() {
	OperatorDeleteContractCmd.Flags().StringVar(&OperatorDeleteContractCmdContractName, "contract-name", "", TRAPI("Contract name."))

	OperatorDeleteContractCmd.Flags().StringVar(&OperatorDeleteContractCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	OperatorDeleteContractCmd.RunE = OperatorDeleteContractCmdRunE

	OperatorCmd.AddCommand(OperatorDeleteContractCmd)
}

// OperatorDeleteContractCmd defines 'delete-contract' subcommand
var OperatorDeleteContractCmd = &cobra.Command{
	Use:   "delete-contract",
	Short: TRAPI("/operators/{operator_id}/contracts/{contract_name}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/contracts/{contract_name}:delete:description`) + "\n\n" + createLinkToAPIReference("Operator", "deleteOperatorContract"),
}

func OperatorDeleteContractCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorDeleteContractCmdParams(ac)
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

func collectOperatorDeleteContractCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	if OperatorDeleteContractCmdOperatorId == "" {
		OperatorDeleteContractCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	err = checkIfRequiredStringParameterIsSupplied("contract_name", "contract-name", "path", parsedBody, OperatorDeleteContractCmdContractName)
	if err != nil {
		return nil, err
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
