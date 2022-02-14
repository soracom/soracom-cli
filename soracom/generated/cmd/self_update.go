package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SelfUpdateCmd)
}

var SelfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: TRCLI("cli.self_update.summary"),
	Long:  TRCLI("cli.self_update.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		latestReleased, err := retrieveLatestReleasedFromGitHub()
		if err != nil {
			return err
		}

		if !isNewerThanCurrentVersion(latestReleased.TagName) {
			fmt.Printf(TRCLI("cli.self_update.already_latest"), version)
			return nil
		}

		fmt.Printf(TRCLI("cli.self_update.prompt_confirmation"), latestReleased.TagName, strings.TrimPrefix(latestReleased.TagName, "v"), runtime.GOARCH)
		yes, err := readDefaultNoConfirmationPrompt()
		if err != nil {
			return err
		}
		if !yes {
			return errors.New("abort")
		}

		downloaded, err := downloadExecutableBinary(latestReleased.TagName, latestReleased.Assets)
		if err != nil {
			return err
		}

		err = swapExecutableBinaryFile(downloaded)
		if err != nil {
			return err
		}

		fmt.Println(TRCLI("cli.self_update.update_finished"))
		return nil
	},
}

type gitHubRelease struct {
	TagName string                `json:"tag_name"`
	Assets  []*gitHubReleaseAsset `json:"assets"`
}

type gitHubReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func retrieveLatestReleasedFromGitHub() (*gitHubRelease, error) {
	hc := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := hc.Get("https://api.github.com/repos/soracom/soracom-cli/releases/latest")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the latest version information of soracom-cli: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the latest version information of soracom-cli: %w", err)
	}

	statusCode := resp.StatusCode
	if statusCode < 200 || statusCode >= 300 {
		return nil, fmt.Errorf("failed to retrieve the latest version information of soracom-cli with the HTTP status code %d; %s", statusCode, body)
	}

	var r gitHubRelease
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the response of the latest released item: %w", err)
	}

	return &r, nil
}

func downloadExecutableBinary(versionToDownload string, assets []*gitHubReleaseAsset) (*os.File, error) {
	asset := grepAssetsByRuntimeInfo(assets, versionToDownload, runtime.GOOS, runtime.GOARCH)
	if asset == nil {
		return nil, errors.New("there is no available executable binary of soracom-cli; we really appreciate if you can report this on https://github.com/soracom/soracom-cli/issues")
	}

	hc := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := hc.Get(asset.BrowserDownloadURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download the executable binary of soracom-cli: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	statusCode := resp.StatusCode
	if statusCode < 200 || statusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to download the executable binary of soracom-cli with the HTTP status code %d; %s", statusCode, body)
	}

	downloaded, err := writeDownloadedContentsToTempFile(resp.Body)
	if err != nil {
		return nil, err
	}
	return downloaded, nil
}

func grepAssetsByRuntimeInfo(assets []*gitHubReleaseAsset, versionToDownload string, goos string, goarch string) *gitHubReleaseAsset {
	nameRe := regexp.MustCompile(fmt.Sprintf("^soracom_%s_%s_%s(?:[.]exe)?$", strings.TrimPrefix(versionToDownload, "v"), goos, goarch))
	for _, asset := range assets {
		if nameRe.MatchString(asset.Name) {
			return asset
		}
	}
	return nil
}

func writeDownloadedContentsToTempFile(contentsBody io.Reader) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "soracom-cli_")
	if err != nil {
		return nil, fmt.Errorf("failed to create a temporary file for soracom-cil downloading: %w", err)
	}
	defer func() {
		_ = tempFile.Close()
	}()
	_ = tempFile.Chmod(0755)

	_, err = io.Copy(tempFile, contentsBody)
	if err != nil {
		return nil, fmt.Errorf("failed to write the soracom-cli binary to a temporary file: %w", err)
	}
	return tempFile, nil
}

func swapExecutableBinaryFile(downloaded *os.File) error {
	currentExecPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get the current executing soracom-cli binary information: %w", err)
	}

	execPath, err := determineActualPathOf(currentExecPath)
	if err != nil {
		return err
	}

	tempFilePath := ""
	if runtime.GOOS == "windows" {
		// In windows, it prohibits to removing/moving the running binary directly, so this accords to the following sequence.
		//   1. rename the current running binary <soracom> to a tempfile <tmp-soracom> (i.e. then tmp-soracom becomes the running binary)
		//   2. rename the downloaded new version binary to the original running binary path <soracom>
		//   3. after that, <soracom> becomes the new version from the next run time
		tempFile, err := os.CreateTemp("", "soracom-cli_")
		if err != nil {
			return fmt.Errorf("failed to create a tempfile to update: %w", err)
		}
		tempFilePath = tempFile.Name()

		err = tempFile.Close() // this close is mandatory to rename it at the bellow
		if err != nil {
			return fmt.Errorf("failed to create a tempfile to update: %w", err)
		}

		err = os.Rename(execPath, tempFilePath)
		if err != nil {
			return fmt.Errorf("failed to swap the current running cli and tempfile: %w", err)
		}
	}

	err = os.Rename(downloaded.Name(), execPath)
	if err != nil {
		if tempFilePath != "" { // restoring
			err = os.Rename(tempFilePath, execPath)
			if err != nil {
				fmt.Printf("error: failed to restore the cli binary; %s", err)
			}
		}
		return fmt.Errorf("failed to swap the binary file between the current executing and downloaded latest one: %w", err)
	}

	return nil
}

func determineActualPathOf(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("failed to refer the lstat information of %s: %w", path, err)
	}

	if fileInfo.Mode()&os.ModeSymlink != os.ModeSymlink {
		return path, nil
	}

	resolvedPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve a symlink for the curent executing path: %w", err)
	}
	return resolvedPath, nil
}
