package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/soracom/soracom-cli/generators/lib"
)

func generateTrunkCommands(apiDef *openapi3.T, templateDir, outputDir string) error {
	subCommandTemplate, err := openTemplateFile(templateDir, "trunk.gotmpl")
	if err != nil {
		return err
	}

	argsSlice := generateArgsForTrunkCommands(apiDef)

	for _, args := range argsSlice {
		f, err := openOutputFile(outputDir, args.FileName)
		if err != nil {
			return err
		}
		err = subCommandTemplate.Execute(f, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateArgsForTrunkCommands(apiDef *openapi3.T) []commandArgs {
	trunkCommands := extractTrunkCommands(apiDef)

	result := make([]commandArgs, 0)

	for _, tc := range trunkCommands {
		s := strings.Split(tc, " ")
		ca := commandArgs{
			Use:                       s[len(s)-1],
			Short:                     fmt.Sprintf("cli.%s.summary", strings.Join(s, ".")),
			Long:                      fmt.Sprintf("cli.%s.description", strings.Join(s, ".")),
			CommandVariableName:       getCommandVariableName(strings.Join(s, " ")),
			ParentCommandVariableName: getParentCommandVariableName(strings.Join(s, " ")),
			FileName:                  fmt.Sprintf("%s.go", lib.SnakeCase(strings.Join(s, "-"))),
		}
		result = append(result, ca)
	}

	return result
}

func extractTrunkCommands(apiDef *openapi3.T) []string {
	commands := map[string]interface{}{}

	for _, path := range apiDef.Paths {
		for _, op := range path.Operations() {
			cliCommands := getCLICommands(op)
			for _, cmd := range cliCommands {
				s := strings.Split(cmd, " ")
				if len(s) <= 1 {
					continue
				}

				for i := 1; i < len(s); i++ {
					ss := s[:i]
					commands[strings.Join(ss, " ")] = true
				}
			}
		}
	}

	result := []string{}
	for k := range commands {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}
