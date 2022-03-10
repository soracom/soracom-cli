package cmd

import (
	"bytes"
	"testing"

	"github.com/tj/assert"
)

func TestPrintStringAsJSONL(t *testing.T) {
	data1 := `[{"a":1},{"b":2},{"c":3}]`
	expected := `{"a":1}
{"b":2}
{"c":3}
`
	out1 := new(bytes.Buffer)
	err := printStringAsJSONLToWriter(data1, out1)
	assert.NoError(t, err)
	assert.Equal(t, expected, out1.String())

	noArray1 := `{"a":1,"b":2}`
	out2 := new(bytes.Buffer)
	err = printStringAsJSONLToWriter(noArray1, out2)
	assert.Error(t, err)
}
