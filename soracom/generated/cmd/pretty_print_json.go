package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func prettyPrintStringAsJSON(rawJSON string) error {
	var obj interface{}

	d := json.NewDecoder(strings.NewReader(rawJSON))
	d.UseNumber()
	err := d.Decode(&obj)
	if err != nil {
		return err
	}
	return prettyPrintObjectAsJSON(obj)
}

func prettyPrintObjectAsJSON(obj interface{}) error {
	b, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(b)
	if err != nil {
		return err
	}

	fmt.Println()
	return nil
}
