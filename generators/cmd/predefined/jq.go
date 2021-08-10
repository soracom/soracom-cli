package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/itchyny/gojq"
)

func processJQ(jqString, responseBody string) error {
	q, err := gojq.Parse(jqString)
	if err != nil {
		return err
	}

	var j interface{}
	err = json.Unmarshal([]byte(responseBody), &j)
	if err != nil {
		return err
	}

	iter := q.Run(j)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}

		if text, err := jsonScalarToString(v); err == nil {
			fmt.Printf(text)
		} else {
			var jsonFragment []byte
			jsonFragment, err = json.Marshal(v)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", jsonFragment)
		}
	}

	return nil
}

func jsonScalarToString(input interface{}) (string, error) {
	switch tt := input.(type) {
	case string:
		return tt, nil
	case float64:
		if math.Trunc(tt) == tt {
			return strconv.FormatFloat(tt, 'f', 0, 64), nil
		} else {
			return strconv.FormatFloat(tt, 'f', 2, 64), nil
		}
	case nil:
		return "", nil
	case bool:
		return fmt.Sprintf("%v", tt), nil
	default:
		return "", fmt.Errorf("cannot convert type to string: %v", tt)
	}
}
