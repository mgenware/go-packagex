package jsonx

import (
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
)

func TestUnmarshalBytesToDict(t *testing.T) {
	json := []byte("{\"a\": 123, \"b\": false}")
	dict, err := UnmarshalBytesToDict(json)
	if err != nil {
		panic(err)
	}
	want := make(map[string]interface{})
	want["a"] = float64(123)
	want["b"] = false
	test.Assert(t, dict, want)
}
