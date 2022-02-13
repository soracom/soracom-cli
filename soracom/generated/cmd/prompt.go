package cmd

import (
	"bufio"
	"os"
	"strings"
)

// XXX: visible for testing
var promptStdin = os.Stdin

// readConfirmationPrompt returns true when the input value is 'y', 'Y', or empty.
func readDefaultYesConfirmationPrompt() (bool, error) {
	s, err := readLine()
	if err != nil {
		return false, err
	}
	return s == "" || strings.ToLower(s) == "y", nil
}

// readConfirmationPrompt returns true when the input value is 'y', 'Y'.
func readDefaultNoConfirmationPrompt() (bool, error) {
	s, err := readLine()
	if err != nil {
		return false, err
	}
	return s != "" && strings.ToLower(s) == "y", nil
}

func readLine() (string, error) {
	reader := bufio.NewReader(promptStdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}
