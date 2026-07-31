package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	DescribeCmd.Flags().BoolVar(&describeAllCommands, "all", false, TRCLI("cli.describe.flags.all"))
	RootCmd.AddCommand(DescribeCmd)
}

// describeAllCommands makes 'describe' (with no arguments) list every command
// instead of just the top-level command groups.
var describeAllCommands bool

// DescribeCmd defines the 'describe' subcommand, which exposes machine-readable
// schema information so that AI agents can discover what a command accepts at
// runtime without parsing --help text.
var DescribeCmd = &cobra.Command{
	Use:   "describe [command...]",
	Short: TRCLI("cli.describe.summary"),
	Long:  TRCLI("cli.describe.description"),
	RunE:  describeRunE,
}

// commandDescription is the machine-readable description of a runnable command.
type commandDescription struct {
	Command     string           `json:"command"`
	Method      string           `json:"method,omitempty"`
	Path        string           `json:"path,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Deprecated  bool             `json:"deprecated,omitempty"`
	Parameters  []parameterDesc  `json:"parameters,omitempty"`
	RequestBody *requestBodyDesc `json:"requestBody,omitempty"`
	Response    *responseDesc    `json:"response,omitempty"`
	// Subcommands lists the commands nested under this one, for the few
	// commands (e.g. 'auth') that are both a runnable command and a group.
	Subcommands []commandSummary `json:"subcommands,omitempty"`
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
	Method  string `json:"method,omitempty"`
	Path    string `json:"path,omitempty"`
	Summary string `json:"summary,omitempty"`
}

// commandGroupSummary is the top-level listing form: one entry per top-level
// command (group) with the number of leaf commands it contains, so an agent can
// get an overview without loading every command into its context.
type commandGroupSummary struct {
	Command  string `json:"command"`
	Summary  string `json:"summary,omitempty"`
	Commands int    `json:"commands"`
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
	command      string
	method       string
	path         string
	op           *oapiOperation
	cobraCommand *cobra.Command
}

func describeRunE(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	entries, err := loadDescribeEntries(cmd)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		if describeAllCommands {
			return prettyPrintObjectAsJSON(summarizeEntries(entries), os.Stdout)
		}
		return prettyPrintObjectAsJSON(topLevelGroups(entries, rootCommandSummary), os.Stdout)
	}

	target := strings.Join(args, " ")
	for _, e := range entries {
		if e.command == target {
			d := buildCommandDescription(e)
			// A few commands (e.g. 'auth') are both runnable and a group;
			// list their nested commands so they stay discoverable.
			d.Subcommands = summarizeEntries(entriesUnderPrefix(entries, target))
			return prettyPrintObjectAsJSON(d, os.Stdout)
		}
	}

	// Not a leaf command; when it matches a command group (e.g. 'describe
	// sims'), list the leaf commands under it instead of erroring out.
	if under := entriesUnderPrefix(entries, target); len(under) > 0 {
		return prettyPrintObjectAsJSON(summarizeEntries(under), os.Stdout)
	}

	return fmt.Errorf("unknown command: '%s'. Run 'soracom describe' to list top-level commands, or 'soracom describe --all' to list all commands", target)
}

func summarizeEntries(entries []describeEntry) []commandSummary {
	summaries := make([]commandSummary, 0, len(entries))
	for _, e := range entries {
		summaries = append(summaries, commandSummary{
			Command: e.command,
			Method:  strings.ToUpper(e.method),
			Path:    e.path,
			Summary: firstLine(entrySummary(e)),
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Command < summaries[j].Command
	})
	return summaries
}

// entriesUnderPrefix returns the commands under a command group, e.g.
// "sims" matches "sims activate" but not "sims" itself nor "sims-x list".
func entriesUnderPrefix(entries []describeEntry, target string) []describeEntry {
	prefix := target + " "
	var under []describeEntry
	for _, e := range entries {
		if strings.HasPrefix(e.command, prefix) {
			under = append(under, e)
		}
	}
	return under
}

// topLevelGroups aggregates commands by their top-level command name.
// summaryOf resolves a top-level name to its human-readable summary (taken from
// the cobra command tree so it matches what 'soracom --help' shows).
func topLevelGroups(entries []describeEntry, summaryOf func(name string) string) []commandGroupSummary {
	counts := map[string]int{}
	for _, e := range entries {
		top := e.command
		if i := strings.Index(top, " "); i >= 0 {
			top = top[:i]
		}
		counts[top]++
	}

	groups := make([]commandGroupSummary, 0, len(counts))
	for name, n := range counts {
		groups = append(groups, commandGroupSummary{
			Command:  name,
			Summary:  summaryOf(name),
			Commands: n,
		})
	}
	sort.Slice(groups, func(i, j int) bool { return groups[i].Command < groups[j].Command })
	return groups
}

// rootCommandSummary returns the one-line summary of a top-level command as
// registered on the root cobra command (already localized via TRCLI/TRAPI).
func rootCommandSummary(name string) string {
	for _, c := range RootCmd.Commands() {
		if c.Name() == name {
			return firstLine(c.Short)
		}
	}
	return ""
}

func loadDescribeEntries(describeCommand *cobra.Command) ([]describeEntry, error) {
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

		apiDef, err := parseAPIDefinition(b)
		if err != nil {
			return nil, fmt.Errorf("unable to parse embedded API definition '%s': %w", f, err)
		}

		for path, pathItem := range apiDef.Paths {
			for _, e := range pathItem.operations() {
				for _, command := range getCLICommandsFromOperation(e.op) {
					entries = append(entries, describeEntry{
						command: command,
						method:  e.method,
						path:    path,
						op:      e.op,
					})
				}
			}
		}
	}

	entries = appendPredefinedDescribeEntries(entries, predefinedDescribeCommands(describeCommand))
	return entries, nil
}

// predefinedDescribeCommands explicitly lists the hand-written commands that
// describe exposes in addition to OpenAPI-backed commands. Keeping this
// separate from cobra's AddCommand calls avoids coupling command-tree
// construction to describe's discovery mechanism. The describe command itself
// is passed in to avoid an initialization cycle through DescribeCmd.RunE.
func predefinedDescribeCommands(describeCommand *cobra.Command) []*cobra.Command {
	return []*cobra.Command{
		CompletionCmd,
		completionBashCmd,
		completionZshCmd,
		ConfigureCmd,
		ConfigureGetCmd,
		ConfigureListCmd,
		ConfigureSandboxCmd,
		describeCommand,
		SelfUpdateCmd,
		UnconfigureCmd,
		VersionCmd,
	}
}

// appendPredefinedDescribeEntries supplements the OpenAPI-derived entries with
// visible hand-written commands. OpenAPI remains authoritative when a command
// exists in both sources because it carries richer request/response schemas.
func appendPredefinedDescribeEntries(entries []describeEntry, commands []*cobra.Command) []describeEntry {
	known := make(map[string]bool, len(entries))
	for _, e := range entries {
		known[e.command] = true
	}

	for _, command := range commands {
		name := relativeCommandName(command)
		if name == "" || known[name] || !commandAndParentsVisible(command) {
			continue
		}
		entries = append(entries, describeEntry{
			command:      name,
			cobraCommand: command,
		})
		known[name] = true
	}

	return entries
}

// relativeCommandName returns a command's space-separated path below RootCmd.
func relativeCommandName(command *cobra.Command) string {
	var parts []string
	for current := command; current != nil && current != RootCmd; current = current.Parent() {
		parts = append(parts, current.Name())
	}
	if len(parts) == 0 {
		return ""
	}
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, " ")
}

func commandAndParentsVisible(command *cobra.Command) bool {
	for current := command; current != nil && current != RootCmd; current = current.Parent() {
		if current.Hidden {
			return false
		}
	}
	return true
}

func buildCommandDescription(e describeEntry) commandDescription {
	d := commandDescription{
		Command: e.command,
		Method:  strings.ToUpper(e.method),
		Path:    e.path,
	}
	if e.op == nil {
		if e.cobraCommand != nil {
			d.Summary = firstLine(e.cobraCommand.Short)
			d.Description = singleLine(e.cobraCommand.Long)
			d.Deprecated = e.cobraCommand.Deprecated != ""
			d.Parameters = buildCobraFlagDescs(e.cobraCommand)
		}
		return d
	}

	d.Summary = operationSummary(e.op)
	d.Description = singleLine(e.op.Description)
	d.Deprecated = e.op.Deprecated
	d.Parameters = buildParameterDescs(e.op.Parameters)
	d.RequestBody = buildRequestBodyDesc(e.op.RequestBody)
	d.Response = buildResponseDesc(e.op.Responses)
	return d
}

func entrySummary(e describeEntry) string {
	if e.op != nil {
		return operationSummary(e.op)
	}
	if e.cobraCommand != nil {
		return e.cobraCommand.Short
	}
	return ""
}

func buildCobraFlagDescs(command *cobra.Command) []parameterDesc {
	var result []parameterDesc
	command.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "help" {
			return
		}
		result = append(result, parameterDesc{
			Name:        flag.Name,
			Option:      flag.Name,
			In:          "flag",
			Type:        cobraFlagType(flag.Value.Type()),
			Required:    false,
			Description: singleLine(flag.Usage),
		})
	})
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func cobraFlagType(flagType string) string {
	switch flagType {
	case "bool":
		return "boolean"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "stringArray", "stringSlice":
		return "array of string"
	default:
		return flagType
	}
}

// buildResponseDesc describes the success response body, fully expanding nested
// object fields so an agent can construct any --fields path (e.g.
// "sessionStatus.cell.mcc") without guessing or reading external docs.
func buildResponseDesc(responses map[string]*oapiResponse) *responseDesc {
	if responses == nil {
		return nil
	}

	var resp *oapiResponse
	for _, code := range []string{"200", "201", "202", "203", "2XX", "default"} {
		if r, ok := responses[code]; ok && r != nil {
			resp = r
			break
		}
	}
	if resp == nil {
		return nil
	}

	for _, contentType := range orderedContentTypes(resp.Content) {
		media := resp.Content[contentType]
		if media == nil || media.Schema == nil || media.Schema.Value == nil {
			continue
		}

		desc := &responseDesc{
			Type:   schemaTypeString(media.Schema),
			Schema: schemaRefName(media.Schema),
		}

		fieldSchema := media.Schema
		if media.Schema.Value.Type == "array" && media.Schema.Value.Items != nil {
			item := media.Schema.Value.Items
			if name := schemaRefName(item); name != "" {
				desc.Type = "array of " + name
				desc.Schema = name
			}
			fieldSchema = item
		}

		desc.Fields = buildResponseFields(fieldSchema, map[*oapiSchema]bool{})
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
func buildResponseFields(s *oapiSchemaRef, seen map[*oapiSchema]bool) []fieldDesc {
	if s == nil || s.Value == nil {
		return nil
	}

	val := s.Value
	if val.Type == "array" && val.Items != nil {
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

func buildParameterDescs(params []*oapiParameter) []parameterDesc {
	result := make([]parameterDesc, 0, len(params))
	for _, p := range params {
		if p == nil {
			continue
		}

		option := ""
		if parameterHasFlag(p) {
			option = lib.OptionCase(p.Name)
		}

		result = append(result, parameterDesc{
			Name:        p.Name,
			Option:      option,
			In:          p.In,
			Type:        schemaTypeString(p.Schema),
			Schema:      schemaRefName(p.Schema),
			Required:    parameterRequired(p),
			Description: singleLine(p.Description),
			Enum:        schemaEnum(p.Schema),
			Default:     schemaDefault(p.Schema),
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
func orderedContentTypes(content map[string]*oapiMediaType) []string {
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
// a schema (and thus produces flags) only when any of the request body's content
// schemas is a $ref; an inline body schema yields no field flags. (The SORACOM
// definition never uses a top-level requestBody $ref, so only the content
// schemas are inspected here.)
func requestBodyHasRefSchema(reqBody *oapiRequestBody) bool {
	if reqBody == nil {
		return false
	}
	for _, media := range reqBody.Content {
		if media != nil && media.Schema != nil && media.Schema.Ref != "" {
			return true
		}
	}
	return false
}

func buildRequestBodyDesc(reqBody *oapiRequestBody) *requestBodyDesc {
	if reqBody == nil {
		return nil
	}

	for _, contentType := range orderedContentTypes(reqBody.Content) {
		media := reqBody.Content[contentType]
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
		if media.Schema.Value.Type == "array" && media.Schema.Value.Items != nil {
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

func buildBodyProperties(schema *oapiSchema, emitOptions bool) []parameterDesc {
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
func getCLICommandsFromOperation(op *oapiOperation) []string {
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
func bodyPropertyOption(propName string, prop *oapiSchemaRef) string {
	if !bodyPropertyHasFlag(prop) {
		return ""
	}
	return lib.OptionCase(cliParamNameOf(propName, prop))
}

func bodyPropertyHasFlag(prop *oapiSchemaRef) bool {
	// The generator only flags fields whose schema has exactly one type (its
	// Types.Is / getTypeOfParam helpers reject multi-type schemas such as
	// ["string","null"]); the OpenAPI 3.0 definition carries a single `type`
	// string, so an empty type is the "no single type" case handled here.
	if prop == nil || prop.Value == nil || prop.Value.Type == "" {
		return false
	}
	switch prop.Value.Type {
	case "string", "integer", "number", "boolean":
		return true
	case "array":
		items := prop.Value.Items
		return items != nil && items.Value != nil && items.Value.Type == "string"
	}
	return false
}

// parameterHasFlag reports whether the generator produces a flag for a
// path/query/header parameter. This mirrors getStringFlags / getStringSliceFlags
// / getIntegerFlags / getFloatFlags / getBoolFlags in
// generators/cmd/src/gen_leaf_cmd.go — the generator is the source of truth, and
// TestDescribeOptionsMatchGeneratedFlags asserts describe stays aligned with it.
func parameterHasFlag(p *oapiParameter) bool {
	// Match the generator, which only flags single-type schemas; in OpenAPI 3.0
	// `type` is a single string, so an empty type is the "no single type" case.
	if p == nil || p.Schema == nil || p.Schema.Value == nil || p.Schema.Value.Type == "" {
		return false
	}
	switch p.Schema.Value.Type {
	case "string", "integer", "number", "boolean":
		return true
	case "array":
		// Only arrays of strings in the query become a (string slice) flag.
		if p.In != "query" {
			return false
		}
		items := p.Schema.Value.Items
		return items != nil && items.Value != nil && items.Value.Type == "string"
	}
	return false
}

// parameterRequired mirrors the generator's special case: operator_id is never a
// required flag because generated commands auto-fill it from the API token when
// omitted (see getStringFlags in generators/cmd/src/gen_leaf_cmd.go).
func parameterRequired(p *oapiParameter) bool {
	if p.Name == "operator_id" {
		return false
	}
	return p.Required
}

func cliParamNameOf(propName string, prop *oapiSchemaRef) string {
	if prop.Value != nil && prop.Value.Extensions != nil {
		if raw, found := prop.Value.Extensions["x-soracom-cli-param-name"]; found {
			if s, ok := raw.(string); ok {
				return s
			}
		}
	}
	return propName
}

func operationSummary(op *oapiOperation) string {
	if op == nil {
		return ""
	}
	return op.Summary
}

func schemaTypeString(s *oapiSchemaRef) string {
	if s == nil || s.Value == nil || s.Value.Type == "" {
		return ""
	}
	t := s.Value.Type
	if t == "array" && s.Value.Items != nil {
		return "array of " + schemaTypeString(s.Value.Items)
	}
	// A free-form object whose keys are arbitrary (additionalProperties) is a
	// map, not a fixed struct. Surface that so an agent knows it must supply
	// caller-defined keys (e.g. tags is "map of string") rather than looking
	// for fixed sub-properties that do not exist. Only a schema-valued
	// additionalProperties counts: for `additionalProperties: false` the Value
	// stays unset, so it is treated as a plain object.
	if t == "object" && s.Value.AdditionalProperties != nil && s.Value.AdditionalProperties.Value != nil {
		valueType := schemaTypeString(s.Value.AdditionalProperties)
		if valueType == "" {
			valueType = "object"
		}
		return "map of " + valueType
	}
	return t
}

// schemaRefName returns the component schema name a property points at, when it
// is defined via $ref (e.g. "#/components/schemas/GroupConfiguration" ->
// "GroupConfiguration"). This gives an agent a stable name to reason about for
// object-typed fields whose structure is not inlined.
func schemaRefName(s *oapiSchemaRef) string {
	if s == nil || s.Ref == "" {
		return ""
	}
	return schemaRefComponentName(s.Ref)
}

// requestBodyExample returns the example payload for a request body, preferring
// the media-type example and falling back to the schema-level example. A
// concrete example is often the most actionable thing an agent can use to build
// a correct --body, especially for free-form/object fields.
func requestBodyExample(media *oapiMediaType) interface{} {
	if media == nil {
		return nil
	}
	if media.Example != nil {
		return toJSONValue(media.Example)
	}
	if media.Schema != nil && media.Schema.Value != nil {
		return toJSONValue(media.Schema.Value.Example)
	}
	return nil
}

func schemaEnum(s *oapiSchemaRef) []interface{} {
	if s == nil || s.Value == nil || len(s.Value.Enum) == 0 {
		return nil
	}
	out := make([]interface{}, len(s.Value.Enum))
	for i, e := range s.Value.Enum {
		out[i] = toJSONValue(e)
	}
	return out
}

func schemaDefault(s *oapiSchemaRef) interface{} {
	if s == nil || s.Value == nil {
		return nil
	}
	return toJSONValue(s.Value.Default)
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
