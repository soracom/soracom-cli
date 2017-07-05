package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgSetRedirectionCmdDescription holds value of 'description' option
var VpgSetRedirectionCmdDescription string

// VpgSetRedirectionCmdGateway holds value of 'gateway' option
var VpgSetRedirectionCmdGateway string

// VpgSetRedirectionCmdId holds value of 'id' option
var VpgSetRedirectionCmdId string

// VpgSetRedirectionCmdEnabled holds value of 'enabled' option
var VpgSetRedirectionCmdEnabled bool

// VpgSetRedirectionCmdBody holds contents of request body to be sent
var VpgSetRedirectionCmdBody string

func init() {
	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdDescription, "description", "", TR(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdGateway, "gateway", "", TR(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdId, "id", "", TR("virtual_private_gateway.junction.redirection.set.post.parameters.id.description"))

	VpgSetRedirectionCmd.Flags().BoolVar(&VpgSetRedirectionCmdEnabled, "enabled", false, TR(""))

	VpgSetRedirectionCmd.Flags().StringVar(&VpgSetRedirectionCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgSetRedirectionCmd)
}

// VpgSetRedirectionCmd defines 'set-redirection' subcommand
var VpgSetRedirectionCmd = &cobra.Command{
	Use:   "set-redirection",
	Short: TR("virtual_private_gateway.junction.redirection.set.post.summary"),
	Long:  TR(`virtual_private_gateway.junction.redirection.set.post.description`),
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

		param, err := collectVpgSetRedirectionCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectVpgSetRedirectionCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgSetRedirectionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgSetRedirectionCmd("/virtual_private_gateways/{id}/junction/set_redirection"),
		query:       buildQueryForVpgSetRedirectionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgSetRedirectionCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgSetRedirectionCmdId, -1)

	return path
}

func buildQueryForVpgSetRedirectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgSetRedirectionCmd() (string, error) {
	if VpgSetRedirectionCmdBody != "" {
		if strings.HasPrefix(VpgSetRedirectionCmdBody, "@") {
			fname := strings.TrimPrefix(VpgSetRedirectionCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgSetRedirectionCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgSetRedirectionCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgSetRedirectionCmdDescription != "" {
		result["description"] = VpgSetRedirectionCmdDescription
	}

	if VpgSetRedirectionCmdGateway != "" {
		result["gateway"] = VpgSetRedirectionCmdGateway
	}

	if VpgSetRedirectionCmdEnabled != false {
		result["enabled"] = VpgSetRedirectionCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
