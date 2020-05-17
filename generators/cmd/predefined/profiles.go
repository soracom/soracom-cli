package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/kennygrant/sanitize"
	"github.com/mitchellh/go-homedir"
	"github.com/soracom/soracom-cli/generators/lib"
	"golang.org/x/crypto/ssh/terminal"
)

type profile struct {
	Sandbox               bool    `json:"sandbox"`
	CoverageType          string  `json:"coverageType"`
	Email                 *string `json:"email,omitempty"`
	Password              *string `json:"password,omitempty"`
	AuthKeyID             *string `json:"authKeyId,omitempty"`
	AuthKey               *string `json:"authKey,omitempty"`
	Username              *string `json:"username,omitempty"`
	OperatorID            *string `json:"operatorId,omitempty"`
	Endpoint              *string `json:"endpoint,omitempty"`
	RegisterPaymentMethod bool    `json:"registerPaymentMethod"`
}

type authInfo struct {
	Email      *string
	Password   *string
	AuthKeyID  *string
	AuthKey    *string
	Username   *string
	OperatorID *string
}

var (
	loadedProfile *profile
)

func getProfile() (*profile, error) {
	if loadedProfile != nil {
		return loadedProfile, nil
	}

	pn := getSpecifiedProfileName()
	if pn == "" {
		pn = "default"
	}

	profile, err := loadProfile(pn)
	if err != nil {
		return nil, err
	}

	loadedProfile = profile
	return loadedProfile, nil
}

func getDefaultProfileName() string {
	return "default"
}

func getDefaultSandboxProfileName() string {
	return "sandbox"
}

func getDefaultCoverageType() string {
	return "g"
}

func getProfileDir() (string, error) {
	profDir := os.Getenv("SORACOM_PROFILE_DIR")
	if profDir == "" {
		dir, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		profDir = filepath.Join(dir, ".soracom")
	}

	return profDir, nil
}

func getSpecifiedProfileName() string {
	if specifiedProfileName == "" {
		return getDefaultProfileName()
	}
	return sanitize.BaseName(specifiedProfileName)
}

func getSpecifiedSandboxProfileName() string {
	if specifiedProfileName == "" {
		return getDefaultSandboxProfileName()
	}
	return sanitize.BaseName(specifiedProfileName)
}

func getSpecifiedCoverageType() string {
	return specifiedCoverageType
}

func loadProfile(profileName string) (*profile, error) {
	dir, err := getProfileDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, profileName+".json")

	// check if permission is less than 0600
	tooOpen, err := lib.IsFilePermissionTooOpen(path)
	if err != nil {
		return nil, err
	}
	if tooOpen {
		msg := fmt.Sprintf(TRCLI("cli.configure.profile.permission_is_too_open"), path)
		if runtime.GOOS != "windows" {
			return nil, errors.New(msg)
		}
		// only warn on windows
		printfStderr("WARN: " + msg + "\n")
	}

	// #nosec
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p profile
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}

	// supply default values for older versions (which support 'jp' coverage type only)
	if p.CoverageType == "" {
		p.CoverageType = "jp"
	}

	return &p, nil
}

func saveProfile(profileName string, prof *profile) error {
	dir, err := getProfileDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, profileName+".json")

	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	// check if profile file already exists
	if _, err := os.Stat(path); err == nil {
		// prompt if overwrites or not when already exist
		fmt.Printf(TRCLI("cli.configure.profile.overwrite"), profileName)
		s, err := readLine()
		if err != nil {
			return err
		}
		if s != "" && strings.ToLower(s) != "y" {
			return errors.New("abort")
		}

		err = lib.ProtectFile(path)
		if err != nil {
			return err
		}
	}

	b, err := json.Marshal(prof)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 0600)
	if err != nil {
		return err
	}

	err = lib.ProtectFile(path)
	if err != nil {
		return err
	}

	return nil
}

func confirmDeleteProfile(profileName string) bool {
	fmt.Printf(TRCLI("cli.unconfigure.prompt"), profileName)
	s, err := readLine()
	if err != nil {
		return false
	}
	if s != "" && strings.ToLower(s) == "y" {
		return true
	}
	return false
}

func deleteProfile(profileName string) error {
	dir, err := getProfileDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, profileName+".json")

	return os.Remove(path)
}

func collectProfileInfo(profileName string) (*profile, error) {
	profDir, err := getProfileDir()
	if err != nil {
		return nil, err
	}

	fmt.Printf(TRCLI("cli.configure.profile.prompt"), profDir, getSpecifiedProfileName())

	ct, err := collectCoverageType()
	if err != nil {
		return nil, err
	}

	ai, err := collectAuthInfo()
	if err != nil {
		return nil, err
	}

	return &profile{
		CoverageType: ct,
		Email:        ai.Email,
		Password:     ai.Password,
		AuthKeyID:    ai.AuthKeyID,
		AuthKey:      ai.AuthKey,
		OperatorID:   ai.OperatorID,
		Username:     ai.Username,
	}, nil
}

func collectSandboxProfileInfo(profileName string, registerPaymentMethod bool) (*profile, error) {
	profDir, err := getProfileDir()
	if err != nil {
		return nil, err
	}

	fmt.Printf(TRCLI("cli.configure_sandbox.profile.prompt"), profDir, getSpecifiedSandboxProfileName())

	ct, err := collectCoverageType()
	if err != nil {
		return nil, err
	}

	ai, err := collectProductionEnvAuthInfoForSandbox()
	if err != nil {
		return nil, err
	}

	sa, err := collectSandboxAccountInfo()
	if err != nil {
		return nil, err
	}

	return &profile{
		Sandbox:               true,
		CoverageType:          ct,
		Email:                 sa.Email,
		Password:              sa.Password,
		AuthKeyID:             ai.AuthKeyID,
		AuthKey:               ai.AuthKey,
		RegisterPaymentMethod: registerPaymentMethod,
	}, nil
}

func readPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func collectCoverageType() (string, error) {
	fmt.Print(TRCLI("cli.configure.profile.coverage_type.prompt"))
	var i int
	for {
		fmt.Print(TRCLI("cli.configure.profile.coverage_type.select"))
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			return "", err
		}
		if i >= 1 && i <= 2 {
			break
		}
	}

	switch i {
	case 1:
		return "g", nil
	case 2:
		return "jp", nil
	}

	return "", errors.New("this line should not be executed")
}

func collectAuthInfo() (*authInfo, error) {
	fmt.Printf(TRCLI("cli.configure.profile.auth.prompt"))
	var i int
	for {
		fmt.Print(TRCLI("cli.configure.profile.auth.select"))
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			return nil, err
		}
		if i >= 1 && i <= 3 {
			break
		}
	}

	switch i {
	case 1:
		var authKeyID, authKey string
		fmt.Print("authKeyId: ")
		_, err := fmt.Scanf("%s\n", &authKeyID)
		if err != nil {
			return nil, err
		}
		authKey, err = readPassword("authKey: ")
		if err != nil {
			return nil, err
		}
		return &authInfo{AuthKeyID: &authKeyID, AuthKey: &authKey}, nil
	case 2:
		var email string
		fmt.Print("email: ")
		_, err := fmt.Scanf("%s\n", &email)
		if err != nil {
			return nil, err
		}
		password, err := readPassword("password: ")
		if err != nil {
			return nil, err
		}
		fmt.Println()
		return &authInfo{Email: &email, Password: &password}, nil
	case 3:
		var operatorID, username string
		fmt.Print("Operator ID (OP00...): ")
		_, err := fmt.Scanf("%s\n", &operatorID)
		if err != nil {
			return nil, err
		}
		fmt.Print("username: ")
		_, err = fmt.Scanf("%s\n", &username)
		if err != nil {
			return nil, err
		}
		password, err := readPassword("password: ")
		if err != nil {
			return nil, err
		}
		fmt.Println()
		return &authInfo{
			OperatorID: &operatorID,
			Username:   &username,
			Password:   &password,
		}, nil
	}

	return nil, errors.New("this line should not be executed")
}

func collectProductionEnvAuthInfoForSandbox() (*authInfo, error) {
	fmt.Printf(TRCLI("cli.configure_sandbox.profile.prod_auth.prompt"))

	var authKeyID, authKey string
	fmt.Print("authKeyId: ")
	_, err := fmt.Scanf("%s\n", &authKeyID)
	if err != nil {
		return nil, err
	}
	authKey, err = readPassword("authKey: ")
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return &authInfo{AuthKeyID: &authKeyID, AuthKey: &authKey}, nil
}

func collectSandboxAccountInfo() (*authInfo, error) {
	fmt.Printf(TRCLI("cli.configure_sandbox.profile.sandbox_account.prompt"))

	var email string
	fmt.Print("email: ")
	_, err := fmt.Scanf("%s\n", &email)
	if err != nil {
		return nil, err
	}
	password, err := readPassword("password: ")
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return &authInfo{Email: &email, Password: &password}, nil
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}
