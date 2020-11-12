package jsonx

import (
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

var sampleDict map[string]interface{}

func init() {
	dict, err := UnmarshalBytesToDict([]byte(`{"pn": 123, "nn": -1, "f": 3.1415, "s": "haha", "b": true, "dict": {"a": true, "b": ["element1", "element2"]}, "arr": [1, 2, "haha"]}`))
	if err != nil {
		panic(err)
	}
	sampleDict = dict
}

func TestGetStringOrDefault(t *testing.T) {
	test.Assert(t, GetStringOrDefault(sampleDict, "s"), "haha")
	test.Assert(t, GetStringOrDefault(sampleDict, "_"), "")
}

func TestGetBoolOrDefault(t *testing.T) {
	test.Assert(t, GetBoolOrDefault(sampleDict, "b"), true)
	test.Assert(t, GetBoolOrDefault(sampleDict, "_"), false)
}

func TestGetIntOrDefault(t *testing.T) {
	test.Assert(t, GetIntOrDefault(sampleDict, "nn"), -1)
	test.Assert(t, GetIntOrDefault(sampleDict, "_"), 0)
}

func TestGetUintOrDefault(t *testing.T) {
	test.Assert(t, GetUintOrDefault(sampleDict, "pn"), uint(123))
	test.Assert(t, GetUintOrDefault(sampleDict, "_"), uint(0))
}

func TestGetInt64OrDefault(t *testing.T) {
	test.Assert(t, GetInt64OrDefault(sampleDict, "nn"), int64(-1))
	test.Assert(t, GetInt64OrDefault(sampleDict, "_"), int64(0))
}

func TestGetUint64OrDefault(t *testing.T) {
	test.Assert(t, GetUint64OrDefault(sampleDict, "pn"), uint64(123))
	test.Assert(t, GetUint64OrDefault(sampleDict, "_"), uint64(0))
}

func TestGetDictOrDefault(t *testing.T) {
	d1 := GetDictOrDefault(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDictOrDefault(sampleDict, "arr"), map[string]interface{}(nil))
}

func TestGetDictOrEmpty(t *testing.T) {
	d1 := GetDictOrEmpty(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetDictOrEmpty(sampleDict, "arr"), make(map[string]interface{}))
}

func TestGetArrayOrDefault(t *testing.T) {
	d1 := GetArrayOrDefault(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArrayOrDefault(sampleDict, "dict"), []interface{}(nil))
}

func TestGetArrayOrEmpty(t *testing.T) {
	d1 := GetArrayOrEmpty(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Assert(t, d1, d2)

	test.Assert(t, GetArrayOrEmpty(sampleDict, "dict"), make([]interface{}, 0))
}
