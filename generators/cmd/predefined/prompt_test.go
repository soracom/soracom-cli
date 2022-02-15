package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDefaultYesConfirmationPrompt(t *testing.T) {
	type test struct {
		input    string
		expected bool
	}

	tests := []*test{
		{input: "y", expected: true},
		{input: "Y", expected: true},
		{input: "", expected: true},
		{input: "n", expected: false},
		{input: "o", expected: false},
	}

	tempFile, err := os.CreateTemp("", "")
	assert.NoError(t, err)
	originalPromptStdin := promptStdin
	promptStdin = tempFile
	defer func() {
		promptStdin = originalPromptStdin
		_ = os.Remove(tempFile.Name())
	}()

	for _, testCase := range tests {
		_, err = tempFile.WriteString(fmt.Sprintf("%s\n", testCase.input))
		assert.NoError(t, err)
		_, err = tempFile.Seek(0, 0)
		assert.NoError(t, err)

		yes, err := readDefaultYesConfirmationPrompt()
		assert.NoError(t, err)
		assert.Equal(t, testCase.expected, yes)

		err = tempFile.Truncate(0)
		assert.NoError(t, err)
		_, err = tempFile.Seek(0, 0)
		assert.NoError(t, err)
	}

	_ = tempFile.Close()
}

func TestReadDefaultNoConfirmationPrompt(t *testing.T) {
	type test struct {
		input    string
		expected bool
	}

	tests := []*test{
		{input: "y", expected: true},
		{input: "Y", expected: true},
		{input: "", expected: false},
		{input: "n", expected: false},
		{input: "o", expected: false},
	}

	tempFile, err := os.CreateTemp("", "")
	assert.NoError(t, err)
	originalPromptStdin := promptStdin
	promptStdin = tempFile
	defer func() {
		promptStdin = originalPromptStdin
		_ = os.Remove(tempFile.Name())
	}()

	for _, testCase := range tests {
		_, err = tempFile.WriteString(fmt.Sprintf("%s\n", testCase.input))
		assert.NoError(t, err)
		_, err = tempFile.Seek(0, 0)
		assert.NoError(t, err)

		yes, err := readDefaultNoConfirmationPrompt()
		assert.NoError(t, err)
		assert.Equal(t, testCase.expected, yes)

		err = tempFile.Truncate(0)
		assert.NoError(t, err)
		_, err = tempFile.Seek(0, 0)
		assert.NoError(t, err)
	}

	_ = tempFile.Close()
}
