// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// CouponsCreateCmdAmount holds value of 'amount' option
var CouponsCreateCmdAmount float64

// CouponsCreateCmdBody holds contents of request body to be sent
var CouponsCreateCmdBody string

func init() {
	CouponsCreateCmd.Flags().Float64Var(&CouponsCreateCmdAmount, "amount", 0, TRAPI("Amount"))

	CouponsCreateCmd.Flags().StringVar(&CouponsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	CouponsCmd.AddCommand(CouponsCreateCmd)
}

// CouponsCreateCmd defines 'create' subcommand
var CouponsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/coupons:post:summary"),
	Long:  TRAPI(`/coupons:post:description`) + "\n\n" + createLinkToAPIReference("Order", "createCouponQuotation"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectCouponsCreateCmdParams(ac)
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

func collectCouponsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForCouponsCreateCmd()
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

	err = checkIfRequiredFloatParameterIsSupplied("amount", "amount", "body", parsedBody, CouponsCreateCmdAmount)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForCouponsCreateCmd("/coupons"),
		query:       buildQueryForCouponsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCouponsCreateCmd(path string) string {

	return path
}

func buildQueryForCouponsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForCouponsCreateCmd() (string, error) {
	var result map[string]interface{}

	if CouponsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(CouponsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(CouponsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if CouponsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(CouponsCreateCmdBody)
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

	if CouponsCreateCmdAmount != 0 {
		result["amount"] = CouponsCreateCmdAmount
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
