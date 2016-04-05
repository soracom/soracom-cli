package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func prettyPrintJSON(rawJSON string) error {
	var obj interface{}

	d := json.NewDecoder(strings.NewReader(rawJSON))
	d.UseNumber()
	err := d.Decode(&obj)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}

	os.Stdout.Write(b)
	fmt.Println()
	return nil
}
