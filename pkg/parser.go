package betterjson

import (
	"encoding/json"
)

// Unmarshal parses the JSON-encoded data with comment support and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	// Replace comments with ' '
	processedData := decomment(data)

	// Parse the JSON
	return json.Unmarshal(processedData, v)
}
