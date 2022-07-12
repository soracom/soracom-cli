package main

import (
	"testing"

	"github.com/tj/assert"
)

func TestRemoveLineBreaks(t *testing.T) {
	testData := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "",
			Expected: "",
		},
		{
			Input:    "a\nb",
			Expected: "ab",
		},
		{
			Input:    "aaaa\r\nbbbbb\r\n",
			Expected: "aaaabbbbb",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Input, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, data.Expected, removeLineBreaks(data.Input))
		})
	}
}
