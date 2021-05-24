// +build !windows

package lib

import "os"

// IsFilePermissionTooOpen returns true only when the `path` doesn't have the expected permission
func IsFilePermissionTooOpen(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if s.Mode()&077 != 0 {
		return true, nil
	}

	return false, nil
}

// ProtectFile changes the mode of the specified `path`
func ProtectFile(path string) error {
	return os.Chmod(path, 0600)
}
