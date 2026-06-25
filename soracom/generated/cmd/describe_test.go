package cmd

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestFirstLine(t *testing.T) {
	if got := firstLine("  one\ntwo\n"); got != "one" {
		t.Errorf("firstLine = %q, want %q", got, "one")
	}
	if got := firstLine("single"); got != "single" {
		t.Errorf("firstLine = %q, want %q", got, "single")
	}
}

func TestSchemaTypeString(t *testing.T) {
	if got := schemaTypeString(openapi3.NewSchemaRef("", openapi3.NewStringSchema())); got != "string" {
		t.Errorf("got %q, want string", got)
	}
	arr := openapi3.NewArraySchema()
	arr.Items = openapi3.NewSchemaRef("", openapi3.NewStringSchema())
	if got := schemaTypeString(openapi3.NewSchemaRef("", arr)); got != "array of string" {
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
	m := openapi3.NewObjectSchema()
	m.AdditionalProperties = openapi3.AdditionalProperties{
		Schema: openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
	}
	if got := schemaTypeString(openapi3.NewSchemaRef("", m)); got != "map of string" {
		t.Errorf("got %q, want 'map of string'", got)
	}

	mo := openapi3.NewObjectSchema()
	mo.AdditionalProperties = openapi3.AdditionalProperties{
		Schema: openapi3.NewSchemaRef("", openapi3.NewObjectSchema()),
	}
	if got := schemaTypeString(openapi3.NewSchemaRef("", mo)); got != "map of object" {
		t.Errorf("got %q, want 'map of object'", got)
	}
}

func TestSchemaRefName(t *testing.T) {
	ref := openapi3.NewSchemaRef("#/components/schemas/GroupConfiguration", openapi3.NewObjectSchema())
	if got := schemaRefName(ref); got != "GroupConfiguration" {
		t.Errorf("got %q, want GroupConfiguration", got)
	}
	if got := schemaRefName(openapi3.NewSchemaRef("", openapi3.NewStringSchema())); got != "" {
		t.Errorf("inline schema should have no ref name, got %q", got)
	}
}

func TestBodyPropertyHasFlag(t *testing.T) {
	arrOfString := openapi3.NewArraySchema()
	arrOfString.Items = openapi3.NewSchemaRef("", openapi3.NewStringSchema())

	mapObj := openapi3.NewObjectSchema()
	mapObj.AdditionalProperties = openapi3.AdditionalProperties{
		Schema: openapi3.NewSchemaRef("", openapi3.NewObjectSchema()),
	}

	cases := []struct {
		name   string
		schema *openapi3.Schema
		want   bool
	}{
		{"string", openapi3.NewStringSchema(), true},
		{"integer", openapi3.NewIntegerSchema(), true},
		{"boolean", openapi3.NewBoolSchema(), true},
		{"array of string", arrOfString, true},
		{"object/map", mapObj, false},
		{"plain object", openapi3.NewObjectSchema(), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := bodyPropertyHasFlag(openapi3.NewSchemaRef("", c.schema)); got != c.want {
				t.Errorf("bodyPropertyHasFlag(%s) = %v, want %v", c.name, got, c.want)
			}
		})
	}
}

func TestOrderedContentTypesPrefersJSON(t *testing.T) {
	content := openapi3.Content{
		"application/xml":  openapi3.NewMediaType(),
		"text/plain":       openapi3.NewMediaType(),
		"application/json": openapi3.NewMediaType(),
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
	item := openapi3.NewObjectSchema()
	item.Properties = openapi3.Schemas{
		"tagName":  openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
		"tagValue": openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
	}
	item.Required = []string{"tagName", "tagValue"}

	arr := openapi3.NewArraySchema()
	arr.Items = openapi3.NewSchemaRef("#/components/schemas/TagUpdateRequest", item)

	rb := &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody().WithJSONSchemaRef(openapi3.NewSchemaRef("", arr)),
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
	session := openapi3.NewObjectSchema()
	session.Properties = openapi3.Schemas{
		"online": openapi3.NewSchemaRef("", openapi3.NewBoolSchema()),
		"imsi":   openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
	}
	sim := openapi3.NewObjectSchema()
	sim.Properties = openapi3.Schemas{
		"simId":         openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
		"sessionStatus": openapi3.NewSchemaRef("#/components/schemas/SessionStatus", session),
	}

	fields := buildResponseFields(openapi3.NewSchemaRef("", sim), map[*openapi3.Schema]bool{})
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
	l3 := openapi3.NewObjectSchema()
	l3.Properties = openapi3.Schemas{"leaf": openapi3.NewSchemaRef("", openapi3.NewStringSchema())}
	l2 := openapi3.NewObjectSchema()
	l2.Properties = openapi3.Schemas{"l3": openapi3.NewSchemaRef("", l3)}
	l1 := openapi3.NewObjectSchema()
	l1.Properties = openapi3.Schemas{"l2": openapi3.NewSchemaRef("", l2)}
	arr := openapi3.NewArraySchema()
	arr.Items = openapi3.NewSchemaRef("", l1)

	fields := buildResponseFields(openapi3.NewSchemaRef("", arr), map[*openapi3.Schema]bool{})
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
	node := openapi3.NewObjectSchema()
	selfRef := openapi3.NewSchemaRef("#/components/schemas/Node", node) // selfRef.Value == node
	node.Properties = openapi3.Schemas{
		"name":  openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
		"child": selfRef,
	}

	// Must terminate (not recurse forever) on a self-referential schema.
	fields := buildResponseFields(openapi3.NewSchemaRef("", node), map[*openapi3.Schema]bool{})
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

func TestGetCLICommandsFromOperation(t *testing.T) {
	op := &openapi3.Operation{
		Extensions: map[string]interface{}{
			"x-soracom-cli": []interface{}{"groups create"},
		},
	}
	got := getCLICommandsFromOperation(op)
	if len(got) != 1 || got[0] != "groups create" {
		t.Errorf("got %#v, want [groups create]", got)
	}

	none := getCLICommandsFromOperation(&openapi3.Operation{})
	if none != nil {
		t.Errorf("operation without extension should return nil, got %#v", none)
	}
}

func TestBuildCommandDescription(t *testing.T) {
	bodySchema := openapi3.NewObjectSchema()
	bodySchema.Properties = openapi3.Schemas{
		"tags": openapi3.NewSchemaRef("", openapi3.NewObjectSchema()),
	}
	bodySchema.Required = []string{"tags"}

	op := &openapi3.Operation{
		Summary:     "Create Group",
		Description: "Create a new group.",
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name:        "group_id",
				In:          "path",
				Required:    true,
				Description: "Group ID.\nsecond line kept",
				Schema:      openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
			}},
		},
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(bodySchema),
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
