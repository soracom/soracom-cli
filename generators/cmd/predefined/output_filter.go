package cmd

import (
	"encoding/json"
	"strings"
)

// applyFieldFilter projects the JSON response so that only the requested fields
// remain. This lets AI agents (and scripts) control the size of the output to
// protect their context window. Each entry in `fields` is a dot-separated path
// (e.g. "imsi" or "sessionStatus.online"). When the top-level value is an
// array, the filter is applied to each element.
func applyFieldFilter(rawJSON string, fields []string) (string, error) {
	var obj interface{}
	d := json.NewDecoder(strings.NewReader(rawJSON))
	d.UseNumber()
	if err := d.Decode(&obj); err != nil {
		return "", err
	}

	filtered := filterValue(obj, fields)

	b, err := marshalJSONUnescaped(filtered)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func filterValue(v interface{}, fields []string) interface{} {
	switch t := v.(type) {
	case []interface{}:
		res := make([]interface{}, 0, len(t))
		for _, e := range t {
			res = append(res, filterValue(e, fields))
		}
		return res
	case map[string]interface{}:
		res := map[string]interface{}{}
		for _, f := range fields {
			f = strings.TrimSpace(f)
			if f == "" {
				continue
			}
			projectField(t, strings.Split(f, "."), res)
		}
		return res
	default:
		return v
	}
}

// projectField copies the value addressed by `path` from `src` into `dst`,
// preserving the nesting structure. Intermediate arrays of objects are
// descended into element by element.
func projectField(src map[string]interface{}, path []string, dst map[string]interface{}) {
	if len(path) == 0 {
		return
	}

	key := path[0]
	val, ok := src[key]
	if !ok {
		return
	}

	if len(path) == 1 {
		dst[key] = val
		return
	}

	switch child := val.(type) {
	case map[string]interface{}:
		next, _ := dst[key].(map[string]interface{})
		if next == nil {
			next = map[string]interface{}{}
			dst[key] = next
		}
		projectField(child, path[1:], next)
	case []interface{}:
		arr := make([]interface{}, 0, len(child))
		for _, e := range child {
			em, ok := e.(map[string]interface{})
			if !ok {
				continue
			}
			elem := map[string]interface{}{}
			projectField(em, path[1:], elem)
			arr = append(arr, elem)
		}
		dst[key] = arr
	}
}
