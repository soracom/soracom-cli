package cmd

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// A minimal OpenAPI 3.0 reader for the 'describe' command: it decodes only the
// subset of the embedded API definition that describe reports, using the yaml
// package the CLI already depends on rather than a full OpenAPI parser.
//
// The SORACOM definition is OpenAPI 3.0.0, so a schema's `type` is a single
// string and the only `$ref` used is `#/components/schemas/<Name>`, resolved by
// resolveSchemaRefs.

// oapiDoc is the top-level OpenAPI document, reduced to what describe reads.
// Component schemas are held as schema-refs, not plain schemas, because a
// component may itself be a `$ref` to another component; such chains are
// followed transitively during resolution.
type oapiDoc struct {
	Paths      map[string]*oapiPathItem `yaml:"paths"`
	Components struct {
		Schemas map[string]*oapiSchemaRef `yaml:"schemas"`
	} `yaml:"components"`
}

// oapiPathItem holds the operations for a single path. Only HTTP-method fields
// are modeled; path-level `parameters`/`summary`/etc. are ignored because
// describe does not use them.
type oapiPathItem struct {
	Get     *oapiOperation `yaml:"get"`
	Put     *oapiOperation `yaml:"put"`
	Post    *oapiOperation `yaml:"post"`
	Delete  *oapiOperation `yaml:"delete"`
	Patch   *oapiOperation `yaml:"patch"`
	Head    *oapiOperation `yaml:"head"`
	Options *oapiOperation `yaml:"options"`
}

// operations returns the defined (method, operation) pairs in a stable order so
// describe output is deterministic. Method names are lower-case, as
// describeEntry.method expects.
func (p *oapiPathItem) operations() []struct {
	method string
	op     *oapiOperation
} {
	all := []struct {
		method string
		op     *oapiOperation
	}{
		{"get", p.Get}, {"put", p.Put}, {"post", p.Post}, {"delete", p.Delete},
		{"patch", p.Patch}, {"head", p.Head}, {"options", p.Options},
	}
	result := all[:0:0]
	for _, e := range all {
		if e.op != nil {
			result = append(result, e)
		}
	}
	return result
}

// oapiOperation is a single API operation. Extensions captures every unmodeled
// key (via yaml inline), which is where the `x-soracom-cli` command mapping and
// other x-* fields live.
type oapiOperation struct {
	Summary     string                   `yaml:"summary"`
	Description string                   `yaml:"description"`
	Deprecated  bool                     `yaml:"deprecated"`
	Parameters  []*oapiParameter         `yaml:"parameters"`
	RequestBody *oapiRequestBody         `yaml:"requestBody"`
	Responses   map[string]*oapiResponse `yaml:"responses"`
	Extensions  map[string]interface{}   `yaml:",inline"`
}

// oapiParameter is a path/query/header parameter. The SORACOM definition never
// uses a `$ref` for parameters (components.parameters is empty), so a parameter
// is always an inline object.
type oapiParameter struct {
	Name        string         `yaml:"name"`
	In          string         `yaml:"in"`
	Required    bool           `yaml:"required"`
	Description string         `yaml:"description"`
	Schema      *oapiSchemaRef `yaml:"schema"`
}

// oapiRequestBody is a request body. The definition never uses a top-level
// requestBody `$ref` (components.requestBodies is not referenced), so only the
// content is modeled; per-content-schema `$ref`s are what carry the payload
// shape and are handled via oapiSchemaRef.
type oapiRequestBody struct {
	Content map[string]*oapiMediaType `yaml:"content"`
}

// oapiResponse is a single response entry (keyed by status code in the map).
type oapiResponse struct {
	Content map[string]*oapiMediaType `yaml:"content"`
}

// oapiMediaType is a content entry (e.g. application/json).
type oapiMediaType struct {
	Schema  *oapiSchemaRef `yaml:"schema"`
	Example interface{}    `yaml:"example"`
}

// oapiSchemaRef is a schema position: either a `$ref` to a component or an
// inline schema. It keeps both the original Ref string (for the component name)
// and a resolved Value (the schema itself). For a `$ref`, Value is filled in by
// resolveSchemaRefs and points at the shared component schema, so multiple refs
// to the same component share one *oapiSchema pointer — the pointer-identity
// cycle guard in buildResponseFields relies on this.
type oapiSchemaRef struct {
	Ref   string
	Value *oapiSchema
}

// UnmarshalYAML distinguishes the three shapes a schema position can take in the
// document: a `$ref` mapping, an inline schema mapping, or a bare bool (only
// `additionalProperties: false/true`, which carries no schema — left empty).
func (r *oapiSchemaRef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// additionalProperties may be a bool; such a node has no schema.
	var b bool
	if err := unmarshal(&b); err == nil {
		return nil
	}

	var probe struct {
		Ref string `yaml:"$ref"`
	}
	if err := unmarshal(&probe); err != nil {
		return err
	}
	if probe.Ref != "" {
		r.Ref = probe.Ref
		return nil
	}

	var s oapiSchema
	if err := unmarshal(&s); err != nil {
		return err
	}
	r.Value = &s
	return nil
}

// oapiSchema is a JSON schema, reduced to the fields describe reports.
type oapiSchema struct {
	Type                 string                    `yaml:"type"`
	Description          string                    `yaml:"description"`
	Properties           map[string]*oapiSchemaRef `yaml:"properties"`
	Items                *oapiSchemaRef            `yaml:"items"`
	Required             []string                  `yaml:"required"`
	Enum                 []interface{}             `yaml:"enum"`
	Default              interface{}               `yaml:"default"`
	Example              interface{}               `yaml:"example"`
	AdditionalProperties *oapiSchemaRef            `yaml:"additionalProperties"`
	Extensions           map[string]interface{}    `yaml:",inline"`
}

// parseAPIDefinition decodes an embedded OpenAPI definition and resolves its
// schema `$ref`s so that every referenced schema position has a Value.
func parseAPIDefinition(b []byte) (*oapiDoc, error) {
	var doc oapiDoc
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	doc.resolveSchemaRefs()
	return &doc, nil
}

// resolveSchemaRefs walks every schema position in the document and, for each
// `$ref`, points its Value at the shared component schema (following component
// `$ref` chains transitively). A visited set of schema-ref pointers keeps a
// self-referential schema graph from looping.
func (doc *oapiDoc) resolveSchemaRefs() {
	seen := map[*oapiSchemaRef]bool{}

	var walkRef func(r *oapiSchemaRef)
	var walkSchema func(s *oapiSchema)

	walkRef = func(r *oapiSchemaRef) {
		if r == nil || seen[r] {
			return
		}
		seen[r] = true
		if r.Ref != "" {
			r.Value = doc.terminalSchema(r, 0)
		}
		walkSchema(r.Value)
	}

	walkSchema = func(s *oapiSchema) {
		if s == nil {
			return
		}
		walkRef(s.Items)
		walkRef(s.AdditionalProperties)
		for _, p := range s.Properties {
			walkRef(p)
		}
	}

	for _, c := range doc.Components.Schemas {
		walkRef(c)
	}
	for _, pi := range doc.Paths {
		if pi == nil {
			continue
		}
		for _, e := range pi.operations() {
			for _, p := range e.op.Parameters {
				if p != nil {
					walkRef(p.Schema)
				}
			}
			if e.op.RequestBody != nil {
				for _, mt := range e.op.RequestBody.Content {
					if mt != nil {
						walkRef(mt.Schema)
					}
				}
			}
			for _, resp := range e.op.Responses {
				if resp != nil {
					for _, mt := range resp.Content {
						if mt != nil {
							walkRef(mt.Schema)
						}
					}
				}
			}
		}
	}
}

// terminalSchema resolves a schema-ref to the concrete schema it denotes,
// following component `$ref` chains (a component whose definition is itself a
// `$ref`). It returns nil if a referenced component is missing; the depth guard
// stops a pathological ref cycle (e.g. A -> B -> A).
func (doc *oapiDoc) terminalSchema(r *oapiSchemaRef, depth int) *oapiSchema {
	if r == nil || depth > 100 {
		return nil
	}
	if r.Ref != "" {
		return doc.terminalSchema(doc.Components.Schemas[schemaRefComponentName(r.Ref)], depth+1)
	}
	return r.Value
}

// schemaRefComponentName returns the component name of a `$ref` string, e.g.
// "#/components/schemas/GroupConfiguration" -> "GroupConfiguration".
func schemaRefComponentName(ref string) string {
	for i := len(ref) - 1; i >= 0; i-- {
		if ref[i] == '/' {
			return ref[i+1:]
		}
	}
	return ref
}

// toJSONValue converts a value decoded by gopkg.in/yaml.v2 into JSON-marshalable
// types. yaml.v2 decodes nested mappings as map[interface{}]interface{}, which
// encoding/json cannot marshal; example/enum/default values flow verbatim into
// describe's JSON output, so they must be normalized to map[string]interface{}.
func toJSONValue(v interface{}) interface{} {
	switch x := v.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(x))
		for k, val := range x {
			m[fmt.Sprintf("%v", k)] = toJSONValue(val)
		}
		return m
	case map[string]interface{}:
		m := make(map[string]interface{}, len(x))
		for k, val := range x {
			m[k] = toJSONValue(val)
		}
		return m
	case []interface{}:
		s := make([]interface{}, len(x))
		for i, val := range x {
			s[i] = toJSONValue(val)
		}
		return s
	default:
		return v
	}
}
