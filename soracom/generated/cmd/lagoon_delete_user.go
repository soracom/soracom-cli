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

// LagoonDeleteUserCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonDeleteUserCmdLagoonUserId int64

func init() {
	LagoonDeleteUserCmd.Flags().Int64Var(&LagoonDeleteUserCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))
	LagoonCmd.AddCommand(LagoonDeleteUserCmd)
}

// LagoonDeleteUserCmd defines 'delete-user' subcommand
var LagoonDeleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}:delete:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}:delete:description`),
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

		param, err := collectLagoonDeleteUserCmdParams(ac)
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

func collectLagoonDeleteUserCmdParams(ac *apiClient) (*apiParams, error) {
	if LagoonDeleteUserCmdLagoonUserId == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "lagoon-user-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLagoonDeleteUserCmd("/lagoon/users/{lagoon_user_id}"),
		query:  buildQueryForLagoonDeleteUserCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonDeleteUserCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonDeleteUserCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonDeleteUserCmd() url.Values {
	result := url.Values{}

	return result
}
