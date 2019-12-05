package jsonwrap

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JSONWrapper struct {
	Verify   bool
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

func NewJSONWrapper() JSONWrapper {
	return JSONWrapper{Verify: true, Prettify: true}
}
