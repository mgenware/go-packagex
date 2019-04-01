package jsonx

import (
	"testing"
)

var sampleDict map[string]interface{}

func init() {
	dict, err := UnmarshalBytesToDict([]byte(`{"pn": 123, "nn": -1, "f": 3.1415, "s": "haha", "b": true}`))
	if err != nil {
		panic(err)
	}
	sampleDict = dict
}

func TestGetStringOrDefault(t *testing.T) {
	if val := GetStringOrDefault(sampleDict, "s"); val != "haha" {
		t.Error("value should be haha")
	}
	if val := GetStringOrDefault(sampleDict, "_"); val != "" {
		t.Error("value should be <empty>")
	}
}

func TestGetBoolOrDefault(t *testing.T) {
	if val := GetBoolOrDefault(sampleDict, "b"); val != true {
		t.Error("value should be true")
	}
	if val := GetBoolOrDefault(sampleDict, "_"); val != false {
		t.Error("value should be false")
	}
}

func TestGetIntOrDefault(t *testing.T) {
	if val := GetIntOrDefault(sampleDict, "nn"); val != -1 {
		t.Error("value should be -1")
	}
	if val := GetIntOrDefault(sampleDict, "_"); val != 0 {
		t.Error("value should be 0")
	}
}

func TestGetUintOrDefault(t *testing.T) {
	if val := GetUintOrDefault(sampleDict, "pn"); val != 123 {
		t.Error("value should be 123")
	}
	if val := GetUintOrDefault(sampleDict, "_"); val != 0 {
		t.Error("value should be 0")
	}
}

func TestGetInt64OrDefault(t *testing.T) {
	if val := GetInt64OrDefault(sampleDict, "nn"); val != -1 {
		t.Error("value should be -1")
	}
	if val := GetInt64OrDefault(sampleDict, "_"); val != 0 {
		t.Error("value should be 0")
	}
}

func TestGetUint64OrDefault(t *testing.T) {
	if val := GetUint64OrDefault(sampleDict, "pn"); val != 123 {
		t.Error("value should be 123")
	}
	if val := GetUint64OrDefault(sampleDict, "_"); val != 0 {
		t.Error("value should be 0")
	}
}
