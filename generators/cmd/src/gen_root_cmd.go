package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type rootCommandArgs struct {
	Use         string
	Short       string
	Long        string
	Subcommands []subcommand
}

type subcommand struct {
	CommandVariableName string
}

func generateRootCommand(apiDef, apiSandboxDef *openapi3.T, templateDir, outputDir string) error {
	t, err := openTemplateFile(templateDir, "root.gotmpl")
	if err != nil {
		return err
	}

	subcommands := getAllSubcommands(apiDef)
	subcommands = append(subcommands, getAllSubcommands(apiSandboxDef)...)

	w, err := openOutputFile(outputDir, "root.go")
	if err != nil {
		return err
	}

	a := rootCommandArgs{
		Use:         "soracom",
		Short:       "soracom command",
		Long:        `A command line tool to invoke SORACOM API`,
		Subcommands: subcommands,
	}
	err = t.Execute(w, a)
	if err != nil {
		return err
	}
	return nil
}

func getAllSubcommands(apiDef *openapi3.T) []subcommand {
	var result []subcommand

	for _, pathItem := range apiDef.Paths {
		for _, op := range pathItem.Operations() {
			for _, commandName := range getCLICommands(op) {
				result = append(result, subcommand{
					CommandVariableName: getCommandVariableName(commandName),
				})
			}
		}
	}

	return result
}
