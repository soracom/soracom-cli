package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/soracom/soracom-cli/generators/lib"
)

func main() {
	outputDir := flag.String("o", "", "output directory")
	apiDefFile := flag.String("a", "", "API definition YAML file")
	templateDir := flag.String("t", "", "template directory")
	predefinedDir := flag.String("p", "", "predefined command files directory")
	flag.Parse()

	if outputDir == nil || *outputDir == "" {
		fmt.Println("-o <output dir> is required")
		os.Exit(1)
	}

	if apiDefFile == nil || *apiDefFile == "" {
		fmt.Println("-a <api definition file> is required")
		os.Exit(1)
	}

	if templateDir == nil || *templateDir == "" {
		fmt.Println("-t <template dir> is required")
		os.Exit(1)
	}

	if predefinedDir == nil || *predefinedDir == "" {
		fmt.Println("-p <predefined command files dir> is required")
		os.Exit(1)
	}

	err := run(*apiDefFile, *templateDir, *predefinedDir, *outputDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(apiDefFile, templateDir, predefinedDir, outputDir string) error {
	err := cleanOutputDir(outputDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(outputDir, 0700)
	if err != nil {
		return err
	}

	apiDef, err := lib.LoadAPIDef(apiDefFile)
	if err != nil {
		return err
	}

	err = generateCommands(apiDef, templateDir, predefinedDir, outputDir)
	if err != nil {
		return err
	}

	return nil
}

func generateCommands(apiDef *lib.APIDefinitions, templateDir, predefinedDir, outputDir string) error {
	err := generateRootCommand(apiDef, templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateTrunkCommands(templateDir, outputDir)
	if err != nil {
		return err
	}

	err = generateLeafCommands(apiDef, templateDir, outputDir)
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
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join(outputDir, filepath.Base(path)), b, 0644)
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
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		b2, err := format.Source(b)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path, b2, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
