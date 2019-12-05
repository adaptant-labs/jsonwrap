// Package jsonwrap provides a simple API for JSON object nesting.
package jsonwrap

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// A JSONWrapper instance provides JSON wrapping functionality with specific formatting and verification options.
type JSONWrapper struct {
	// Verify the provided JSON input string before attempting to wrap it
	Verify   bool

	// Pretty-format the resulting JSON output
	Prettify bool
}

func validJson(jsonStr string) bool {
	var input interface{}
	return json.Unmarshal([]byte(jsonStr), &input) == nil
}

func prettifyJson(jsonStr string) string {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(jsonStr), "", "    ")
	if err != nil {
		fmt.Println("failed to prettify JSON")
		return jsonStr
	}

	return string(prettyJson.Bytes())
}

// Wrap handles the actual wrapping of the complete child JSON string within the parent object.
func (j JSONWrapper) Wrap(parent string, child string) string {
	if j.Verify && !validJson(child) {
		fmt.Println("invalid input JSON")
		return child
	}

	wrapped := fmt.Sprintf("{ \"%s\": %s }", parent, child)
	if j.Prettify {
		return prettifyJson(wrapped)
	}

	return wrapped
}

// NewJSONWrapper instantiates a new JSONWrapper instance with the default configuration.
func NewJSONWrapper() JSONWrapper {
	return JSONWrapper{Verify: true, Prettify: true}
}
