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

// PortMappingsCreateCmdDuration holds value of 'duration' option
var PortMappingsCreateCmdDuration float64

// PortMappingsCreateCmdTlsRequired holds value of 'tlsRequired' option
var PortMappingsCreateCmdTlsRequired bool

// PortMappingsCreateCmdBody holds contents of request body to be sent
var PortMappingsCreateCmdBody string

func InitPortMappingsCreateCmd() {
	PortMappingsCreateCmd.Flags().Float64Var(&PortMappingsCreateCmdDuration, "duration", 0, TRAPI("The duration (in seconds) to maintain the On-Demand Remote Access entry (the time to allow remote access). After the specified time has passed, the On-Demand Remote Access entry will be automatically deleted. The maximum duration is 8 hours."))

	PortMappingsCreateCmd.Flags().BoolVar(&PortMappingsCreateCmdTlsRequired, "tls-required", false, TRAPI("Whether to encrypt the connection from the source to Soracom using TLS.- 'true': Encrypt using TLS. Specify this if the device is listening on HTTP.- 'false': Do not use TLS. Specify this when connecting to the device via SSH or if the device is listening on HTTPS. Note that communication from the source to the device is encrypted in SSH and HTTPS."))

	PortMappingsCreateCmd.Flags().StringVar(&PortMappingsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	PortMappingsCreateCmd.RunE = PortMappingsCreateCmdRunE

	PortMappingsCmd.AddCommand(PortMappingsCreateCmd)
}

// PortMappingsCreateCmd defines 'create' subcommand
var PortMappingsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/port_mappings:post:summary"),
	Long:  TRAPI(`/port_mappings:post:description`) + "\n\n" + createLinkToAPIReference("PortMapping", "createPortMapping"),
}

func PortMappingsCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectPortMappingsCreateCmdParams(ac)
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

func collectPortMappingsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForPortMappingsCreateCmd()
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

	return &apiParams{
		method:      "POST",
		path:        buildPathForPortMappingsCreateCmd("/port_mappings"),
		query:       buildQueryForPortMappingsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPortMappingsCreateCmd(path string) string {

	return path
}

func buildQueryForPortMappingsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForPortMappingsCreateCmd() (string, error) {
	var result map[string]interface{}

	if PortMappingsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(PortMappingsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(PortMappingsCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if PortMappingsCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(PortMappingsCreateCmdBody)
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

	if PortMappingsCreateCmd.Flags().Lookup("duration").Changed {
		result["duration"] = PortMappingsCreateCmdDuration
	}

	if PortMappingsCreateCmd.Flags().Lookup("tls-required").Changed {
		result["tlsRequired"] = PortMappingsCreateCmdTlsRequired
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
