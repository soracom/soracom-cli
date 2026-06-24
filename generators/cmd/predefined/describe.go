package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DescribeCmd)
}

// DescribeCmd defines the 'describe' subcommand, which exposes machine-readable
// schema information so that AI agents can discover what a command accepts at
// runtime without parsing --help text.
var DescribeCmd = &cobra.Command{
	Use:   "describe [command...]",
	Short: TRCLI("cli.describe.summary"),
	Long:  TRCLI("cli.describe.description"),
	RunE:  describeRunE,
}

// commandDescription is the machine-readable description of a single leaf command.
type commandDescription struct {
	Command     string           `json:"command"`
	Method      string           `json:"method"`
	Path        string           `json:"path"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Deprecated  bool             `json:"deprecated,omitempty"`
	Parameters  []parameterDesc  `json:"parameters,omitempty"`
	RequestBody *requestBodyDesc `json:"requestBody,omitempty"`
	Response    *responseDesc    `json:"response,omitempty"`
}

// responseDesc describes the shape of a command's success response so that an
// agent can discover the exact field names to pass to --fields without guessing.
type responseDesc struct {
	Type   string      `json:"type,omitempty"`
	Schema string      `json:"schema,omitempty"`
	Fields []fieldDesc `json:"fields,omitempty"`
}

type fieldDesc struct {
	Name        string      `json:"name"`
	Type        string      `json:"type,omitempty"`
	Schema      string      `json:"schema,omitempty"`
	Description string      `json:"description,omitempty"`
	Fields      []fieldDesc `json:"fields,omitempty"`
}

// commandSummary is the compact form used when listing all commands.
type commandSummary struct {
	Command string `json:"command"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Summary string `json:"summary,omitempty"`
}

type parameterDesc struct {
	Name        string        `json:"name"`
	Option      string        `json:"option,omitempty"`
	In          string        `json:"in"`
	Type        string        `json:"type,omitempty"`
	Schema      string        `json:"schema,omitempty"`
	Required    bool          `json:"required"`
	Description string        `json:"description,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`
	Default     interface{}   `json:"default,omitempty"`
}

type requestBodyDesc struct {
	ContentType string          `json:"contentType,omitempty"`
	Type        string          `json:"type,omitempty"`
	Schema      string          `json:"schema,omitempty"`
	Required    []string        `json:"required,omitempty"`
	Properties  []parameterDesc `json:"properties,omitempty"`
	Example     interface{}     `json:"example,omitempty"`
}

type describeEntry struct {
	command string
	method  string
	path    string
	op      *openapi3.Operation
}

func describeRunE(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	entries, err := loadDescribeEntries()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return describeAll(entries)
	}

	target := strings.Join(args, " ")
	for _, e := range entries {
		if e.command == target {
			return prettyPrintObjectAsJSON(buildCommandDescription(e), os.Stdout)
		}
	}

	return fmt.Errorf("unknown command: '%s'. Run 'soracom describe' to list all commands", target)
}

func describeAll(entries []describeEntry) error {
	summaries := make([]commandSummary, 0, len(entries))
	for _, e := range entries {
		summaries = append(summaries, commandSummary{
			Command: e.command,
			Method:  strings.ToUpper(e.method),
			Path:    e.path,
			Summary: firstLine(operationSummary(e.op)),
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Command < summaries[j].Command
	})
	return prettyPrintObjectAsJSON(summaries, os.Stdout)
}

func loadDescribeEntries() ([]describeEntry, error) {
	lang := getSelectedLanguage()
	if !supportedLanguages[lang] {
		lang = defaultLang
	}

	files := []string{
		"assets/soracom-api." + lang + ".yaml",
		"assets/sandbox/soracom-sandbox-api." + lang + ".yaml",
	}

	var entries []describeEntry
	for _, f := range files {
		b, err := assets.ReadFile(f)
		if err != nil {
			// sandbox definition may be optional in some builds; skip silently
			continue
		}

		loader := openapi3.NewLoader()
		apiDef, err := loader.LoadFromData(b)
		if err != nil {
			return nil, fmt.Errorf("unable to parse embedded API definition '%s': %w", f, err)
		}

		for path, pathItem := range apiDef.Paths.Map() {
			for method, op := range pathItem.Operations() {
				for _, command := range getCLICommandsFromOperation(op) {
					entries = append(entries, describeEntry{
						command: command,
						method:  method,
						path:    path,
						op:      op,
					})
				}
			}
		}
	}

	return entries, nil
}

func buildCommandDescription(e describeEntry) commandDescription {
	d := commandDescription{
		Command:     e.command,
		Method:      strings.ToUpper(e.method),
		Path:        e.path,
		Summary:     operationSummary(e.op),
		Description: singleLine(e.op.Description),
		Deprecated:  e.op.Deprecated,
		Parameters:  buildParameterDescs(e.op.Parameters),
	}
	d.RequestBody = buildRequestBodyDesc(e.op.RequestBody)
	d.Response = buildResponseDesc(e.op.Responses)
	return d
}

// buildResponseDesc describes the success response body, fully expanding nested
// object fields so an agent can construct any --fields path (e.g.
// "sessionStatus.cell.mcc") without guessing or reading external docs.
func buildResponseDesc(responses *openapi3.Responses) *responseDesc {
	if responses == nil {
		return nil
	}

	m := responses.Map()
	var resp *openapi3.ResponseRef
	for _, code := range []string{"200", "201", "202", "203", "2XX", "default"} {
		if r, ok := m[code]; ok && r != nil {
			resp = r
			break
		}
	}
	if resp == nil || resp.Value == nil {
		return nil
	}

	for _, contentType := range orderedContentTypes(resp.Value.Content) {
		media := resp.Value.Content[contentType]
		if media == nil || media.Schema == nil || media.Schema.Value == nil {
			continue
		}

		desc := &responseDesc{
			Type:   schemaTypeString(media.Schema),
			Schema: schemaRefName(media.Schema),
		}

		fieldSchema := media.Schema
		if media.Schema.Value.Type.Is("array") && media.Schema.Value.Items != nil {
			item := media.Schema.Value.Items
			if name := schemaRefName(item); name != "" {
				desc.Type = "array of " + name
				desc.Schema = name
			}
			fieldSchema = item
		}

		desc.Fields = buildResponseFields(fieldSchema, map[*openapi3.Schema]bool{})
		return desc
	}

	return nil
}

// buildResponseFields lists the properties of a response schema, recursing into
// nested object/array-of-object fields. `seen` holds the schemas on the current
// ancestor path so a self-referential schema stops instead of recursing
// forever; the same schema may still appear in sibling branches. Expansion is
// naturally bounded by the schema (the deepest SORACOM response is a handful of
// levels and well under ~100 fields).
func buildResponseFields(s *openapi3.SchemaRef, seen map[*openapi3.Schema]bool) []fieldDesc {
	if s == nil || s.Value == nil {
		return nil
	}

	val := s.Value
	if val.Type.Is("array") && val.Items != nil {
		return buildResponseFields(val.Items, seen)
	}

	if seen[val] {
		return nil
	}
	seen[val] = true
	defer delete(seen, val)

	fields := make([]fieldDesc, 0, len(val.Properties))
	for name, prop := range val.Properties {
		if prop == nil || prop.Value == nil {
			continue
		}
		fields = append(fields, fieldDesc{
			Name:        name,
			Type:        schemaTypeString(prop),
			Schema:      schemaRefName(prop),
			Description: singleLine(prop.Value.Description),
			Fields:      buildResponseFields(prop, seen),
		})
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	return fields
}

func buildParameterDescs(params openapi3.Parameters) []parameterDesc {
	result := make([]parameterDesc, 0, len(params))
	for _, p := range params {
		if p.Value == nil {
			continue
		}

		option := ""
		if parameterHasFlag(p.Value) {
			option = lib.OptionCase(p.Value.Name)
		}

		result = append(result, parameterDesc{
			Name:        p.Value.Name,
			Option:      option,
			In:          p.Value.In,
			Type:        schemaTypeString(p.Value.Schema),
			Schema:      schemaRefName(p.Value.Schema),
			Required:    parameterRequired(p.Value),
			Description: singleLine(p.Value.Description),
			Enum:        schemaEnum(p.Value.Schema),
			Default:     schemaDefault(p.Value.Schema),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

// orderedContentTypes returns the request body's content types in a stable
// order so describe output is deterministic. application/json is preferred
// because it is the CLI's default content type and the structured form that
// carries named properties; the rest follow lexicographically. Without this the
// Go map iteration order would make describe output for the few multi
// content-type operations change between runs.
func orderedContentTypes(content openapi3.Content) []string {
	keys := make([]string, 0, len(content))
	for k := range content {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if (keys[i] == "application/json") != (keys[j] == "application/json") {
			return keys[i] == "application/json"
		}
		return keys[i] < keys[j]
	})
	return keys
}

// requestBodyHasRefSchema reports whether the generator would create per-field
// body flags. It mirrors getRequestBodySchema in gen_leaf_cmd.go, which returns
// a schema (and thus produces flags) only when the request body itself or any of
// its content schemas is a $ref; an inline body schema yields no field flags.
func requestBodyHasRefSchema(reqBody *openapi3.RequestBodyRef) bool {
	if reqBody == nil || reqBody.Value == nil {
		return false
	}
	if reqBody.Ref != "" {
		return true
	}
	for _, media := range reqBody.Value.Content {
		if media != nil && media.Schema != nil && media.Schema.Ref != "" {
			return true
		}
	}
	return false
}

func buildRequestBodyDesc(reqBody *openapi3.RequestBodyRef) *requestBodyDesc {
	if reqBody == nil || reqBody.Value == nil {
		return nil
	}

	for _, contentType := range orderedContentTypes(reqBody.Value.Content) {
		media := reqBody.Value.Content[contentType]
		if media == nil || media.Schema == nil || media.Schema.Value == nil {
			continue
		}

		desc := &requestBodyDesc{
			ContentType: contentType,
			Type:        schemaTypeString(media.Schema),
			Schema:      schemaRefName(media.Schema),
			Example:     requestBodyExample(media),
		}

		// The schema whose properties describe the payload shape: for an array
		// body (e.g. "put-tags" takes an array of TagUpdateRequest) it is the
		// item schema, so an agent can see each element's fields; otherwise it
		// is the body schema itself.
		propSchema := media.Schema
		// The generator only creates per-field body flags when the body schema is
		// referenced via $ref (getRequestBodySchema returns nil for an inline
		// schema); an array body is likewise sent raw. In those cases the fields
		// have no flag and must go through --body, so do not advertise an option.
		emitOptions := requestBodyHasRefSchema(reqBody)
		if media.Schema.Value.Type.Is("array") && media.Schema.Value.Items != nil {
			item := media.Schema.Value.Items
			if name := schemaRefName(item); name != "" {
				desc.Type = "array of " + name
				desc.Schema = name
			}
			propSchema = item
			emitOptions = false
		}

		if propSchema.Value != nil {
			desc.Required = propSchema.Value.Required
			desc.Properties = buildBodyProperties(propSchema.Value, emitOptions)
		}

		return desc
	}

	return nil
}

func buildBodyProperties(schema *openapi3.Schema, emitOptions bool) []parameterDesc {
	props := make([]parameterDesc, 0, len(schema.Properties))
	for propName, prop := range schema.Properties {
		if prop == nil || prop.Value == nil {
			continue
		}
		option := ""
		if emitOptions {
			option = bodyPropertyOption(propName, prop)
		}
		props = append(props, parameterDesc{
			Name:        propName,
			Option:      option,
			In:          "body",
			Type:        schemaTypeString(prop),
			Schema:      schemaRefName(prop),
			Required:    containsStringValue(schema.Required, propName),
			Description: singleLine(prop.Value.Description),
			Enum:        schemaEnum(prop),
			Default:     schemaDefault(prop),
		})
	}
	sort.Slice(props, func(i, j int) bool { return props[i].Name < props[j].Name })
	return props
}

// getCLICommandsFromOperation reads the x-soracom-cli extension that maps an
// operation to one or more CLI command names (e.g. "subscribers create").
func getCLICommandsFromOperation(op *openapi3.Operation) []string {
	raw, found := op.Extensions["x-soracom-cli"]
	if !found {
		return nil
	}

	b, err := json.Marshal(raw)
	if err != nil {
		return nil
	}

	var result []string
	if err := json.Unmarshal(b, &result); err != nil {
		return nil
	}
	return result
}

// bodyPropertyOption returns the CLI flag name for a request-body property, but
// only when the generator actually produces a flag for it. The generator emits
// flags for scalar body fields (string/integer/number/boolean) and arrays of
// strings; object/map and other complex fields get no flag and must be supplied
// through --body. Returning "" for those avoids telling an agent to use a
// --flag that does not exist.
func bodyPropertyOption(propName string, prop *openapi3.SchemaRef) string {
	if !bodyPropertyHasFlag(prop) {
		return ""
	}
	return lib.OptionCase(cliParamNameOf(propName, prop))
}

func bodyPropertyHasFlag(prop *openapi3.SchemaRef) bool {
	if prop == nil || prop.Value == nil || prop.Value.Type == nil {
		return false
	}
	types := prop.Value.Type.Slice()
	// The generator only flags fields whose schema has exactly one type (its
	// Types.Is / getTypeOfParam helpers reject multi-type schemas such as
	// ["string","null"]), so require a single type here too.
	if len(types) != 1 {
		return false
	}
	switch types[0] {
	case "string", "integer", "number", "boolean":
		return true
	case "array":
		items := prop.Value.Items
		return items != nil && items.Value != nil && items.Value.Type.Is("string")
	}
	return false
}

// parameterHasFlag reports whether the generator produces a flag for a
// path/query/header parameter. This mirrors getStringFlags / getStringSliceFlags
// / getIntegerFlags / getFloatFlags / getBoolFlags in
// generators/cmd/src/gen_leaf_cmd.go — the generator is the source of truth, and
// TestDescribeOptionsMatchGeneratedFlags asserts describe stays aligned with it.
func parameterHasFlag(p *openapi3.Parameter) bool {
	if p == nil || p.Schema == nil || p.Schema.Value == nil || p.Schema.Value.Type == nil {
		return false
	}
	types := p.Schema.Value.Type.Slice()
	// Match the generator, which only flags single-type schemas.
	if len(types) != 1 {
		return false
	}
	switch types[0] {
	case "string", "integer", "number", "boolean":
		return true
	case "array":
		// Only arrays of strings in the query become a (string slice) flag.
		if p.In != "query" {
			return false
		}
		items := p.Schema.Value.Items
		return items != nil && items.Value != nil && items.Value.Type.Is("string")
	}
	return false
}

// parameterRequired mirrors the generator's special case: operator_id is never a
// required flag because generated commands auto-fill it from the API token when
// omitted (see getStringFlags in generators/cmd/src/gen_leaf_cmd.go).
func parameterRequired(p *openapi3.Parameter) bool {
	if p.Name == "operator_id" {
		return false
	}
	return p.Required
}

func cliParamNameOf(propName string, prop *openapi3.SchemaRef) string {
	if prop.Value.Extensions != nil {
		if raw, found := prop.Value.Extensions["x-soracom-cli-param-name"]; found {
			if s, ok := raw.(string); ok {
				return s
			}
		}
	}
	return propName
}

func operationSummary(op *openapi3.Operation) string {
	if op == nil {
		return ""
	}
	return op.Summary
}

func schemaTypeString(s *openapi3.SchemaRef) string {
	if s == nil || s.Value == nil || s.Value.Type == nil {
		return ""
	}
	types := s.Value.Type.Slice()
	if len(types) == 0 {
		return ""
	}
	if types[0] == "array" && s.Value.Items != nil {
		return "array of " + schemaTypeString(s.Value.Items)
	}
	// A free-form object whose keys are arbitrary (additionalProperties) is a
	// map, not a fixed struct. Surface that so an agent knows it must supply
	// caller-defined keys (e.g. tags is "map of string") rather than looking
	// for fixed sub-properties that do not exist.
	if types[0] == "object" && s.Value.AdditionalProperties.Schema != nil {
		valueType := schemaTypeString(s.Value.AdditionalProperties.Schema)
		if valueType == "" {
			valueType = "object"
		}
		return "map of " + valueType
	}
	return types[0]
}

// schemaRefName returns the component schema name a property points at, when it
// is defined via $ref (e.g. "#/components/schemas/GroupConfiguration" ->
// "GroupConfiguration"). This gives an agent a stable name to reason about for
// object-typed fields whose structure is not inlined.
func schemaRefName(s *openapi3.SchemaRef) string {
	if s == nil || s.Ref == "" {
		return ""
	}
	ref := s.Ref
	if i := strings.LastIndex(ref, "/"); i >= 0 {
		return ref[i+1:]
	}
	return ref
}

// requestBodyExample returns the example payload for a request body, preferring
// the media-type example and falling back to the schema-level example. A
// concrete example is often the most actionable thing an agent can use to build
// a correct --body, especially for free-form/object fields.
func requestBodyExample(media *openapi3.MediaType) interface{} {
	if media == nil {
		return nil
	}
	if media.Example != nil {
		return media.Example
	}
	if media.Schema != nil && media.Schema.Value != nil {
		return media.Schema.Value.Example
	}
	return nil
}

func schemaEnum(s *openapi3.SchemaRef) []interface{} {
	if s == nil || s.Value == nil {
		return nil
	}
	return s.Value.Enum
}

func schemaDefault(s *openapi3.SchemaRef) interface{} {
	if s == nil || s.Value == nil {
		return nil
	}
	return s.Value.Default
}

func firstLine(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.IndexAny(s, "\r\n"); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return s
}

var whitespaceRun = regexp.MustCompile(`\s+`)

// singleLine collapses a (possibly multi-line) description into one line by
// replacing any run of whitespace with a single space. Unlike firstLine it
// keeps the full text, so multi-line descriptions that carry the meaning of
// each enum value (e.g. "VPG Type. - 14 : Type-E - 15 : Type-F ...") are not
// lost — matching the richness of the command's --help text.
func singleLine(s string) string {
	return strings.TrimSpace(whitespaceRun.ReplaceAllString(s, " "))
}

func containsStringValue(ss []string, target string) bool {
	for _, s := range ss {
		if s == target {
			return true
		}
	}
	return false
}
