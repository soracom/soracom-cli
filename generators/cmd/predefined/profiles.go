package cmd

import (
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
	"golang.org/x/crypto/ssh/terminal"
)

type profile struct {
	Email      *string `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	AuthKeyID  *string `json:"authKeyId,omitempty"`
	AuthKey    *string `json:"authKey,omitempty"`
	Username   *string `json:"username,omitempty"`
	OperatorID *string `json:"operatorId,omitempty"`
}

func getDefaultProfileName() string {
	return "default"
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

func loadProfile(profileName string) (*profile, error) {
	dir, err := getProfileDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, profileName+".json")

	// check if permission is 0600
	if runtime.GOOS != "windows" {
		s, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		if s.Mode()&077 != 0 {
			return nil, fmt.Errorf("permission for %s is too open", path)
		}
	} else {
		// TODO: handle ACL on windows env
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p profile
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func saveProfile(profileName string, prof *profile) error {
	dir, err := getProfileDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, profileName+".json")

	// check if profile dir exists
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	// check if profile file already exists
	if _, err := os.Stat(path); err == nil {
		// prompt if overwrites or not when already exist
		fmt.Printf(TR("configure.cli.profile.overwrite"), profileName)
		var s string
		fmt.Scanf("%s\n", &s)
		if s != "" && strings.ToLower(s) != "y" {
			return errors.New("abort")
		}

		os.Chmod(path, 0600)

		if runtime.GOOS == "windows" {
			// TODO: handle ACL on windows
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

	if runtime.GOOS == "windows" {
		// TODO: handle ACL on windows
	}

	return nil
}

func collectProfileInfo(profileName string) (*profile, error) {
	profDir, err := getProfileDir()
	if err != nil {
		return nil, err
	}

	fmt.Printf(TR("configure.cli.profile.prompt"), profDir, getSpecifiedProfileName())

	var i int
	for {
		fmt.Print(TR("configure.cli.profile.select"))
		fmt.Scanf("%d\n", &i)
		if i >= 1 && i <= 3 {
			break
		}
	}

	switch i {
	case 1:
		var authKeyID, authKey string
		fmt.Print("authKeyId: ")
		fmt.Scanf("%s\n", &authKeyID)
		fmt.Print("authKey: ")
		fmt.Scanf("%s\n", &authKey)
		return &profile{AuthKeyID: &authKeyID, AuthKey: &authKey}, nil
	case 2:
		var email string
		fmt.Print("email: ")
		fmt.Scanf("%s\n", &email)
		password, err := readPassword("password: ")
		if err != nil {
			return nil, err
		}
		fmt.Println()
		return &profile{Email: &email, Password: &password}, nil
	case 3:
		var operatorID, username string
		fmt.Print("Operator ID (OP00...): ")
		fmt.Scanf("%s\n", &operatorID)
		fmt.Print("username: ")
		fmt.Scanf("%s\n", &username)
		password, err := readPassword("password: ")
		if err != nil {
			return nil, err
		}
		fmt.Println()
		return &profile{
			OperatorID: &operatorID,
			Username:   &username,
			Password:   &password,
		}, nil
	}

	return nil, errors.New("this line should not be executed")
}

func readPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(password), nil
}
