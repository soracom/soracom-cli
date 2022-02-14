package cmd

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrepAssetsByRuntimeInfo(t *testing.T) {
	assets := []*gitHubReleaseAsset{
		{
			Name:               "soracom_0.11.0_amd64.deb",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_amd64.deb",
		},
		{
			Name:               "soracom_0.11.0_arm64.deb",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_arm64.deb",
		},
		{
			Name:               "soracom_0.11.0_armhf.deb",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_armhf.deb",
		},
		{
			Name:               "soracom_0.11.0_darwin_amd64",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_amd64",
		},
		{
			Name:               "soracom_0.11.0_darwin_amd64.zip",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_amd64.zip",
		},
		{
			Name:               "soracom_0.11.0_darwin_arm64",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_arm64",
		},
		{
			Name:               "soracom_0.11.0_darwin_arm64.zip",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_arm64.zip",
		},
		{
			Name:               "soracom_0.11.0_freebsd_386",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_386",
		},
		{
			Name:               "soracom_0.11.0_freebsd_386.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_386.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_freebsd_amd64",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_amd64",
		},
		{
			Name:               "soracom_0.11.0_freebsd_amd64.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_amd64.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_freebsd_arm",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_arm",
		},
		{
			Name:               "soracom_0.11.0_freebsd_arm.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_freebsd_arm.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_i386.deb",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_i386.deb",
		},
		{
			Name:               "soracom_0.11.0_linux_386",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_386",
		},
		{
			Name:               "soracom_0.11.0_linux_386.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_386.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_linux_amd64",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_amd64",
		},
		{
			Name:               "soracom_0.11.0_linux_amd64.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_amd64.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_linux_arm",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_arm",
		},
		{
			Name:               "soracom_0.11.0_linux_arm.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_arm.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_linux_arm64",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_arm64",
		},
		{
			Name:               "soracom_0.11.0_linux_arm64.tar.gz",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_arm64.tar.gz",
		},
		{
			Name:               "soracom_0.11.0_windows_386.exe",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_windows_386.exe",
		},
		{
			Name:               "soracom_0.11.0_windows_386.zip",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_windows_386.zip",
		},
		{
			Name:               "soracom_0.11.0_windows_amd64.exe",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_windows_amd64.exe",
		},
		{
			Name:               "soracom_0.11.0_windows_amd64.zip",
			BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_windows_amd64.zip",
		},
	}

	type test struct {
		goos     string
		goarch   string
		expected *gitHubReleaseAsset
	}

	tests := []*test{
		{
			goos:   "darwin",
			goarch: "amd64",
			expected: &gitHubReleaseAsset{
				Name:               "soracom_0.11.0_darwin_amd64",
				BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_amd64",
			},
		},
		{
			goos:   "darwin",
			goarch: "arm64",
			expected: &gitHubReleaseAsset{
				Name:               "soracom_0.11.0_darwin_arm64",
				BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_darwin_arm64",
			},
		},
		{
			goos:   "linux",
			goarch: "arm",
			expected: &gitHubReleaseAsset{
				Name:               "soracom_0.11.0_linux_arm",
				BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_linux_arm",
			},
		},
		{
			goos:   "windows",
			goarch: "386",
			expected: &gitHubReleaseAsset{
				Name:               "soracom_0.11.0_windows_386.exe",
				BrowserDownloadURL: "https://github.com/soracom/soracom-cli/releases/download/v0.11.0/soracom_0.11.0_windows_386.exe",
			},
		},
		{
			goos:     "__not_exists__",
			goarch:   "amd64",
			expected: nil,
		},
		{
			goos:     "linux",
			goarch:   "__not_exists__",
			expected: nil,
		},
		{
			goos:     "__not_exists__",
			goarch:   "__not_exists__",
			expected: nil,
		},
	}

	for _, testCase := range tests {
		got := grepAssetsByRuntimeInfo(assets, "0.11.0", testCase.goos, testCase.goarch)
		assert.EqualValues(t, testCase.expected, got)
	}

	got := grepAssetsByRuntimeInfo(assets, "0.0.0", runtime.GOOS, runtime.GOARCH)
	assert.Nil(t, got)
}

func TestDetermineActualPathOf(t *testing.T) {
	tempFile, err := os.CreateTemp("", "")
	assert.NoError(t, err)
	tempFilePath := tempFile.Name()
	defer func() {
		_ = os.Remove(tempFilePath)
	}()

	tempSymlinkPath := tempFilePath + ".symlink"
	err = os.Symlink(tempFilePath, tempSymlinkPath)
	assert.NoError(t, err)
	defer func() {
		_ = os.Remove(tempSymlinkPath)
	}()

	actualPath, err := determineActualPathOf(tempFilePath)
	assert.NoError(t, err)
	assert.Equal(t, tempFilePath, actualPath)

	actualPath, err = determineActualPathOf(tempSymlinkPath)
	assert.NoError(t, err)
	assert.Equal(t, tempFilePath, actualPath)
}
