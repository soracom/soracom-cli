package cmd

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// strSchema and friends build the minimal schema-ref values that describe's
// build* helpers operate on.
func strSchema() *oapiSchemaRef  { return &oapiSchemaRef{Value: &oapiSchema{Type: "string"}} }
func boolSchema() *oapiSchemaRef { return &oapiSchemaRef{Value: &oapiSchema{Type: "boolean"}} }

func TestFirstLine(t *testing.T) {
	if got := firstLine("  one\ntwo\n"); got != "one" {
		t.Errorf("firstLine = %q, want %q", got, "one")
	}
	if got := firstLine("single"); got != "single" {
		t.Errorf("firstLine = %q, want %q", got, "single")
	}
}

func TestSchemaTypeString(t *testing.T) {
	if got := schemaTypeString(strSchema()); got != "string" {
		t.Errorf("got %q, want string", got)
	}
	arr := &oapiSchemaRef{Value: &oapiSchema{Type: "array", Items: strSchema()}}
	if got := schemaTypeString(arr); got != "array of string" {
		t.Errorf("got %q, want 'array of string'", got)
	}
	if got := schemaTypeString(nil); got != "" {
		t.Errorf("nil schema should give empty type, got %q", got)
	}
}

func TestSingleLine(t *testing.T) {
	in := "VPG Type.\n\n- `14` : Type-E\n- `15` : Type-F\n- `242` : Type-F2\n"
	want := "VPG Type. - `14` : Type-E - `15` : Type-F - `242` : Type-F2"
	if got := singleLine(in); got != want {
		t.Errorf("singleLine collapsed wrong:\n got: %q\nwant: %q", got, want)
	}
	if got := singleLine("   one   line   "); got != "one line" {
		t.Errorf("got %q, want 'one line'", got)
	}
}

func TestSchemaTypeStringMap(t *testing.T) {
	m := &oapiSchemaRef{Value: &oapiSchema{
		Type:                 "object",
		AdditionalProperties: strSchema(),
	}}
	if got := schemaTypeString(m); got != "map of string" {
		t.Errorf("got %q, want 'map of string'", got)
	}

	mo := &oapiSchemaRef{Value: &oapiSchema{
		Type:                 "object",
		AdditionalProperties: &oapiSchemaRef{Value: &oapiSchema{Type: "object"}},
	}}
	if got := schemaTypeString(mo); got != "map of object" {
		t.Errorf("got %q, want 'map of object'", got)
	}

	// additionalProperties: false carries no schema (Value nil) and must not be
	// reported as a map.
	mf := &oapiSchemaRef{Value: &oapiSchema{
		Type:                 "object",
		AdditionalProperties: &oapiSchemaRef{},
	}}
	if got := schemaTypeString(mf); got != "object" {
		t.Errorf("got %q, want 'object'", got)
	}
}

func TestSchemaRefName(t *testing.T) {
	ref := &oapiSchemaRef{Ref: "#/components/schemas/GroupConfiguration", Value: &oapiSchema{Type: "object"}}
	if got := schemaRefName(ref); got != "GroupConfiguration" {
		t.Errorf("got %q, want GroupConfiguration", got)
	}
	if got := schemaRefName(strSchema()); got != "" {
		t.Errorf("inline schema should have no ref name, got %q", got)
	}
}

func TestBodyPropertyHasFlag(t *testing.T) {
	arrOfString := &oapiSchema{Type: "array", Items: strSchema()}
	mapObj := &oapiSchema{Type: "object", AdditionalProperties: &oapiSchemaRef{Value: &oapiSchema{Type: "object"}}}

	cases := []struct {
		name   string
		schema *oapiSchema
		want   bool
	}{
		{"string", &oapiSchema{Type: "string"}, true},
		{"integer", &oapiSchema{Type: "integer"}, true},
		{"boolean", &oapiSchema{Type: "boolean"}, true},
		{"array of string", arrOfString, true},
		{"object/map", mapObj, false},
		{"plain object", &oapiSchema{Type: "object"}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := bodyPropertyHasFlag(&oapiSchemaRef{Value: c.schema}); got != c.want {
				t.Errorf("bodyPropertyHasFlag(%s) = %v, want %v", c.name, got, c.want)
			}
		})
	}
}

func TestOrderedContentTypesPrefersJSON(t *testing.T) {
	content := map[string]*oapiMediaType{
		"application/xml":  {},
		"text/plain":       {},
		"application/json": {},
	}
	got := orderedContentTypes(content)
	want := []string{"application/json", "application/xml", "text/plain"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("position %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestBuildRequestBodyDescArray(t *testing.T) {
	item := &oapiSchema{
		Type: "object",
		Properties: map[string]*oapiSchemaRef{
			"tagName":  strSchema(),
			"tagValue": strSchema(),
		},
		Required: []string{"tagName", "tagValue"},
	}
	arr := &oapiSchemaRef{Value: &oapiSchema{
		Type:  "array",
		Items: &oapiSchemaRef{Ref: "#/components/schemas/TagUpdateRequest", Value: item},
	}}

	rb := &oapiRequestBody{
		Content: map[string]*oapiMediaType{
			"application/json": {Schema: arr},
		},
	}

	d := buildRequestBodyDesc(rb)
	if d == nil {
		t.Fatal("expected a request body description")
	}
	if d.Type != "array of TagUpdateRequest" {
		t.Errorf("type = %q, want 'array of TagUpdateRequest'", d.Type)
	}
	if d.Schema != "TagUpdateRequest" {
		t.Errorf("schema = %q, want TagUpdateRequest", d.Schema)
	}
	if len(d.Properties) != 2 {
		t.Fatalf("expected 2 element properties, got %d", len(d.Properties))
	}
	// Element fields of an array body have no per-field flag.
	for _, p := range d.Properties {
		if p.Option != "" {
			t.Errorf("array element field %q should not advertise an option, got %q", p.Name, p.Option)
		}
		if !p.Required {
			t.Errorf("field %q should be required", p.Name)
		}
	}
}

func TestBuildResponseFieldsNested(t *testing.T) {
	session := &oapiSchema{
		Type: "object",
		Properties: map[string]*oapiSchemaRef{
			"online": boolSchema(),
			"imsi":   strSchema(),
		},
	}
	sim := &oapiSchemaRef{Value: &oapiSchema{
		Type: "object",
		Properties: map[string]*oapiSchemaRef{
			"simId":         strSchema(),
			"sessionStatus": {Ref: "#/components/schemas/SessionStatus", Value: session},
		},
	}}

	fields := buildResponseFields(sim, map[*oapiSchema]bool{})
	byName := map[string]fieldDesc{}
	for _, f := range fields {
		byName[f.Name] = f
	}

	if _, ok := byName["simId"]; !ok {
		t.Fatal("missing simId")
	}
	if len(byName["simId"].Fields) != 0 {
		t.Error("scalar field should have no nested fields")
	}

	ss, ok := byName["sessionStatus"]
	if !ok {
		t.Fatal("missing sessionStatus")
	}
	if ss.Schema != "SessionStatus" {
		t.Errorf("sessionStatus schema = %q, want SessionStatus", ss.Schema)
	}
	sub := map[string]bool{}
	for _, f := range ss.Fields {
		sub[f.Name] = true
	}
	if !sub["online"] || !sub["imsi"] {
		t.Errorf("nested sessionStatus fields should include online and imsi, got %v", ss.Fields)
	}
}

func TestBuildResponseFieldsFullyExpandsDeepNesting(t *testing.T) {
	l3 := &oapiSchema{Type: "object", Properties: map[string]*oapiSchemaRef{"leaf": strSchema()}}
	l2 := &oapiSchema{Type: "object", Properties: map[string]*oapiSchemaRef{"l3": {Value: l3}}}
	l1 := &oapiSchema{Type: "object", Properties: map[string]*oapiSchemaRef{"l2": {Value: l2}}}
	arr := &oapiSchemaRef{Value: &oapiSchema{Type: "array", Items: &oapiSchemaRef{Value: l1}}}

	fields := buildResponseFields(arr, map[*oapiSchema]bool{})
	// Walk l2 -> l3 -> leaf; all three levels must be present (not just one).
	if len(fields) != 1 || fields[0].Name != "l2" {
		t.Fatalf("level 1 = %v, want [l2]", fields)
	}
	if len(fields[0].Fields) != 1 || fields[0].Fields[0].Name != "l3" {
		t.Fatalf("level 2 = %v, want [l3]", fields[0].Fields)
	}
	if len(fields[0].Fields[0].Fields) != 1 || fields[0].Fields[0].Fields[0].Name != "leaf" {
		t.Fatalf("level 3 = %v, want [leaf]", fields[0].Fields[0].Fields)
	}
}

func TestBuildResponseFieldsCycleGuard(t *testing.T) {
	node := &oapiSchema{Type: "object"}
	selfRef := &oapiSchemaRef{Ref: "#/components/schemas/Node", Value: node} // selfRef.Value == node
	node.Properties = map[string]*oapiSchemaRef{
		"name":  strSchema(),
		"child": selfRef,
	}

	// Must terminate (not recurse forever) on a self-referential schema.
	fields := buildResponseFields(&oapiSchemaRef{Value: node}, map[*oapiSchema]bool{})
	byName := map[string]fieldDesc{}
	for _, f := range fields {
		byName[f.Name] = f
	}
	if _, ok := byName["child"]; !ok {
		t.Fatal("child field should be listed")
	}
	if len(byName["child"].Fields) != 0 {
		t.Errorf("cycle guard should stop expansion of a self-referential field, got %v", byName["child"].Fields)
	}
}

func TestEntriesUnderPrefix(t *testing.T) {
	entries := []describeEntry{
		{command: "sims activate", method: "post", path: "/sims/{sim_id}/activate"},
		{command: "sims deactivate", method: "post", path: "/sims/{sim_id}/deactivate"},
		{command: "sims-x list", method: "get", path: "/sims-x"},
		{command: "auth", method: "post", path: "/auth"},
		{command: "analysis queries start", method: "post", path: "/analysis/queries"},
		{command: "analysis queries get", method: "get", path: "/analysis/queries/{id}"},
	}

	got := entriesUnderPrefix(entries, "sims")
	if len(got) != 2 {
		t.Fatalf("expected 2 entries under 'sims', got %d: %+v", len(got), got)
	}
	for _, e := range got {
		if !strings.HasPrefix(e.command, "sims ") {
			t.Errorf("unexpected entry %q under 'sims'", e.command)
		}
	}

	// A nested group prefix works too.
	if got := entriesUnderPrefix(entries, "analysis queries"); len(got) != 2 {
		t.Errorf("expected 2 entries under 'analysis queries', got %d", len(got))
	}

	// A leaf command has nothing under it.
	if got := entriesUnderPrefix(entries, "auth"); len(got) != 0 {
		t.Errorf("expected no entries under leaf 'auth', got %d", len(got))
	}

	// An unknown name matches nothing.
	if got := entriesUnderPrefix(entries, "nope"); len(got) != 0 {
		t.Errorf("expected no entries under 'nope', got %d", len(got))
	}
}

func TestTopLevelGroups(t *testing.T) {
	entries := []describeEntry{
		{command: "sims activate"},
		{command: "sims deactivate"},
		{command: "auth"},
		{command: "analysis queries start"},
	}
	summaries := map[string]string{
		"sims": "Manage SIMs.",
		"auth": "Authenticate.",
	}

	got := topLevelGroups(entries, func(name string) string { return summaries[name] })
	want := []commandGroupSummary{
		{Command: "analysis", Summary: "", Commands: 1},
		{Command: "auth", Summary: "Authenticate.", Commands: 1},
		{Command: "sims", Summary: "Manage SIMs.", Commands: 2},
	}
	if len(got) != len(want) {
		t.Fatalf("expected %d groups, got %d: %+v", len(want), len(got), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("group %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestAppendPredefinedDescribeEntries(t *testing.T) {
	openAPIEntry := describeEntry{
		command: "configure list",
		op:      &oapiOperation{Summary: "OpenAPI summary wins"},
	}

	got := appendPredefinedDescribeEntries(
		[]describeEntry{openAPIEntry},
		[]*cobra.Command{ConfigureCmd, ConfigureListCmd, DescribeCmd, Test500Cmd},
	)

	byName := map[string]describeEntry{}
	for _, e := range got {
		byName[e.command] = e
	}

	if len(byName) != 3 {
		t.Fatalf("got commands %v, want configure, configure list, and describe", byName)
	}
	if byName["configure"].cobraCommand != ConfigureCmd {
		t.Error("configure should be added from the predefined command registry")
	}
	if byName["describe"].cobraCommand != DescribeCmd {
		t.Error("describe should be added from the predefined command registry")
	}
	if byName["configure list"].op != openAPIEntry.op || byName["configure list"].cobraCommand != nil {
		t.Error("an existing OpenAPI entry should not be replaced by a predefined entry")
	}
	if _, ok := byName["test 500"]; ok {
		t.Error("a command below a hidden parent should not be described")
	}
}

func TestBuildCommandDescriptionForPredefinedCommand(t *testing.T) {
	command := &cobra.Command{
		Use:   "sample",
		Short: "Sample command.",
		Long:  "Sample command.\nWith more detail.",
	}
	command.Flags().Bool("overwrite", false, "Overwrite an existing value.")

	d := buildCommandDescription(describeEntry{
		command:      "sample",
		cobraCommand: command,
	})

	if d.Method != "" || d.Path != "" {
		t.Errorf("predefined command should not have an HTTP method or path: %+v", d)
	}
	if d.Summary != "Sample command." || d.Description != "Sample command. With more detail." {
		t.Errorf("unexpected command text: %+v", d)
	}
	if len(d.Parameters) != 1 {
		t.Fatalf("got %d parameters, want 1: %+v", len(d.Parameters), d.Parameters)
	}
	p := d.Parameters[0]
	if p.Name != "overwrite" || p.Option != "overwrite" || p.In != "flag" || p.Type != "boolean" {
		t.Errorf("unexpected flag description: %+v", p)
	}

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if strings.Contains(string(b), `"method"`) || strings.Contains(string(b), `"path"`) {
		t.Errorf("empty HTTP fields should be omitted for predefined commands: %s", b)
	}
}

func TestGetCLICommandsFromOperation(t *testing.T) {
	op := &oapiOperation{
		Extensions: map[string]interface{}{
			"x-soracom-cli": []interface{}{"groups create"},
		},
	}
	got := getCLICommandsFromOperation(op)
	if len(got) != 1 || got[0] != "groups create" {
		t.Errorf("got %#v, want [groups create]", got)
	}

	none := getCLICommandsFromOperation(&oapiOperation{})
	if none != nil {
		t.Errorf("operation without extension should return nil, got %#v", none)
	}
}

func TestBuildCommandDescription(t *testing.T) {
	bodySchema := &oapiSchema{
		Type: "object",
		Properties: map[string]*oapiSchemaRef{
			"tags": {Value: &oapiSchema{Type: "object"}},
		},
		Required: []string{"tags"},
	}

	op := &oapiOperation{
		Summary:     "Create Group",
		Description: "Create a new group.",
		Parameters: []*oapiParameter{
			{
				Name:        "group_id",
				In:          "path",
				Required:    true,
				Description: "Group ID.\nsecond line kept",
				Schema:      strSchema(),
			},
		},
		RequestBody: &oapiRequestBody{
			Content: map[string]*oapiMediaType{
				"application/json": {Schema: &oapiSchemaRef{Value: bodySchema}},
			},
		},
	}

	d := buildCommandDescription(describeEntry{
		command: "groups create",
		method:  "post",
		path:    "/groups/{group_id}",
		op:      op,
	})

	if d.Command != "groups create" || d.Method != "POST" || d.Path != "/groups/{group_id}" {
		t.Errorf("unexpected header fields: %+v", d)
	}
	if d.Summary != "Create Group" {
		t.Errorf("summary = %q", d.Summary)
	}
	if len(d.Parameters) != 1 {
		t.Fatalf("expected 1 parameter, got %d", len(d.Parameters))
	}
	p := d.Parameters[0]
	if p.Name != "group_id" || p.In != "path" || !p.Required || p.Type != "string" {
		t.Errorf("unexpected parameter: %+v", p)
	}
	if p.Description != "Group ID. second line kept" {
		t.Errorf("description should be the full text collapsed to one line, got %q", p.Description)
	}
	if d.RequestBody == nil {
		t.Fatalf("expected request body description")
	}
	if d.RequestBody.ContentType != "application/json" {
		t.Errorf("content type = %q", d.RequestBody.ContentType)
	}
	if len(d.RequestBody.Properties) != 1 || d.RequestBody.Properties[0].Name != "tags" {
		t.Fatalf("unexpected body properties: %+v", d.RequestBody.Properties)
	}
	if !d.RequestBody.Properties[0].Required {
		t.Errorf("'tags' should be required")
	}
}

// TestParseAPIDefinitionResolvesRefs exercises the yaml.v2-based loader and the
// $ref resolution pass end-to-end on a small inline document.
func TestParseAPIDefinitionResolvesRefs(t *testing.T) {
	doc, err := parseAPIDefinition([]byte(`
openapi: 3.0.0
paths:
  /groups/{group_id}:
    post:
      summary: Create Group
      x-soracom-cli:
      - groups create
      parameters:
      - name: group_id
        in: path
        required: true
        schema:
          type: string
      requestBody:
        content:
          application/json:
            example:
              tags:
                name: value
            schema:
              $ref: '#/components/schemas/GroupCreateRequest'
      responses:
        "200":
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Group'
components:
  schemas:
    GroupCreateRequest:
      type: object
      properties:
        tags:
          type: object
          additionalProperties:
            type: string
      required:
      - tags
    Group:
      type: object
      properties:
        groupId:
          type: string
        configuration:
          $ref: '#/components/schemas/Group'
`))
	if err != nil {
		t.Fatalf("parseAPIDefinition: %v", err)
	}

	pi := doc.Paths["/groups/{group_id}"]
	if pi == nil || pi.Post == nil {
		t.Fatal("expected a POST operation on /groups/{group_id}")
	}
	op := pi.Post

	if cmds := getCLICommandsFromOperation(op); len(cmds) != 1 || cmds[0] != "groups create" {
		t.Errorf("x-soracom-cli = %v, want [groups create]", cmds)
	}

	// requestBody schema $ref resolved and its property (map type) expanded.
	rb := buildRequestBodyDesc(op.RequestBody)
	if rb == nil || rb.Schema != "GroupCreateRequest" {
		t.Fatalf("request body schema = %+v, want GroupCreateRequest", rb)
	}
	if len(rb.Properties) != 1 || rb.Properties[0].Name != "tags" || rb.Properties[0].Type != "map of string" {
		t.Errorf("body property = %+v, want tags: map of string", rb.Properties)
	}
	// The example (a nested map) must survive as JSON-marshalable output.
	if _, err := json.Marshal(rb.Example); err != nil {
		t.Errorf("request body example is not JSON-marshalable: %v", err)
	}

	// response schema $ref resolved, and the self-referential field terminates.
	resp := buildResponseDesc(op.Responses)
	if resp == nil || resp.Schema != "Group" {
		t.Fatalf("response schema = %+v, want Group", resp)
	}
	byName := map[string]fieldDesc{}
	for _, f := range resp.Fields {
		byName[f.Name] = f
	}
	if _, ok := byName["groupId"]; !ok {
		t.Error("expected groupId in response fields")
	}
	if _, ok := byName["configuration"]; !ok {
		t.Error("expected self-referential configuration field to be listed")
	}
}
