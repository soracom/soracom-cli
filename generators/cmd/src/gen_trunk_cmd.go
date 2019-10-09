package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/soracom/soracom-cli/generators/lib"
)

func generateTrunkCommands(apiDef *lib.APIDefinitions, templateDir, outputDir string) error {
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

func generateArgsForTrunkCommands(apiDef *lib.APIDefinitions) []commandArgs {
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

func extractTrunkCommands(apiDef *lib.APIDefinitions) []string {
	commands := map[string]interface{}{}

	for _, m := range apiDef.Methods {
		if m.CLI == nil || len(m.CLI) == 0 {
			continue
		}

		for _, cli := range m.CLI {
			s := strings.Split(cli, " ")
			if len(s) <= 1 {
				continue
			}

			for i := 1; i < len(s); i++ {
				ss := s[:i]
				commands[strings.Join(ss, " ")] = true
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
