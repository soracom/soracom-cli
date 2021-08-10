// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorAddCoverageTypeCmdCoverageType holds value of 'coverage_type' option
var OperatorAddCoverageTypeCmdCoverageType string

// OperatorAddCoverageTypeCmdOperatorId holds value of 'operator_id' option
var OperatorAddCoverageTypeCmdOperatorId string

func init() {
	OperatorAddCoverageTypeCmd.Flags().StringVar(&OperatorAddCoverageTypeCmdCoverageType, "coverage-type", "", TRAPI("coverage_type"))

	OperatorAddCoverageTypeCmd.Flags().StringVar(&OperatorAddCoverageTypeCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorAddCoverageTypeCmd)
}

// OperatorAddCoverageTypeCmd defines 'add-coverage-type' subcommand
var OperatorAddCoverageTypeCmd = &cobra.Command{
	Use:   "add-coverage-type",
	Short: TRAPI("/operators/{operator_id}/coverage_type/{coverage_type}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/coverage_type/{coverage_type}:post:description`),
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

		param, err := collectOperatorAddCoverageTypeCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectOperatorAddCoverageTypeCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorAddCoverageTypeCmdOperatorId == "" {
		OperatorAddCoverageTypeCmdOperatorId = ac.OperatorID
	}

	if OperatorAddCoverageTypeCmdCoverageType == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "coverage-type")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorAddCoverageTypeCmd("/operators/{operator_id}/coverage_type/{coverage_type}"),
		query:  buildQueryForOperatorAddCoverageTypeCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorAddCoverageTypeCmd(path string) string {

	escapedCoverageType := url.PathEscape(OperatorAddCoverageTypeCmdCoverageType)

	path = strReplace(path, "{"+"coverage_type"+"}", escapedCoverageType, -1)

	escapedOperatorId := url.PathEscape(OperatorAddCoverageTypeCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorAddCoverageTypeCmd() url.Values {
	result := url.Values{}

	return result
}
