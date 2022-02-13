package cmd

import "testing"

func TestVersionInt(t *testing.T) {
	var testData = []struct {
		Name   string
		VerStr string
		VerInt uint32
	}{
		{
			Name:   "pattern 1",
			VerStr: "v0.1.2",
			VerInt: 0x00010200,
		},
		{
			Name:   "pattern 2",
			VerStr: "v1.22.333", // 333 == 0x14d but only 0x4d will be stored in the 3rd place
			VerInt: 0x01164d00,
		},
		{
			Name:   "pattern 3",
			VerStr: "1.2.3", // no "v" prefix
			VerInt: 0x01020300,
		},
		{
			Name:   "pattern 4",
			VerStr: "v1.2.3-special",
			VerInt: 0x01020300,
		},
		{
			Name:   "pattern 5",
			VerStr: "v1.2.3-special1",
			VerInt: 0x01020301,
		},
		{
			Name:   "pattern 6",
			VerStr: "v1.2.3.4",
			VerInt: 0x01020304,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			v := versionInt(data.VerStr)
			if v != data.VerInt {
				t.Errorf("result of versionInt() is unmatched with expected.\nArg: %v\nExpected: %#08x\nActual:   %#08x", data.VerStr, data.VerInt, v)
			}
		})
	}
}
