package cmd

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestApplyFieldFilterObject(t *testing.T) {
	in := `{"imsi":"440101","status":"active","extra":{"a":1,"b":2}}`
	out, err := applyFieldFilter(in, []string{"imsi", "extra.a"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := mustUnmarshal(t, out)
	want := map[string]interface{}{
		"imsi":  "440101",
		"extra": map[string]interface{}{"a": json.Number("1")},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestApplyFieldFilterArray(t *testing.T) {
	in := `[{"imsi":"1","x":9},{"imsi":"2","x":8}]`
	out, err := applyFieldFilter(in, []string{"imsi"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := mustUnmarshal(t, out)
	want := []interface{}{
		map[string]interface{}{"imsi": "1"},
		map[string]interface{}{"imsi": "2"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestApplyFieldFilterNestedArray(t *testing.T) {
	in := `{"sessions":[{"online":true,"imei":"a"},{"online":false,"imei":"b"}]}`
	out, err := applyFieldFilter(in, []string{"sessions.online"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := mustUnmarshal(t, out)
	want := map[string]interface{}{
		"sessions": []interface{}{
			map[string]interface{}{"online": true},
			map[string]interface{}{"online": false},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestApplyFieldFilterMissingField(t *testing.T) {
	in := `{"imsi":"1"}`
	out, err := applyFieldFilter(in, []string{"nonexistent"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "{}" {
		t.Errorf("got %q, want %q", out, "{}")
	}
}

func mustUnmarshal(t *testing.T, s string) interface{} {
	t.Helper()
	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()
	var v interface{}
	if err := d.Decode(&v); err != nil {
		t.Fatalf("unable to unmarshal %q: %v", s, err)
	}
	return v
}
