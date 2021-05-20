package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var configureSandboxAuthKeyID string
var configureSandboxAuthKey string
var configureSandboxEmail string
var configureSandboxPassword string
var configureSandboxRegisterPaymentMethod bool
var configureSandboxOverwrite bool

func init() {
	ConfigureSandboxCmd.Flags().StringVar(&configureSandboxAuthKeyID, "auth-key-id", "", TRCLI("cli.configure_sandbox.auth_key_id"))
	ConfigureSandboxCmd.Flags().StringVar(&configureSandboxAuthKey, "auth-key", "", TRCLI("cli.configure_sandbox.auth_key"))
	ConfigureSandboxCmd.Flags().StringVar(&configureSandboxEmail, "email", "", TRCLI("cli.configure_sandbox.email"))
	ConfigureSandboxCmd.Flags().StringVar(&configureSandboxPassword, "password", "", TRCLI("cli.configure_sandbox.password"))
	ConfigureSandboxCmd.Flags().BoolVar(&configureSandboxRegisterPaymentMethod, "register-payment-method", true, TRCLI("cli.configure_sandbox.register_payment_method"))
	ConfigureSandboxCmd.Flags().BoolVar(&configureSandboxOverwrite, "overwrite", true, TRCLI("cli.configure_sandbox.overwrite"))
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
		if configureSandboxAuthKeyID == "" || configureSandboxAuthKey == "" || configureSandboxEmail == "" || configureSandboxPassword == "" {
			p, err = collectSandboxProfileInfo(pn, configureSandboxRegisterPaymentMethod)
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
				AuthKeyID:             &configureSandboxAuthKeyID,
				AuthKey:               &configureSandboxAuthKey,
				Email:                 &configureSandboxEmail,
				Password:              &configureSandboxPassword,
				RegisterPaymentMethod: configureSandboxRegisterPaymentMethod,
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

		err = saveProfile(pn, p, configureSandboxOverwrite)
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

	respBody, err := ac.callAPI(param)
	if err != nil {
		return nil, err
	}

	var ar authResult
	err = json.Unmarshal([]byte(respBody), &ar)
	if err != nil {
		return nil, err
	}

	return &ar, err
}
