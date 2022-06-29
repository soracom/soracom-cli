package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
)

func prettyPrintStringAsJSON(rawJSON string) error {
	return prettyPrintStringAsJSONToWriter(rawJSON, os.Stdout)
}

func prettyPrintStringAsJSONToWriter(rawJSON string, w io.Writer) error {
	var obj interface{}

	d := json.NewDecoder(strings.NewReader(rawJSON))
	d.UseNumber()
	err := d.Decode(&obj)
	if err != nil {
		return err
	}
	return prettyPrintObjectAsJSON(obj, w)
}

func prettyPrintObjectAsJSON(obj interface{}, w io.Writer) error {
	var bb bytes.Buffer
	e := json.NewEncoder(&bb)
	e.SetEscapeHTML(false)
	e.SetIndent("", "\t")
	err := e.Encode(obj)
	if err != nil {
		return err
	}

	_, err = bb.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}

func printStringAsJSONL(rawJSON string) error {
	return printStringAsJSONLToWriter(rawJSON, os.Stdout)
}

func printStringAsJSONLToWriter(rawJSON string, w io.Writer) error {
	var arr []interface{}

	d := json.NewDecoder(strings.NewReader(rawJSON))
	d.UseNumber()
	err := d.Decode(&arr)
	if err != nil {
		return err
	}

	for _, obj := range arr {
		err = printObjectOneLine(obj, w)
		if err != nil {
			return err
		}
	}

	return nil
}

func printObjectOneLine(obj interface{}, w io.Writer) error {
	var bb bytes.Buffer
	e := json.NewEncoder(&bb)
	e.SetEscapeHTML(false)
	err := e.Encode(obj)
	if err != nil {
		return err
	}

	_, err = bb.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}
