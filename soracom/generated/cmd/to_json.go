package cmd

import "encoding/json"

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
