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

// SimsDisableTerminationCmdSimId holds value of 'sim_id' option
var SimsDisableTerminationCmdSimId string

func init() {
	SimsDisableTerminationCmd.Flags().StringVar(&SimsDisableTerminationCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsDisableTerminationCmd)
}

// SimsDisableTerminationCmd defines 'disable-termination' subcommand
var SimsDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TRAPI("/sims/{sim_id}/disable_termination:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/disable_termination:post:description`),
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

		param, err := collectSimsDisableTerminationCmdParams(ac)
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

func collectSimsDisableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsDisableTerminationCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsDisableTerminationCmd("/sims/{sim_id}/disable_termination"),
		query:  buildQueryForSimsDisableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsDisableTerminationCmd(path string) string {

	escapedSimId := url.PathEscape(SimsDisableTerminationCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsDisableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
