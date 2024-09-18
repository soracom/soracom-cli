package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/soracom/soracom-cli/generators/lib"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run() error {
	outputDir := flag.String("o", "", "output directory")
	apiDefFile := flag.String("a", "", "API definition YAML file")
	apiSandboxDefFile := flag.String("s", "", "API Sandbox definition YAML file")
	templateDir := flag.String("t", "", "template directory")
	predefinedDir := flag.String("p", "", "predefined command files directory")
	flag.Parse()

	if outputDir == nil || *outputDir == "" {
		return errors.New("-o <output dir> is required")
	}

	if apiDefFile == nil || *apiDefFile == "" {
		return errors.New("-a <api definition file> is required")
	}

	if apiSandboxDefFile == nil || *apiSandboxDefFile == "" {
		return errors.New("-s <api sandbox definition file> is required")
	}

	if templateDir == nil || *templateDir == "" {
		return errors.New("-t <template dir> is required")
	}

	if predefinedDir == nil || *predefinedDir == "" {
		return errors.New("-p <predefined command files dir> is required")
	}

	err := cleanOutputDir(*outputDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(*outputDir, 0700)
	if err != nil {
		return err
	}

	loader := openapi3.NewLoader()
	apiDef, err := loader.LoadFromFile(*apiDefFile)
	if err != nil {
		return err
	}

	apiSandboxDef, err := loader.LoadFromFile(*apiSandboxDefFile)
	if err != nil {
		return err
	}

	return generateCommands(apiDef, apiSandboxDef, *templateDir, *predefinedDir, *outputDir)
}

func generateCommands(apiDef, apiSandboxDef *openapi3.T, templateDir, predefinedDir, outputDir string) error {
	err := generateRootCommand(apiDef, apiSandboxDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateTrunkCommands(apiDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateTrunkCommands(apiSandboxDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateLeafCommands(apiDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateLeafCommands(apiSandboxDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = copyPredefinedCommands(predefinedDir, outputDir)
	if err != nil {
		return err
	}

	/*
		err = formatGeneratedFiles(outputDir)
		if err != nil {
			return err
		}
	*/
	return nil
}

func cleanOutputDir(outputDir string) error {
	if outputDir == "." || outputDir == ".." || outputDir == "/" {
		return fmt.Errorf("cowardly rejected removing '%s'", outputDir)
	}
	return os.RemoveAll(outputDir)
}

func openTemplateFile(templateDir, filename string) (*template.Template, error) {
	return template.Must(template.ParseFiles(filepath.Join(templateDir, filename))), nil
}

func openOutputFile(dir, file string) (*os.File, error) {
	return os.Create(filepath.Join(dir, file))
}

func copyPredefinedCommands(predefinedDir, outputDir string) error {
	paths, err := filepath.Glob(filepath.Join(predefinedDir, "*.go"))
	if err != nil {
		return err
	}

	for _, path := range paths {
		// #nosec
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(outputDir, filepath.Base(path)), b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func formatGeneratedFiles(outputDir string) error {
	paths, err := filepath.Glob(filepath.Join(outputDir, "*.go"))
	if err != nil {
		return err
	}
	for _, path := range paths {
		fmt.Printf("formatting %s\n", path)
		// #nosec
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		b2, err := format.Source(b)
		if err != nil {
			return err
		}

		err = os.WriteFile(path, b2, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCLICommands(op *openapi3.Operation) []string {
	xSoracomCliField, found := op.Extensions["x-soracom-cli"]
	if !found {
		return nil
	}

	sliceOfAny, ok := xSoracomCliField.([]any)
	if !ok {
		lib.WarnfStderr("invalid x-soracom-cli: %v (%T)\n", op.Extensions["x-soracom-cli"], op.Extensions["x-soracom-cli"])
		return nil
	}

	var result []string
	for _, v := range sliceOfAny {
		s, ok := v.(string)
		if !ok {
			lib.WarnfStderr("invalid value of x-soracom-cli: %v (%T) in %v\n", v, v, sliceOfAny)
			continue
		}
		result = append(result, s)
	}

	return result
}

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
