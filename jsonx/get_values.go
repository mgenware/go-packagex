package jsonx

// GetStringOrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetStringOrDefault(dict map[string]interface{}, key string) string {
	val, _ := dict[key].(string)
	return val
}

// GetFloat64OrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetFloat64OrDefault(dict map[string]interface{}, key string) float64 {
	val, _ := dict[key].(float64)
	return val
}

// GetIntOrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetIntOrDefault(dict map[string]interface{}, key string) int {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return int(val)
}

// GetUintOrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetUintOrDefault(dict map[string]interface{}, key string) uint {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return uint(val)
}

// GetInt64OrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetInt64OrDefault(dict map[string]interface{}, key string) int64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return int64(val)
}

// GetUint64OrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetUint64OrDefault(dict map[string]interface{}, key string) uint64 {
	// All numbers parsed from JSON are float64
	val := GetFloat64OrDefault(dict, key)
	return uint64(val)
}

// GetBoolOrDefault retrieves the value for the given key in the dictionary, or a default value if it does not exist.
func GetBoolOrDefault(dict map[string]interface{}, key string) bool {
	val, _ := dict[key].(bool)
	return val
}

// GetDictOrDefault retrieves the value for the given key in the dictionary, or nil if it does not exist.
func GetDictOrDefault(dict map[string]interface{}, key string) map[string]interface{} {
	val, _ := dict[key].(map[string]interface{})
	return val
}

// GetDictOrEmpty retrieves the value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetDictOrEmpty(dict map[string]interface{}, key string) map[string]interface{} {
	val := GetDictOrDefault(dict, key)
	if val == nil {
		return make(map[string]interface{})
	}
	return val
}

// GetArrayOrDefault retrieves the value for the given key in the dictionary, or nil if it does not exist.
func GetArrayOrDefault(dict map[string]interface{}, key string) []interface{} {
	val, _ := dict[key].([]interface{})
	return val
}

// GetArrayOrEmpty retrieves the value for the given key in the dictionary, or an empty dictionary if it does not exist.
func GetArrayOrEmpty(dict map[string]interface{}, key string) []interface{} {
	val := GetArrayOrDefault(dict, key)
	if val == nil {
		return make([]interface{}, 0)
	}
	return val
}
