package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func generateRootCommand(apiDef *openapi3.T, templateDir, outputDir string) error {
	t, err := openTemplateFile(templateDir, "root.gotmpl")
	if err != nil {
		return err
	}

	w, err := openOutputFile(outputDir, "root.go")
	if err != nil {
		return err
	}

	a := commandArgs{
		Use:   "soracom",
		Short: "soracom command",
		Long:  `A command line tool to invoke SORACOM API`,
	}
	err = t.Execute(w, a)
	if err != nil {
		return err
	}
	return nil
}
