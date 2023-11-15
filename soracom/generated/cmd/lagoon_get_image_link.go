// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LagoonGetImageLinkCmdV2 holds value of 'v2' option
var LagoonGetImageLinkCmdV2 bool

func InitLagoonGetImageLinkCmd() {
	LagoonGetImageLinkCmd.Flags().BoolVar(&LagoonGetImageLinkCmdV2, "v2", false, TRAPI("If the value is true, a request will be issued to Lagoon 2.This is only valid if both Lagoon 3 and Lagoon 2 are enabled."))

	LagoonGetImageLinkCmd.RunE = LagoonGetImageLinkCmdRunE

	LagoonCmd.AddCommand(LagoonGetImageLinkCmd)
}

// LagoonGetImageLinkCmd defines 'get-image-link' subcommand
var LagoonGetImageLinkCmd = &cobra.Command{
	Use:   "get-image-link",
	Short: TRAPI("/lagoon/image/link:get:summary"),
	Long:  TRAPI(`/lagoon/image/link:get:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "getImageLink"),
}

func LagoonGetImageLinkCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
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

	param, err := collectLagoonGetImageLinkCmdParams(ac)
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

func collectLagoonGetImageLinkCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLagoonGetImageLinkCmd("/lagoon/image/link"),
		query:  buildQueryForLagoonGetImageLinkCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonGetImageLinkCmd(path string) string {

	return path
}

func buildQueryForLagoonGetImageLinkCmd() url.Values {
	result := url.Values{}

	if LagoonGetImageLinkCmdV2 != false {
		result.Add("v2", sprintf("%t", LagoonGetImageLinkCmdV2))
	}

	return result
}
