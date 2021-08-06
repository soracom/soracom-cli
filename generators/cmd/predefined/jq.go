package cmd

import (
	"encoding/json"
	"fmt"

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
		fmt.Printf("%#v\n", v)
	}

	return nil
}
