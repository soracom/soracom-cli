package cmd

import (
	"github.com/pkg/errors"
	"github.com/soracom/soracom-cli/generators/lib"
)

type APICredentialsSource interface {
	GetAPICredentials(ac *apiClient) (*APICredentials, error)
}

var (
	apiCredentialsSources = []APICredentialsSource{
		newCredentialsSourceWithAPIKeyAndToken(),
		newCredentialsSourceWithAuthKeyIDAndAuthKey(),
		newCredentialsSourceWithProfileCommand(),
		newCredentialsSourceWithProfile(),
	}
)

type credentialsSourceWithAPIKeyAndToken struct {
}

func newCredentialsSourceWithAPIKeyAndToken() *credentialsSourceWithAPIKeyAndToken {
	return &credentialsSourceWithAPIKeyAndToken{}
}

func (s *credentialsSourceWithAPIKeyAndToken) GetAPICredentials(_ *apiClient) (*APICredentials, error) {
	if providedAPIKey == "" && providedAPIToken == "" {
		return nil, nil
	}

	if (providedAPIKey != "" && providedAPIToken == "") || (providedAPIKey == "" && providedAPIToken != "") {
		return nil, errors.New("both --api-key and --api-token must be specified")
	}

	return &APICredentials{
		APIKey:   providedAPIKey,
		APIToken: providedAPIToken,
	}, nil
}

type credentialsSourceWithAuthKeyIDAndAuthKey struct {
}

func newCredentialsSourceWithAuthKeyIDAndAuthKey() *credentialsSourceWithAuthKeyIDAndAuthKey {
	return &credentialsSourceWithAuthKeyIDAndAuthKey{}
}

func (s *credentialsSourceWithAuthKeyIDAndAuthKey) GetAPICredentials(ac *apiClient) (*APICredentials, error) {
	if providedAuthKeyID == "" && providedAuthKey == "" {
		return nil, nil
	}

	if (providedAuthKeyID != "" && providedAuthKey == "") || (providedAuthKeyID == "" && providedAuthKey != "") {
		return nil, errors.New("both --auth-key-id and --auth-key must be specified")
	}

	areq := &authRequest{
		AuthKeyID: &providedAuthKeyID,
		AuthKey:   &providedAuthKey,
	}
	ares, err := ac.authenticate(areq)
	if err != nil {
		return nil, err
	}

	return apiCredentialsFromAuthResult(ares), nil
}

type credentialsSourceWithProfileCommand struct {
}

func newCredentialsSourceWithProfileCommand() *credentialsSourceWithProfileCommand {
	return &credentialsSourceWithProfileCommand{}
}

func (s *credentialsSourceWithProfileCommand) GetAPICredentials(ac *apiClient) (*APICredentials, error) {
	if providedProfileCommand == "" {
		return nil, nil
	}

	p, err := getProfileFromExternalCommand(providedProfileCommand)
	if err != nil {
		return nil, err
	}

	if p.SourceProfile != nil && *p.SourceProfile != "" {
		sourceProfile, err := loadProfile(*p.SourceProfile)
		if err != nil {
			lib.PrintfStderr("unable to load the specified source profile: %s\n", *p.SourceProfile)
			return nil, err
		}
		ares, err := ac.authenticateWithSwitchUser(p, sourceProfile)
		if err != nil {
			return nil, err
		}
		return apiCredentialsFromAuthResult(ares), nil
	}

	areq := authRequestFromProfile(p)
	ares, err := ac.authenticate(areq)
	if err != nil {
		return nil, err
	}

	return apiCredentialsFromAuthResult(ares), nil
}

type credentialsSourceWithProfile struct {
}

func newCredentialsSourceWithProfile() *credentialsSourceWithProfile {
	return &credentialsSourceWithProfile{}
}

func (s *credentialsSourceWithProfile) GetAPICredentials(ac *apiClient) (*APICredentials, error) {
	p, err := getProfile()
	if err != nil {
		return nil, err
	}

	if p.ProfileCommand != nil && *p.ProfileCommand != "" {
		p, err = getProfileFromExternalCommand(*p.ProfileCommand)
		if err != nil {
			return nil, err
		}
	}

	if p.SourceProfile != nil && *p.SourceProfile != "" {
		sourceProfile, err := loadProfile(*p.SourceProfile)
		if err != nil {
			lib.PrintfStderr("unable to load the specified source profile: %s\n", *p.SourceProfile)
			return nil, err
		}
		ares, err := ac.authenticateWithSwitchUser(p, sourceProfile)
		if err != nil {
			return nil, err
		}
		return apiCredentialsFromAuthResult(ares), nil
	}

	areq := authRequestFromProfile(p)
	ares, err := ac.authenticate(areq)
	if err != nil {
		return nil, err
	}

	return apiCredentialsFromAuthResult(ares), nil
}
