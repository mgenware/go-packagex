package jsonx

import (
	"testing"

	"github.com/mgenware/go-packagex/test"
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
	test.Compare(t, "haha", GetStringOrDefault(sampleDict, "s"))
	test.Compare(t, "", GetStringOrDefault(sampleDict, "_"))
}

func TestGetBoolOrDefault(t *testing.T) {
	test.Compare(t, true, GetBoolOrDefault(sampleDict, "b"))
	test.Compare(t, false, GetBoolOrDefault(sampleDict, "_"))
}

func TestGetIntOrDefault(t *testing.T) {
	test.Compare(t, -1, GetIntOrDefault(sampleDict, "nn"))
	test.Compare(t, 0, GetIntOrDefault(sampleDict, "_"))
}

func TestGetUintOrDefault(t *testing.T) {
	test.Compare(t, uint(123), GetUintOrDefault(sampleDict, "pn"))
	test.Compare(t, uint(0), GetUintOrDefault(sampleDict, "_"))
}

func TestGetInt64OrDefault(t *testing.T) {
	test.Compare(t, int64(-1), GetInt64OrDefault(sampleDict, "nn"))
	test.Compare(t, int64(0), GetInt64OrDefault(sampleDict, "_"))
}

func TestGetUint64OrDefault(t *testing.T) {
	test.Compare(t, uint64(123), GetUint64OrDefault(sampleDict, "pn"))
	test.Compare(t, uint64(0), GetUint64OrDefault(sampleDict, "_"))
}

func TestGetDictOrDefault(t *testing.T) {
	d1 := GetDictOrDefault(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Compare(t, d2, d1)

	test.Compare(t, map[string]interface{}(nil), GetDictOrDefault(sampleDict, "arr"))
}

func TestGetDictOrEmpty(t *testing.T) {
	d1 := GetDictOrEmpty(sampleDict, "dict")
	d2 := make(map[string]interface{})
	d2["a"] = true
	d2["b"] = []interface{}{"element1", "element2"}
	test.Compare(t, d2, d1)

	test.Compare(t, make(map[string]interface{}), GetDictOrEmpty(sampleDict, "arr"))
}

func TestGetArrayOrDefault(t *testing.T) {
	d1 := GetArrayOrDefault(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Compare(t, d2, d1)

	test.Compare(t, []interface{}(nil), GetArrayOrDefault(sampleDict, "dict"))
}

func TestGetArrayOrEmpty(t *testing.T) {
	d1 := GetArrayOrEmpty(sampleDict, "arr")
	d2 := []interface{}{float64(1), float64(2), "haha"}
	test.Compare(t, d2, d1)

	test.Compare(t, make([]interface{}, 0), GetArrayOrEmpty(sampleDict, "dict"))
}
