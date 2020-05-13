package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/soracom/soracom-cli/generators/lib"
)

func generateLeafCommands(apiDef *lib.APIDefinitions, templateDir, outputDir string) error {
	var err error
	subCommandTemplate, err := openTemplateFile(templateDir, "leaf.gotmpl")
	if err != nil {
		return err
	}

	for _, m := range apiDef.Methods {
		//fmt.Printf("%6s: %s => %s\n", m.Method, m.Path, m.CLI)
		err = generateCommandFiles(apiDef, m, subCommandTemplate, outputDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateCommandFiles(apiDef *lib.APIDefinitions, m lib.APIMethod, tmpl *template.Template, outputDir string) error {
	// m.CLI is an array of cli subcommands list.
	// e.g. [ "subscribers list", "sim list" ]
	for _, commandName := range m.CLI {
		// commandName is a space separated subcommands list. e.g. "subscribers list"
		filename := lib.SnakeCase(commandName)
		f, err := openOutputFile(outputDir, filename+".go")
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Printf("[WARN] unable to close a file '%s'", f.Name())
			}
		}()

		a := commandArgs{
			Use:                               getLast(commandName),
			Short:                             m.Path + ":" + m.Method + ":summary",
			Long:                              m.Path + ":" + m.Method + ":description",
			CommandVariableName:               getCommandVariableName(commandName),
			ParentCommandVariableName:         getParentCommandVariableName(commandName),
			RequireAuth:                       m.Security != nil,
			RequireOperatorID:                 isOperatorIDRequired(m.Parameters, apiDef.StructDefs),
			BodyExists:                        doesRequestBodyExist(m.Parameters),
			SendBodyRaw:                       isBodyArray(m.Parameters) || isBodyBinary(m.Parameters),
			ResponseBodyRaw:                   isResponseBodyRaw(m),
			Method:                            strings.ToUpper(m.Method),
			BasePath:                          apiDef.BasePath,
			Path:                              m.Path,
			PathParamsExist:                   doPathParamsExist(m.Parameters),
			QueryParamsExist:                  doQueryParamsExist(m.Parameters),
			StringFlags:                       getStringFlags(m, m.Parameters, apiDef.StructDefs),
			StringSliceFlags:                  getStringSliceFlags(m.Parameters, apiDef.StructDefs),
			IntegerFlags:                      getIntegerFlags(m.Parameters, apiDef.StructDefs),
			FloatFlags:                        getFloatFlags(m.Parameters, apiDef.StructDefs),
			BoolFlags:                         getBoolFlags(m.Parameters, apiDef.StructDefs),
			RequiredFlagExists:                doesRequiredFlagExist(m, m.Parameters, apiDef.StructDefs),
			PaginationAvailable:               m.Pagination != nil,
			PaginationKeyHeaderInResponse:     getPaginationKeyHeaderInResponse(m.Pagination),
			PaginationRequestParameterInQuery: getPaginationRequestparameterInQuery(m.Pagination),
		}
		if a.Method == "POST" || a.Method == "PUT" {
			if doesContentTypeParamExist(m.Parameters) {
				a.ContentTypeFromArg = true
				a.ContentTypeVarName = getContentTypeVarName(a.StringFlags)
			} else {
				a.ContentType = "application/json"
			}
		}

		err = tmpl.Execute(f, a)
		if err != nil {
			return err
		}
	}

	return nil
}

var templateRegex = regexp.MustCompile("^\\s*\\${(.+)}\\s*$")

func trimTemplate(s string) string {
	b1 := templateRegex.ReplaceAll([]byte(s), []byte("$1"))
	return string(b1)
}

func doesResponseHavePayload(props *lib.APIMethod) bool {
	return props.Responses["200"].Schema != nil
}

func getLast(s string) string {
	ss := strings.Split(s, " ")
	return ss[len(ss)-1]
}

func escapeBackquote(s string) string {
	return strings.Replace(s, "`", "'", -1)
}

// if commandName == "subscriber list"
// then this func returns "SubscriberListCmd"
func getCommandVariableName(commandName string) string {
	return lib.TitleCase(commandName) + "Cmd"
}

// if commandName == "subscriber list"
// then this func returns "SubscriberCmd"
func getParentCommandVariableName(commandName string) string {
	s := strings.Split(commandName, " ")
	if len(s) == 1 {
		return "RootCmd"
	}

	return lib.TitleCase(strings.Join(s[0:len(s)-1], "-")) + "Cmd"
}

// if commandName == "subscriber list"
// then this func returns "subscriber.go"
func getParentCommandFileName(commandName string) string {
	s := strings.Split(commandName, " ")
	if len(s) == 1 {
		return "root.go"
	}

	return lib.SnakeCase(strings.Join(s[0:len(s)-1], "-")) + ".go"
}

func doesRequestBodyExist(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.In == "body" {
			return true
		}
	}
	return false
}

func isBodyArray(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.In == "body" {
			return (param.Type == "array") || (param.Schema.Type == "array")
		}
	}
	return false
}

func isBodyBinary(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.In == "body" {
			return (param.Schema.Type == "string") && (param.Schema.Format == "binary")
		}
	}
	return false
}

func isResponseBodyRaw(m lib.APIMethod) bool {
	if strings.ToUpper(m.Method) == "GET" && m.Path == "/files/{scope}/{path}" {
		return true
	}
	return false
}

func doPathParamsExist(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.In == "path" {
			return true
		}
	}
	return false
}

func doQueryParamsExist(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.In == "query" {
			return true
		}
	}
	return false
}

func doesContentTypeParamExist(parameters []lib.APIParam) bool {
	for _, param := range parameters {
		if param.Name == "content-type" {
			return true
		}
	}
	return false
}

func getContentTypeVarName(stringFlags []stringFlag) string {
	for _, sf := range stringFlags {
		if sf.Name == "content-type" {
			return sf.VarName
		}
	}
	return ""
}

func getStringFlags(m lib.APIMethod, parameters []lib.APIParam, definitions map[string]lib.StructDef) []stringFlag {
	result := []stringFlag{}
	for _, param := range parameters {
		switch param.In {
		case "path", "query", "header":
			if param.Type != "string" {
				continue
			}
			f := stringFlag{
				VarName:                lib.TitleCase(param.Name),
				LongOption:             lib.OptionCase(param.Name),
				DefaultValueSpecified:  param.Default != nil,
				DefaultValue:           param.GetDefaultValueAsString(),
				ShortHelp:              trimTemplate(param.Description),
				Name:                   param.Name,
				In:                     param.In,
				HarvestFilesPathEscape: isHarvestFilesPathEscapeRequired(m, param),
			}
			if param.Name != "operator_id" {
				f.Required = param.Required
			}
			result = append(result, f)
		case "body":
			var s []stringFlag
			if param.Schema.Ref != "" {
				s = getStringFlagsFromStruct(param, definitions[getStructNameFromReference(param.Schema.Ref)])
			} else if param.Schema.Type == "array" {
				//fmt.Println("[WARN] array is not supported yet.")
			}
			result = append(result, s...)
		default:
			fmt.Printf("[WARN] parameters in '%s' is not supported\n", param.In)
		}
	}

	sort.Sort(stringFlagsByName(result))
	return result
}

func isHarvestFilesPathEscapeRequired(m lib.APIMethod, param lib.APIParam) bool {
	if m.Path == "/files/{scope}/{path}" || m.Path == "/files/{scope}/{path}/" {
		if param.Name == "path" {
			return true
		}
	}
	return false
}

func isOperatorIDRequired(parameters []lib.APIParam, definitions map[string]lib.StructDef) bool {
	for _, param := range parameters {
		switch param.In {
		case "path":
			if param.Name != "operator_id" {
				continue
			}
			if param.Type != "string" {
				continue
			}
			return true
		case "body":
			if param.Schema.Ref != "" {
				s := definitions[getStructNameFromReference(param.Schema.Ref)]
				for _, p := range s.Properties {
					if p.Name == "operatorId" {
						return p.Required
					}
				}
			}
		}
	}
	return false
}

func getStringSliceFlags(parameters []lib.APIParam, definitions map[string]lib.StructDef) []stringFlag {
	result := []stringFlag{}
	for _, param := range parameters {
		switch param.In {
		case "query":
			if param.Type != "array" || param.Items.Type != "string" {
				continue
			}
			var f stringFlag
			f.VarName = lib.TitleCase(param.Name)
			f.LongOption = lib.OptionCase(param.Name)
			f.DefaultValueSpecified = param.Default != nil
			f.DefaultValue = param.GetDefaultValueAsString()
			f.ShortHelp = trimTemplate(param.Description)
			f.Name = param.Name
			f.In = param.In
			f.Required = param.Required
			result = append(result, f)
		case "path", "body", "header":
			continue
		default:
			fmt.Printf("[WARN] parameters in '%s' is not supported\n", param.In)
		}
	}

	sort.Sort(stringFlagsByName(result))
	return result
}

func getStringFlagsFromStruct(param lib.APIParam, definition lib.StructDef) []stringFlag {
	//fmt.Println("        getting string flags")
	result := []stringFlag{}
	for _, prop := range definition.Properties {
		if prop.Type != "string" {
			continue
		}
		//fmt.Printf("          %s\n", prop.Name)
		var f stringFlag
		f.VarName = lib.TitleCase(prop.Name)
		f.LongOption = lib.OptionCase(prop.Name)
		f.DefaultValueSpecified = prop.Default != nil
		f.DefaultValue = prop.GetDefaultValueAsString()
		f.ShortHelp = trimTemplate(prop.Description)
		f.In = "body"
		f.Name = prop.Name
		f.Required = prop.Required
		result = append(result, f)
	}
	return result
}

func getIntegerFlags(parameters []lib.APIParam, definitions map[string]lib.StructDef) []integerFlag {
	result := []integerFlag{}
	for _, param := range parameters {
		switch param.In {
		case "path", "query", "header":
			if param.Type != "integer" {
				continue
			}
			var f integerFlag
			f.VarName = lib.TitleCase(param.Name)
			f.LongOption = lib.OptionCase(param.Name)
			f.DefaultValueSpecified = param.Default != nil
			f.DefaultValue = param.GetDefaultValueAsInt64()
			f.ShortHelp = trimTemplate(param.Description)
			f.Name = param.Name
			f.In = param.In
			f.Required = param.Required
			result = append(result, f)
		case "body":
			var s []integerFlag
			if param.Schema.Ref != "" {
				s = getIntegerFlagsFromStruct(param, definitions[getStructNameFromReference(param.Schema.Ref)])
			} else if param.Schema.Type == "array" {
				//fmt.Println("[WARN] array is not supported yet.")
			}
			result = append(result, s...)
		default:
			fmt.Printf("[WARN] parameters in '%s' is not supported\n", param.In)
		}
	}

	sort.Sort(integerFlagsByName(result))
	return result
}

func getIntegerFlagsFromStruct(param lib.APIParam, definition lib.StructDef) []integerFlag {
	//fmt.Println("        getting integer flags")
	result := []integerFlag{}
	for _, prop := range definition.Properties {
		if prop.Type != "integer" {
			continue
		}
		//fmt.Printf("          %s\n", prop.Name)
		var f integerFlag
		f.VarName = lib.TitleCase(prop.Name)
		f.LongOption = lib.OptionCase(prop.Name)
		f.DefaultValueSpecified = prop.Default != nil
		f.DefaultValue = prop.GetDefaultValueAsInt64()
		f.Format = prop.Format
		f.ShortHelp = trimTemplate(prop.Description)
		f.In = "body"
		f.Name = prop.Name
		f.Required = prop.Required
		result = append(result, f)
	}
	return result
}

func getFloatFlags(parameters []lib.APIParam, definitions map[string]lib.StructDef) []floatFlag {
	result := []floatFlag{}
	for _, param := range parameters {
		switch param.In {
		case "path", "query", "header":
			if param.Type != "number" {
				continue
			}
			var f floatFlag
			f.VarName = lib.TitleCase(param.Name)
			f.LongOption = lib.OptionCase(param.Name)
			f.DefaultValueSpecified = param.Default != nil
			f.DefaultValue = param.GetDefaultValueAsFloat()
			f.ShortHelp = trimTemplate(param.Description)
			f.Name = param.Name
			f.In = param.In
			f.Required = param.Required
			result = append(result, f)
		case "body":
			var s []floatFlag
			if param.Schema.Ref != "" {
				s = getFloatFlagsFromStruct(param, definitions[getStructNameFromReference(param.Schema.Ref)])
			} else if param.Schema.Type == "array" {
				//fmt.Println("[WARN] array is not supported yet.")
			}
			result = append(result, s...)
		default:
			fmt.Printf("[WARN] parameters in '%s' is not supported\n", param.In)
		}
	}

	sort.Sort(floatFlagsByName(result))
	return result
}

func getFloatFlagsFromStruct(param lib.APIParam, definition lib.StructDef) []floatFlag {
	//fmt.Println("        getting number flags")
	result := []floatFlag{}
	for _, prop := range definition.Properties {
		if prop.Type != "number" {
			continue
		}
		//fmt.Printf("          %s\n", prop.Name)
		var f floatFlag
		f.VarName = lib.TitleCase(prop.Name)
		f.LongOption = lib.OptionCase(prop.Name)
		f.DefaultValueSpecified = prop.Default != nil
		f.DefaultValue = prop.GetDefaultValueAsFloat()
		f.Format = prop.Format
		f.ShortHelp = trimTemplate(prop.Description)
		f.In = "body"
		f.Name = prop.Name
		f.Required = prop.Required
		result = append(result, f)
	}
	return result
}

func getBoolFlags(parameters []lib.APIParam, definitions map[string]lib.StructDef) []boolFlag {
	result := []boolFlag{}
	for _, param := range parameters {
		switch param.In {
		case "path", "query", "header":
			if param.Type != "boolean" {
				continue
			}
			var f boolFlag
			f.VarName = lib.TitleCase(param.Name)
			f.LongOption = lib.OptionCase(param.Name)
			f.DefaultValueSpecified = param.Default != nil
			f.DefaultValue = param.GetDefaultValueAsBool()
			f.ShortHelp = trimTemplate(param.Description)
			f.Name = param.Name
			f.In = param.In
			f.Required = param.Required
			result = append(result, f)
		case "body":
			var s []boolFlag
			if param.Schema.Ref != "" {
				s = getBoolFlagsFromStruct(param, definitions[getStructNameFromReference(param.Schema.Ref)])
			} else if param.Schema.Type == "array" {
				//fmt.Println("[WARN] array is not supported yet.")
			}
			result = append(result, s...)
		default:
			fmt.Printf("[WARN] parameters in '%s' is not supported\n", param.In)
		}
	}

	sort.Sort(boolFlagsByName(result))
	return result
}

func getBoolFlagsFromStruct(param lib.APIParam, definition lib.StructDef) []boolFlag {
	//fmt.Println("        getting number flags")
	result := []boolFlag{}
	for _, prop := range definition.Properties {
		if prop.Type != "boolean" {
			continue
		}
		//fmt.Printf("          %s\n", prop.Name)
		var f boolFlag
		f.VarName = lib.TitleCase(prop.Name)
		f.LongOption = lib.OptionCase(prop.Name)
		f.DefaultValueSpecified = prop.Default != nil
		f.DefaultValue = prop.GetDefaultValueAsBool()
		f.Format = prop.Format
		f.ShortHelp = trimTemplate(prop.Description)
		f.In = "body"
		f.Name = prop.Name
		f.Required = prop.Required
		result = append(result, f)
	}
	return result
}

func doesRequiredFlagExist(m lib.APIMethod, params []lib.APIParam, definitions map[string]lib.StructDef) bool {
	for _, f := range getStringFlags(m, params, definitions) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getStringSliceFlags(params, definitions) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getIntegerFlags(params, definitions) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getFloatFlags(params, definitions) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getBoolFlags(params, definitions) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	return false
}

// convert "#/definitions/StructName" to "StructName"
func getStructNameFromReference(ref string) string {
	var defPrefix = "#/definitions/"
	if strings.Index(ref, defPrefix) == 0 {
		return strings.Replace(ref, defPrefix, "", 1)
	}
	return ref
}

func getPaginationKeyHeaderInResponse(p *lib.Pagination) string {
	if p == nil {
		return ""
	}

	return p.Response.Header
}

func getPaginationRequestparameterInQuery(p *lib.Pagination) string {
	if p == nil {
		return ""
	}

	return p.Request.Param
}
