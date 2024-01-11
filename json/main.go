package mz_json

import (
	"encoding/json"
	"io"
	"strings"
)

// The function Decode decodes a JSON string into a map[string]interface{} and optionally checks if the
// JSON format is valid.
func Decode(value string, isCheckJsonFormat bool) (map[string]interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(value))
	var result map[string]interface{}
	if isCheckJsonFormat {
		for {
			if err := decoder.Decode(&result); err == io.EOF { // end of file
				break
			} else if err != nil {
				// panic(err)
				return nil, err
			}
		}
	} else {
		err := decoder.Decode(&result)
		if err != nil {
			// panic(err)
			return nil, err
		}
	}
	return result, nil
}
