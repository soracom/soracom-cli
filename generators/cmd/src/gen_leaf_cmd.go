package main

import (
	"encoding/json"
	"math"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/soracom/soracom-cli/generators/lib"
)

func generateLeafCommands(apiDef *openapi3.T, templateDir, outputDir string) error {
	subCommandTemplate, err := openTemplateFile(templateDir, "leaf.gotmpl")
	if err != nil {
		return err
	}

	for path, pathItem := range apiDef.Paths {
		for method, op := range pathItem.Operations() {
			err := generateCommandFiles(apiDef, path, method, op, subCommandTemplate, outputDir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generateCommandFiles(apiDef *openapi3.T, path, method string, op *openapi3.Operation, tmpl *template.Template, outputDir string) error {
	for _, commandName := range getCLICommands(op) {
		// commandName is a space separated subcommands list. e.g. "subscribers list"
		filename := lib.SnakeCase(commandName)
		f, err := openOutputFile(outputDir, filename+".go")
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				lib.WarnfStderr("unable to close a file '%s'", f.Name())
			}
		}()

		pagination := getXSoracomCliPagination(op)
		tag := ""
		if len(op.Tags) > 0 {
			tag = op.Tags[0]
		}

		a := commandArgs{
			Use:                               getLast(commandName),
			Short:                             path + ":" + strings.ToLower(method) + ":summary",
			Long:                              path + ":" + strings.ToLower(method) + ":description",
			CommandVariableName:               getCommandVariableName(commandName),
			ParentCommandVariableName:         getParentCommandVariableName(commandName),
			RequireAuth:                       op.Security != nil,
			RequireOperatorID:                 isOperatorIDRequired(apiDef, op.Parameters, op.RequestBody),
			BodyExists:                        op.RequestBody != nil,
			SendBodyRaw:                       isBodyArray(op.RequestBody) || isBodyBinary(op.RequestBody),
			ResponseBodyRaw:                   isResponseBodyRaw(path, method),
			Method:                            strings.ToUpper(method),
			BasePath:                          getBasePath(apiDef.Servers),
			Path:                              path,
			PathParamsExist:                   doPathParamsExist(op.Parameters),
			QueryParamsExist:                  doQueryParamsExist(op.Parameters),
			StringFlags:                       getStringFlags(apiDef, path, op.Parameters, op.RequestBody),
			StringSliceFlags:                  getStringSliceFlags(apiDef, path, op.Parameters, op.RequestBody),
			IntegerFlags:                      getIntegerFlags(apiDef, op.Parameters, op.RequestBody),
			FloatFlags:                        getFloatFlags(apiDef, op.Parameters, op.RequestBody),
			BoolFlags:                         getBoolFlags(apiDef, op.Parameters, op.RequestBody),
			RequiredFlagExists:                doesRequiredFlagExist(apiDef, path, op.Parameters, op.RequestBody),
			PaginationAvailable:               pagination != nil,
			PaginationKeyHeaderInResponse:     getPaginationResponseHeader(pagination),
			PaginationRequestParameterInQuery: getPaginationRequestParam(pagination),
			Deprecated:                        op.Deprecated,
			AlternativeCommand:                getXSoracomAlternativeCli(op),
			HasArrayResponse:                  hasArrayResponse(op.Responses),
			Tag:                               tag,
			OperationID:                       op.OperationID,
		}
		if a.Method == "POST" || a.Method == "PUT" {
			if doesContentTypeParamExist(op.Parameters) {
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

func getLast(s string) string {
	ss := strings.Split(s, " ")
	return ss[len(ss)-1]
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

func isOperatorIDRequired(apiDef *openapi3.T, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) bool {
	for _, param := range parameters {
		if param.Value == nil {
			lib.WarnfStderr("param.Value == nil. We might need to support param.Ref?\n")
			continue
		}

		switch param.Value.In {
		case "path":
			if param.Value.Name != "operator_id" {
				continue
			}
			if getTypeOfParam(param) != "string" {
				continue
			}
			return true
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			if isRequiredField(schema, "operatorId") {
				return true
			}
		}
	}

	return false
}

func getTypeOfParam(param *openapi3.ParameterRef) string {
	if param == nil {
		lib.WarnfStderr("param == nil\n")
		return ""
	}
	if param.Value == nil {
		lib.WarnfStderr("param.Value == nil. param == %#v\n", param)
		return ""
	}
	if param.Value.Schema == nil {
		lib.WarnfStderr("param.Value.Schema == nil. param.Value == %#v\n", param.Value)
		return ""
	}
	if param.Value.Schema.Value == nil {
		lib.WarnfStderr("param.Value.Schema.Value == nil. param.Value.Schema == %#v\n", param.Value.Schema)
		return ""
	}

	return param.Value.Schema.Value.Type
}

func getTypeOfArrayItem(param *openapi3.ParameterRef) string {
	if getTypeOfParam(param) != "array" {
		return ""
	}

	if param.Value.Schema.Value.Items == nil || param.Value.Schema.Value.Items.Value == nil {
		return ""
	}

	return param.Value.Schema.Value.Items.Value.Type
}

func isTypeOfParamArrayOfString(param *openapi3.ParameterRef) bool {
	if getTypeOfParam(param) != "array" {
		return false
	}

	if getTypeOfArrayItem(param) != "string" {
		return false
	}

	return true
}

func getRequestBodySchema(apiDef *openapi3.T, reqBody *openapi3.RequestBodyRef) *openapi3.SchemaRef {
	if apiDef == nil || reqBody == nil {
		return nil
	}

	if reqBody.Ref != "" {
		return dereferenceSchema(apiDef, reqBody.Ref)
	}
	for _, content := range reqBody.Value.Content {
		if content.Schema == nil {
			continue
		}

		if content.Schema.Ref != "" {
			return dereferenceSchema(apiDef, content.Schema.Ref)
		}
	}
	return nil
}

func dereferenceSchema(apiDef *openapi3.T, ref string) *openapi3.SchemaRef {
	name := getStructNameFromReference(ref)
	schema := apiDef.Components.Schemas[name]
	if schema == nil && strings.HasPrefix(ref, "#/components/requestBodies") {
		lib.WarnfStderr("dereferenceSchema(): schema %s not found. trying to find in requestBodies with name %s\n", ref, name)
		rbr, found := apiDef.Components.RequestBodies[name]
		if !found {
			lib.WarnfStderr("dereferenceSchema(): schema %s not found in requestBodies.\n", ref)
			lib.WarnfStderr("dereferenceSchema(): requestBodies: %s\n", toJSON(apiDef.Components.RequestBodies))
			return nil
		}
		for _, content := range rbr.Value.Content {
			lib.WarnfStderr("dereferenceSchema(): content: %#v (%s)\n", content, toJSON(content))
			if content.Schema != nil && content.Schema.Ref != "" {
				name = getStructNameFromReference(content.Schema.Ref)
				schema = apiDef.Components.Schemas[name]
			}
		}
	}
	return schema
}

// convert "#/components/schemas/StructName" or "#/components/requestBodies/StructName" to "StructName"
func getStructNameFromReference(ref string) string {
	prefixes := []string{"#/components/schemas/", "#/components/requestBodies/"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(ref, prefix) {
			return strings.Replace(ref, prefix, "", 1)
		}
	}
	return ref
}

func isBodyArray(body *openapi3.RequestBodyRef) bool {
	if body == nil {
		return false
	}
	for _, content := range body.Value.Content {
		if content.Schema.Value.Type == "array" {
			return true
		}
	}
	return false
}

func isBodyBinary(body *openapi3.RequestBodyRef) bool {
	if body == nil {
		return false
	}
	for _, content := range body.Value.Content {
		if content.Schema.Value.Type == "string" && content.Schema.Value.Format == "binary" {
			return true
		}
	}
	return false
}

func isResponseBodyRaw(path, method string) bool {
	if strings.ToUpper(method) == "GET" && path == "/files/{scope}/{path}" {
		return true
	}

	// response contains signed URL, which should not be modified while Go json serializer prettifies '&' to '\u0026'
	if strings.ToUpper(method) == "POST" && strings.HasSuffix(path, "/export") {
		return true
	}

	return false
}

func getBasePath(servers openapi3.Servers) string {
	if len(servers) <= 0 {
		return "/"
	}
	u, err := url.Parse(servers[0].URL)
	if err != nil {
		lib.WarnfStderr("unable to parse server url: %s\n", servers[0].URL)
		return "/"
	}

	return u.Path
}

func doPathParamsExist(parameters openapi3.Parameters) bool {
	for _, param := range parameters {
		if param.Value.In == "path" {
			return true
		}
	}
	return false
}

func doQueryParamsExist(parameters openapi3.Parameters) bool {
	for _, param := range parameters {
		if param.Value.In == "query" {
			return true
		}
	}
	return false
}

func getStringFlags(apiDef *openapi3.T, path string, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) []stringFlag {
	result := []stringFlag{}
	for _, param := range parameters {
		switch param.Value.In {
		case "path", "query", "header":
			if getTypeOfParam(param) != "string" {
				continue
			}
			required := param.Value.Required
			if param.Value.Name == "operator_id" {
				required = false
			}
			f := stringFlagFromParameter(param, path, required)
			result = append(result, f)
		default:
			lib.WarnfStderr("parameters in '%s' is not supported\n", param.Value.In)
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			s := getStringFlagsFromStruct(schema)
			result = append(result, s...)
		}
	}

	sort.Sort(stringFlagsByName(result))

	return result
}

func getStringFlagsFromStruct(schema *openapi3.SchemaRef) []stringFlag {
	result := []stringFlag{}

	for propName, prop := range getPropertiesFromSchema(schema) {
		if prop.Value.Type != "string" {
			continue
		}
		f := stringFlagFromProperty(propName, prop, "body", containsString(schema.Value.Required, propName))
		result = append(result, f)
	}
	return result
}

func getStringSliceFlags(apiDef *openapi3.T, path string, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) []stringFlag {
	result := []stringFlag{}
	for _, param := range parameters {
		switch param.Value.In {
		case "query":
			if !isTypeOfParamArrayOfString(param) {
				continue
			}
			f := stringFlagFromParameter(param, path, param.Value.Required)
			result = append(result, f)
		case "path", "body", "header":
			continue
		default:
			lib.WarnfStderr("parameters in '%s' is not supported\n", param.Value.In)
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			s := getStringSliceFlagsFromStruct(schema)
			result = append(result, s...)
		}
	}

	sort.Sort(stringFlagsByName(result))
	return result
}

func getStringSliceFlagsFromStruct(schema *openapi3.SchemaRef) []stringFlag {
	result := []stringFlag{}
	for propName, prop := range getPropertiesFromSchema(schema) {
		if prop.Value.Type != "array" || prop.Value.Items == nil || prop.Value.Items.Value.Type != "string" {
			continue
		}
		f := stringFlagFromProperty(propName, prop, "body", containsString(schema.Value.Required, propName))
		result = append(result, f)
	}
	return result
}

func getIntegerFlags(apiDef *openapi3.T, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) []integerFlag {
	result := []integerFlag{}
	for _, param := range parameters {
		switch param.Value.In {
		case "path", "query", "header":
			if getTypeOfParam(param) != "integer" {
				continue
			}
			f := integerFlagFromParameter(param)
			result = append(result, f)
		default:
			lib.WarnfStderr("parameters in '%s' is not supported\n", param.Value.In)
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			s := getIntegerFlagsFromStruct(schema)
			result = append(result, s...)
		}
	}

	sort.Sort(integerFlagsByName(result))
	return result
}

func getIntegerFlagsFromStruct(schema *openapi3.SchemaRef) []integerFlag {
	result := []integerFlag{}
	for propName, prop := range getPropertiesFromSchema(schema) {
		if prop.Value.Type != "integer" {
			continue
		}
		f := integerFlagFromProperty(propName, prop, "body", containsString(schema.Value.Required, propName))
		result = append(result, f)
	}
	return result
}

func getFloatFlags(apiDef *openapi3.T, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) []floatFlag {
	result := []floatFlag{}
	for _, param := range parameters {
		switch param.Value.In {
		case "path", "query", "header":
			if getTypeOfParam(param) != "number" {
				continue
			}
			f := floatFlagFromParameter(param)
			result = append(result, f)
		default:
			lib.WarnfStderr("parameters in '%s' is not supported\n", param.Value.In)
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			s := getFloatFlagsFromStruct(schema)
			result = append(result, s...)
		}
	}

	sort.Sort(floatFlagsByName(result))
	return result
}

func getFloatFlagsFromStruct(schema *openapi3.SchemaRef) []floatFlag {
	result := []floatFlag{}
	for propName, prop := range getPropertiesFromSchema(schema) {
		if prop.Value.Type != "number" {
			continue
		}
		f := floatFlagFromProperty(propName, prop, "body", containsString(schema.Value.Required, propName))
		result = append(result, f)
	}
	return result
}

func getBoolFlags(apiDef *openapi3.T, parameters openapi3.Parameters, reqBody *openapi3.RequestBodyRef) []boolFlag {
	result := []boolFlag{}
	for _, param := range parameters {
		switch param.Value.In {
		case "path", "query", "header":
			if getTypeOfParam(param) != "boolean" {
				continue
			}
			f := boolFlagFromParameter(param)
			result = append(result, f)
		default:
			lib.WarnfStderr("parameters in '%s' is not supported\n", param.Value.In)
		}
	}

	if reqBody != nil {
		schema := getRequestBodySchema(apiDef, reqBody)
		if schema != nil {
			s := getBoolFlagsFromStruct(schema)
			result = append(result, s...)
		}
	}

	sort.Sort(boolFlagsByName(result))
	return result
}

func getBoolFlagsFromStruct(schema *openapi3.SchemaRef) []boolFlag {
	result := []boolFlag{}
	for propName, prop := range getPropertiesFromSchema(schema) {
		if prop.Value.Type != "boolean" {
			continue
		}
		f := boolFlagFromProperty(propName, prop, "body", containsString(schema.Value.Required, propName))
		result = append(result, f)
	}
	return result
}

func getPropertiesFromSchema(schema *openapi3.SchemaRef) openapi3.Schemas {
	if schema == nil {
		return nil
	}
	if schema.Value == nil {
		return nil
	}
	if schema.Value.Properties == nil {
		return nil
	}
	return schema.Value.Properties
}

func stringFlagFromParameter(param *openapi3.ParameterRef, path string, required bool) stringFlag {
	return stringFlag{
		VarName:                lib.TitleCase(param.Value.Name),
		LongOption:             lib.OptionCase(param.Value.Name),
		DefaultValueSpecified:  param.Value.Schema.Value.Default != nil,
		DefaultValue:           getDefaultValueAsString(param.Value.Schema),
		ShortHelp:              convertDescriptionToShortHelp(param.Value.Description),
		In:                     param.Value.In,
		Name:                   param.Value.Name,
		Required:               required,
		HarvestFilesPathEscape: isHarvestFilesPathEscapeRequired(path, param),
	}
}

func stringFlagFromProperty(propName string, prop *openapi3.SchemaRef, in string, required bool) stringFlag {
	return stringFlag{
		VarName:               lib.TitleCase(propName),
		LongOption:            lib.OptionCase(propName),
		DefaultValueSpecified: prop.Value.Default != nil,
		DefaultValue:          getDefaultValueAsString(prop),
		ShortHelp:             convertDescriptionToShortHelp(prop.Value.Description),
		In:                    in,
		Name:                  propName,
		Required:              required,
	}
}

func integerFlagFromParameter(param *openapi3.ParameterRef) integerFlag {
	return integerFlag{
		VarName:               lib.TitleCase(param.Value.Name),
		LongOption:            lib.OptionCase(param.Value.Name),
		DefaultValueSpecified: param.Value.Schema.Value.Default != nil,
		DefaultValue:          getDefaultValueAsInt64(param.Value.Schema),
		ShortHelp:             convertDescriptionToShortHelp(param.Value.Description),
		In:                    param.Value.In,
		Name:                  param.Value.Name,
		Required:              param.Value.Required,
	}
}

func integerFlagFromProperty(propName string, prop *openapi3.SchemaRef, in string, required bool) integerFlag {
	return integerFlag{
		VarName:               lib.TitleCase(propName),
		LongOption:            lib.OptionCase(propName),
		DefaultValueSpecified: prop.Value.Default != nil,
		DefaultValue:          getDefaultValueAsInt64(prop),
		Format:                prop.Value.Format,
		ShortHelp:             convertDescriptionToShortHelp(prop.Value.Description),
		In:                    in,
		Name:                  propName,
		Required:              required,
	}
}

func floatFlagFromParameter(param *openapi3.ParameterRef) floatFlag {
	return floatFlag{
		VarName:               lib.TitleCase(param.Value.Name),
		LongOption:            lib.OptionCase(param.Value.Name),
		DefaultValueSpecified: param.Value.Schema.Value.Default != nil,
		DefaultValue:          getDefaultValueAsFloat(param.Value.Schema),
		ShortHelp:             convertDescriptionToShortHelp(param.Value.Description),
		In:                    param.Value.In,
		Name:                  param.Value.Name,
		Required:              param.Value.Required,
	}
}

func floatFlagFromProperty(propName string, prop *openapi3.SchemaRef, in string, required bool) floatFlag {
	return floatFlag{
		VarName:               lib.TitleCase(propName),
		LongOption:            lib.OptionCase(propName),
		DefaultValueSpecified: prop.Value.Default != nil,
		DefaultValue:          getDefaultValueAsFloat(prop),
		Format:                prop.Value.Format,
		ShortHelp:             convertDescriptionToShortHelp(prop.Value.Description),
		In:                    in,
		Name:                  propName,
		Required:              required,
	}
}

func boolFlagFromParameter(param *openapi3.ParameterRef) boolFlag {
	return boolFlag{
		VarName:               lib.TitleCase(param.Value.Name),
		LongOption:            lib.OptionCase(param.Value.Name),
		DefaultValueSpecified: param.Value.Schema.Value.Default != nil,
		DefaultValue:          getDefaultValueAsBool(param.Value.Schema),
		ShortHelp:             convertDescriptionToShortHelp(param.Value.Description),
		In:                    param.Value.In,
		Name:                  param.Value.Name,
		Required:              param.Value.Required,
	}
}

func boolFlagFromProperty(propName string, prop *openapi3.SchemaRef, in string, required bool) boolFlag {
	return boolFlag{
		VarName:               lib.TitleCase(propName),
		LongOption:            lib.OptionCase(propName),
		DefaultValueSpecified: prop.Value.Default != nil,
		DefaultValue:          getDefaultValueAsBool(prop),
		Format:                prop.Value.Format,
		ShortHelp:             convertDescriptionToShortHelp(prop.Value.Description),
		In:                    in,
		Name:                  propName,
		Required:              required,
	}
}

func doesRequiredFlagExist(apiDef *openapi3.T, path string, params openapi3.Parameters, reqBody *openapi3.RequestBodyRef) bool {
	for _, f := range getStringFlags(apiDef, path, params, reqBody) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getStringSliceFlags(apiDef, path, params, reqBody) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getIntegerFlags(apiDef, params, reqBody) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getFloatFlags(apiDef, params, reqBody) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	for _, f := range getBoolFlags(apiDef, params, reqBody) {
		if f.Required && !f.DefaultValueSpecified {
			return true
		}
	}

	return false
}

func getXSoracomCliPagination(op *openapi3.Operation) *Pagination {
	paginationRaw, found := op.ExtensionProps.Extensions["x-soracom-cli-pagination"]
	if !found {
		return nil
	}

	jr, ok := paginationRaw.(json.RawMessage)
	if !ok {
		return nil
	}

	var p Pagination
	err := json.Unmarshal(jr, &p)
	if err != nil {
		lib.WarnfStderr("expected pagination info is defined in `x-soracom-cli-pagination`, but it was not\n")
		return nil
	}
	return &p
}

func getPaginationResponseHeader(p *Pagination) string {
	if p == nil {
		return ""
	}
	return p.Response.Header
}

func getPaginationRequestParam(p *Pagination) string {
	if p == nil {
		return ""
	}
	return p.Request.Param
}

func getXSoracomAlternativeCli(op *openapi3.Operation) string {
	altCLIRaw, found := op.ExtensionProps.Extensions["x-soracom-alternative-cli"]
	if !found {
		return ""
	}

	jr, ok := altCLIRaw.(json.RawMessage)
	if !ok {
		return ""
	}

	var altCLI string
	err := json.Unmarshal(jr, &altCLI)
	if err != nil {
		lib.WarnfStderr("expected string in `x-soracom-alternative-cli`, but it was not\n")
		return ""
	}

	return altCLI
}

func hasArrayResponse(responses openapi3.Responses) bool {
	for _, res := range responses {
		for _, content := range res.Value.Content {
			if content.Schema != nil && content.Schema.Value.Type == "array" {
				return true
			}
		}
	}

	return false
}

func doesContentTypeParamExist(parameters openapi3.Parameters) bool {
	for _, param := range parameters {
		if param.Value.Name == "content-type" {
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

func getDefaultValueAsString(schema *openapi3.SchemaRef) string {
	if schema.Value.Default == nil {
		return ""
	}
	s, ok := schema.Value.Default.(string)
	if !ok {
		return ""
	}
	return s
}

func getDefaultValueAsInt64(schema *openapi3.SchemaRef) int64 {
	if schema.Value.Default == nil {
		return 0
	}
	s, ok := schema.Value.Default.(float64)
	if !ok {
		return 0
	}
	return int64(math.Round(s))
}

func getDefaultValueAsFloat(schema *openapi3.SchemaRef) float64 {
	if schema.Value.Default == nil {
		return 0.0
	}
	s, ok := schema.Value.Default.(float64)
	if !ok {
		return 0.0
	}
	return s
}

func getDefaultValueAsBool(schema *openapi3.SchemaRef) bool {
	if schema.Value.Default == nil {
		return false
	}
	s, ok := schema.Value.Default.(bool)
	if !ok {
		return false
	}
	return s
}

func isHarvestFilesPathEscapeRequired(path string, param *openapi3.ParameterRef) bool {
	if path == "/files/{scope}/{path}" || path == "/files/{scope}/{path}/" {
		if param.Value.Name == "path" {
			return true
		}
	}
	return false
}

func isRequiredField(schema *openapi3.SchemaRef, fieldName string) bool {
	if schema == nil {
		lib.WarnfStderr("isRequiredField(): schema is nil\n")
		return false
	}
	if schema.Value == nil {
		lib.WarnfStderr("isRequiredField(): schema.Value is nil: %#v\n", schema)
		return false
	}
	if schema.Value.Required == nil {
		//lib.WarnfStderr("isRequiredField(): schema.Value.Required is nil: %#v\n", schema.Value)
		return false
	}

	return containsString(schema.Value.Required, fieldName)
}

func containsString(ss []string, target string) bool {
	for _, s := range ss {
		if s == target {
			return true
		}
	}
	return false
}

func convertDescriptionToShortHelp(description string) string {
	return replaceBackticksWithSingleQuotationMarks(trimTemplate(escapeDoubleQuotationMarks(removeLineBreaks(description))))
}

var lineBreaks = regexp.MustCompile(`[\r\n]`)

func removeLineBreaks(s string) string {
	return lineBreaks.ReplaceAllString(s, "")
}

func escapeDoubleQuotationMarks(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

var templateRegex = regexp.MustCompile(`^\s*\${(.+)}\s*$`)

func trimTemplate(s string) string {
	b1 := templateRegex.ReplaceAll([]byte(s), []byte("$1"))
	return string(b1)
}

func replaceBackticksWithSingleQuotationMarks(s string) string {
	return strings.ReplaceAll(s, "`", "'")
}
