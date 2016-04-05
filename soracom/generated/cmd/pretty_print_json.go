package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func prettyPrintJSON(rawJSON string) error {
	var obj interface{}

	err := json.Unmarshal(bytes.NewBufferString(rawJSON).Bytes(), &obj)
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
