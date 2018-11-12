package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var ConfigureSandboxAuthKeyID string
var ConfigureSandboxAuthKey string
var ConfigureSandboxEmail string
var ConfigureSandboxPassword string
var ConfigureSandboxRegisterPaymentMethod bool

func init() {
	ConfigureSandboxCmd.Flags().StringVar(&ConfigureSandboxAuthKeyID, "auth-key-id", "", TRCLI("cli.configure_sandbox.auth_key_id"))
	ConfigureSandboxCmd.Flags().StringVar(&ConfigureSandboxAuthKey, "auth-key", "", TRCLI("cli.configure_sandbox.auth_key"))
	ConfigureSandboxCmd.Flags().StringVar(&ConfigureSandboxEmail, "email", "", TRCLI("cli.configure_sandbox.email"))
	ConfigureSandboxCmd.Flags().StringVar(&ConfigureSandboxPassword, "password", "", TRCLI("cli.configure_sandbox.password"))
	ConfigureSandboxCmd.Flags().BoolVar(&ConfigureSandboxRegisterPaymentMethod, "register-payment-method", false, TRCLI("cli.configure_sandbox.register_payment_method"))
	RootCmd.AddCommand(ConfigureSandboxCmd)
}

// ConfigureSandboxCmd defines 'configure-sandbox' subcommand
var ConfigureSandboxCmd = &cobra.Command{
	Use:   "configure-sandbox",
	Short: TRCLI("cli.configure_sandbox.summary"),
	Long:  TRCLI("cli.configure_sandbox.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pn := getSpecifiedSandboxProfileName()

		var p *profile
		var err error
		if ConfigureSandboxAuthKeyID == "" || ConfigureSandboxAuthKey == "" || ConfigureSandboxEmail == "" || ConfigureSandboxPassword == "" {
			p, err = collectSandboxProfileInfo(pn)
			if err != nil {
				cmd.SilenceUsage = true
				return err
			}
			ep := getSpecifiedSandboxEndpoint(p.CoverageType)
			p.Endpoint = &ep
		} else {
			ct := getSpecifiedCoverageType()
			ep := getSpecifiedSandboxEndpoint(ct)
			p = &profile{
				Sandbox:               true,
				CoverageType:          ct,
				Endpoint:              &ep,
				AuthKeyID:             &ConfigureSandboxAuthKeyID,
				AuthKey:               &ConfigureSandboxAuthKey,
				Email:                 &ConfigureSandboxEmail,
				Password:              &ConfigureSandboxPassword,
				RegisterPaymentMethod: ConfigureSandboxRegisterPaymentMethod,
			}
		}

		_, err = sandboxInit(p)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// AuthKeyID and AuthKey are not needed to be saved to a file. They are needed only when creating an account on sandbox
		p.AuthKeyID = nil
		p.AuthKey = nil

		err = saveProfile(pn, p)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}

func sandboxInit(profile *profile) (*authResult, error) {
	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Endpoint: *profile.Endpoint,
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}

	reqBodyBytes, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}

	param := &apiParams{
		method:         "POST",
		path:           "/sandbox/init",
		contentType:    "application/json",
		noVersionCheck: true,
		body:           string(reqBodyBytes),
	}

	_, respBody, err := ac.callAPI(param)

	var ar authResult
	err = json.Unmarshal([]byte(respBody), &ar)
	return &ar, err
}
