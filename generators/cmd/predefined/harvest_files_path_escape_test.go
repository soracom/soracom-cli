package cmd

import (
	"testing"
)

func TestHarvestFilesPathEscape(t *testing.T) {
	var testData = []struct {
		Name     string
		Path     string
		Expected string
	}{
		{
			Name: "pattern 1",
			Path: "hoge",
			Expected: "hoge",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			v := harvestFilesPathEscape(data.Path)
			if v != data.Expected {
				t.Errorf("result of harvestFilesPathEscape() is unmatched with expected.\nArg: %s\nExpected: %s\nActual:   %s", data.Path, data.Expected, v)
			}
		})
	}
}