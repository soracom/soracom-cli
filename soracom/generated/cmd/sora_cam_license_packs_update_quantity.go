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

// SoraCamLicensePacksUpdateQuantityCmdLicensePackId holds value of 'license_pack_id' option
var SoraCamLicensePacksUpdateQuantityCmdLicensePackId string

// SoraCamLicensePacksUpdateQuantityCmdCurrentQuantity holds value of 'currentQuantity' option
var SoraCamLicensePacksUpdateQuantityCmdCurrentQuantity int64

// SoraCamLicensePacksUpdateQuantityCmdDesiredQuantity holds value of 'desiredQuantity' option
var SoraCamLicensePacksUpdateQuantityCmdDesiredQuantity int64

// SoraCamLicensePacksUpdateQuantityCmdBody holds contents of request body to be sent
var SoraCamLicensePacksUpdateQuantityCmdBody string

func init() {
	SoraCamLicensePacksUpdateQuantityCmd.Flags().StringVar(&SoraCamLicensePacksUpdateQuantityCmdLicensePackId, "license-pack-id", "", TRAPI("ID of the license pack"))

	SoraCamLicensePacksUpdateQuantityCmd.Flags().Int64Var(&SoraCamLicensePacksUpdateQuantityCmdCurrentQuantity, "current-quantity", 0, TRAPI("Current license quantity of the license pack"))

	SoraCamLicensePacksUpdateQuantityCmd.Flags().Int64Var(&SoraCamLicensePacksUpdateQuantityCmdDesiredQuantity, "desired-quantity", 0, TRAPI("Desired license quantity of the license pack"))

	SoraCamLicensePacksUpdateQuantityCmd.Flags().StringVar(&SoraCamLicensePacksUpdateQuantityCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SoraCamLicensePacksCmd.AddCommand(SoraCamLicensePacksUpdateQuantityCmd)
}

// SoraCamLicensePacksUpdateQuantityCmd defines 'update-quantity' subcommand
var SoraCamLicensePacksUpdateQuantityCmd = &cobra.Command{
	Use:   "update-quantity",
	Short: TRAPI("/sora_cam/license_packs/{license_pack_id}/quantity:put:summary"),
	Long:  TRAPI(`/sora_cam/license_packs/{license_pack_id}/quantity:put:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "updateSoraCamLicensePackQuantity"),
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

		param, err := collectSoraCamLicensePacksUpdateQuantityCmdParams(ac)
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

func collectSoraCamLicensePacksUpdateQuantityCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamLicensePacksUpdateQuantityCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("license_pack_id", "license-pack-id", "path", parsedBody, SoraCamLicensePacksUpdateQuantityCmdLicensePackId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSoraCamLicensePacksUpdateQuantityCmd("/sora_cam/license_packs/{license_pack_id}/quantity"),
		query:       buildQueryForSoraCamLicensePacksUpdateQuantityCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamLicensePacksUpdateQuantityCmd(path string) string {

	escapedLicensePackId := url.PathEscape(SoraCamLicensePacksUpdateQuantityCmdLicensePackId)

	path = strReplace(path, "{"+"license_pack_id"+"}", escapedLicensePackId, -1)

	return path
}

func buildQueryForSoraCamLicensePacksUpdateQuantityCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamLicensePacksUpdateQuantityCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamLicensePacksUpdateQuantityCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamLicensePacksUpdateQuantityCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamLicensePacksUpdateQuantityCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SoraCamLicensePacksUpdateQuantityCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamLicensePacksUpdateQuantityCmdBody)
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

	if SoraCamLicensePacksUpdateQuantityCmdCurrentQuantity != 0 {
		result["currentQuantity"] = SoraCamLicensePacksUpdateQuantityCmdCurrentQuantity
	}

	if SoraCamLicensePacksUpdateQuantityCmdDesiredQuantity != 0 {
		result["desiredQuantity"] = SoraCamLicensePacksUpdateQuantityCmdDesiredQuantity
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}