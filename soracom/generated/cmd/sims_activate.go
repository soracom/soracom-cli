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

// SimsActivateCmdSimId holds value of 'sim_id' option
var SimsActivateCmdSimId string

func init() {
	SimsActivateCmd.Flags().StringVar(&SimsActivateCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsActivateCmd)
}

// SimsActivateCmd defines 'activate' subcommand
var SimsActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: TRAPI("/sims/{sim_id}/activate:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/activate:post:description`),
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

		param, err := collectSimsActivateCmdParams(ac)
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

func collectSimsActivateCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsActivateCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsActivateCmd("/sims/{sim_id}/activate"),
		query:  buildQueryForSimsActivateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsActivateCmd(path string) string {

	escapedSimId := url.PathEscape(SimsActivateCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsActivateCmd() url.Values {
	result := url.Values{}

	return result
}
