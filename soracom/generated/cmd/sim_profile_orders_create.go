// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SimProfileOrdersCreateCmdDescription holds value of 'description' option
var SimProfileOrdersCreateCmdDescription string

// SimProfileOrdersCreateCmdSubscription holds value of 'subscription' option
var SimProfileOrdersCreateCmdSubscription string

// SimProfileOrdersCreateCmdQuantity holds value of 'quantity' option
var SimProfileOrdersCreateCmdQuantity int64

// SimProfileOrdersCreateCmdBody holds contents of request body to be sent
var SimProfileOrdersCreateCmdBody string

func InitSimProfileOrdersCreateCmd() {
	SimProfileOrdersCreateCmd.Flags().StringVar(&SimProfileOrdersCreateCmdDescription, "description", "", TRAPI("The description of the order. You can use this to identify the order in the order history."))

	SimProfileOrdersCreateCmd.Flags().StringVar(&SimProfileOrdersCreateCmdSubscription, "subscription", "", TRAPI("The subscription of the eSIM profile to be ordered.- 'plan01s'- 'planP1'- 'planX1'"))

	SimProfileOrdersCreateCmd.Flags().Int64Var(&SimProfileOrdersCreateCmdQuantity, "quantity", 0, TRAPI("The quantity of the eSIM profile to be ordered."))

	SimProfileOrdersCreateCmd.Flags().StringVar(&SimProfileOrdersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SimProfileOrdersCreateCmd.RunE = SimProfileOrdersCreateCmdRunE

	SimProfileOrdersCmd.AddCommand(SimProfileOrdersCreateCmd)
}

// SimProfileOrdersCreateCmd defines 'create' subcommand
var SimProfileOrdersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sim_profile_orders:post:summary"),
	Long:  TRAPI(`/sim_profile_orders:post:description`) + "\n\n" + createLinkToAPIReference("SimProfileOrder", "createProfileOrder"),
}

func SimProfileOrdersCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimProfileOrdersCreateCmdParams(ac)
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

func collectSimProfileOrdersCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimProfileOrdersCreateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("subscription", "subscription", "body", parsedBody, SimProfileOrdersCreateCmdSubscription)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("quantity", "quantity", "body", parsedBody, SimProfileOrdersCreateCmdQuantity)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimProfileOrdersCreateCmd("/sim_profile_orders"),
		query:       buildQueryForSimProfileOrdersCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimProfileOrdersCreateCmd(path string) string {

	return path
}

func buildQueryForSimProfileOrdersCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimProfileOrdersCreateCmd() (string, error) {
	var result map[string]interface{}

	if SimProfileOrdersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimProfileOrdersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SimProfileOrdersCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SimProfileOrdersCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SimProfileOrdersCreateCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if SimProfileOrdersCreateCmdDescription != "" {
		result["description"] = SimProfileOrdersCreateCmdDescription
	}

	if SimProfileOrdersCreateCmdSubscription != "" {
		result["subscription"] = SimProfileOrdersCreateCmdSubscription
	}

	if SimProfileOrdersCreateCmd.Flags().Lookup("quantity").Changed {
		result["quantity"] = SimProfileOrdersCreateCmdQuantity
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
