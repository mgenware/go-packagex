package jsonx

import (
	"testing"
)

func TestUnmarshalBytesToDict(t *testing.T) {
	json := []byte("{\"a\": 123, \"b\": false}")
	dict, err := UnmarshalBytesToDict(json)
	if err != nil {
		panic(err)
	}
	if len(dict) != 2 || dict["a"] != float64(123) || dict["b"] != false {
		t.Errorf("Wrong dict")
	}
}
